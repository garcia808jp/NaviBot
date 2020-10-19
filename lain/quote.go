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
	quoteDoc := Doc{
		name:        "quote - WIP",
		synopsis:    "quote",
		description: "WIP",
		example:     "WIP",
		origin:      "built-in, lain",
		Exec:        quote,
	}

	CommandList["quote"] = quoteDoc
	CommandList["quote"] = quoteDoc
}

// quote command
// returns a string containing a random URL in the quoteSlice
func quote([]string) (msgOut string) {
	// Seed the rand package using current time in Unix format
	rand.Seed(time.Now().UnixNano())
	// Choose a rondom integer using the length of quoteSlice
	randEntry := rand.Intn(len(quoteSlice))
	// Prevent an error if the integer is out of bounds
	if randEntry == len(quoteSlice) {
		randEntry = randEntry - 1
	}

	// Create a string using a random entry in quoteSlice
	msgOut = quoteSlice[randEntry]
	return
}

// quote slice
// contains a list of quotes
var quoteSlice = [...]string{
	"http://housequote.com/menagerie/photos/birds/cowboygosling.jpg",
	"http://i.imgur.com/Nyv6k1S.jpg",
	"http://media.advance.net/newyorkupstatecom_national_desk_blog/photo/2018/09/12/captain-jason-big-jay-barnes-of-frontenac-fowlers-guide-service-sets-up-for-a-canada-quote-hunt-bfe8525bbaf0895a.jpg",
	"http://media.pixcove.com/B/5/5/quote-Gosling-quote-White-Free-Image-Bird-Swimming-9381.jpg",
	"https://api.culture.pl/sites/default/files/styles/1920_auto/public/images/imported/KUCHNIA/GESINA/full_gesina_forum_770.jpg?itok=o5V8Ncz7",
	"https://d1ia71hq4oe7pn.cloudfront.net/photo/59938311-480px.jpg",
	"https://d1ia71hq4oe7pn.cloudfront.net/photo/59940721-480px.jpg",
	"https://d1ia71hq4oe7pn.cloudfront.net/photo/59950681-720px.jpg",
	"https://d1ia71hq4oe7pn.cloudfront.net/photo/59950701-480px.jpg",
	"https://download.ams.birds.cornell.edu/api/v1/asset/59953191/1800",
	"https://i.dailymail.co.uk/i/pix/2010/05/22/article-1280370-09B0E58D000005DC-919_964x577.jpg",
	"https://i2.wp.com/www.wildlifehotline.com/wp-content/uploads/2011/10/lilducks.jpg",
	"https://image.iol.co.za/image/1/process/620x349?source=https://inm-baobab-prod-eu-west-1.s3.amazonaws.com/public/inm/media/file/1/1734662/1407919967/image/525269601.jpg",
	"https://insider.si.edu/wp-content/uploads/2018/05/Canada_quote_in_pond.jpg",
	"https://kywildliferemovalpros.com/wp-content/uploads/2018/03/quote.jpg",
	"https://media-cdn.tripadvisor.com/media/photo-s/0b/11/97/6d/canada-quote-and-gosling.jpg",
	"https://media.mnn.com/assets/images/2018/04/canada_quote_chicks.jpg.838x0_q80.jpg",
	"https://nature.mdc.mo.gov/sites/default/files/styles/centered_full/public/media/images/2013/10/canada-quote_swimming_10-30-13.jpg?itok=1uPvCy4_",
	"https://png.pngtree.com/element_origin_min_pic/16/10/27/a5dfdc24f753f5500a4dfe432f4b24e5.jpg",
	"https://s3-us-west-2.amazonaws.com/hobbyfarms.com/wp-content/uploads/2017/12/29195949/quote-ducks-resolutions.jpg",
	"https://st.depositphotos.com/2752123/4495/i/950/depositphotos_44954913-stock-photo-wild-quote-with-puppy-quotes.jpg",
	"https://thenypost.files.wordpress.com/2017/05/shutterstock_633482957.jpg?quality=90&strip=all&w=618&h=410&crop=1",
	"https://upload.wikimedia.org/wikipedia/commons/thumb/e/e4/Canada_quote_flight_cropped_and_NR.jpg/1200px-Canada_quote_flight_cropped_and_NR.jpg",
	"https://www.aflac.com/_global-assets/images/mockup-images-do-not-use-for-final/new-final-images/product-cost-ducks/2015-03-12-cost-of-policy-milk-duck-300x300.png",
	"https://www.eastcountyzoo.com/images/quote%20gosling%20african%20(4).jpg",
}
