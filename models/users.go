package models

type Users struct {
	Id        uint   `json:"id"`
	Name      string `json:"name" `
	Email     string `json:"email"`
	CreatedAt string `json:"created_at"`
}
