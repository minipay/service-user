package models

type User struct {
	Id   int64  `json:"id_user"`
	Name string `json:"name" form:"name" query:"name" validate:"required"`
}