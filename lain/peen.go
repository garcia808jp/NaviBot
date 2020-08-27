// NaviBot: Discord bot for digital assistance
// LainBot commands

package lain

import (
	// Standard packages
	"math/rand"
	"time"
)

// Register the command for the CommandList
func init() {
	peenDoc := Doc{
		Name:        "peen - bully a member",
		Synopsis:    "peen __user__",
		Description: "WIP",
		Example:     "WIP",
		Origin:      "built-in, lain",
	}

	CommandList["penis"] = peenDoc
}

// Peen command
// returns a string containing a random URL in peenSlice
func Peen() (msgOut string) {
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

// Peen slice
// insults to bully with
var peenSlice = [...]string{
	"hi <@98773013649977344>",
	"<@219936359236894731> gay",
}
