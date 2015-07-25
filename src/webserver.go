package main

import (
    "fmt"
    "html"
    "log"
    "net/http"
)

func main() {
    http.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {

      if html.EscapeString(r.URL.Path)=="/welcome" {
        fmt.Fprintf(w, "<b>Welcome to the welcome part</b>")
      }

      if html.EscapeString(r.URL.Path)=="/test" {
        fmt.Fprintf(w, "Welcome to the test Application")
      }

    })

    log.Fatal(http.ListenAndServe(":8082", nil))

}
