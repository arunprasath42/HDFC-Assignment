// Package service - contains all the business logic for the qed request handler
package service

import (
	"encoding/json"
	"fmt"
	"hdfc-assignment/src/models"
	"hdfc-assignment/utils/constant"
	"strconv"
	"sync"
	"time"

	log "github.com/sirupsen/logrus"
)

// Handler - struct for qed request handler
type Handler struct{}

var (
	productMap = make(map[int]string)
	orderMap   = make(map[int]string)
	mutex      = &sync.Mutex{}
)

// Add 10 different products to the map.
func init() {

	// Create a product map to store the product data.
	productMap = make(map[int]string)

	// product data to the map.
	productMap[1] = `{"id":1,"name":"Product 1","price":9.99,"availability":true,"category":"Premium","quantity":10}`
	productMap[2] = `{"id":2,"name":"Product 2","price":19.99,"availability":false,"category":"Premium","quantity":10}`
	productMap[3] = `{"id":3,"name":"Product 3","price":4.99,"availability":true,"category":"Budget","quantity":10}`
	productMap[4] = `{"id":4,"name":"Product 4","price":29.99,"availability":true,"category":"Premium","quantity":10}`
	productMap[5] = `{"id":5,"name":"Product 5","price":14.99,"availability":false,"category":"Regular","quantity":10}`
	productMap[6] = `{"id":6,"name":"Product 6","price":7.99,"availability":true,"category":"Budget","quantity":10}`
	productMap[7] = `{"id":7,"name":"Product 7","price":39.99,"availability":true,"category":"Premium","quantity":10}`
	productMap[8] = `{"id":8,"name":"Product 8","price":24.99,"availability":true,"category":"Regular","quantity":10}`
	productMap[9] = `{"id":9,"name":"Product 9","price":8.99,"availability":true,"category":"Budget","quantity":10}`
	productMap[10] = `{"id":10,"name":"Product 10","price":49.99,"availability":false,"category":"Premium","quantity":10}`

	// sync map to store the product data.
	var productSyncMap sync.Map

	// Add the product data to the sync map.
	for key, value := range productMap {
		productSyncMap.Store(key, value)
	}

	// Print the data from the sync map.
	productSyncMap.Range(func(key, value interface{}) bool {
		log.Println("Key:", key, "Value:", value)
		return true
	})
}

// GetAllProducts - returns all the products in the catalogue - using sync.Map
func (h *Handler) GetAllProducts() []models.Product {

	//get the product data from the sync map.
	var productSyncMap sync.Map
	productSyncMap.Range(func(key, value interface{}) bool {
		log.Println("Key:", key, "Value:", value)
		return true
	})

	// Create a slice of products.
	var products []models.Product

	// Iterate through the product map and add the product data to the slice.
	for _, value := range productMap {
		var product models.Product
		json.Unmarshal([]byte(value), &product)
		products = append(products, product)
	}

	var sortedProducts []models.Product
	for i := 1; i <= len(products); i++ {
		for _, product := range products {
			if product.ID == i {
				sortedProducts = append(sortedProducts, product)
			}
		}
	}

	// Return the slice of products.
	return sortedProducts
}

// GetProductDetails - returns the details of a product
func (h *Handler) GetProductDetails(id int) (*models.Product, error) {

	// Check if the product exists in the map.
	p, ok := productMap[id]
	log.Println("p: ", p)
	if !ok {
		return nil, fmt.Errorf("product not found")
	}

	// Unmarshal the product data into a Go struct and return it as JSON.
	var product models.Product
	err := json.Unmarshal([]byte(p), &product)
	if err != nil {
		log.Println("Error unmarshalling product data: ", err)
		return nil, err
	}

	return &product, nil
}

// PlaceOrder - places an order for a product and returns the order details
func (h *Handler) PlaceOrder(req models.OrderRequest) (*models.Order, error) {
	var orderValue float64
	var premiumCount int

	for i, id := range req.ProductID {
		// Check if the product exists in the map.
		p, ok := productMap[id]
		log.Println("p: ", p)
		if !ok {
			return nil, fmt.Errorf("product not found")
		}

		//confirm the product is available and has the required quantity in stock before placing the order
		var product models.Product
		err := json.Unmarshal([]byte(p), &product)
		if err != nil {
			log.Println("Error unmarshalling product data: ", err)
			return nil, err
		}

		if product.Quantity < req.Quantity[i] {
			return nil, fmt.Errorf("product quantity not available")
		}

		if req.Category[i] == constant.PREMIUM {
			premiumCount++
		}

		orderValue += product.Price * float64(req.Quantity[i])
	}

	log.Println("orderValue before discount: ", orderValue)

	//check if the order is eligible for a discount
	if premiumCount >= 3 {
		log.Println("Order is eligible for a discount" + strconv.Itoa(premiumCount))
		orderValue *= 0.9 // apply 10% discount
	}

	log.Println("orderValue after discount: ", orderValue)
	//create a new order for all the products in the request
	order := models.Order{
		ID:            len(orderMap) + 1,
		ProductID:     req.ProductID,
		Quantity:      req.Quantity,
		OrderValue:    orderValue,
		OrderStatus:   "Placed",
		OrderDateTime: time.Now(),
	}

	//add the order to the map of orders

	orderJSON, err := json.Marshal(order)
	if err != nil {
		log.Println("Error marshalling order data: ", err)
		return nil, err
	}

	orderMap[order.ID] = string(orderJSON)

	return &order, nil
}

// GetAllOrders - returns all orders
func (h *Handler) GetAllOrders() ([]models.Order, error) {

	// Create a slice of orders.
	var orders []models.Order

	log.Println("orderMap: ", orderMap)
	// Iterate over the map and unmarshal the data into a Go struct.
	for _, o := range orderMap {
		var order models.Order
		err := json.Unmarshal([]byte(o), &order)
		if err != nil {
			log.Println("Error unmarshalling order data: ", err)
			return nil, err
		}

		// Add the order to the slice.
		orders = append(orders, order)
	}

	log.Println("orders: ", orders)
	return orders, nil
}

// GetOrderDetails - returns the details of an specific order
func (h *Handler) GetOrderDetails(id int) (*models.Order, error) {

	// Check if the order exists in the map.
	o, ok := orderMap[id]
	if !ok {
		return nil, fmt.Errorf("order not found")
	}

	// Unmarshal the order data into a Go struct and return it as JSON.
	var order models.Order
	err := json.Unmarshal([]byte(o), &order)
	if err != nil {
		log.Println("Error unmarshalling order data: ", err)
		return nil, err
	}

	return &order, nil
}

// UpdateOrderStatus - updates the order status
func (h *Handler) UpdateOrderStatus(req models.OrderUpdateRequest) (*models.Order, error) {

	// Check if the order exists in the map.
	o, ok := orderMap[req.OrderID]
	if !ok {
		return nil, fmt.Errorf("order not found")
	}

	// Unmarshal the order data.
	var order models.Order
	err := json.Unmarshal([]byte(o), &order)
	if err != nil {
		log.Println("Error unmarshalling order data: ", err)
		return nil, err
	}

	// Update the order status.
	order.OrderStatus = req.OrderStatus

	// Update the dispatch date.
	if req.OrderStatus == constant.DISPATCHED {
		tym := time.Now()
		order.DispatchDate = &tym
	}
	// Marshal the order data into JSON.
	b, err := json.Marshal(order)
	if err != nil {
		log.Println("Error marshalling order data: ", err)
		return nil, err
	}

	// Update the order in the map.
	mutex.Lock()
	orderMap[order.ID] = string(b)
	mutex.Unlock()

	return &order, nil
}
