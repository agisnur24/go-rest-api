package domain

import "time"

type Product struct {
	Id         int
	Name       string
	Price      int
	CategoryId int
	CreatedAt  time.Time
	UpdatedAt  time.Time
}
