package main

import (
	"fmt"
	"net/http"
)

func main() {
	http.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {
		fmt.Fprint(w, "<h1>asdasd111231d23bbbbbb</h1>")
	})
	fmt.Println("Starting server at port 3000...")
	http.ListenAndServe(":3000", nil)
}
