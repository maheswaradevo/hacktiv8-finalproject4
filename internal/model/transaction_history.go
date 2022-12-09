package model

import "time"

type TransactionHistory struct {
	TransactionID uint64    `db:"id"`
	ProductID     uint64    `db:"product_id"`
	UserID        uint64    `db:"user_id"`
	Quantity      uint64    `db:"quantity"`
	TotalPrice    uint64    `db:"total_price"`
	CreatedAt     time.Time `db:"created_at"`
	UpdatedAt     time.Time `db:"updated_at"`
}

type TransactionProductJoined struct {
	Product
	TransactionHistory
}
