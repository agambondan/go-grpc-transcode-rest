package model

import "github.com/google/uuid"

type Like struct {
	BaseInt
	Like      *int64     `json:"like,omitempty" gorm:"default:1"`
	UserID    *uuid.UUID `json:"userID,omitempty"`
	User      *User      `json:"user,omitempty"`
	ArticleID *int64     `json:"articleID,omitempty"`
	Article   *Article   `json:"article,omitempty"`
	//CommentID *int64     `json:"commentID,omitempty"`
	//Comment   *Comment   `json:"comment,omitempty"`
}

/*
1 Article bisa punya banyak like
1 user bisa like banyak article
1 user bisa like banyak comment

1 like hanya punya 1 user
1 like hanya punya 1 article
1 like hanya punya 1 comment
*/
