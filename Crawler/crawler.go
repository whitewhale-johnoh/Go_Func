package main

import (
	"fmt"
	"golang.org/x/net/html"
	"log"
	"net/http"
	"runtime"
	"sync"
)

func fetch(url string) (*html.Node, error) {
	res, err := http.Get(url)
	if err != nil {
		log.Println(err)
		return nil, err
	}

	doc, err := html.Parse(res.Body)
	if err != nil {
		log.Println(err)
		return nil, err
	}
	return doc, nil
}

func parseFollowing(doc *html.Node) []string {
	var urls = make([]string, 0)

	var f func(*html.Node)
	f = func(n *html.Node) {
		if n.Type == html.ElementNode && n.Data == "img" { //img 태그
			for _, a := range n.Attr {
				if a.Key == "class" && a.Val == "avatar left" {
					//class가 avartar left인 요소
					for _, a := range n.Attr {
						if a.Key == "alt" {
							fmt.Println(a.Val) //사용자 이름 출력
							break
						}
					}
				}

				if a.Key == "class" && a.Val == "gravatar" { //class가 gravatar인 요소
					user := n.Parent.Attr[0].Val //부모요소의 첫번째 속성(href)

					//사용자 이름으로 팔로잉 URL 조합
					urls = append(urls, "https://github.com"+user+"/following")
					break
				}
			}
		}
		for c := n.FirstChild; c != nil; c = c.NextSibling {
			f(c) //재귀호출로 자식과 형제를 모두 탐색
		}
	}
	f(doc)

	return urls
}

var fetched = struct {
	m map[string]error
	sync.Mutex
}{m: make(map[string]error)}

func crawl(url string) {
	fetched.Lock()                   //map은 뮤텍스로 보호
	if _, ok := fetched.m[url]; ok { // URL 중복 처리 여부를 검사
		fetched.Unlock()
		return
	}
	fetched.Unlock()

	doc, err := fetch(url) //url에서 파싱된 HTML 데이터를 가져옴

	fetched.Lock()
	fetched.m[url] = err //가져온 URL은 맵에 URL과 에러 값 저장
	fetched.Unlock()

	urls := parseFollowing(doc) //사용자 정보 출력, 팔오잉 URL을 구함

	done := make(chan bool)
	for _, u := range urls { //URL 개수만큼 고루틴 생성, 재귀호출
		go func(url string) {
			crawl(url)
			done <- true
		}(u)
	}

	for i := 0; i < len(urls); i++ {
		<-done //고루틴이 모두 실행될때까지 대기
	}
}

func main() {
	numCPUs := runtime.NumCPU()
	runtime.GOMAXPROCS(numCPUs)

	crawl("https://github.com/pyrasis/following")
}
