package main

import (
	"fmt"
	"log"
	"strings"
	//"io"
	//"os"
	"github.com/PuerkitoBio/goquery"
	"net/http"
)

func githubTrendingScraper() {
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
	doc.Find("ol li").Each(func(i int, s *goquery.Selection) {
		repositoryName := strings.TrimSpace(s.Find("h3").Text())
		totalStarsToday := strings.TrimSpace(s.Find(".float-sm-right").Text())
		href, hasAttr := s.Find("a").Attr("href")
		if !hasAttr {
			href = "No valid url found"
		}
		fmt.Println(repositoryName, "\t", totalStarsToday, "\t", "https://github.com"+href)
	})

}

func main() {
	githubTrendingScraper()
}
