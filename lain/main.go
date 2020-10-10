// NaviBot: Discord bot for digital assistance
// LainBot commands

package lain

// Doc struct
// provides access to commands and information about them
type Doc struct {
	Name        string
	Synopsis    string
	Description string
	Example     string
	Origin      string
	Exec        func([]string) string
}

// HelpDoc documentation
// Define the help command
var helpDoc = Doc{
	Name:     "help - display help content",
	Synopsis: "help __command__",
	Description: "The help command provides documentation on other commands available " +
		"within the bot. It is vaguely modelled in the style of *nix man pages.",
	Example: "`man`, like the other built-in commands, is called after the configured prefix:" +
		"\t`$PREFIX help help`",
	Origin: "built-in, lain",
}

// Module project site
const codeURL string = "https://github.com/phossil/NaviBot"

// CodeDoc documentation
// Define the help command
var codeDoc = Doc{
	Name:        "code - link to module project site",
	Synopsis:    "code",
	Description: "WIP",
	Example:     "WIP",
	Origin:      "built-in, lain",
}

// CommandList map
// contains strings and Doc structs to register commands for the handler
var CommandList = map[string]Doc{"help": helpDoc, "code": codeDoc}

// Help command
// provides documentation of commands available in the bot
func Help(msgArray []string) (msgOut string) {
	// If there are no arguments notify the user
	if len(msgArray) == 2 {
		return helpDoc.Name
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

// Code command
// links to module project site
func Code([]string) (msgOut string) {
	return codeURL
}
