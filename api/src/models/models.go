// Models is used to define the request and response structs
package models

import (
	"time"
)

// Product represents a product in the product catalog.
type Product struct {
	ID           int     `json:"id"`
	Name         string  `json:"name"`
	Price        float64 `json:"price"`
	Availability bool    `json:"availability"`
	Category     string  `json:"category"`
	Quantity     int     `json:"quantity"`
}

// Order represents an order for a product.
type Order struct {
	ID            int        `json:"id"`
	ProductID     []int      `json:"productId"`
	Quantity      []int      `json:"quantity"`
	OrderValue    float64    `json:"orderValue"`
	OrderStatus   string     `json:"orderStatus" binding:"required,oneof=Placed Dispatched Completed Cancelled"`
	DispatchDate  *time.Time `json:"dispatchDate,omitempty"`
	PremiumCount  int        `json:"premiumCount,omitempty"`
	OrderDateTime time.Time  `json:"orderDateTime,omitempty"`
}

type Quantity struct {
	Quantity int `json:"quantity" binding:"required,min=1,max=10"`
}

type OrderRequest struct {
	ProductID    []int    `json:"id" binding:"required"`
	Quantity     []int    `json:"quantity" binding:"required,min=1,max=10"`
	PremiumCount int      `json:"premiumCount"`
}

type OrderUpdateRequest struct {
	OrderID     int    `json:"id" binding:"required"`
	OrderStatus string `json:"orderStatus" binding:"required,oneof=Placed Dispatched Completed Cancelled"`
}
