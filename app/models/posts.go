package models

import "time"

type Posts struct {
	ID         uint64    `json:"user_id,omitempty"`
	Title      string    `json:"title,omitempty"`
	Contet     string    `json:"content,omitempty"`
	AuthorID   uint64    `json:"author_id,omitempty"`
	AuthorNick string    `json:"autor_nick,omitempty"`
	Like       uint64    `json:"like"`
	CreatedAt  time.Time `json:"createdAt,omitempty"`
}
