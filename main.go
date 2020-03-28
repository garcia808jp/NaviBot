// NaviBot: Discord bot for digital assistance

// Defines this source file as the main one
package main

// Import depency packages
import (
	// Standard packages
	"fmt"
	"os"
	"os/signal"
	"strings"
	"time"

	// Third-party packages
	"github.com/bwmarrin/discordgo"
	"github.com/joho/godotenv"
)

var (
	token     string
	prefix    string
	startTime time.Time
)

// Init function
func init() {
	// Grab the time on startup
	startTime = time.Now()

	// Welcome message
	fmt.Println("NaviBot: Discord bot for digital assistance")

	// Load the .nenv file
	err := godotenv.Load(".nenv")
	if err != nil {
		fmt.Println("error loading .nenv file,")
		return
	}

	// Load environment variables from the .nenv file
	token = os.Getenv("TOKEN")
	prefix = os.Getenv("PREFIX")
}

// Main function
func main() {
	// Output the values from nenv
	fmt.Printf("\n\tTOKEN=%T\n\tPREFIX=\"%v\"\n", token, prefix)

	// Output the startTime to console
	fmt.Printf("\n\tstart time: %v\n", startTime)

	// Create a Discord session
	dg, err := discordgo.New("Bot " + token)
	if err != nil {
		fmt.Println("error creating Discord session,", err)
		return
	}

	// Open the Discord session
	err = dg.Open()
	if err != nil {
		fmt.Println("error opening connection,", err)
		return
	}

	// Do something in messageCreate when a message is created
	dg.AddHandler(messageCreate)

	// Run until a signal to terminate is received
	fmt.Println("\nNaviBot running and ready :3")
	sc := make(chan os.Signal, 1)
	signal.Notify(sc, os.Interrupt, os.Kill)
	<-sc

	// Close the Discord session
	dg.Close()
}

// Respond to messages
func messageCreate(s *discordgo.Session, m *discordgo.MessageCreate) {
	// Ignore all messages created by the bot itself
	if m.Author.ID == s.State.User.ID {
		return
	}

	// Parse the message as an array
	msgArray := strings.Fields(m.Content)

	// Respond to messages containing the prefix
	if strings.HasPrefix(m.Content, prefix) {
		// Check the  built-in commands against
		switch msgArray[1] {
		// If the message is "ping" reply with "pong"
		case "ping":
			s.ChannelMessageSend(m.ChannelID, "pong")
		// If the message is "pong" reply with "ping"
		case "pong":
			s.ChannelMessageSend(m.ChannelID, "ping")
		// If the message is "uptime" reply with the uptime in string format
		case "uptime":
			// Calculate the time since startup to the command received
			uptime := time.Since(startTime)
			// Reply with a string
			s.ChannelMessageSend(m.ChannelID, uptime.String())
		// Query the requested man page from online
		case "man":
			s.ChannelMessageSend(m.ChannelID, "command not implemented yet")
		// Notify the user if the command is not recognised
		default:
			s.ChannelMessageSend(m.ChannelID, "command not recognized")
		}
	}
}
