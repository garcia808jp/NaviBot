// NaviBot: Discord bot for digital assistance
// Bult-in commands

package commands

// Doc struct
// contains documentaion about commands, *nix man style
type Doc struct {
	Name        string
	Synopsis    string
	Description string
	Example     string
	Origin      string
}

// HelpDoc documentaion
// Define the help command
var HelpDoc = Doc{
	Name:     "help - display help content",
	Synopsis: "help __command__",
	Description: "The help command provides documentaion on other commands available " +
		"within the bot. It is vaguely modelled in the style of *nix man pages.",
	Example: "`man`, like the other built-in commands, is called after the configured prefix:" +
		"\t`$PREFIX help help`",
	Origin: "built-in",
}

// CommandList map
// contains strings and Doc structs to register commands for the handler
var CommandList = map[string]Doc{"help": HelpDoc}

// Help command
// provides documentaion of commands available in the bot
func Help(msgArray []string) (msgOut string) {
	// If there are no arguments notify the user
	if len(msgArray) == 2 {
		return HelpDoc.Name
	}
	if len(msgArray) >= 3 {
		_, err := CommandList[msgArray[2]]
		if err == false {
			msgOut = "command not found"
			return
		}
		return CommandList[msgArray[2]].Name
	}
	return
}
