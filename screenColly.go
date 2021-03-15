package main

import (
	"fmt"
	"github.com/gocolly/colly"
)

type screenPlay struct {
	FileName string
	FileText []string
}

func main() {
	// Instantiate default collector
	c := colly.NewCollector(
		// Visit only domains: hackerspaces.org, wiki.hackerspaces.org
		colly.AllowedDomains("imsdb.com/TV/Futurama"),
	)

	myScreenPlays := []screenPlay{}
	// setting a hard limit on how many links the collector will explore.

	numVisited := 0
	c.OnRequest(func(r *colly.Request) {
		if numVisited > 200 {
			r.Abort()
		}
		numVisited++
	})

	// On every a element which has href attribute call callback
	c.OnHTML("a[href]", func(e *colly.HTMLElement) {
		link := e.Attr("href")

		// If the link in question begins with either of these strings, visit it.
		if strings.HasPrefix(link, "/TV Transcripts/Futurama") || strings.HasPrefix(link, "/transcripts/Futurama") {
			e.Request.Visit(link)
		}

	})

	// Before making a request print "Visiting ..."
	c.OnRequest(func(r *colly.Request) {
		fmt.Println("Visiting", r.URL.String())
	})

	c.OnHTML("td", func(e *colly.HTMLElement) {
		if e.Attr("class") == "scrtext" {
			log.Println("Futurama script found!")
			// e.ChildText
			// How to iterate over all lines
			// contained in between <pre> </pre>
			// tags

			fullText := []string{}

		}
	})

	// TODO: if the line in question reads td class = "scrtext" then
	// we need to proceed with extracting the wrapped text

	c.Visit("imsdb.com/TV/Futurama.html")
}
