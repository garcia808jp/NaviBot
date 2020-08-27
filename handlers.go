// NaviBot: Discord bot for digital assistance
// event handling

// Defines this source file as part of the main package
package main

// Import depency packages
import (
	// Standard
	"strings"
	"time"

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
			s.ChannelMessageSend(m.ChannelID, greetingMsg)
			return
		}

		// Check the user input against the  built-in commands
		switch msgArray[1] {
		// If the message is "ping" reply with "pong"
		case "ping":
			s.ChannelMessageSend(m.ChannelID, "pong")
		// If the message is "pong" reply with "ping"
		case "pong":
			s.ChannelMessageSend(m.ChannelID, "ping")
		// If the message is "uptime" reply with the uptime in string format
		case "uptime":
			s.ChannelMessageSend(m.ChannelID, time.Since(startTime).String())
		// Query the requested man page from online
		case "man":
			s.ChannelMessageSend(m.ChannelID, commands.Man(msgArray))
		// Search for requested images from *booru
		case "le":
			s.ChannelMessageSend(m.ChannelID, "command not implemented yet")

		// Help command
		case "help":
			s.ChannelMessageSend(m.ChannelID, commands.Help(msgArray))
		// Link to the git repo
		case "code":
			s.ChannelMessageSend(m.ChannelID, codeURL)
		// Notify the user if the command is not recognised
		default:
			s.ChannelMessageSend(m.ChannelID, "command not recognized")
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
	// msgArray[0] should be the prefix
	// entries after that should be arguments
	msgArray := strings.Fields(m.Content)

	// Respond to messages containing the prefix
	if msgArray[0] == lainPrefix {
		// Make sure there is a second entry in the array
		// ie is there a command in the message
		if len(msgArray) == 1 {
			s.ChannelMessageSend(m.ChannelID, greetingMsg)
			return
		}

		// Check the user input against the  built-in commands
		switch msgArray[1] {
		// Return a wholesome image if the user wants a hug
		case "hug":
			s.ChannelMessageSend(m.ChannelID, lain.Hug())
		// Return a wholesome image if the user wants a pat
		case "pat":
			s.ChannelMessageSend(m.ChannelID, lain.Pat())
		// Return a random fortune
		case "8ball":
			s.ChannelMessageSend(m.ChannelID, lain.EightBall(msgArray))
		// Return a random wired site
		case "site":
			s.ChannelMessageSend(m.ChannelID, lain.Site())
		// Return a random gif
		case "gif":
			s.ChannelMessageSend(m.ChannelID, lain.Gif())
		// Return a random image
		case "image":
			s.ChannelMessageSend(m.ChannelID, lain.Image())
		// Help the user
		case "help":
			s.ChannelMessageSend(m.ChannelID, lain.Help(msgArray))
		// Link to the git repo
		case "code":
			s.ChannelMessageSend(m.ChannelID, lain.Code())
		// Ping a pour soul
		case "penis":
			s.ChannelMessageSend(m.ChannelID, lain.Peen())
		// Notify the user if the command is not recognised
		default:
			s.ChannelMessageSend(m.ChannelID, "command not recognized")
		}
	}
}
