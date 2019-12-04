// NaviBot: Discord bot for digital assistance

// Defines this source file as the main one
package main

// Import depency packages
import (
        // Standard packages
        "fmt"
        "os"
        "log"

        // Third-party packages
        "github.com/joho/godotenv"
        "github.com/bwmarrin/discordgo"
)

// Init function
func init() {
        // Welcome message when the Go package loads
        fmt.Println("NaviBot: Discord bot for digital assistance")

        // Load the .nenv file
        err := godotenv.Load(".nenv")
        if err != nil {
            log.Fatal("Error loading .nenv file")
        }
}

// Main function
func main() {
        // Get the Discord token and the bot prefix
        var (
            token string = os.Getenv("TOKEN")
            prefix string = os.Getenv("PREFIX")
        )

        // Create a Discord session
        dg, err := discordgo.New("Bot " + Token)
            if err != nil {
            fmt.Println("Error creating Discord session,", err)
            return
        }
        
        // Output the values from the variables above
        fmt.Printf("TOKEN=\"%v\"\nPREFIX=\"%v\"\n", token, prefix)
        
        // Close the Discord session
        dg.Close()
}
