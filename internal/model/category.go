package model

import "time"

type Category struct {
	CategoryID        uint64    `db:"id"`
	Type              string    `db:"type"`
	SoldProductAmount uint64    `db:"sold_product_amount"`
	CreatedAt         time.Time `db:"created_at"`
	UpdatedAt         time.Time `db:"updated_at"`
}
