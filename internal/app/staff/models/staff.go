package models

import "time"

type Staff struct {
	StaffID  int       `json:"id"`
	Name     string    `json:"name" validate:"required,min=4,max=100"`
	Email    string    `json:"email" validate:"required,email"`
	Password string    `json:"password" validate:"required,min=8,max=100"`
	EditedAt time.Time `json:"editedAt" validate:"required"`
	Photo    string    `json:"photo"`
	IsOwner  bool      `json:"isOwner"`
	CafeId   int       `json:"CafeId"`
}

type SafeStaff struct {
	StaffID  int       `json:"id"`
	Name     string    `json:"name" validate:"required,min=4,max=100"`
	Email    string    `json:"email" validate:"required,email"`
	EditedAt time.Time `json:"editedAt" validate:"required"`
	Photo    string    `json:"photo"`
	IsOwner  bool      `json:"isOwner"`
	CafeId   int       `json:"CafeId"`
}
