// NaviBot: Discord bot for digital assistance
// LainBot commands

package lain

import (
	// Standard packages
	"math/rand"
	"time"
)

// sorri for the ugly code ._.
// it's here until i can makde it better

// Hug command
// returns a string containing a random URL in hugSlice
func Hug() (msgOut string) {
	// Seed the rand package using current time in Unix format
	rand.Seed(time.Now().UnixNano())
	// Choose a rondom integer using the length of the Hug slice
	randEntry := rand.Intn(len(hugSlice))
	// Prevent an error if the integer is out of bounds
	if randEntry == len(hugSlice) {
		randEntry = randEntry - 1
	}

	// Create a string using a random entry in hugSlice
	msgOut = "much yay, very hug !!!\n" + hugSlice[randEntry]
	return
}

// Pat command
// returns a string containing a random URL in patSlice
func Pat() (msgOut string) {
	// Seed the rand package using current time in Unix format
	rand.Seed(time.Now().UnixNano())
	// Choose a rondom integer using the length of patSlice
	randEntry := rand.Intn(len(patSlice))
	// Prevent an error if the integer is out of bounds
	if randEntry == len(patSlice) {
		randEntry = randEntry - 1
	}

	// Create a string using a random entry in patSlice
	msgOut = "much yay, very uwu !!!\n" + patSlice[randEntry]
	return
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

// Site command
// returns a string containing a random URL in siteSlice
func Site() (msgOut string) {
	// Seed the rand package using current time in Unix format
	rand.Seed(time.Now().UnixNano())
	// Choose a rondom integer using the length of siteSlice
	randEntry := rand.Intn(len(siteSlice))
	// Prevent an error if the integer is out of bounds
	if randEntry == len(siteSlice) {
		randEntry = randEntry - 1
	}

	// Create a string using a random entry in siteSlice
	msgOut = "become one with the wired, faithful follower\n" + siteSlice[randEntry]
	return
}

// Gif command
// returns a string containing a random URL in gifSlice
func Gif() (msgOut string) {
	// Seed the rand package using current time in Unix format
	rand.Seed(time.Now().UnixNano())
	// Choose a rondom integer using the length of gifSlice
	randEntry := rand.Intn(len(gifSlice))
	// Prevent an error if the integer is out of bounds
	if randEntry == len(gifSlice) {
		randEntry = randEntry - 1
	}

	// Create a string using a random entry in gifSlice
	msgOut = "mlem\n" + gifSlice[randEntry]
	return
}

// Image command
// returns a string containing a random URL in the imageSlice
func Image() (msgOut string) {
	// Seed the rand package using current time in Unix format
	rand.Seed(time.Now().UnixNano())
	// Choose a rondom integer using the length of imageSlice
	randEntry := rand.Intn(len(imageSlice))
	// Prevent an error if the integer is out of bounds
	if randEntry == len(imageSlice) {
		randEntry = randEntry - 1
	}

	// Create a string using a random entry in imageSlice
	msgOut = "owo\n" + imageSlice[randEntry]
	return
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
