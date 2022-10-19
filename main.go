package main

import (
	"encoding/json"
	"fmt"
	"log"
	"net/http"

	"github.com/PuerkitoBio/goquery"
)

type Product struct {
	ImagePath string `json:"imagePath"`
	Name      string `json:"name"`
	Price     string `json:"price"`
}

func main() {

	webPage := "https://tryguys.com"
	response, err := http.Get(webPage)

	if err != nil {
		log.Fatal(err)
	}

	defer response.Body.Close()

	if response.StatusCode != 200 {
		log.Fatal("Failed to fetch data")
	}

	doc, err2 := goquery.NewDocumentFromReader(response.Body)

	if err2 != nil {
		log.Fatal(err2)
	}

	var products []Product

	doc.Find("div.product").Each(func(index int, selector *goquery.Selection) {
		productImage, _ := selector.Find("img.product_card__image").Attr("src")
		productName := selector.Find("a.title").Text()
		productPrice := selector.Find("span.money").Text()

		newProduct := Product{ImagePath: productImage, Name: productName, Price: productPrice}
		products = append(products, newProduct)
	})

	j, _ := json.Marshal(products)

	if j != nil {
		j, _ = json.MarshalIndent(products[1:], "", "  ")
		fmt.Println(string(j))
	}

}
