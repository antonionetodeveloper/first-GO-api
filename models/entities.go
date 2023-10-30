package models

type User struct {
	ID     int64  `json:"id"`
	Name   string `json:"name"`
	Age    int    `json:"age"`
	Gender bool   `json:"deleted"`
}
