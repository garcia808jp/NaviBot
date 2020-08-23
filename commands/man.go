// NaviBot: Discord bot for digital assistance
// man commands

package commands

// Register the command for the CommandList
func init() {
	ManDoc := Doc{
		Name:        "man - search for *nix manual pages from online sources",
		Synopsis:    "man __query__",
		Description: "WIP",
		Example:     "WIP",
		Origin:      "built-in",
	}

	CommandList["man"] = ManDoc
}

// Man command
// returns a string containing the requested man page from the message array
func Man(msgArray []string) (msgOut string) {
	// If the message contains arguments, complete the task; notify the user otherwise
	if len(msgArray) >= 3 {
		// The default URL for the online man pages
		manSrc := "https://jlk.fjfi.cvut.cz/arch/manpages/search?q="
		// Concatenate the URL with the requested man page
		msgOut = manSrc + msgArray[2]
	} else {
		msgOut = "command not implemented yet"
	}
	return
}
