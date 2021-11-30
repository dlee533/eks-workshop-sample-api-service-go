package main

import (
	"fmt"
	"net/http"
)

const IndexHTML = `
<h1>hello world! - updated</h1>
`

func main() {
	http.HandleFunc("/", Index)
	http.ListenAndServe(":8080", nil)
}

func Index(w http.ResponseWriter, r *http.Request) {
    url := r.FormValue("url")
    if url == "" {
        fmt.Fprint(w, IndexHTML)
        return
    }
    key := "Placeholder"
    fmt.Fprintf(w, "http://localhost:8080/%s", key)
}
