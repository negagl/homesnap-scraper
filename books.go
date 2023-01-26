package main

import (
	"encoding/json"
	"fmt"
	"os"

	"github.com/gocolly/colly"
)

type book struct {
	Name   string `json:"name"`
	Price  string `json:"price"`
	ImgUrl string `json:"imgurl"`
}

func main2() {
	// Instantiate default collector
	c := colly.NewCollector(
		colly.AllowedDomains("books.toscrape.com"),
	)

	var books []book

	// On every a element which has href attribute call callback
	c.OnHTML("article[class=product_pod]", func(e *colly.HTMLElement) {
		book := book{
			Name:   e.ChildText("h3"),
			Price:  e.ChildText("p.price_color"),
			ImgUrl: e.ChildAttr("a", "href"),
		}

		fmt.Printf("Book: %q | Price: %s\n", book.Name, book.Price)
		books = append(books, book)
	})

	// Next Page
	c.OnHTML("li[class=next]", func(e *colly.HTMLElement) {
		next_page := e.Request.AbsoluteURL(e.ChildAttr("a", "href"))
		c.Visit(next_page)
	})

	c.Visit("http://books.toscrape.com/catalogue/page-1.html")
	// fmt.Println(books)

	content, err := json.Marshal(books)

	if err != nil {
		fmt.Println(err.Error())
	}

	os.WriteFile("books.json", content, 0644)
}
