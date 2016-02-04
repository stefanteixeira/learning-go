package url

import (
  "math/rand"
  "net/url"
  "time"
)

const (
  size = 5
  symbols = "abcdefghijklmnopqrstuvxwyzABCDEFGHIJKLMNOPQRSTUVXWYZ1234567890_-+"
)

type Repository interface {
  IdExists(id string) bool
  FindById(id string) *Url
  FindByUrl(url string) *Url
  Save(url Url) error
  RegisterClick(id string)
  FindClicks(id string) int
}

var repo Repository

func init() {
  rand.Seed(time.Now().UnixNano())
}

func ConfigRepository(r Repository) {
  repo = r
}

type Url struct {
  Id string `json:"id"`
  CreationTime time.Time `json:"creation"`
  OriginalUrl string `json:"origUrl"`
}

type Stats struct {
  Url *Url `json:"url"`
  Clicks int `json:"clicks"`
}

func Find(id string) *Url {
  return repo.FindById(id)
}

func RegisterClick(id string) {
  repo.RegisterClick(id)
}

func FindOrCreateUrl(origUrl string) (u *Url, new bool, err error) {
  if u = repo.FindByUrl(origUrl); u != nil {
    return u, false, nil
  }

  if _, err = url.ParseRequestURI(origUrl); err != nil {
    return nil, false, err
  }

  url := Url{generateId(), time.Now(), origUrl}
  repo.Save(url)
  return &url, true, nil
}

func (u *Url) Stats() *Stats {
  clicks := repo.FindClicks(u.Id)
  return &Stats{u, clicks}
}

func generateId() string {
  newId := func() string {
    id := make([]byte, size, size)
    for i := range id {
      id[i] = symbols[rand.Intn(len(symbols))]
    }
    return string(id)
  }

  for {
    if id := newId(); !repo.IdExists(id) {
      return id
    }
  }
}
