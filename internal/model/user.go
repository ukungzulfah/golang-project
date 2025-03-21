package model

type User struct {
	ID           int    `json:"id"`
	FullName     string `json:"user_fullname"`
	UserName     string `json:"user_name"`
	UserPassword string `json:"user_password"`
	UserMenu     int    `json:"user_menu"`
}
