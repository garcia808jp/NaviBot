// NaviBot: Discord bot for digital assistance
// man command

package commands

// Register the command for the CommandList
func init() {
	manDoc := command{
		name:        "man - search for *nix manual pages from online sources",
		synopsis:    "man __query__",
		description: "WIP",
		example:     "WIP",
		origin:      "built-in",
		Exec:        man,
	}

	CommandList["man"] = manDoc
}

// Man command
// returns a string containing the requested man page from the message array
func man(msgArray []string) (msgOut string) {
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
