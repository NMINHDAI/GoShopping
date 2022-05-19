package models

type Shop struct {
	ProductId   uint   `json:"id"`
	Quantity    int32  `json:"quantity"`
	OwnBy       uint   `json:"ownby"`
	PaymentInfo string `json:"paymentInfo" default:"1"`
}
