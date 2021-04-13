package main

import (
	"fmt"
	"github.com/gocolly/colly"
	"io"
	"net/http"
	"os"
	"strings"
)

func download(url, filename string) (err error) {
	fmt.Println("Downloading ", url, " to ", filename)

	resp, err := http.Get(url)
	if err != nil {
		return
	}
	defer resp.Body.Close()

	f, err := os.Create(filename)
	if err != nil {
		return
	}
	defer f.Close()

	_, err = io.Copy(f, resp.Body)
	return
}

func main() {

	show := "Seinfeld"

	wd, _ := os.Getwd()
	directory := wd + "/" + show + "/"

	c := colly.NewCollector(
		// This prevents Colly from wandering off into
		// irrelevant places
		colly.AllowedDomains("imsdb.com"),
	)

	fmt.Println("Here we go!")
	// setting a hard limit on how many links the collector will explore.

	numVisited := 0
	c.OnRequest(func(r *colly.Request) {
		fmt.Println("Visited ", numVisited, " pages so far!")
		if numVisited > 600 {
			r.Abort()
		}
		numVisited++
	})

	// On every a element which has href attribute call callback
	c.OnHTML("a[href]", func(e *colly.HTMLElement) {
		link := e.Attr("href")
		fmt.Println(link)

		// If the link in question begins with the given string, visit it.
		if strings.HasPrefix(link, "/TV Transcripts/"+show) {
			e.Request.Visit(link)
		}
		// If the a href element begins with this string, then trim the string to get the name and download the linked html file
		if strings.HasPrefix(link, "/transcripts/"+show+"-") {
			myLink := "https://imsdb.com" + link

			name := strings.TrimLeft(link, "/transcripts/"+show+"-")
			fileTitle := directory + name

			fmt.Println("Checking if " + fileTitle + " exists ...")
			if _, err := os.Stat(fileTitle); os.IsNotExist(err) {
				err := download(myLink, fileTitle)
				if err != nil {
					panic(err)
				}
				fmt.Println(fileTitle + " saved!")
			} else {
				fmt.Println(fileTitle + " already exists!")
			}
		}
	})

	c.OnRequest(func(r *colly.Request) {
		fmt.Println("Visiting", r.URL)
	})

	c.Visit("https://imsdb.com/TV/" + show + ".html")
}
