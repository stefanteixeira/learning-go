package main

import (
  "fmt"
  "log"
  "flag"
  "net/http"
  "strings"
  "encoding/json"
  "github.com/stefanteixeira/learning-go/shortener/url"
)

var (
  port *int
  logging *bool
  baseUrl string
)

func init() {
  port = flag.Int("p", 8888, "port")
  logging = flag.Bool("l", true, "logging")

  flag.Parse()

  baseUrl = fmt.Sprintf("http://localhost:%d", *port)
}

type Headers map[string]string

type Redirect struct{
  stats chan string
}

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
  sendResponseWith(w, status, Headers{
    "Location": shortUrl,
    "Link": fmt.Sprintf("<%s/api/stats/%s>; rel=\"stats\"", baseUrl, url.Id),
  })

  logEvent("URL %s shortened to %s", url.OriginalUrl, shortUrl)
}

func (red *Redirect) ServeHTTP(w http.ResponseWriter, r *http.Request) {
  findUrlAndExecute(w, r, func(url *url.Url) {
    http.Redirect(w, r, url.OriginalUrl, http.StatusMovedPermanently)
    red.stats <- url.Id
  })
}

func GenerateStats(w http.ResponseWriter, r *http.Request) {
  findUrlAndExecute(w, r, func(url *url.Url) {
    json, err := json.Marshal(url.Stats())

    if err != nil {
      w.WriteHeader(http.StatusInternalServerError)
      return
    }

    sendJSONResponse(w, string(json))
  })
}

func findUrlAndExecute(w http.ResponseWriter, r *http.Request, executor func(*url.Url)) {
  path := strings.Split(r.URL.Path, "/")
  id := path[len(path)-1]

  if url := url.Find(id); url != nil {
    executor(url)
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

func sendJSONResponse(w http.ResponseWriter, response string) {
  sendResponseWith(w, http.StatusOK, Headers{
    "Content-Type": "application/json",
  })
  fmt.Fprintf(w, response)
}

func extractUrl(r *http.Request) string {
  rawBody := make([]byte, r.ContentLength, r.ContentLength)
  r.Body.Read(rawBody)
  return string(rawBody)
}

func main() {
  stats := make(chan string)
  defer close(stats)
  go registerStats(stats)

  url.ConfigRepository(url.NewInMemoryRepository())

  http.HandleFunc("/api/shorten", Shorten)
  http.Handle("/r/", &Redirect{stats})
  http.HandleFunc("/api/stats/", GenerateStats)

  logEvent("Server initializing on port %d", *port)
  log.Fatal(http.ListenAndServe(fmt.Sprintf(":%d", *port), nil))
}

func registerStats(ids <-chan string) {
  for id := range ids {
    url.RegisterClick(id)
    logEvent("Click registered for id %s", id)
  }
}

func logEvent(format string, values ...interface{}) {
  if *logging {
    log.Printf(fmt.Sprintf("%s\n", format), values...)
  }
}
