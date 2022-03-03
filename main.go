package main

import (
	"encoding/json"
	"fmt"
	"log"
	"os"
	"strconv"
	"strings"

	"github.com/gocolly/colly"
	"github.com/gocolly/colly/extensions"
)

type Article struct {
	Title string `json:"title"`
	Date string `json:"date"`
}

func main() {
	articles := []Article{}

	numberOfPages := 0

	// Instantiate default collector
	collector := colly.NewCollector(
		colly.AllowedDomains("darknetlive.com"),
		colly.CacheDir("./.cache"),
	)

	extensions.RandomUserAgent(collector)

	articleCollector := collector.Clone()

	// On every a article grab the title and date
	// collector.OnHTML("article.a1e > a", func(e *colly.HTMLElement) {
	// 	link := e.Attr("href")
	// 	articleCollector.Visit(e.Request.AbsoluteURL(link))
	// }) 

	collector.OnHTML("li.a2j:last-child > a", func(e *colly.HTMLElement){
		numberOfPages += 1
		fmt.Println(e.Attr("href"))
	})

	articleCollector.OnHTML(".h-entry", func(e *colly.HTMLElement) {
		e.ForEachWithBreak(".a2u > p", func(i int, p *colly.HTMLElement) bool {
			paragraph := p.Text
			if strings.Contains(strings.ToLower(paragraph), "fentanyl") {
				title := e.ChildText(".a1i")
				date := e.ChildText(".a2t")

				newArticle := Article{}
				newArticle.Title = title
				newArticle.Date = date

				articles = append(articles, newArticle)
				return false
			}
			return true
		})
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

	json, _ := json.MarshalIndent(articles, "", "  ")

	writeError := os.WriteFile("data.json", json, 0644)
	if writeError != nil {
		panic(writeError)
	}

}