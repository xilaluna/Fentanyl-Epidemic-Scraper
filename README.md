# Fentanyl Epidemic Scraper

[![Go Report Card](https://goreportcard.com/badge/github.com/xilaluna/Fentanyl-Epidemic-Scraper)](https://goreportcard.com/report/github.com/xilaluna/Fentanyl-Epidemic-Scraper)

The mission for this project was to create a scraper that would pull articles pertaining to fentanyl, for the goal of mapping the trend of the fentanyl epidemic & the increase of fentanyl distrubtion through illegal sites.

[graph image](/assets/graph.png)

## How it Works

How it works is quite simple, there is a simple go scraper built using colly which visits https://darknetlive.com/post/ and its next pages. The scraper then looks into each article and searches all the paragraphs for the word "fentanyl", if the article includes the word it will save the article name and the date it was published.

## Tech Stack

- [GoLang](https://go.dev/)
- [Colly](http://go-colly.org/)
- JSON
- HTML
- [Chart.js](https://www.chartjs.org/)
