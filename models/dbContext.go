package models

import (
  "database/sql"
  "github.com/coopernurse/gorp"
  _ "github.com/lib/pq"
  "log"
)

type DbContext struct {
  dbmap *gorp.DbMap
}

func NewDbContext() DbContext {
  db, err := sql.Open("postgres", "user=goblog_dev password=password dbname=goblog_dev")
  if err != nil { log.Fatal(err) }

  dbmap := &gorp.DbMap{Db: db, Dialect: gorp.SqliteDialect{}}
  dbmap.AddTableWithName(Post{}, "posts").SetKeys(true, "Id")

  return DbContext{dbmap: dbmap}
}

func (self DbContext) FetchPosts() []Post{
  var posts []Post
  _, err := self.dbmap.Select(&posts, "SELECT * FROM POSTS ORDER BY CreatedAt DESC")
  if err != nil { panic(err) }

  return posts
}