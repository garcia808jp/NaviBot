// NaviBot: Discord bot for digital assistance
// event handling

// Defines this source file as part of the main package
package main

// Import depency packages
import (
	// Standard
	"strings"

	// Third-party
	"github.com/bwmarrin/discordgo"

	// In-house
	"github.com/phossil/NaviBot/commands"
)

// Respond to messages
func commandHandler(s *discordgo.Session, m *discordgo.MessageCreate) {
	// Ignore all messages created by the bot itself or are empty
	if m.Author.ID == s.State.User.ID || m.Content == "" {
		return
	}

	// Respond to messages containing the prefix
	if strings.HasPrefix(m.Content, prefix) == true {
		// Parse the message as a string array
		msgArray := strings.Fields(m.Content)

		// Make sure there is a second entry in the array
		// ie is there a command in the message
		if len(msgArray) == 1 {
			// s.ChannelMessageSend(m.ChannelID, greetingMsg)
			return
		} else if len(msgArray) >= 2 {
			_, err := commands.CommandList[msgArray[1]]
			if err == false {
				// s.ChannelMessageSend(m.ChannelID, "command not recognized")
				return
			}
			s.ChannelMessageSend(m.ChannelID, commands.CommandList[msgArray[1]].Exec(msgArray))
			return
		} else {
			s.ChannelMessageSend(m.ChannelID, "unknown error?: end of 'else if'")
		}
	}
}
