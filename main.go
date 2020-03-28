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

    // Respond to messages containing the prefix
	if strings.HasPrefix(m.Content, prefix) {
		// Parse the message as an array
		msgArray := strings.Fields(m.Content)

		// If the message is "ping" reply with "pongo"
		if msgArray[1] == "ping" {
			s.ChannelMessageSend(m.ChannelID, "pongo")
		}
		// If the message is "pong" reply with "pingo"
		if msgArray[1] == "pong" {
			s.ChannelMessageSend(m.ChannelID, "pingo")
		}
		// If the message is "uptime" reply with the uptime in string format
		if msgArray[1] == "uptime" {
			uptime := time.Since(startTime)
			s.ChannelMessageSend(m.ChannelID, uptime.String())
		}
		// If the message is "pong" reply with "pingo"
		if msgArray[1] == "man" {
			s.ChannelMessageSend(m.ChannelID, "man is not implemented yet")
		}
	}
}
