package main

import (
	"log"
	"encoding/json"
	"fmt"
	"io"
	"net/http"
	"os"
	"sort"
	"strings"
)

func main() {
	fs := http.FileServer(http.Dir("./static"))
	http.Handle("/", fs)

	log.Println("Listening on :3000...")
	err := http.ListenAndServe(":3000", nil)
	if err != nil {
		log.Fatal(err)
	}
}
func main() {
	fs := http.FileServer(http.Dir("./index.html"))
	http.Handle("/", fs)

	log.Println("Listening on :8080...")
	http.ListenAndServe(":8080", nil)

	// staticHandler := http.FileServer(http.Dir("./assets"))
	// mux.Handle("/assets/", http.StripPrefix("/assets/", staticHandler))
	// http.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {
  //  	        w.Write([]byte("<h1>Hello World</h1> <br> <p>Cloud Project</p>"))
	//
	// 	fmt.Println("Hello world - the log message")
	// })
	// http.ListenAndServe(":8080", nil)
}

type response struct {
	Message string   `json:"message"`
	EnvVars []string `json:"env"`
	Fib     []int    `json:"fib"`
}

func fib() func() int {
	a, b := 0, 1
	return func() int {
		a, b = b, a+b
		return a
	}
}
