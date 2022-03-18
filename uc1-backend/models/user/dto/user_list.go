package dto

type UserList struct {
	Id         string `json:"id"`
	Name  	   string    `json:"name" binding:"required"`
	Role  	   Role    	 `json:"role"  binding:"required"`
	Active     bool      `json:"active" binding:"required"`
}