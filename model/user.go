package model

import "time"

type User struct {
	ID        int       `xorm:"id pk autoincr"`
	Name      string    `xorm:"name"`
	Age       int       `xorm:"age"`
	CreatedAt time.Time `xorm:"created_at"`
	UpdatedAt time.Time `xorm:"updated_at"`
}

type UserInput struct {
	ID   int    `json:"id" binding:"required"`
	Name string `json:"name" binding:"required"`
	Age  int    `json:"age" binding:"required"`
}
