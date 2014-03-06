package models

import (
  "database/sql"
  "github.com/coopernurse/gorp"
  _ "github.com/lib/pq"
  "log"
  "time"
)

type DbContext struct {
  dbmap *gorp.DbMap
}

func NewDbContext() DbContext {
  db, err := sql.Open("postgres", "user=goblog_dev password=password dbname=goblog_dev")
  if err != nil { log.Fatal(err) }

  dbmap := &gorp.DbMap{Db: db, Dialect: gorp.PostgresDialect{}}
  dbmap.AddTableWithName(Post{}, "posts").SetKeys(true, "Id")

  return DbContext{dbmap: dbmap}
}

func (self DbContext) FetchPosts() []Post{
  var posts []Post
  _, err := self.dbmap.Select(&posts, "SELECT * FROM POSTS ORDER BY CreatedAt DESC")
  if err != nil { panic(err) }

  return posts
}

func (self DbContext) InsertPost(post Post) {

  post.CreatedAt    = time.Now()
  post.PublishedAt  = time.Now()
  post.UpdatedAt    = time.Now()

  err := self.dbmap.Insert(&post)
  if err != nil { panic(err) }
}