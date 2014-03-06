package models

import (
  "time"
)

type Post struct {
  Id             int64 `db:"Id"`
  Content        string
  Title          string

  CreatedAt      time.Time
  PublishedAt    time.Time
  UpdatedAt      time.Time
}