package tistory

import (
	"fmt"
	"net/http"
	"time"

	"github.com/PuerkitoBio/goquery"
)

func findTitle(url string, title_c chan string) {
	res, _ := http.Get(url)

	doc, _ := goquery.NewDocumentFromResponse(res)
	title_c <- doc.Find(".title_view").Text()
}

func Crawler_tistory() {
	check_start_time := time.Now()

	title_c := make(chan string)

	url := "https://naithar01.tistory.com/"

	for i := 90; i < 101; i++ {
		change_url := fmt.Sprintf("%s%d", url, i)
		go findTitle(change_url, title_c)
	}

	for j := 90; j < 101; j++ {
		fmt.Println(<-title_c)
	}

	fmt.Println(time.Since(check_start_time))
}
