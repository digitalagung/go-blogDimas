package controllers

import (
  "github.com/phantomdata/go-blog/models"
)

type PostsController struct {
}

func (controller *PostsController) Index(appContext *models.AppContext) {
  posts := appContext.DbContext.FetchPosts()
  appContext.Renderer.HTML(200, "posts/index", posts)
}

func (controller *PostsController) New(appContext *models.AppContext) {
  post := models.Post{}
  appContext.Renderer.HTML(200, "posts/new", post)
}

func (controller *PostsController) Create(appContext *models.AppContext, post models.Post) {
  appContext.DbContext.InsertPost(post)

  appContext.Renderer.Redirect("/")
}