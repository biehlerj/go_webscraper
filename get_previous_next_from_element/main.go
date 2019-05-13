package main

import (
	"fmt"
	"github.com/PuerkitoBio/goquery"
	"log"
	"net/http"
	"strings"
)

func navigateSiblings() {
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
	liSelection := doc.Find("ol li")
	fifthElement := liSelection.Eq(4) // using Eq() and passing the index we can navigate to the element with given index
	fmt.Println(strings.TrimSpace(fifthElement.Find("h3").Text()))
	fourthElement := fifthElement.Prev()
	fmt.Println(strings.TrimSpace(fourthElement.Find("h3").Text()))
	sixthElement := fifthElement.Next()
	fmt.Println(strings.TrimSpace(sixthElement.Find("h3").Text()))
}

func main() {
	navigateSiblings()
}
