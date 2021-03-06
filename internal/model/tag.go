package model

import "github.com/bianxm/blog-service/pkg/app"

type Tag struct {
	*Model
	Name  string `json:"name"`
	State uint8  `json:"state"`
}

func (a Tag) TableName() string {
	return "blog_tag"
}

// tag.go
type TagSwagger struct {
	List  []*Tag
	Pager *app.Pager
}
