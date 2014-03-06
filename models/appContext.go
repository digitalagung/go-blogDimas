package models

import (
  "github.com/codegangsta/martini"
  "github.com/codegangsta/martini-contrib/render"
  "net/http"
)

type AppContext struct {
  DbContext      DbContext
  Request        *http.Request
  Renderer       render.Render
  MartiniContext martini.Context
}