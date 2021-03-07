// NaviBot: Lain resource Discord bot


package commands

import (
	// Standard packages
	"math/rand"
	"time"
)

// Register the command for the CommandList
func init() {
	hugCommand := Command{
		name:        "hug - WIP",
		synopsis:    "hug",
		description: "WIP",
		example:     "WIP",
		origin:      "built-in, lain",
		Exec:        hug,
	}

	CommandList["hug"] = hugCommand
}

// hug command
// returns a string containing a random URL in hugSlice
func hug([]string) (msgOut string) {
	// Seed the rand package using current time in Unix format
	rand.Seed(time.Now().UnixNano())
	// Choose a rondom integer using the length of the Hug slice
	randEntry := rand.Intn(len(hugSlice))
	// Prevent an error if the integer is out of bounds
	if randEntry == len(hugSlice) {
		randEntry = randEntry - 1
	}

	// Create a string using a random entry in hugSlice
	msgOut = hugSlice[randEntry]
	return
}

// Hug slice
// contains a list of images with hugs
var hugSlice = [...]string{
	"http://i.imgur.com/I8LyQ9L.gif",
	"http://i.imgur.com/tuH4gqZ.gif",
	"https://37.media.tumblr.com/f2a878657add13aa09a5e089378ec43d/tumblr_n5uovjOi931tp7433o1_500.gif",
	"https://66.media.tumblr.com/18fdf4adcb5ad89f5469a91e860f80ba/tumblr_oltayyHynP1sy5k7wo1_400.gif",
	"https://66.media.tumblr.com/291c8b98b219283f9e21927e7ef6c3f2/tumblr_mzscklfLYx1tptsy9o1_400.gif",
	"https://i.imgur.com/DBV5Ac9.gif",
	"https://i.imgur.com/Ed4DlVM.gif",
	"https://i.imgur.com/IESFOWD.gif",
	"https://i.imgur.com/Ltmb8aa.gif",
	"https://i.imgur.com/UMm95sV.gif",
	"https://i.imgur.com/c3WzMZu.gif",
	"https://i.imgur.com/fhCVWgN.gif",
	"https://i.imgur.com/ntqYLGl.gif",
	"https://i.imgur.com/r9aU2xv.gif",
	"https://i.imgur.com/x4koMxC.gif",
	"https://i.pinimg.com/originals/f2/80/5f/f2805f274471676c96aff2bc9fbedd70.gif",
	"https://media.tenor.com/images/2e1d34d002d73459b6119d57e6a795d6/tenor.gif",
	"https://media0.giphy.com/media/kvKFM3UWg2P04/giphy.gif?cid=3640f6095c8719d76e614e7932321eae",
	"https://media1.tenor.com/images/1069921ddcf38ff722125c8f65401c28/tenor.gif?itemid=11074788",
	"https://media1.tenor.com/images/40aed63f5bc795ed7a980d0ad5c387f2/tenor.gif?itemid=11098589",
	"https://media1.tenor.com/images/4d89d7f963b41a416ec8a55230dab31b/tenor.gif?itemid=5166500",
	"https://media1.tenor.com/images/5845f40e535e00e753c7931dd77e4896/tenor.gif?itemid=9920978",
	"https://media1.tenor.com/images/6088a5ea989ddf19836f655b8555710a/tenor.gif?itemid=12159970",
	"https://media1.tenor.com/images/62048cf3073b2670e176c470aa1d2714/tenor.gif?itemid=12668675",
	"https://media1.tenor.com/images/75a607663feb91c59bb4a84ab803fba6/tenor.gif?itemid=11144378",
	"https://media1.tenor.com/images/7db5f172665f5a64c1a5ebe0fd4cfec8/tenor.gif?itemid=9200935",
	"https://media1.tenor.com/images/fdefc2134e17de3bb15bc398ff66c6ca/tenor.gif?itemid=9469917",
	"https://thumbs.gfycat.com/BlindOblongAmurratsnake-small.gif",
	"https://uploads.disquscdn.com/images/f5ee0928f31e1d867c85a1965005c54b3446595d2ff31989dc06a213679cc272.gif",
}
