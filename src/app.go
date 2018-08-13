package main

import (
  "fmt"
  "net/http"
)


const PORT = "8080"

func main() {
  fmt.Printf("Listening on %s", PORT)
  http.HandleFunc("/", handler)
  http.ListenAndServe(":"+PORT, nil)
}

func handler(w http.ResponseWriter, r *http.Request) {
  fmt.Fprintf(w, "Your url path is %s", r.URL.Path)
}
