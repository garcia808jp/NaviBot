// NaviBot: Discord bot for digital assistance
// LainBot commands

package lain

import (
	"strings"
	// standard
	"net/url"

	// third-party
	"cgt.name/pkg/go-mwclient"
	"cgt.name/pkg/go-mwclient/params"
)

// Register the command for the CommandList
func init() {
	wikiDoc := Doc{
		name:        "wiki - link to pages in the lain wiki",
		synopsis:    "wiki __query__",
		description: "WIP, only accepts one word for the query",
		example:     "WIP",
		origin:      "built-in",
		Exec:        wiki,
	}

	CommandList["wiki"] = wikiDoc
}

// wiki command
// interacts with the MediaWiki API to make a URL of the requested page
func wiki(msgArray []string) (msgOut string) {
	// quit if there are not enough arguments
	if len(msgArray) == 2 {
		msgOut = "command accepts one word for the query at the moment"
		return
	}

	// remove prefix and command from msgArray and join the fields
	msgArray[0], msgArray[1] = "", ""
	userQuery := strings.Join(msgArray, " ")

	// lain.wiki URLs for the MediaWiki API and article path
	// DOES NOT ACCOUNT FOR SPECIAL NAMESPACE PATHS
	const wikiAPI, wikiIndex string = "https://lain.wiki/api.php", "https://lain.wiki/wiki/"

	// Initialize a *Client with New(), specifying the wiki's API URL
	// and your HTTP User-Agent. Try to use a meaningful User-Agent.
	w, err := mwclient.New(wikiAPI, "phossil-NaviBot/0.0.1")
	if err != nil {
		panic(err) // Malformed URL
	}

	parameters := params.Values{
		"action":   "query",
		"list":     "search",
		"srlimit":  "1",
		"srsearch": userQuery,
	}
	response, err := w.Get(parameters)
	if err != nil {
		panic(err)
	}

	// here lies ugly

	// open json object "query"
	query, err := response.GetObject("query")
	if err != nil {
		panic(err)
	}
	// open json object "searchinfo"
	searchinfo, err := query.GetObject("searchinfo")
	if err != nil {
		panic(err)
	}
	// open "totalhits" from "searchinfo" to know the amount of hits for the given query
	totalhits, err := searchinfo.GetNumber("totalhits")
	if err != nil {
		panic(err)
	} else if totalhits == "0" {
		// return if there are no hits
		return "no results found"
	}
	// open json object of arrays (?) "search"
	search, err := query.GetObjectArray("search")
	if err != nil {
		panic(err)
	}
	// open the only array and find the page "title"
	title, err := search[0].GetString("title")
	if err != nil {
		panic(err)
	}
	// make string url friendly
	result := url.PathEscape(title)

	msgOut = wikiIndex + result
	return
}
