// NaviBot: Discord bot for digital assistance

// Defines this source file as the main one
package main

// Import depency packages
import (
	// Standard packages
	"fmt"
	"log"
	"os"

	// Third-party packages
	"github.com/bwmarrin/discordgo"
	"github.com/joho/godotenv"
)

// Init function
func init() {
	// Welcome message
	fmt.Println("NaviBot: Discord bot for digital assistance")

	// Load the .nenv file
	err := godotenv.Load(".nenv")
	if err != nil {
		log.Fatal("error loading .nenv file")
	}
}

// Main function
func main() {
	// Get the Discord token and the bot prefix
	var (
		token  string = os.Getenv("TOKEN")
		prefix string = os.Getenv("PREFIX")
	)

	// Create a Discord session
	dg, err := discordgo.New("Bot " + token)
	if err != nil {
		fmt.Println("error creating Discord session x-x,", err)
		return
	}

	// Open the Discord session
	err = dg.Open()
	if err != nil {
		fmt.Println("error opening connection,", err)
		return
	}

	// Output the values from the variables above
	fmt.Printf("\tTOKEN=\"%v\"\n\tPREFIX=\"%v\"\n", token, prefix)

	// Close the Discord session
	dg.Close()
}
