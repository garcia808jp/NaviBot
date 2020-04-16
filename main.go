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

	// In-house packages
	"github.com/phossil/NaviBot/commands"
	"github.com/phossil/NaviBot/lain"
)

// Initialize some global variables
var (
	token       string
	prefix      string
	lainPrefix  string
	startTime   time.Time
	greetingMsg string = "NaviBot: Discord bot for digital assistance"
)

// Init function
func init() {
	// Grab the time on startup
	startTime = time.Now()

	// Welcome message
	fmt.Println(greetingMsg)

	// Load the .nenv file
	err := godotenv.Load(".nenv")
	if err != nil {
		fmt.Println("error loading .nenv file,")
		return
	}

	// Load environment variables into the respective global variables
	token = os.Getenv("TOKEN")
	prefix = os.Getenv("PREFIX")
	lainPrefix = os.Getenv("LAIN_PREFIX")
}

// Main function
func main() {
	// Output the data type of 'token' and the value of 'prefix'
	fmt.Printf("\n\tTOKEN=%T\n\tPREFIX=\"%v\"\n\tLAIN_PREFIX=\"%v\"\n", token, prefix, lainPrefix)

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
	// Close the Discord session on exit
	defer dg.Close()

	// Handle messages with mainHandler
	dg.AddHandler(mainHandler)
	// Handle messages with lainHandler
	dg.AddHandler(lainHandler)

	// Run until a signal to terminate is received
	// this portion is a mystery to me :c
	fmt.Println("\nNaviBot running and ready :3")
	sc := make(chan os.Signal, 1)
	signal.Notify(sc, os.Interrupt, os.Kill)
	<-sc
}

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
			uptime := time.Since(startTime)
			s.ChannelMessageSend(m.ChannelID, uptime.String())
		// Query the requested man page from online
		case "man":
			s.ChannelMessageSend(m.ChannelID, commands.Man(msgArray))
		// Search for requested images from *booru
		case "le":
			s.ChannelMessageSend(m.ChannelID, "command not implemented yet")

		// Help command
		case "help":
			s.ChannelMessageSend(m.ChannelID, "command not implemented yet")
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
			s.ChannelMessageSend(m.ChannelID, "command not implemented yet")
		// Notify the user if the command is not recognised
		default:
			s.ChannelMessageSend(m.ChannelID, "command not recognized")
		}
	}
}
