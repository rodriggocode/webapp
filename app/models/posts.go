package models

import "time"

type Posts struct {
	ID         uint64    `json:"id,omitempty"`
	Title      string    `json:"title,omitempty"`
	Contet     string    `json:"content,omitempty"`
	AuthorID   uint64    `json:"autor_id,omitempty"`
	AuthorNick string    `json:"author_nick,omitempty"`
	Like       uint64    `json:"like"`
	CreatedAt  time.Time `json:"createAt,omitempty"`
}
