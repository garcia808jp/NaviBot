// NaviBot: Lain resource Discord bot


package commands

import (
	// Standard packages
	"math/rand"
	"time"
)

// Register the command for the CommandList
func init() {
	patCommand := Command{
		name:        "pat - WIP",
		synopsis:    "pat",
		description: "WIP",
		example:     "WIP",
		origin:      "built-in, lain",
		Exec:        pat,
	}

	CommandList["pat"] = patCommand
}

// pat command
// returns a string containing a random URL in patSlice
func pat([]string) (msgOut string) {
	// Seed the rand package using current time in Unix format
	rand.Seed(time.Now().UnixNano())
	// Choose a rondom integer using the length of patSlice
	randEntry := rand.Intn(len(patSlice))
	// Prevent an error if the integer is out of bounds
	if randEntry == len(patSlice) {
		randEntry = randEntry - 1
	}

	// Create a string using a random entry in patSlice
	msgOut = patSlice[randEntry]
	return
}

// pat slice
// contains a list of images that have pats
var patSlice = [...]string{
	"https://encrypted-tbn0.gstatic.com/images?q=tbn:ANd9GcSR_B2M0Yc9qyqXjKrqLRAJngXEpM1XUxb9lYeOr9eJdndts0tB",
	"https://i.gifer.com/7MPC.gif",
	"https://i.imgur.com/0znUWqT.gif",
	"https://i.imgur.com/2k0MFIr.gif",
	"https://i.imgur.com/2lacG7l.gif",
	"https://i.imgur.com/4ssddEQ.gif",
	"https://i.imgur.com/F3cjr3n.gif",
	"https://i.imgur.com/LUypjw3.gif",
	"https://i.imgur.com/TPqMPka.gif",
	"https://i.imgur.com/UWbKpx8.gif",
	"https://i.imgur.com/fp9XJZO.gif",
	"https://i.imgur.com/sLwoifL.gif",
	"https://media.tenor.com/images/098689061fc2b850aa29fd4410fa97e7/tenor.gif",
	"https://media.tenor.com/images/e549c61c9bc3d8defdb0559b358b92a7/tenor.gif",
	"https://media0.giphy.com/media/BxdqaKAfdcCsg/giphy.gif",
	"https://thumbs.gfycat.com/ImpurePleasantArthropods-small.gif",
}
