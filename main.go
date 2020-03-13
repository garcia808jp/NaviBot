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
	"syscall"

	// Third-party packages
	"github.com/bwmarrin/discordgo"
	"github.com/joho/godotenv"
)

var (
	token  string
	prefix string
)

// Init function
func init() {
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
	// Create a Discord session
	dg, err := discordgo.New("Bot " + token)
	if err != nil {
		fmt.Println("error creating Discord session,", err)
		return
	}

	// Do something in messageCreate when a message is created
	dg.AddHandler(messageCreate)

	// Open the Discord session
	err = dg.Open()
	if err != nil {
		fmt.Println("error opening connection,", err)
		return
	}

	// Output the values from nenv
	fmt.Printf("\n\tTOKEN=%T\n\tPREFIX=\"%v\"\n", token, prefix)

	// Run until a signal to terminate is received
	fmt.Println("\nNavibot ready and running :3")
	sc := make(chan os.Signal, 1)
	signal.Notify(sc, syscall.SIGINT, syscall.SIGTERM, os.Interrupt, os.Kill)
	<-sc

	// Close the Discord session
	dg.Close()
}

// Respond to messages
func messageCreate(s *discordgo.Session, m *discordgo.MessageCreate) {
	// Respond to messages containing the prefix
	if strings.HasPrefix(m.Content, prefix) {
		// Ignore all messages created by the bot itself
		if m.Author.ID == s.State.User.ID {
			return
		}

		// If the message is "ping" reply with "pongo"
		if m.Content == prefix+"ping" {
			s.ChannelMessageSend(m.ChannelID, "pongo")
		}
		// If the message is "pong" reply with "pingo"
		if m.Content == prefix+"pong" {
			s.ChannelMessageSend(m.ChannelID, "pingo")
		}
	}
}
