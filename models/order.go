package models

import "time"

type Order struct {
	Id       uint      `json:"id"`
	UserId   uint      `json:"userId"`
	CreateAt time.Time `json:"createAt"`
}
