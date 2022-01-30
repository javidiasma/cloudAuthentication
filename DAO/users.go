package DAO

type Users struct {
	ID       int    `json:"id" gorm:"primary_key"`
	Username string `json:"username"`
	Password string `json:"password"`
}
