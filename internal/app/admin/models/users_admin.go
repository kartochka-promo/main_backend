package models

type GetAdmin struct {
	AdminId  int    `json:"id"`
	Username string `json:"user"`
	Password string `json:"password"`
}

type GetAdmins struct {
	Users []GetAdmin `json:"users"`
}

type DeleteAdmin struct {
	AdminId int `json:"id"`
}

type CreateOrUpdateAdmin struct {
	Username       string `json:"user" db:"username"`
	Password       string `json:"password"`
	HashedPassword string `db:"password"`
}

type LogAdmin struct {
	Username       string `json:"user" db:"username"`
	Password       string `json:"password"`
}