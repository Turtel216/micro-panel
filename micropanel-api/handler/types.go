package handler

import "time"

type ProductReq struct {
	ID           int64   `json:"id"`
	Name         string  `json:"name"`
	Image        string  `json:"image"`
	Category     string  `json:"category"`
	Description  string  `json:"description"`
	Rating       int64   `json:"rating"`
	NumReviews   int64   `json:"num_reviews"`
	Price        float32 `json:"price"`
	CountInStock int64   `json:"count_in_stock"`
}

type ProductRes struct {
	ID           int64      `json:"id"`
	Name         string     `json:"name"`
	Image        string     `json:"image"`
	Category     string     `json:"category"`
	Description  string     `json:"description"`
	Rating       int64      `json:"rating"`
	NumReviews   int64      `json:"num_reviews"`
	Price        float32    `json:"price"`
	CountInStock int64      `json:"count_in_stock"`
	CreatedAt    time.Time  `json:"created_at"`
	UpdatedAt    *time.Time `json:"updated_at"`
}

type OrderItem struct {
	Name      string  `json:"name"`
	Quantity  int64   `json:"quantity"`
	Image     string  `json:"image"`
	Price     float32 `json:"price"`
	ProductID int64   `json:"product_id"`
}

type OrderReq struct {
	ID            int64       `json:"id"`
	Items         []OrderItem `json:"items"`
	PaymentMethod string      `json:"payment_method"`
	TaxPrice      float32     `json:"tax_price"`
	ShippingPrice float32     `json:"shipping_price"`
	TotalPrice    float32     `json:"total_price"`
	Status        string      `json:"status"`
}

type OrderRes struct {
	ID            int64       `json:"id"`
	Items         []OrderItem `json:"items"`
	PaymentMethod string      `json:"payment_method"`
	TaxPrice      float32     `json:"tax_price"`
	ShippingPrice float32     `json:"shipping_price"`
	TotalPrice    float32     `json:"total_price"`
	Status        string      `json:"status"`
	CreatedAt     time.Time   `json:"created_at"`
	UpdatedAt     *time.Time  `json:"updated_at"`
}
