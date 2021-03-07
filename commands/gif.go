// NaviBot: Lain resource Discord bot


package commands

import (
	// Standard packages
	"math/rand"
	"time"
)

// Register the command for the CommandList
func init() {
	gifCommand := Command{
		name:        "gif - WIP",
		synopsis:    "gif",
		description: "WIP",
		example:     "WIP",
		origin:      "built-in, lain",
		Exec:        gif,
	}

	CommandList["gif"] = gifCommand
}

// gif command
// returns a string containing a random URL in gifSlice
func gif([]string) (msgOut string) {
	// Seed the rand package using current time in Unix format
	rand.Seed(time.Now().UnixNano())
	// Choose a rondom integer using the length of gifSlice
	randEntry := rand.Intn(len(gifSlice))
	// Prevent an error if the integer is out of bounds
	if randEntry == len(gifSlice) {
		randEntry = randEntry - 1
	}

	// Create a string using a random entry in gifSlice
	msgOut = gifSlice[randEntry]
	return
}

// gif slice
// contains a list of online GIFs
var gifSlice = [...]string{
	"http://25.media.tumblr.com/tumblr_m2hpn4XlZQ1r73plvo1_500.gif",
	"http://33.media.tumblr.com/931f551c7c5cd9a0e6a0a558775d81f6/tumblr_mytlj0llZV1rzn9vfo1_250.gif",
	"http://38.media.tumblr.com/931f551c7c5cd9a0e6a0a558775d81f6/tumblr_mytlj0llZV1rzn9vfo1_500.gif",
	"http://78.media.tumblr.com/4c9d850a8e5ac7ade46e286142f878f7/tumblr_nw2adbFCyf1uorwlqo1_400.gif",
	"http://cs.gettysburg.edu/~duncjo01/archive/patterns/lain/1st_day.gif",
	"http://i.4cdn.org/c/1547054999646.gif",
	"http://marrowproductions.com/Twisted/wiki/images/4/42/LainTwist.gif",
	"https://66.media.tumblr.com/6e9d6b5af133b1b3884491d7f1f82983/tumblr_oog4yzg5u01vlb6q0o1_540.gif",
	"https://66.media.tumblr.com/d4b2745356f425ddbd4e7de8975ecad7/tumblr_phon3a66b51w67s65o1_400.gif",
	"https://67.media.tumblr.com/5f8ca77f56da64b33145879cfcf73675/tumblr_nt2o8qQQS21tes8zmo1_500.gif",
	"https://78.media.tumblr.com/bd9272f9b0568794c7dad8ec52c01346/tumblr_n92s34HGM31qjlwa8o1_r1_500.gif",
	"https://afinde-production.s3.amazonaws.com/uploads/2eee0567-1b6d-4553-bebc-358003433f17.gif",
	"https://arisuchan.jp/cult/src/1504745451093.gif",
	"https://arisuchan.jp/cult/src/1504745451093.gif",
	"https://bakanoweeby.files.wordpress.com/2017/10/lain-2.gif?w=739",
	"https://cdn.4archive.org/img/6a6KxdZ.gif",
	"https://cdn.discordapp.com/attachments/528841759879331872/536465067508498432/we_be_walking.gif",
	"https://data.whicdn.com/images/243034617/original.gif",
	"https://desu-usergeneratedcontent.xyz/mu/image/1525/23/1525239710532.gif",
	"https://gifimage.net/wp-content/uploads/2017/08/serial-experiments-lain-gif-21.gif",
	"https://i.4pcdn.org/pol/1445669927592.gif",
	"https://i.4pcdn.org/s4s/1532546585395.gif",
	"https://i.4pcdn.org/tg/1516038669389.gif",
	"https://i.gifer.com/8LV3.gif",
	"https://i.gifer.com/C2pQ.gif",
	"https://i.gifer.com/VvKY.gif",
	"https://i.imgur.com/6UxYXDJ.gif",
	"https://i.imgur.com/V4vSSo2.gif",
	"https://i.imgur.com/eqkzFu0.gif",
	"https://i.kym-cdn.com/photos/images/original/001/130/238/951.gif",
	"https://i.makeagif.com/media/10-10-2015/a4Whsh.gif",
	"https://img.fireden.net/a/image/1457/15/1457158385482.gif",
	"https://img.fireden.net/a/image/1532/92/1532920863393.gif",
	"https://img.fireden.net/v/image/1450/34/1450340196726.gif",
	"https://img.fireden.net/v/image/1454/27/1454275799295.gif",
	"https://imgur.com/oSNisPS",
	"https://media.giphy.com/media/112CeAWuyhQX1C/giphy.gif",
	"https://media.giphy.com/media/5vAuoSpIK7uh2/source.gif",
	"https://media.giphy.com/media/cz1fqXcuMymgo/giphy.gif",
	"https://media.giphy.com/media/h7J6rblcKKPDi/giphy.gif",
	"https://media1.tenor.com/images/83eeeeeb64813c10236111e63f2080a8/tenor.gif?itemid=11578021",
	"https://media1.tenor.com/images/8d397680ad0f8fce375ac7c7bd8b4fad/tenor.gif?itemid=11598206",
	"https://orig00.deviantart.net/868b/f/2018/226/4/a/lain_gif_remaster_final_by_deznak-dck4l98.gif",
	"https://pa1.narvii.com/6318/77ba9762f081a004831cbdbc2fc2c11cc4a3a602_hq.gif",
	"https://steamusercontent-a.akamaihd.net/ugc/110732351534804953/6170F4AC4C6080A353792551E2D9F7E61719713D/",
	"https://steamusercontent-a.akamaihd.net/ugc/859477345548989008/23073CCB3B1FE199008FB2A2F1CC7A4328A3F958/",
	"https://steamusercontent-a.akamaihd.net/ugc/882000754018213605/C5280F51D7FEF548A21B7FDE639F8E949E9052FF/",
	"https://thumbs.gfycat.com/KindheartedIdleHoneycreeper-max-1mb.gif",
}
