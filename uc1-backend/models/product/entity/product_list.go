package entity

type ProductList struct {
	ID          string `json:"id" gorm:"primaryKey, type:varchar(255)"`
	Name        string
	Description string
	Status      string
}
