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
	songCommand := Command{
		name:        "song - WIP",
		synopsis:    "song",
		description: "WIP",
		example:     "WIP",
		origin:      "built-in, lain",
		Exec:        song,
	}

	CommandList["song"] = songCommand
}

// song command
// returns a string containing a random URL in the songSlice
func song([]string) (msgOut string) {
	// Seed the rand package using current time in Unix format
	rand.Seed(time.Now().UnixNano())
	// Choose a rondom integer using the length of songSlice
	randEntry := rand.Intn(len(songSlice))
	// Prevent an error if the integer is out of bounds
	if randEntry == len(songSlice) {
		randEntry = randEntry - 1
	}

	// Create a string using a random entry in songSlice
	msgOut = songSlice[randEntry]
	return
}

// song slice
// contains a list of music
var songSlice = [...]string{
	"https://youtu.be/28mTETQMRs4",
	"https://youtu.be/4-PkAQcuZOw",
	"https://youtu.be/4DnCETKeHkc",
	"https://youtu.be/5dbi4N6NGn4",
	"https://youtu.be/H221MRRgFZs?list=PLLiA91TyEc8tL1jj8RZ0zCkG1QCDUd8YX",
	"https://youtu.be/H40u9ufvSVo",
	"https://youtu.be/IQicDkntMVA",
	"https://youtu.be/K00pcctFIuY",
	"https://youtu.be/MFcpcwdBYrk",
	"https://youtu.be/N6Jn98ktFw0",
	"https://youtu.be/RVIreGd1-NA",
	"https://youtu.be/SQ02E7qgZ_E",
	"https://youtu.be/YSecmwH_AoQ",
	"https://youtu.be/bEHUFRRK9Sk",
	"https://youtu.be/dvozrz9zWBg",
	"https://youtu.be/eYMQnb2igTs",
	"https://youtu.be/gZISvTviaj8",
	"https://youtu.be/o0ndkiL5ivU",
	"https://youtu.be/pTzumRQKGgM",
	"https://youtu.be/qHT6eEKP3jI",
	"https://youtu.be/rLiyFaLs8PY",
	"https://youtu.be/yO4myLCfN-Y",
}
