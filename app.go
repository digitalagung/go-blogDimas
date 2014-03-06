package main

import (
  "PD.GoBlog/controllers"
  "PD.GoBlog/models"
  "github.com/codegangsta/martini"
  "github.com/codegangsta/martini-contrib/binding"
  "github.com/codegangsta/martini-contrib/render"
  "log"
  "net/http"
)

func main() {
  m := martini.Classic()

  postsController := new(controllers.PostsController)

  m.Use(render.Renderer(render.Options{
    Directory: "templates",
    Layout:    "layout",
  }))
  m.Use(PopulateAppContext)
  m.Use(martini.Static("public"))

  m.Get("/",                                        postsController.Index)
  m.Get("/posts/new",                               postsController.New)
  m.Post("/posts/new", binding.Form(models.Post{}), postsController.Create)

  log.Fatal(http.ListenAndServe("192.168.1.60:8080", m))
}

func PopulateAppContext(martiniContext martini.Context, w http.ResponseWriter, request *http.Request, renderer render.Render) {
  dbContext := models.NewDbContext()
  appContext := &models.AppContext{Request: request, Renderer: renderer, MartiniContext: martiniContext, DbContext: dbContext}

  martiniContext.Map(appContext)
}