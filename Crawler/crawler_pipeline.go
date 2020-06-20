package main

import (
	"fmt"
	"golang.org/x/net/html"
	"log"
	"net/http"
	"runtime"
	"sync"
)

var fetched = struct {
	m map[string]error //중복 검사를 위한 URL과 에러값 저장
	sync.Mutex
}{m: make(map[string]error)} // 변수를 선언하면서 이름이 없는 구조체를 정의하고
//초기값을 생성하여 대입

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

type result struct { //결과값 저장할 구조체
	url  string // 가져온 URL
	name string // 사용자 이름
}

func worker(done <-chan struct{}, urls chan string, c chan<- result) {
	for url := range urls { // urls 채널에서 URL을 가져옴
		select {
		case <-done: //채널이 닫히면 worker함수를 빠져나옴
			return
		default:
			crawl(url, urls, c) //URL 처리
		}
	}
}

func parseFollowing(doc *html.Node, urls chan string) <-chan string {
	name := make(chan string)

	go func() { /*교착상태가 되지 않도록 고루틴으로 만듬 ->
		  Cawl, parseFollowing 함수도 worker함수의 for url := range urls {}안에서 실행되고 있으므로
		  urls 채널에 바로 값을 보내면 교착 상태가 됩니다.
		  그러므로 반드시 고루틴을 생성하여 채널에 값을 보내야합니다.
		  html 정보를 탐색하는 부분은 반복되는 부분이 많고, 나중에라도 로직이 복잡해질 수 있다.
		  따라서 시간이 약간이나마 오래걸릴수 있으므로 고루틴으로 만들어서
		  worker함수와는 별개로 실행*/
		var f func(*html.Node)
		f = func(n *html.Node) {
			if n.Type == html.ElementNode && n.Data == "img" {
				for _, a := range n.Attr {
					if a.Key == "class" && a.Val == "avartar left" {
						for _, a := range n.Attr {
							if a.Key == "alt" {
								name <- a.Val
								break
							}
						}
					}

					if a.Key == "class" && a.Val == "gravatar" {
						user := n.Parent.Attr[0].Val

						urls <- "https://github.com" + user + "/following"
						break
					}
				}
			}
			for c := n.FirstChild; c != nil; c = c.NextSibling {
				f(c) //재귀호출로 자식과 형제를 모두 탐색
			}
		}
		f(doc)
	}()

	return name //채널을 리턴
}

func crawl(url string, urls chan string, c chan<- result) {
	fetched.Lock() // map은 Mutex로 보호
	if _, ok := fetched.m[url]; ok {
		fetched.Unlock()
		return
	}
	fetched.Unlock()

	doc, err := fetch(url) //url에서 파싱된 데이터를 가져옴
	if err != nil {        //url을 가져오지 못했을 때
		go func(u string) { //교착상태가 되지 않도록 고루틴을 생성
			urls <- u //채널에 URL을 보냄
		}(url)
		return
	}

	fetched.Lock()
	fetched.m[url] = err //가져온 URL은 맵에 URL과 에러 값 저장
	fetched.Unlock()
	name := <-parseFollowing(doc, urls) //사용자 정보, 팔로잉 URL을 구함
	c <- result{url, name}              //가져온 URL과 사용자 이름을 구조체 인스턴스로 생성하여 채널 c로 보냄
}

func main() {
	numCPUs := runtime.NumCPU()
	runtime.GOMAXPROCS(numCPUs) // 모든 CPU를 사용하도록 설정

	urls := make(chan string)   // 작업을 요청할 채널
	done := make(chan struct{}) //작업 고루틴에 정지신호를 보낼 채널
	c := make(chan result)      //결과값을 저장할 채널

	var wg sync.WaitGroup
	const numWorkers = 10
	wg.Add(numWorkers)
	for i := 0; i < numWorkers; i++ { //작업을 처리할 고루틴을 10개 생성
		go func() {
			worker(done, urls, c)
			wg.Done()
		}()
	}

	go func() {
		wg.Wait() //고루틴이 끝날때 까지 대기
		close(c)  // 대기가 끝나면 결과값 채널을 닫음
	}()

	urls <- "https://github.com/pyrasis/following" //최초 URL을 보냄

	count := 0
	for r := range c { //결과 채널에 값이 들어올 때까지 대기한 뒤 값을 가져옴
		fmt.Println(r.name)

		count++
		if count > 100 { // 100명만 출력한 뒤
			close(done) //done을 닫아서 worker고루틴을 종료
			break
		}
	}
}
