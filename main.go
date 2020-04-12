// NaviBot: Discord bot for digital assistance

// Defines this source file as the main one
package main

// Import depency packages
import (
	// Standard packages
	"fmt"
	"math/rand"
	"os"
	"os/signal"
	"strings"
	"time"

	// Third-party packages
	"github.com/bwmarrin/discordgo"
	"github.com/joho/godotenv"

	// In-house packages
	"github.com/phossil/NaviBot/lain"
	//"github.com/phossil/NaviBot/commands"
)

// Initialize some global variables
var (
	token       string
	prefix      string
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
}

// Main function
func main() {
	// Output the data type of 'token' and the value of 'prefix'
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
	// Close the Discord session on exit
	defer dg.Close()

	// Do something in messageCreate when a message is created
	dg.AddHandler(messageCreate)

	// Run until a signal to terminate is received
	// this portion is a mystery to me :c
	fmt.Println("\nNaviBot running and ready :3")
	sc := make(chan os.Signal, 1)
	signal.Notify(sc, os.Interrupt, os.Kill)
	<-sc
}

// Respond to messages
func messageCreate(s *discordgo.Session, m *discordgo.MessageCreate) {
	// Ignore all messages created by the bot itself
	if m.Author.ID == s.State.User.ID {
		return
	}

	// Parse the message as a string array
	// the first entry in the array, msgArray[0], should be the prefix
	// the second should be  the command and the rest should be arguments
	msgArray := strings.Fields(m.Content)

	// Respond to messages containing the prefix
	if strings.HasPrefix(m.Content, prefix) {
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
			s.ChannelMessageSend(m.ChannelID, manComm(msgArray))
		// Search for requested images from *booru
		case "le":
			s.ChannelMessageSend(m.ChannelID, "command not implemented yet")

		// Return a wholesome image if the user wants a hug
		case "lainh":
			s.ChannelMessageSend(m.ChannelID, lainhComm())
		// Return a wholesome image if the user wants a pat
		case "lainp":
			s.ChannelMessageSend(m.ChannelID, lainpComm())
			// Return a random fortune
		case "lain8":
			s.ChannelMessageSend(m.ChannelID, lain8Comm(msgArray))
		// Notify the user if the command is not recognised
		default:
			s.ChannelMessageSend(m.ChannelID, "command not recognized")
		}
	}
}

// Man command
// the command returns a string containing the requested man page from the message array
func manComm(msgArray []string) (msgOut string) {
	// If the message contains arguments, complete the task; notify the user otherwise
	if len(msgArray) >= 3 {
		// The default URL for the online man pages
		manSrc := "https://jlk.fjfi.cvut.cz/arch/manpages/search?q="
		// Concatenate the URL with the requested man page
		msgOut = manSrc + msgArray[2]
	} else {
		msgOut = "command not implemented yet"
	}
	return msgOut
}

// LainBot commands
// sorri for the ugly code ._.
// it's here until i can makde it better

// Lain hug command
// the command returns a string containing a random URL in the lain.Hug slice
func lainhComm() (msgOut string) {
	// Seed the rand package using current time in Unix format
	rand.Seed(time.Now().UnixNano())
	// Choose a rondom integer using the length of the lain.Hug slice
	randEntry := rand.Intn(len(lain.Hug))
	// Prevent an error if the integer is out of bounds
	if randEntry == len(lain.Hug) {
		randEntry = randEntry - 1
	}

	// Create a string using a random entry in the lain.Hug slice
	msgOut = "much yay, very hug !!!\n" + lain.Hug[randEntry]
	return msgOut
}

// Lain pat command
// the command returns a string containing a random URL in the lain.Pat slice
func lainpComm() (msgOut string) {
	// Seed the rand package using current time in Unix format
	rand.Seed(time.Now().UnixNano())
	// Choose a rondom integer using the length of the lain.Pat slice
	randEntry := rand.Intn(len(lain.Pat))
	// Prevent an error if the integer is out of bounds
	if randEntry == len(lain.Pat) {
		randEntry = randEntry - 1
	}

	// Create a string using a random entry in the lain.Pat slice
	msgOut = "much yay, very uwu !!!\n" + lain.Pat[randEntry]
	return msgOut
}

// Lain 8-ball command
// the command returns a string from the lain.EightBall slice
func lain8Comm(msgArray []string) (msgOut string) {
	// If there are no arguments notify the user
	if len(msgArray) == 2 {
		return "can't tell ur fortune without a queston -_-"
	}
	// Seed the rand package using current time in Unix format
	rand.Seed(time.Now().UnixNano())
	// Choose a rondom integer using the length of the lain.EightBall slice
	randEntry := rand.Intn(len(lain.EightBall))
	// Prevent an error if the integer is out of bounds
	if randEntry == len(lain.EightBall) {
		randEntry = randEntry - 1
	}

	// Create a string using a random entry in the lain.EightBall slice
	msgOut = ":8ball: Lain has a message from beyond~~\n" + lain.EightBall[randEntry]
	return msgOut
}
