// NaviBot: Lain resource Discord bot


package commands

import "time"

// Register the command for the CommandList
func init() {
	uptimeDoc := Command{
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
	msgOut = time.Since(startTime).String()
	return
}
