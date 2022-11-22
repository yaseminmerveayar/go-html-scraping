package goquery

import (
	"encoding/json"
	"fmt"
	"log"
	"net/http"
	"os"

	"github.com/PuerkitoBio/goquery"
)

type Product struct {
	ImagePath string `json:"imagePath"`
	Name      string `json:"name"`
	Price     string `json:"price"`
}

func GoQuery() {

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

	products := make([]Product, 0)

	doc.Find("div.product").Each(func(index int, selector *goquery.Selection) {
		productImage, _ := selector.Find("img.product_card__image").Attr("src")
		productName := selector.Find("a.title").Text()
		productPrice := selector.Find("span.money").Text()

		newProduct := Product{ImagePath: productImage, Name: productName, Price: productPrice}
		products = append(products, newProduct)
	})

	js, err := json.MarshalIndent(products[1:], "", "   ")
	if err != nil {
		log.Fatal(err)
	} else {
		fmt.Println(string(js))
	}

	if err := os.WriteFile("products.json", js, 0664); err == nil {
		fmt.Println("Data written to file successfully")
	}

}
