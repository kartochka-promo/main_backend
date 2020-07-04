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
	Username string `json:"user"`
	Password string `json:"password"`
}
