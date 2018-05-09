package application

import (
	"time"
	"os"
	"fmt"
	"net/http"
	"io/ioutil"
	"io"
)

//func fetch() {
//	for _, url := range os.Args[2:] {
//		resp, err := http.Get(url)
//		if err != nil {
//			fmt.Fprintf(os.Stderr, "fetch: %v\n", err)
//			os.Exit(1)
//		}
//
//		b, err := ioutil.ReadAll(resp.Body)
//		resp.Body.Close()
//		if err != nil {
//			fmt.Fprintf(os.Stderr, "read body: %v\n", err)
//			os.Exit(1)
//		}
//
//		fmt.Printf("%s\n", b)
//	}
//}

// Fetch given url and count times
func fetch() {
	start := time.Now()
	ch := make(chan string)
	for _, url := range os.Args[2:] {
		go process_url(url, ch)
	}

	for range os.Args[2:] {
		fmt.Println(<-ch)
	}

	fmt.Printf("%.2fs elapsed\n", time.Since(start).Seconds())
}

func process_url(url string, ch chan<- string) {
	start := time.Now()

	resp, err := http.Get(url)
	if err != nil {
		ch <- fmt.Sprintf("fetch: %v\n", err)
		return
	}
	nbytes, err := io.Copy(ioutil.Discard, resp.Body)
	resp.Body.Close()
	if err != nil {
		ch <- fmt.Sprintf("read body: %v\n", err)
		return
	}
	secs := time.Since(start).Seconds()
	ch <- fmt.Sprintf("%.2fs %7d %s", secs, nbytes, url)
}
