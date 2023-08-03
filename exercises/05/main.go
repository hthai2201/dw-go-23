package main

import (
	"fmt"

	"github.com/hthai2201/dw-go-23/exercises/05/crawler/storycrawler"
	"github.com/hthai2201/dw-go-23/exercises/05/sitemapreader"
)

func main() {

	result, error := sitemapreader.ReadSitemapFromGZLink("https://truyenfull.vn/sitemap/truyen_sitemap.xml.gz")
	if error != nil {
		fmt.Println(error)
	}
	first5Stories := result.Urls[:5]

	var urls []string
	for _, urlObj := range first5Stories {
		urls = append(urls, urlObj.Loc)
	}
	resultChan := storycrawler.CrawlStories(urls)
	for data := range resultChan {
		fmt.Println("Fetched Data:", data)
	}

	fmt.Println("Crawling done.")
}
