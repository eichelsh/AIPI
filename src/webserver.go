package main

import (
    "fmt"
    "io/ioutil"
    "html"
    "log"
    "net/http"
)

func main() {
    http.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {

      if html.EscapeString(r.URL.Path)=="/welcome" {
        resp,err:=http.Get("http://cnn.nl/")
        if err != nil {
	// handle error
  // space
}
    defer resp.Body.Close()
    body, err := ioutil.ReadAll(resp.Body)
    fmt.Fprintf(w, string(body[:]))
      }

      if html.EscapeString(r.URL.Path)=="/test" {
        fmt.Fprintf(w, "Welcome to the test Application")
      }

    })

    log.Fatal(http.ListenAndServe(":8082", nil))

}
