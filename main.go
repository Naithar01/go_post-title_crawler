package main

import (
	"fmt"
	"net/http"
	"os"
	"time"

	"github.com/PuerkitoBio/goquery"
)

func getPostInfo(url string, g_title chan string, save_titles []string) {
	res, err := http.Get(url)

	if err != nil {
		panic(err.Error())
	}

	defer res.Body.Close()

	doc, _ := goquery.NewDocumentFromResponse(res)

	// tr > td > a
	posts := doc.Find(".us-post")

	posts.Each(func(i int, s *goquery.Selection) {
		g_title <- fmt.Sprintln(s.Find("td > a").Text())
	})

}

func main() {
	//check time
	set_time := time.Now()

	// 글 제목 채널
	g_title := make(chan string)
	// 사이트 페이지
	page := "https://gall.dcinside.com/board/lists/?id=ohmygirl&page="

	// 몇 페이지 검색할지, 시작은 1 페이지 부터
	end_page := 20

	// 검색한 제목을 저장해줄 문자열 변수
	save_titles := make([]string, end_page*50)

	// 1 ~ 5 페이지
	for i := 1; i <= end_page; i++ {
		cnt_page := fmt.Sprintf("%s%d", page, i)

		go getPostInfo(cnt_page, g_title, save_titles)
	}

	// 검색이 모두 끝나고 종료해주는 조건이 필요
	for j := 0; j < end_page*50-1; j++ {
		save_titles = append(save_titles, <-g_title)
	}

	file, err := os.Create("test.txt")

	if err != nil {
		panic(err.Error())
	}

	defer file.Close()

	for i := range save_titles {
		file.WriteString(save_titles[i])
	}

	fmt.Println(time.Since(set_time))

}
