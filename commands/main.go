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

var helpDoc command

func init() {
	// HelpDoc documentation
	// Define the help command
	helpDoc = command{
		name:        "help - display help content",
		synopsis:    "help __command__",
		description: "The help command provides documentation on other commands available within the bot. It is vaguely modelled in the style of *nix man pages.",
		example:     "`man`, like the other built-in commands, is called after the configured prefix:\t`$PREFIX help help`",
		origin:      "built-in",
		Exec:        help,
	}
}

// CommandList map
// contains strings and Doc structs to register commands for the handler
var CommandList = map[string]command{"help": helpDoc}

// Help command
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
