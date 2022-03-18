package entity

import "time"

type Product struct {
	ID          string    `json:"id" gorm:"primaryKey, type:varchar(255)"`
	Name        string    `json:"name" binding:"required"`
	Description string    `json:"description" binding:"required"`
	Status      string    `json:"status" binding:"required"`
	MakerID     string    `json:"makerId" binding:"required"`
	SignerID    string    `json:"signerId"`
	CheckerID   string    `json:"checkerId"`
	CreatedAt   time.Time `json:"createdAt"`
	UpdatedAt   time.Time `json:"updatedAt"`
	DeletedAt   time.Time `json:"deletedAt"`
}
