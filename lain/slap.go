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
	slapDoc := Doc{
		name:        "slap - WIP",
		synopsis:    "slap",
		description: "WIP",
		example:     "WIP",
		origin:      "built-in, lain",
		Exec:        slap,
	}

	CommandList["slap"] = slapDoc
}

// slap command
// returns a string containing a random URL in the slapSlice
func slap([]string) (msgOut string) {
	// Seed the rand package using current time in Unix format
	rand.Seed(time.Now().UnixNano())
	// Choose a rondom integer using the length of slapSlice
	randEntry := rand.Intn(len(slapSlice))
	// Prevent an error if the integer is out of bounds
	if randEntry == len(slapSlice) {
		randEntry = randEntry - 1
	}

	// Create a string using a random entry in slapSlice
	msgOut = slapSlice[randEntry]
	return
}

// slap slice
// contains a list of images with slaps
var slapSlice = [...]string{
	"https://i.imgur.com/4MQkDKm.gif",
	"https://i.imgur.com/Agwwaj6.gif",
	"https://i.imgur.com/CwbYjBX.gif",
	"https://i.imgur.com/VW0cOyL.gif",
	"https://i.imgur.com/YA7g7h7.gif",
	"https://i.imgur.com/fm49srQ.gif",
	"https://i.imgur.com/jNaAaxn.gif",
	"https://i.imgur.com/kSLODXO.gif",
	"https://i.imgur.com/mIg8erJ.gif",
	"https://i.imgur.com/o2SJYUS.gif",
	"https://i.imgur.com/oOCq3Bt.gif",
	"https://i.pinimg.com/originals/1c/8f/0f/1c8f0f43c75c11bf504b25340ddd912d.gif",
	"https://media.giphy.com/media/10Am8idu3qBYRy/giphy.gif",
	"https://media.giphy.com/media/VEmm8ngZxwJ9K/giphy.gif",
	"https://media.giphy.com/media/iUgoB9zOO0QkU/giphy.gif",
	"https://media.giphy.com/media/zRlGxKCCkatIQ/giphy.gif",
	"https://media1.tenor.com/images/477821d58203a6786abea01d8cf1030e/tenor.gif?itemid=7958720",
	"https://media1.tenor.com/images/fb17a25b86d80e55ceb5153f08e79385/tenor.gif?itemid=7919028",
	"https://pa1.narvii.com/6807/ac91cef2e5ae98f598665193f37bba223301d75c_hq.gif",
}
