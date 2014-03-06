package models

import (
  "time"
)

type Post struct {
  Id             int64  `db:"Id"`
  Content        string `form:"content"`
  Title          string `form:"title"`

  CreatedAt      time.Time
  PublishedAt    time.Time
  UpdatedAt      time.Time
}