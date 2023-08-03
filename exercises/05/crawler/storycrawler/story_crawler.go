package storycrawler

import (
	"fmt"
	"log"
	"strconv"
	"strings"
	"sync"

	"github.com/gocolly/colly/v2"
)

type Story struct {
	Slug        string
	Title       string
	Thumbnail   string
	Author      string
	Categories  []string
	Source      string
	Status      string
	Rate        StoryRate
	Description string
}
type StoryRate struct {
	Value     float32
	BestValue float32
	Count     int
}

func CrawlStory(url string) (Story, error) {
	resultStory := Story{}
	// slug
	urlParts := strings.Split(strings.TrimSuffix(url, "/"), "/")
	resultStory.Slug = urlParts[len(urlParts)-1]

	var wg sync.WaitGroup
	wg.Add(1)
	c := colly.NewCollector()
	c.OnHTML(".col-truyen-main", func(e *colly.HTMLElement) {
		resultStory.Title = e.ChildText(".col-truyen-main h3.title")
		resultStory.Thumbnail = e.ChildAttr(".col-truyen-main .book img", "src")
		resultStory.Author = e.ChildText(".col-truyen-main [itemprop='author']")
		resultStory.Categories = e.ChildTexts(".col-truyen-main a[itemprop='genre']")
		resultStory.Source = e.ChildText(".col-truyen-main .source")
		resultStory.Status = e.ChildText(".col-truyen-main .info:last-child span")
		description, descriptionError := e.DOM.Find(".col-truyen-main [itemprop='description']").First().Html()
		if descriptionError == nil {
			resultStory.Description = description
		}

		strRateValue, strRateValueError := strconv.ParseFloat(e.ChildText(".col-truyen-main [itemprop='ratingValue']"), 32)
		if strRateValueError == nil {
			resultStory.Rate.Value = float32(strRateValue)
		}

		strRateBestValue, strRateBestValueError := strconv.ParseFloat(e.ChildText(".col-truyen-main [itemprop='bestRating']"), 32)
		if strRateBestValueError == nil {
			resultStory.Rate.BestValue = float32(strRateBestValue)
		}
		strRateCount, strRateCountError := strconv.Atoi(e.ChildText(".col-truyen-main [itemprop='ratingCount']"))
		if strRateCountError == nil {
			resultStory.Rate.Count = strRateCount
		}

		wg.Done()
	})
	c.OnRequest(func(r *colly.Request) {
		fmt.Println("Crawling... ", r.URL)
	})

	c.Visit(url)
	wg.Wait()
	return resultStory, nil
}

func CrawlStories(urls []string) chan Story {
	resultChan := make(chan Story, len(urls))
	var wg sync.WaitGroup
	maxConcurrent := 5
	semaphore := make(chan struct{}, maxConcurrent)

	for _, url := range urls {
		wg.Add(1)
		go func(url string) {
			defer wg.Done()
			semaphore <- struct{}{}

			data, err := CrawlStory(url)
			if err != nil {
				log.Println("Error fetching data from URL:", err)
				<-semaphore
				return
			}
			resultChan <- data
			<-semaphore

		}(url)
	}
	wg.Wait()
	close(resultChan)
	return resultChan

}
