package entity

type User struct {
	// gorm.Model
	ID              string `json:"id" gorm:"primaryKey, type:varchar(50)"`
	Personal_number string `json:"personalNumber" binding:"required"`
	Name            string `json:"name" binding:"required"`
	Password        string `json:"-" binding:"required"`
	Email           string `json:"email" binding:"required"`
	Active          bool   `json:"active" binding:"required"`
	RoleID          string `json:"roleId" binding:"required"`
	Title           string `json:"title" gorm:"-"`
	Role            Role   `gorm:"foreignKey:RoleID;references:ID"`
}
