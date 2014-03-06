DROP TABLE posts;

CREATE TABLE posts (
  Id          serial primary key,
  Title       varchar(255),
  Content     text,
  CreatedAt   timestamp not null default now(),
  PublishedAt timestamp not null default now(),
  UpdatedAt   timestamp not null default now()
);