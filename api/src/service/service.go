// Package service - contains all the business logic for the qed request handler
package service

import (
	"errors"
	"fmt"
	"hdfc-assignment/src/models"
	"hdfc-assignment/utils/constant"
	"time"

	log "github.com/sirupsen/logrus"
)

// Handler - struct for the service
type Handler struct{}

var product []*models.Product
var orders []*models.Order

func init() {
	product = append(product, &models.Product{ID: 1, Name: "Product 1", Price: 100, Availability: true, Category: "Premium", Quantity: 12})
	product = append(product, &models.Product{ID: 2, Name: "Product 2", Price: 200, Availability: true, Category: "Regular", Quantity: 20})
	product = append(product, &models.Product{ID: 3, Name: "Product 3", Price: 300, Availability: true, Category: "Budget", Quantity: 15})
	product = append(product, &models.Product{ID: 4, Name: "Product 4", Price: 400, Availability: true, Category: "Premium", Quantity: 10})
	product = append(product, &models.Product{ID: 5, Name: "Product 5", Price: 500, Availability: true, Category: "Regular", Quantity: 10})
	product = append(product, &models.Product{ID: 6, Name: "Product 6", Price: 600, Availability: true, Category: "Budget", Quantity: 15})
	product = append(product, &models.Product{ID: 7, Name: "Product 7", Price: 700, Availability: true, Category: "Premium", Quantity: 18})
	product = append(product, &models.Product{ID: 8, Name: "Product 8", Price: 800, Availability: true, Category: "Regular", Quantity: 20})
	product = append(product, &models.Product{ID: 9, Name: "Product 9", Price: 900, Availability: true, Category: "Budget", Quantity: 10})
	product = append(product, &models.Product{ID: 10, Name: "Product 10", Price: 1000, Availability: true, Category: "Premium", Quantity: 13})
}

// GetAllProducts - returns all the products in the catalogue
func (h *Handler) GetAllProducts() []*models.Product {
	return product
}

// GetProductDetails - returns the details of a specific product
func (h *Handler) GetProductDetails(id int) (*models.Product, error) {
	var prod *models.Product
	for _, val := range product {
		if val.ID == id {
			prod = val
		}
	}
	if prod == nil {
		return nil, constant.ErrProductNotFound
	}
	return prod, nil
}

func (h *Handler) PlaceOrder(request []*models.OrderReq) (string, error) {
	var orderValue float64
	var premium int
	for _, req := range request {
		var found bool
		for _, prod := range product {
			if req.ProductId == prod.ID {
				found = true
				// Add additional validation to ensure quantity is valid
				if req.Quantity <= 0 {
					log.Printf("Invalid quantity: %d", req.Quantity)
					return "", errors.New("requested quantity should be greater than 0")
				}

				// Check if requested quantity is available
				if req.Quantity > prod.Quantity {
					log.Printf("Requested quantity exceeds available quantity: %d > %d", req.Quantity, prod.Quantity)
					return "", fmt.Errorf(constant.QUANTITYEXCEEDS, req.Quantity, prod.Quantity)
				}

				// Updating product quantity by decrementing the requested quantity
				prod.Quantity = prod.Quantity - req.Quantity

				// If product quantity is 0, then availability is set to false
				if prod.Quantity == 0 {
					prod.Availability = false
				}

				// If product is premium, incrementing the premium count
				if prod.Category == constant.PREMIUM {
					premium = premium + 1
				}
				orderValue = orderValue + (float64(req.Quantity) * prod.Price)
			}

		}
		if !found {
			return "", constant.ErrProductNotFound
		}

	}
	if premium >= 3 {
		log.Printf("Order is eligible for 10%% discount and the order value is: %f and the discount is: %f", orderValue*0.9, orderValue*0.1)
		orderValue *= 0.9 // apply 10% discount
	}

	// Placing the order and adding it to the orders slice
	var order models.Order
	order.ID = 1
	order.ProductDetails = request
	order.OrderStatus = constant.PLACED
	order.OrderDateTime = time.Now()
	order.OrderValue = orderValue
	order.PremiumCount = premium
	orders = append(orders, &order)

	return constant.ORDERPLACED, nil
}

// GetAllOrders - returns all the orders placed
func (h *Handler) GetAllOrders() ([]*models.Order, error) {
	return orders, nil
}

// GetOrderDetailsByID - returns the details of a specific order
func (h *Handler) GetOrderDetailsByID(id int) (*models.Order, error) {
	var ord *models.Order
	for _, val := range orders {
		if val.ID == id {
			ord = val
		}
	}
	if ord == nil {
		return nil, constant.ErrOrderNotFound
	}
	return ord, nil
}

// UpdateOrderStatus - updates the status of a specific order
func (h *Handler) UpdateOrderStatus(req models.OrderUpdateRequest) (string, error) {
	var ord *models.Order
	for _, val := range orders {
		if val.ID == req.OrderID {
			ord = val
		}
	}
	if ord == nil {
		return "", constant.ErrOrderNotFound
	}
	ord.OrderStatus = req.OrderStatus
	return constant.ORDERUPDATED, nil
}
