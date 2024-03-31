package main

import (
	"fmt"
)

type Fetcher interface {
	// Fetch 返回 URL 的 body 内容，并且将在这个页面上找到的 URL 放到一个 slice 中。
	Fetch(url string) (body string, urls []string, err error)
}

// Crawl 使用 fetcher 从某个 URL 开始递归的爬取页面，直到达到最大深度。
func Crawl(url string, depth int, fetcher Fetcher) {
	// 使用 map 来跟踪已经访问过的 URL
	visited := make(map[string]bool)

	// 创建一个 channel 来协调 goroutine
	ch := make(chan string)

	// 启动一个 goroutine 开始爬取
	go func() {
		crawlHelper(url, depth, fetcher, visited, ch)
		close(ch)
	}()

	// 处理 channel 中的结果
	for url := range ch {
		fmt.Printf("found: %s\n", url)
	}
}

func crawlHelper(url string, depth int, fetcher Fetcher, visited map[string]bool, ch chan<- string) {
	// 如果达到最大深度或者 URL 已经访问过，则返回
	if depth <= 0 || visited[url] {
		return
	}

	// 标记当前 URL 为已访问
	visited[url] = true

	// 抓取 URL
	body, urls, err := fetcher.Fetch(url)
	if err != nil {
		fmt.Println(err)
		return
	}

	// 将找到的 URL 发送到 channel
	ch <- fmt.Sprintf("%s %q", url, body)

	// 并行抓取找到的 URL
	for _, u := range urls {
		go crawlHelper(u, depth-1, fetcher, visited, ch)
	}
}

func main() {
	Crawl("https://golang.org/", 4, fetcher)
}

// fakeFetcher 是返回若干结果的 Fetcher。
type fakeFetcher map[string]*fakeResult

type fakeResult struct {
	body string
	urls []string
}

func (f fakeFetcher) Fetch(url string) (string, []string, error) {
	if res, ok := f[url]; ok {
		return res.body, res.urls, nil
	}
	return "", nil, fmt.Errorf("not found: %s", url)
}

// fetcher 是填充后的 fakeFetcher。
var fetcher = fakeFetcher{
	"https://golang.org/": &fakeResult{
		"The Go Programming Language",
		[]string{
			"https://golang.org/pkg/",
			"https://golang.org/cmd/",
		},
	},
	"https://golang.org/pkg/": &fakeResult{
		"Packages",
		[]string{
			"https://golang.org/",
			"https://golang.org/cmd/",
			"https://golang.org/pkg/fmt/",
			"https://golang.org/pkg/os/",
		},
	},
	"https://golang.org/pkg/fmt/": &fakeResult{
		"Package fmt",
		[]string{
			"https://golang.org/",
			"https://golang.org/pkg/",
		},
	},
	"https://golang.org/pkg/os/": &fakeResult{
		"Package os",
		[]string{
			"https://golang.org/",
			"https://golang.org/pkg/",
		},
	},
}
