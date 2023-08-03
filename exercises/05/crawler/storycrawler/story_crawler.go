package storycrawler

import (
	"fmt"
	"log"
	"strings"
	"sync"

	"github.com/gocolly/colly"
)

type Story struct {
	Slug       string
	Title      string
	Thumbnail  string
	Author     string
	Categories []string
	Status     string
	Rate       StoryRate
}
type StoryRate struct {
	Value       float32
	BestValue   float32
	Count       int
	Description string
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
