package models

import "time"

type Posts struct {
	ID         uint64    `json:"id_posts,omitempty"`
	Title      string    `json:"title,omitempty"`
	Content    string    `json:"content,omitempty"`
	AuthorID   uint64    `json:"author_id,omitempty"`
	AuthorNick string    `json:"author_nick,omitempty"`
	Like       uint64    `json:"like"`
	CreatedAt  time.Time `json:"createdAt,omitempty"`
}
