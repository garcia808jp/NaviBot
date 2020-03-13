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
// load environment variables from the .nenv file
func init() {
	// Welcome message
	fmt.Println("NaviBot: Discord bot for digital assistance")

	// Load the .nenv file
	err := godotenv.Load(".nenv")
	if err != nil {
		fmt.Println("error loading .nenv file,")

	}

	// Get the Discord token and the bot prefix
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

	// Register the messageCreate func as a callback for MessageCreate events.
	dg.AddHandler(messageCreate)

	// Open the Discord session
	err = dg.Open()
	if err != nil {
		fmt.Println("error opening connection,", err)
		return
	}

	// Output the values from the variables above
	fmt.Printf("\n\tTOKEN=\"%v\"\n\tPREFIX=\"%v\"\n", token, prefix)

	// Run until a signal to terminal is received
	fmt.Println("\nNavibot ready and running :3")
	sc := make(chan os.Signal, 1)
	signal.Notify(sc, syscall.SIGINT, syscall.SIGTERM, os.Interrupt, os.Kill)
	<-sc

	// Close the Discord session
	dg.Close()
}

// This function will be called (due to AddHandler above) every time a new
// message is created on any channel that the autenticated bot has access to.
func messageCreate(s *discordgo.Session, m *discordgo.MessageCreate) {

	if strings.HasPrefix(m.Content, prefix) {
		// Ignore all messages created by the bot itself
		// This isn't required in this specific example but it's a good practice.
		if m.Author.ID == s.State.User.ID {
			return
		}
		// If the message is "ping" reply with "Pong!"
		if m.Content == "ping" {
			s.ChannelMessageSend(m.ChannelID, "Pong!")
		}

		// If the message is "pong" reply with "Ping!"
		if m.Content == "pong" {
			s.ChannelMessageSend(m.ChannelID, "Ping!")
		}
	}
}
