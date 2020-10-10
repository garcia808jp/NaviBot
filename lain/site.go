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
	siteDoc := Doc{
		name:        "site - WIP",
		synopsis:    "site",
		description: "WIP",
		example:     "WIP",
		origin:      "built-in, lain",
		Exec:        site,
	}

	CommandList["site"] = siteDoc
}

// site command
// returns a string containing a random URL in siteSlice
func site([]string) (msgOut string) {
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

// site slice
// contains a list of Lain related sites
var siteSlice = [...]string{
	"http://lain.angelic-trust.net/wired.html",
	"http://navi.solutions/",
	"http://sel.wikia.com/wiki/Serial_Experiments_Lain_Wiki",
	"http://www.cjas.org/~leng/lain.htm",
	"https://arisuchan.jp/",
	"https://arisuchan.jp/cyb/res/1210.html",
	"https://asphyxia.su/",
	"https://blackwings.neocities.org/",
	"https://fauux.neocities.org/",
	"https://mebious.neocities.org/Layer/Wierd.html",
	"https://systemspace.link/warning.php",
	"https://systemspace.network/",
}
