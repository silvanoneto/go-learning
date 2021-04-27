// Package chapter01 06_fetchall contains the code studied on section 06
package chapter01

import (
	"bufio"
	"fmt"
	"io"
	"io/ioutil"
	"net/http"
	"os"
	"strings"
	"time"
)

// FetchAll search URLs in parallel and display time spent and sizes
func FetchAll() {
	start := time.Now()
	ch := make(chan string)
	for _, url := range os.Args[1:] {
		go fetch(url, ch)
	}
	for range os.Args[1:] {
		fmt.Println(<-ch)
	}
	fmt.Printf("%.2fs elapsed\n", time.Since(start).Seconds())
}

func fetch(url string, ch chan<- string) {
	start := time.Now()
	resp, err := http.Get(url)
	if err != nil {
		ch <- fmt.Sprint(err)
		return
	}

	nbytes, err := io.Copy(ioutil.Discard, resp.Body)
	resp.Body.Close()
	if err != nil {
		ch <- fmt.Sprintf("while reading %s: %v", url, err)
		return
	}
	secs := time.Since(start).Seconds()
	ch <- fmt.Sprintf("%.2fs %7d %s", secs, nbytes, url)
}

// Exercise10 search URLs in parallel, save the body content in a file and
// display time spent and sizes
func Exercise10() {
	fetch := func(url string, ch chan<- string) {
		start := time.Now()
		resp, err := http.Get(url)
		if err != nil {
			ch <- fmt.Sprint(err)
			return
		}

		filename := url
		filename = strings.TrimPrefix(filename, "http://")
		filename = strings.TrimPrefix(filename, "https://")
		filename = strings.TrimSuffix(filename, "/")
		os.Mkdir("Exercise10", 0755)
		file, err := os.Create("Exercise10/" + filename)
		if err != nil {
			ch <- fmt.Sprint(err)
			return
		}

		nbytes, err := io.Copy(file, resp.Body)
		file.Close()
		resp.Body.Close()
		if err != nil {
			ch <- fmt.Sprintf("while reading %s: %v", url, err)
			return
		}
		secs := time.Since(start).Seconds()
		ch <- fmt.Sprintf("%.2fs %7d %s", secs, nbytes, url)
	}

	start := time.Now()
	ch := make(chan string)
	for _, url := range os.Args[1:] {
		go fetch(url, ch)
	}
	for range os.Args[1:] {
		fmt.Println(<-ch)
	}
	fmt.Printf("%.2fs elapsed\n", time.Since(start).Seconds())
}

// Exercise11 search URLs from an text file and display time spent and sizes
func Exercise11() {
	file, err := os.Open(os.Args[1])
	if err != nil {
		fmt.Fprintf(os.Stderr, "dup3: %v\n", err)
		return
	}
	defer file.Close()

	var linesCount int
	urlList := bufio.NewScanner(file)
	for urlList.Scan() {
		linesCount++
	}
	file.Seek(0, io.SeekStart)
	urlList = bufio.NewScanner(file)

	start := time.Now()
	ch := make(chan string)
	for urlList.Scan() {
		url := urlList.Text()
		hasHttpPrefix := strings.HasPrefix(url, "http://")
		hasHttpsPrefix := strings.HasPrefix(url, "https://")
		if !hasHttpPrefix && !hasHttpsPrefix {
			url = "http://" + url
		}

		go fetch(url, ch)
	}
	for i := 0; i < linesCount; i++ {
		fmt.Println(<-ch)
	}
	fmt.Printf("%.2fs elapsed\n", time.Since(start).Seconds())
}
