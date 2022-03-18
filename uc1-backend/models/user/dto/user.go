package dto

type User struct {
	Id              string `json:"id"`
	Name            string `json:"name" binding:"required"`
	Personal_number string `json:"personalNumber"  binding:"required"`
	Email           string `json:"email" binding:"required"`
	Role            Role   `json:"role"`
	Active          bool   `json:"active"`
	Password        string `json:"password" binding:"required"`
}