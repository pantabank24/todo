package models

import "time"

type TodoReq struct {
	Title       string `json:"title"`
	Description string `json:"description"`
	Image       string `json:"image"`
}

type TodoRes struct {
	ID          string    `json:"id" db:"id"`
	Title       string    `json:"title" db:"title"`
	Description string    `json:"description" db:"description"`
	CreatedAt   time.Time `json:"created_at" db:"created_at"`
	UpdatedAt   time.Time `json:"updated_at" db:"updated_at"`
	Image       string    `json:"image" db:"image"`
	Status      string    `json:"status" db:"status"`
}

type TodosFilter struct {
	Order   string `json:"order" query:"order"`
	OrderBy string `json:"order_by" query:"orderBy"`
	Search  string `json:"search" query:"search"`
}
