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
	"github.com/phossil/NaviBot/lain"
)

// Respond to messages
func mainHandler(s *discordgo.Session, m *discordgo.MessageCreate) {
	// Ignore all messages created by the bot itself or are empty
	if m.Author.ID == s.State.User.ID || m.Content == "" {
		return
	}

	// Parse the message as a string array
	// msgArray[0] should be the prefix
	// entries after that should be arguments
	msgArray := strings.Fields(m.Content)

	// Respond to messages containing the prefix
	if msgArray[0] == prefix {
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

// Respond to messages
func lainHandler(s *discordgo.Session, m *discordgo.MessageCreate) {
	// Ignore all messages created by the bot itself or are empty
	if m.Author.ID == s.State.User.ID || m.Content == "" {
		return
	}

	// Parse the message as a string array
	msgArray := strings.Fields(m.Content)

	// Respond to messages containing the prefix
	if msgArray[0] == lainPrefix {
		// Make sure there is a second entry in the array
		// ie is there a command in the message
		if len(msgArray) == 1 {
			// s.ChannelMessageSend(m.ChannelID, greetingMsg)
			return
		} else if len(msgArray) >= 2 {
			_, err := lain.CommandList[msgArray[1]]
			if err == false {
				// s.ChannelMessageSend(m.ChannelID, "command not recognized")
				return
			}
			s.ChannelMessageSend(m.ChannelID, lain.CommandList[msgArray[1]].Exec(msgArray))
			return
		} else {
			s.ChannelMessageSend(m.ChannelID, "unknown error?: end of 'else if'")
		}
	}
}
