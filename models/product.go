package models

type Product struct {
	Id          uint    `json:"id"`
	Name        string  `json:"name" gorm:"unique"`
	Image       string  `json:"image"`
	Description string  `json:"description"`
	Price       float64 `json:"price" `
	QuantityInv int     `json:"quantityInv"`
}
