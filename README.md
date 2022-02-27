# Fentanyl Epidemic Scraper

The mission for this project was to create a scraper that will pull articles in relation to fentanyl for the goal of mapping the exponential growth of the fentanyl epidemic & the increase of fentanyl distrubtion through illegal sites.

## How it Works

How it works is quite simple, there is a simple go scraper built using colly which visits https://darknetlive.com/post/ and its next pages. The scraper then grabs the title and checks to see if the keyword fent is within the title, if it is it then saves the title alongside the date.

## Tech Stack

- GoLang
- Colly
- JSON
