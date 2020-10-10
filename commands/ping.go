// NaviBot: Discord bot for digital assistance
// ping command

package commands

// Register the command for the CommandList
func init() {
	pingDoc := command{
		name:        "ping - classic bot command",
		synopsis:    "ping **|** pong",
		description: "Returns *pong* if the command was *ping*, *ping* if *pong*, otherwise the bot questions its existence",
		example:     "ping",
		origin:      "built-in",
		Exec:        ping,
	}

	CommandList["ping"] = pingDoc
	CommandList["pong"] = pingDoc
}

// ping command
// return ping or pong
func ping(msgArray []string) (msgOut string) {
	switch msgArray[1] {
	case "ping":
		msgOut = "pong"
	case "pong":
		msgOut = "ping"
	default:
		msgOut = "is the ping command nested elsewhere?"
	}
	return
}
