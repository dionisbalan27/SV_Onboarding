package entity

type CheckLogin struct {
	// gorm.Model
	ID       string `json:"id" gorm:"primaryKey, type:varchar(50)"`
	Password string `json:"-" binding:"required"`
}
