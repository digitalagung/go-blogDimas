package controllers

import (
  "PD.GoBlog/models"
)

type PostsController struct {
}

func (controller *PostsController) Index(appContext *models.AppContext) {
  posts := appContext.DbContext.FetchPosts()
  appContext.Renderer.HTML(200, "posts/index", posts)
}