package models

import "time"

type GetCafe struct {
	CafeID      int       `json:"id"`
	CafeName    string    `json:"name"`
	Address     string    `json:"address"`
	Description string    `json:"description"`
	StaffID     int       `json:"staffID"`
	OpenTime    time.Time `json:"openTime"`
	CloseTime   time.Time `json:"closeTime"`
	Photo       string    `json:"photo"`
	Location    string    `json:"location"`
}

type GetCafes struct {
	Cafes []GetCafe `json:"cafes"`
}

type CreateOrUpdateCafe struct {
	CafeID      int       `json:"id"`
	CafeName    string    `json:"name"`
	Address     string    `json:"address"`
	Description string    `json:"description"`
	StaffID     int       `json:"staffID"`
	OpenTime    time.Time `json:"openTime"`
	CloseTime   time.Time `json:"closeTime"`
	Photo       string    `json:"photo"`
	Location    string    `json:"location"`
}
