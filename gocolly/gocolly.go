package gocolly

import (
	"encoding/json"
	"fmt"
	"log"
	"os"

	"github.com/gocolly/colly"
)

type Product struct {
	ImagePath string `json:"imagePath"`
	Name      string `json:"name"`
	Price     string `json:"price"`
}

func GoColly() {
	c := colly.NewCollector(
		colly.AllowedDomains("tryguys.com"),
	)

	products := make([]Product, 0)

	c.OnHTML("div.product", func(e *colly.HTMLElement) {

		product := Product{}
		product.Name = e.ChildText(".title")
		product.ImagePath = e.ChildAttr(".product_card__image", "src")
		product.Price = e.ChildText(".money")

		products = append(products, product)

	})

	c.OnResponse(func(r *colly.Response) {
		fmt.Println("Status code : ", r.StatusCode)
	})

	c.OnRequest(func(r *colly.Request) {
		fmt.Println("\nVisiting", r.URL)
	})

	c.OnScraped(func(r *colly.Response) {

		fmt.Println("Finished\n", r.Request.URL)

		js, err := json.MarshalIndent(products, "", "    ")
		if err != nil {
			log.Fatal(err)
		}
		fmt.Println(string(js))

		if err := os.WriteFile("products.json", js, 0664); err == nil {
			fmt.Println("Data written to file successfully")
		}

	})

	c.Visit("https://tryguys.com")
}
