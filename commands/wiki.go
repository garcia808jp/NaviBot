// NaviBot: Discord bot for digital assistance
// lw command

package commands

import (
	// standard
	"net/url"

	// third-party
	"cgt.name/pkg/go-mwclient"
	"cgt.name/pkg/go-mwclient/params"
)

// Register the command for the CommandList
func init() {
	lwDoc := command{
		name:        "lw - link to pages in the lain wiki",
		synopsis:    "lw __query__",
		description: "WIP, only accepts one word for the query",
		example:     "WIP",
		origin:      "built-in",
		Exec:        lw,
	}

	CommandList["lw"] = lwDoc
}

// lw command
// interacts with the MediaWiki API to make a URL of the requested page
func lw(msgArray []string) (msgOut string) {
	// quit if there are not enough arguments
	if len(msgArray) == 2 {
		msgOut = "command accepts one word for the query at the moment"
		return
	}

	// lain.wiki URLs for the MediaWiki API and article path
	// DOES NOT ACCOUNT FOR SPECIAL NAMESPACE PATHS
	const lainAPI, lainIndex string = "https://lain.wiki/api.php", "https://lain.wiki/wiki/"

	// Initialize a *Client with New(), specifying the wiki's API URL
	// and your HTTP User-Agent. Try to use a meaningful User-Agent.
	w, err := mwclient.New(lainAPI, "phossil-NaviBot/0.0.1")
	if err != nil {
		panic(err) // Malformed URL
	}

	parameters := params.Values{
		"action":   "query",
		"list":     "search",
		"srlimit":  "1",
		"srsearch": msgArray[2],
	}
	response, err := w.Get(parameters)
	if err != nil {
		panic(err)
	}

	// USE "TOTALHITS" TO AVOID ERROR IF NO RESULT IS FOUND

	// open json object "query"
	query, err := response.GetObject("query")
	if err != nil {
		panic(err)
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

	msgOut = lainIndex + result
	return
}
