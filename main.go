// NaviBot: Discord bot for digital assistance

// Defines this source file as part of the main package
package main

// Import depency packages
import (
	// Standard
	"fmt"
	"log"
	"os"
	"os/signal"
	"time"

	// Third-party
	"github.com/bwmarrin/discordgo"
	"github.com/joho/godotenv"
)

// Declare some global variables
var (
	token      string
	prefix     string
	lainPrefix string
	startTime  time.Time
)

// Initialize some global constants
const (
	codeURL     string = "https://github.com/phossil/NaviBot/"
	greetingMsg string = "NaviBot: Discord bot for digital assistance"
	version     string = "0.0.1-alpha-20201006"
)

// Init function
func init() {
	// Grab the time on startup
	startTime = time.Now()

	// Welcome message
	fmt.Println(greetingMsg)

	// Load the .nenv file
	err := godotenv.Load("nenv")
	if err != nil {
		log.Print(err)
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
		log.Print(err)
		return
	}

	// Open the Discord session
	err = dg.Open()
	if err != nil {
		log.Print(err)
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
