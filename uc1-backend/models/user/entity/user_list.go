package entity

type UserList struct {
	ID     string `json:"id"`
	Name   string `json:"name" binding:"required"`
	RoleID string `json:"roleId"  binding:"required"`
	Title  string `json:"title" binding:"required"`
	Active bool   `json:"active" binding:"required"`
}