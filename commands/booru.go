// NaviBot: Discord bot for digital assistance
// le command

package commands

// Register the command for the CommandList
func init() {
	leDoc := command{
		name:        "le - search for images from *booru sites",
		synopsis:    "le __query__",
		description: "WIP",
		example:     "WIP",
		origin:      "built-in",
		Exec:        le,
	}

	CommandList["le"] = leDoc
}

// Man command
// returns a string containing the requested man page from the message array
func le(msgArray []string) (msgOut string) {
	return "command not implemented yet"
}
