package test_server

import (
	"encoding/json"
	"fmt"
	"io"
	"log"
	"net/http"
	"os"
	"html/template"
)

const url = "https://jsonplaceholder.typicode.com/"


type todo struct {
	ID int `json:"id"`
	UserID int `json:"userId"`
	Title string `json:"title"`
	Completed bool `json:"completed"`
}

var form = `
<h1>Todo #{{.ID}}</h1>
<div>{{printf "User %d"  .UserID}}</div>
<div>{{printf "%s (completed: %t)" .Title .Completed}}</div>
`

func handler(w http.ResponseWriter, r *http.Request) {
	response, err := http.Get(url + r.URL.Path[1:])

	if err != nil {
		http.Error(w, err.Error(), http.StatusServiceUnavailable)
		return
	}

	defer response.Body.Close()

	if response.StatusCode != http.StatusOK {
		http.Error(w, fmt.Sprintf("received non-200 response: %d", response.StatusCode), http.StatusServiceUnavailable)
		return
	}

	body, err := io.ReadAll(response.Body)
	if err != nil {
		http.Error(w, err.Error(), http.StatusServiceUnavailable)
		return
	}

	fmt.Println("Raw response body:", string(body)) 



	var item todo
	if err = json.Unmarshal(body, &item); err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	tmpl := template.New("Kirill")
	tmpl, err = tmpl.Parse(form)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	err = tmpl.Execute(w, item)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}
}



func Up() {
	http.HandleFunc("/", handler)
	fmt.Println("Test server is up")
	log.Fatal((http.ListenAndServe(":8083", nil)))
}

func ClientUp() {
	response, err := http.Get(url + "/todos/1")

	if err != nil {
		fmt.Fprintln(os.Stderr, err)
		os.Exit(-1)
	}

	defer response.Body.Close()

	if response.StatusCode == http.StatusOK {
		body, err := io.ReadAll(response.Body)

		if err != nil {
			fmt.Fprintln(os.Stderr, err)
			os.Exit(-1)
		}

		var item todo

		err = json.Unmarshal(body, &item)

		if err != nil {
			fmt.Fprintln(os.Stderr, err)
			os.Exit(-1)
		}

		fmt.Printf("%#v\n", item)
		fmt.Printf("%v\n", item)

	}
}