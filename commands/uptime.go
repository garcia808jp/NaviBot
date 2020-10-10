// NaviBot: Discord bot for digital assistance
// uptime command

package commands

// Register the command for the CommandList
func init() {
	uptimeDoc := command{
		name:        "uptime - report the bot's uptime",
		synopsis:    "uptime",
		description: "WIP",
		example:     "WIP",
		origin:      "built-in",
		Exec:        uptime,
	}

	CommandList["uptime"] = uptimeDoc
}

// uptime command
// return the bot's uptime using an ugly solution via msgArray
func uptime(msgArray []string) (msgOut string) {
	// time.Since(startTime).String()
	return "command currently broken"
}
