// +build lainbot
// NaviBot: Lain resource Discord bot

package commands

import (
	// Standard packages
	"math/rand"
	"time"
)

// Register the command for the CommandList
func init() {
	peenCommand := Command{
		name:        "peen - bully a member",
		synopsis:    "peen __user__",
		description: "WIP",
		example:     "WIP",
		origin:      "built-in, lain",
		Exec:        peen,
	}

	CommandList["penis"] = peenCommand
}

// peen command
// returns a string containing a random URL in peenSlice
func peen([]string) (msgOut string) {
	// Seed the rand package using current time in Unix format
	rand.Seed(time.Now().UnixNano())
	// Choose a rondom integer using the length of peenSlice
	randEntry := rand.Intn(len(peenSlice))
	// Prevent an error if the integer is out of bounds
	if randEntry == len(peenSlice) {
		randEntry = randEntry - 1
	}

	// Create a string using a random entry in peenSlice
	msgOut = peenSlice[randEntry]
	return
}

// peen slice
// insults to bully with
var peenSlice = [...]string{
	"hi <@98773013649977344>",
	"<@219936359236894731> gay",
}
