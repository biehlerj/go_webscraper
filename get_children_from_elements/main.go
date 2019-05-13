package main

import (
	"fmt"
	"github.com/PuerkitoBio/goquery"
	"log"
	"net/http"
	"strings"
)

func navigateChildrens() {
	resp, err := http.Get("https://github.com/trending")
	if err != nil {
		log.Fatal(err)
	}
	defer resp.Body.Close()

	if resp.StatusCode != 200 {
		log.Fatalf("Status code error: %d %s", resp.StatusCode, resp.Status)
	}
	doc, err := goquery.NewDocumentFromReader(resp.Body)
	if err != nil {
		log.Fatal(err)
	}
	fmt.Println(doc.Find("title").Text())
	olSelection := doc.Find("ol")
	olSelection.Children().Each(func(i int, s *goquery.Selection) { // using .Children() on the ol selection to get all li
		fmt.Println(strings.TrimSpace(s.Find("h3").Text()))
	})
}

func main() {
	navigateChildrens()
}
