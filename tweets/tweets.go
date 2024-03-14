package tweets

import (
	"fmt"
	"net/http"

	"github.com/gocolly/colly"
)

func Hello(name string) string {
    // Return a greeting that embeds the name in a message.
    message := fmt.Sprintf("Hi, %v. Welcome!", name)
    return message
}

func ViewURL1(url string){
    doc, err := http.Get(url)
    if err != nil {
        fmt.Println(err)
    }
    fmt.Println(doc)
}

func ViewURL2(url string){
    c := colly.NewCollector()
    c.OnHTML("a[href]", func(e *colly.HTMLElement) {
        link := e.Attr("href")
        fmt.Println(link)
    })
    c.Visit(url)
}

