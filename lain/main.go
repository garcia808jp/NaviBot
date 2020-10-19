// NaviBot: Discord bot for digital assistance
// LainBot commands

package lain

// module project site
const codeURL string = "https://github.com/phossil/NaviBot/"

// Doc struct
// provides access to commands and information about them
type Doc struct {
	name        string
	synopsis    string
	description string
	example     string
	origin      string
	Exec        func([]string) string
}

// CodeDoc documentation
// Define the help command
var codeDoc = Doc{
	name:        "code - link to module project site",
	synopsis:    "code",
	description: "WIP",
	example:     "WIP",
	origin:      "built-in, lain",
	Exec:        code,
}

var (
	// CommandList map
	// contains strings and Doc structs to register commands for the handler
	CommandList = map[string]Doc{"code": codeDoc}

	// empty helpDoc struct
	// placed here to prevent an initializition loop
	helpDoc Doc
)

func init() {
	// helpDoc struct
	// Define the help command
	helpDoc = Doc{
		name:        "help - display help content",
		synopsis:    "help __command__",
		description: "The help command provides information on other commands available within the bot. It is vaguely modelled in the style of *nix man pages.",
		example:     "`man`, like the other built-in commands, is called after the configured prefix:\t`$PREFIX help help`",
		origin:      "built-in, lain",
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
