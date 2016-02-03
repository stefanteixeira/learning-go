package main

import (
  "net/http"
  "testing"
  "github.com/stretchr/testify/assert"
  "github.com/emicklei/forest"
)

func TestAPI(t *testing.T) {
  go main()

  api := forest.NewClient("http://localhost:8888", new(http.Client))

  r := api.POST(t, forest.Path("/api/shorten").Body("https://github.com/stefanteixeira"))
  r2 := api.POST(t, forest.Path("/api/shorten").Body("https://github.com/stefanteixeira"))

  assert.Equal(t, r.StatusCode, 201)
  assert.Equal(t, r2.StatusCode, 200)
}
