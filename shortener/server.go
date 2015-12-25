package main

import (
  "fmt"
  "log"
  "net/http"
  "strings"
  "github.com/stefanteixeira/learning-go/shortener/url"
)

var (
  port int
  baseUrl string
)

func init() {
  port = 8888
  baseUrl = fmt.Sprintf("http://localhost:%d", port)
}

type Headers map[string]string

func Shorten(w http.ResponseWriter, r *http.Request) {
  if r.Method != "POST" {
    sendResponseWith(w, http.StatusMethodNotAllowed, Headers{"Allow": "POST"})
    return
  }

  url, new, err := url.FindOrCreateUrl(extractUrl(r))

  if err != nil {
    sendResponseWith(w, http.StatusBadRequest, nil)
    return
  }

  var status int
  if new {
    status = http.StatusCreated
  } else {
    status = http.StatusOK
  }

  shortUrl := fmt.Sprintf("%s/r/%s", baseUrl, url.Id)
  sendResponseWith(w, status, Headers{"Location": shortUrl})
}

func Redirect(w http.ResponseWriter, r *http.Request) {
  path := strings.Split(r.URL.Path, "/")
  id := path[len(path)-1]

  if url := url.Find(id); url != nil {
    http.Redirect(w, r, url.OriginalUrl, http.StatusMovedPermanently)
  } else {
    http.NotFound(w, r)
  }
}

func sendResponseWith(w http.ResponseWriter, status int, headers Headers) {
  for k, v := range headers {
    w.Header().Set(k, v)
  }
  w.WriteHeader(status)
}

func extractUrl(r *http.Request) string {
  rawBody := make([]byte, r.ContentLength, r.ContentLength)
  r.Body.Read(rawBody)
  return string(rawBody)
}

func main() {
  http.HandleFunc("/api/shorten", Shorten)
  http.HandleFunc("/r/", Redirect)

  log.Fatal(http.ListenAndServe(fmt.Sprintf(":%d", port), nil))
}
