package main

import (
	"fmt"
	"log"

	"github.com/gocolly/colly"
)

type Article struct {
	Title string
	Date string
}

func scrape(url string) {
	// Instantiate default collector
	c := colly.NewCollector(
		colly.AllowedDomains("darknetlive.com"),
	)

	// Before making a request print "Visiting ..."
	c.OnRequest(func(r *colly.Request) {
		fmt.Println("Visiting", r.URL.String())
	})

	c.OnError(func(_ *colly.Response, err error) {
    	log.Println("Something went wrong:", err)
	})

	c.OnScraped(func(r *colly.Response) {
    	fmt.Println("Finished", r.Request.URL)
	})

	// On every a article grab the title and date
	c.OnHTML(".a1e", func(e *colly.HTMLElement) {
		fmt.Printf("Title: %v \n", e.ChildText(".a1f"))
		fmt.Printf("Date: %v \n", e.ChildText(".a1g > time"))
	}) 

	// Start scraping url
	c.Visit(url)
}


func main() {
	scrape("https://darknetlive.com/post/")
}