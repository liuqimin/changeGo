package api

import (
	"changeGo/jsontime"
)

type AuthorAddPost struct {
	Name     string            `json:"name"`
	Country  string            `json:"country"`
	Birthday jsontime.JsonDate `json"birthdy"`
}

type AuthorDelPost struct {
	Id int `json:"id"`
}

type PressAddPost struct {
	Name string `json:"name"`
}
