package entity

type ProductDetailID struct {
	ID          string `json:"id"`
	Name        string `json:"name"`
	Description string `json:"description"`
	Status      string `json:"status"`
	Maker       Action `json:"maker"`
	Checker     Action `json:"checker"`
	Signer      Action `json:"signer"`
}
