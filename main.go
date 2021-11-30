package main

import (
	"encoding/json"
	"fmt"
	"io"
	"net/http"
	"os"
	"sort"
	"strings"
)

func main() {
	http.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {
   	        w.Write([]byte("<h1>Hello World</h1> <br> <p>Cloud Project</p>"))

		fmt.Println("Hello world - the log message")
	})
	http.ListenAndServe(":8080", nil)
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
