package model

type UserNameFilter struct {
	Name string `json:"name" form:"name"`
}

type CVFilter struct {
	Name string `json:"name" form:"name"`
}
