package main

import (
	"fmt"
	"log"
	"strconv"

	"github.com/gocolly/colly"
)

type Article struct {
	Title string `json:"title"`
	Date string `json:"date"`
}

func main() {
	// articles := []Article{}

	// Instantiate default collector
	collector := colly.NewCollector(
		colly.AllowedDomains("darknetlive.com"),
	)

	articleCollector := collector.Clone()

	// On every a article grab the title and date
	collector.OnHTML("article.a1e > a", func(e *colly.HTMLElement) {
		link := e.Attr("href")
		link_title := e.Attr("title")
		fmt.Println(link_title)
		articleCollector.Visit(e.Request.AbsoluteURL(link))
	}) 

	articleCollector.OnHTML(".h-entry", func(e *colly.HTMLElement) {
		title := e.ChildText(".a1i")
		fmt.Printf("%v", title)
	})

	// Before making a request print "Visiting ..."
	collector.OnRequest(func(r *colly.Request) {
		fmt.Println("Visiting", r.URL.String())
	})

	collector.OnError(func(_ *colly.Response, err error) {
    	log.Println("Something went wrong:", err)
	})

	// Start scraping url and its numbered pages
	collector.Visit("https://darknetlive.com/post/")
	for i := 2; i < 3; i++ {
		collector.Visit("https://darknetlive.com/post/page/" + strconv.Itoa(i) + "/")
	}

	// json, _ := json.MarshalIndent(articles, "", "  ")

	// writeError := os.WriteFile("data.json", json, 0644)
	// if writeError != nil {
	// 	panic(writeError)
	// }

}