package model

import "time"

type Product struct {
	ProductID  uint64    `db:"id"`
	Title      string    `db:"title"`
	Price      uint64    `db:"price"`
	Stock      uint64    `db:"stock"`
	CategoryID uint64    `db:"category_id"`
	CreatedAt  time.Time `db:"created_at"`
	UpdatedAt  time.Time `db:"updated_at"`
}

type ProductCategoryJoined struct {
	Product    Product
	Categories Categories
}

type ProductJoined []*ProductCategoryJoined

type Products []*Product
