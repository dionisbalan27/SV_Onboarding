package dto

type Login struct {
	Personal_number string `json:"personalNumber"  binding:"required"`
	Password        string `json:"password" binding:"required"`
}
