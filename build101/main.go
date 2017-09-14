package main

import (
	"fmt"
	"net/http"
)

func main() {
	http.HandleFunc("/", func(resp http.ResponseWriter, req *http.Request) {
		defer req.Body.Close()
		resp.Write([]byte("Hello from Golang!"))
	})
	fmt.Println("Listening on port 80")
	http.ListenAndServe(":80", http.DefaultServeMux)
}
