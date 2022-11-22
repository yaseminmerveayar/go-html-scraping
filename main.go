package main

import (
	"fmt"

	"example.com/packages/gocolly"
	"example.com/packages/goquery"
)

func main() {
	fmt.Println("\n--------------------------------------------------------Scraping with goquery--------------------------------------------------------")
	goquery.GoQuery()

	fmt.Println("\n--------------------------------------------------------Scraping with gocolly--------------------------------------------------------")
	gocolly.GoColly()
}
