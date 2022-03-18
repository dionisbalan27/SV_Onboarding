package entity

type UserID struct {
	// gorm.Model
	ID              string `json:"id" gorm:"primaryKey, type:varchar(50)"`
	Personal_number string `json:"personalNumber" binding:"required"`
	Name            string `json:"name" binding:"required"`
	Email           string `json:"email" binding:"required"`
	Active          bool   `json:"active" binding:"required"`
	RoleID          string `json:"role"  binding:"required"`
	Title           string `json:"title" binding:"required"`
	// Role            Role   `gorm:"foreignKey:RoleID"`
}
