package main

import (
	"fmt"
	"net/http"
)

func main() {
	google := "https://www.google.com"
	yahoo := "https://www.yahoo.com"

	resp, err := http.Get(google)
	if err != nil {
		panic(err)
	}
	defer close(resp, "google")

	if resp.StatusCode < http.StatusBadRequest {

		fmt.Printf("google was good, trying a second request to yahoo ...")

		resp, err = http.Get(yahoo)
		if err != nil {
			panic(err)
		}
		defer close(resp, "yahoo")
	}
}

func close(w *http.Response, s string) {
	fmt.Println("")
	fmt.Printf("%p\n", w)
	fmt.Println(s)
	w.Body.Close()
	fmt.Println("")
}
