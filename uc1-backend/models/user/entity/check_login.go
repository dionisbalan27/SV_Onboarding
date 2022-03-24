package entity

type CheckLogin struct {
	// gorm.Model
	ID       string `json:"id" gorm:"primaryKey, type:varchar(50)"`
	Name     string `json:"-" binding:"required"`
	Password string `json:"-" binding:"required"`
	RoleID   string `json:"roleId"  binding:"required"`
}
