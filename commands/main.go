// NaviBot: Discord bot for digital assistance
// Bult-in commands

package commands

// command struct
// provides access to commands and information about them
type command struct {
	name        string
	synopsis    string
	description string
	example     string
	origin      string
	Exec        func([]string) string
}

// module project site
const codeURL string = "https://github.com/phossil/NaviBot"

// CodeDoc documentation
// Define the help command
var codeDoc = command{
	name:        "code - link to project site",
	synopsis:    "code",
	description: "WIP",
	example:     "WIP",
	origin:      "built-in",
	Exec:        code,
}

// CommandList map
// contains strings and Doc structs to register commands for the handler
var CommandList = map[string]command{"code": codeDoc}

// empty helpDoc struct
// placed here to prevent an initializition loop
var helpDoc command

func init() {
	// helpDoc documentation
	// Define the help command
	helpDoc = command{
		name:        "help - display help content",
		synopsis:    "help __command__",
		description: "The help command provides documentation on other commands available within the bot. It is vaguely modelled in the style of *nix man pages.",
		example:     "`man`, like the other built-in commands, is called after the configured prefix:\t`$PREFIX help help`",
		origin:      "built-in",
		Exec:        help,
	}
	CommandList["help"] = helpDoc
}

// help command
// provides documentation of commands available in the bot
func help(msgArray []string) (msgOut string) {
	// If there are no arguments notify the user
	if len(msgArray) == 2 {
		return helpDoc.name
	}
	if len(msgArray) >= 3 {
		_, err := CommandList[msgArray[2]]
		if err == false {
			msgOut = "command not found"
			return
		}
		return CommandList[msgArray[2]].name
	}
	return
}

// code command
// links to module project site
func code([]string) (msgOut string) {
	return codeURL
}
