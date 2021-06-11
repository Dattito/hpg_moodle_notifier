package main

import (
	"fmt"
	"log"
	"net/http"
	"net/url"

	"github.com/PuerkitoBio/goquery"
)

func ExampleScrape() {
	// Request the HTML page.

	authid := "20052003DavSir"
	authpw := "a0a179"

	res, err := http.PostForm(fmt.Sprintf("https://vertretungsplan.hpg-speyer.de/pmwiki/pmwiki.php?n=Main.%s", authid),
		url.Values{
			"authid": {authid},
			"authpw": {authpw},
		},
	)
	if err != nil {
		log.Fatal(err)
	}
	defer res.Body.Close()
	if res.StatusCode != 200 {
		log.Fatalf("status code error: %d %s", res.StatusCode, res.Status)
	}

	// Load the HTML document
	doc, err := goquery.NewDocumentFromReader(res.Body)
	if err != nil {
		log.Fatal(err)
	}

	// Find the review items
	s := doc.Find("table").First()

	fmt.Println(s.Html())
}

func main() {
	ExampleScrape()
}
