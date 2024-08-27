package csp

import (
	"fmt"
	"log"
	"net/http"
	"time"
	"context"
)

type result struct {
	url string
	err error
	latency time.Duration
	status_code int
}

func get(ctx context.Context, url string, ch chan <- result) {
	start := time.Now()
	req, _ := http.NewRequestWithContext(ctx, http.MethodGet, url, nil)

	if response, err := http.DefaultClient.Do(req); err != nil {
		ch <- result{url, err, 0, 500}
	} else {
		t := time.Since(start).Round(time.Millisecond)
		ch <- result{url, nil, t, response.StatusCode}
		response.Body.Close()
	}
}
func CSP() {
	results := make(chan result)
	list := []string{
		"https://amazon.com",
		"https://google.com",
		"https://lokalise.com",
		"https://wsj.com",
		"https://abra_google_lok.com",
		"http://localhost:8080",

	}
	ctx, cancel := context.WithTimeout(context.Background(), 3* time.Second)

	defer cancel()

	for _, url := range list {
		go get(ctx, url, results)
	}

	for range list {
		r := <- results

		if r.err != nil {
			log.Printf("%-20s %s http code %d\n", r.url, r.err, r.status_code)
		} else {
			log.Printf("%-20s %s http code %d\n", r.url, r.latency, r.status_code)
		}
	}
}









type nextCh chan int

func (ch nextCh) handler(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintf(w, "<h1>You got %d<h1>", <- ch)
	
}

func counter(ch chan <- int) {
	for i:= 0;; i++ {
		ch <- i
	}
}

func RaceResolver() {
	var nextID nextCh = make(chan int)

	go counter(nextID)

	http.HandleFunc("/", nextID.handler)
	log.Fatal(http.ListenAndServe(":8080", nil))
}


func generator(limit int, ch chan <- int) {
	for i :=2; i <limit; i++ {
		ch <- i
	}
	close(ch)
}

func filter(src <- chan int, dst chan <- int, prime int) {
	 for i := range src {
		if i % prime != 0 {
			dst <- i
		}

	 }
	 close(dst)
}


func sieve(limit int) {
	ch := make(chan int)

	go generator(limit, ch)

	for  {
		prime, ok := <-ch

		if !ok {
			break
		}

		ch1 := make(chan int)
		go filter(ch,ch1, prime)

		ch = ch1

		fmt.Print(prime, " ")
	}
	fmt.Println()
}
func PrimeFilter() {
	sieve(300)
}