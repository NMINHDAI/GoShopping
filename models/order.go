package models

import "time"

type Order struct {
	Id       uint      `json:"id"`
	UserId   uint      `json:"userId"`
	Status   string    `json:"status"`
	CreateAt time.Time `json:"createAt"`
}
