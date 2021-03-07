// NaviBot: Lain resource Discord bot


package commands

import (
	// Standard packages
	"math/rand"
	"time"
)

// Register the command for the CommandList
func init() {
	memeCommand := Command{
		name:        "meme - WIP",
		synopsis:    "meme",
		description: "WIP",
		example:     "WIP",
		origin:      "built-in, lain",
		Exec:        meme,
	}

	CommandList["meme"] = memeCommand
}

// meme command
// returns a string containing a random URL in the memeSlice
func meme([]string) (msgOut string) {
	// Seed the rand package using current time in Unix format
	rand.Seed(time.Now().UnixNano())
	// Choose a rondom integer using the length of memeSlice
	randEntry := rand.Intn(len(memeSlice))
	// Prevent an error if the integer is out of bounds
	if randEntry == len(memeSlice) {
		randEntry = randEntry - 1
	}

	// Create a string using a random entry in memeSlice
	msgOut = memeSlice[randEntry]
	return
}

// meme slice
// contains a list of memes with Lain
var memeSlice = [...]string{
	"https://2static.fjcdn.com/pictures/And+you+dont+seem+to+understand+lain+is+for+hugs_f5d3f0_5980863.jpg",
	"https://cdn.discordapp.com/attachments/528841759879331872/535565438663524352/Linux_708676_4900166.jpg",
	"https://cdn.discordapp.com/attachments/528841759879331872/535569708502482954/image_search_1547760155063.jpg",
	"https://ci.memecdn.com/9579043.jpg",
	"https://i.chzbgr.com/full/2490778880/h50C02810/",
	"https://i.imgur.com/x8Ep0j5.jpg",
	"https://i.kym-cdn.com/photos/images/original/001/043/266/9bf.jpg",
	"https://i.kym-cdn.com/photos/images/original/001/104/212/326.jpg",
	"https://i.kym-cdn.com/photos/images/original/001/130/384/ff7.jpg",
	"https://i.pinimg.com/originals/54/ab/8a/54ab8adb4d5312534f54b976263245b0.jpg",
	"https://i.ytimg.com/vi/ysDc35yXqQk/maxresdefault.jpg",
	"https://img.fireden.net/a/image/1466/73/1466736357302.png",
	"https://img.fireden.net/a/image/1523/99/1523995070828.jpg",
	"https://img.memecdn.com/eat-up-and-become-a-good-nerd-serial-experiments-lain_o_6681023.jpg",
	"https://img.memecdn.com/hell-yeah-high-tech-cubic-monitors-serial-experiments-lain_o_6673309.jpg",
	"https://img.memecdn.com/inglorious-bastards-serial-experiments-lain_o_1676629.jpg",
	"https://img.memecdn.com/why-the-heck-would-you-wear-something-like-that-indoors-serial-experiments-lain_o_7178838.jpg",
	"https://pics.awwmemes.com/aprantozaho-serial-experiment-lain-yuri-on-ice-kill-la-kill-21843017.png",
	"https://pics.me.me/hey-mom-and-dad-i-dont-need-to-goto-i-30039967.png",
}
