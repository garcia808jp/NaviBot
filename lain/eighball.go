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
	ebDoc := Doc{
		Name:        "8ball - divination with the goddess herself",
		Synopsis:    "8ball __question__",
		Description: "WIP",
		Example:     "WIP",
		Origin:      "built-in, lain",
	}

	CommandList["8ball"] = ebDoc
}

// EightBall command
// returns a string from eightBallSlice
func EightBall(msgArray []string) (msgOut string) {
	// If there are no arguments notify the user
	if len(msgArray) == 2 {
		return "can't tell ur fortune without a queston -_-"
	}
	// Seed the rand package using current time in Unix format
	rand.Seed(time.Now().UnixNano())
	// Choose a rondom integer using the length of eightBallSlice
	randEntry := rand.Intn(len(eightBallSlice))
	// Prevent an error if the integer is out of bounds
	if randEntry == len(eightBallSlice) {
		randEntry = randEntry - 1
	}

	// Create a string using a random entry in the eightBallSlice
	msgOut = ":8ball: Lain has a message from beyond~~\n" + eightBallSlice[randEntry]
	return
}

// EightBall slice
// a slice of strings that tell your fortune
var eightBallSlice = [...]string{
	"It is certain.",
	"As I see it yes.",
	"Ask again later.",
	"Better not tell you now.",
	"Cannot predict now.",
	"Concentrate and ask again.",
	"Don't count on it.",
	"It is decidedly so.",
	"Most likely.",
	"My reply is no.",
	"My sources say no.",
	"Outlook good.",
	"Outlook not so good.",
	"Reply hazy try again.",
	"Signs point to yes.",
	"Very doubtful.",
	"Without a doubt.",
	"Yes - definitely.",
	"Yes.",
	"You may rely on it.",
}
