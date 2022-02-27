package main

import (
	"encoding/json"
	"fmt"
	"log"
	"os"
	"strconv"
	"strings"

	"github.com/gocolly/colly"
)

type Article struct {
	Title string `json:"title"`
	Date string `json:"date"`
}

func main() {
	// Instantiate default collector
	collector := colly.NewCollector(
		colly.AllowedDomains("darknetlive.com"),
	)

	articles := []Article{}

	// On every a article grab the title and date
	collector.OnHTML(".a1e", func(e *colly.HTMLElement) {
		title := e.ChildText(".a1f")
		date := e.ChildText(".a1g > time")
		if strings.Contains(strings.ToLower(title), "fent") {
			newArticle := Article{}

			newArticle.Title = title
			newArticle.Date = date

			articles = append(articles, newArticle)
		}
		
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
	for i := 2; i < 20; i++ {
		collector.Visit("https://darknetlive.com/post/page/" + strconv.Itoa(i) + "/")
	}

	json, _ := json.MarshalIndent(articles, "", "  ")

	writeError := os.WriteFile("data.json", json, 0644)
	if writeError != nil {
		panic(writeError)
	}

}