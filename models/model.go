package models

type User struct {
	ID       int64  `json:"id"`
	Nama     string `json:"nama"`
	Email    string `json:"email"`
	Alamat   string `json:"alamat"`
	Password string `json:"password"`
}
