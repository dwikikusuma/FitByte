package models

type User struct {
	ID       uint   `json:"id" gorm:"primaryKey;autoIncrement"`
	Email    string `json:"email" validate:"required,email"`
	Password string `json:"password"`
}

type RegisterResponse struct {
	Token string `json:"token"`
	Email string `json:"email"`
}
