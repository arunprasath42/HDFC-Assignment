// Package controllers - This package contains the controllers for the qed request
package controllers

import (
	"fmt"
	"net/http"
	"strconv"

	"hdfc-assignment/src/models"
	"hdfc-assignment/src/service"

	"hdfc-assignment/utils/validator"

	"github.com/gin-gonic/gin"
)

// GetAllProducts - returns all the products in the catalogue
func GetAllProducts(c *gin.Context) {

	var service = service.Handler{}

	// Get the product details from the catalogue
	catalogue := service.GetAllProducts()
	c.JSON(http.StatusOK, catalogue)
}

// GetProductDetails - returns the details of a specific product
func GetProductDetails(c *gin.Context) {

	var service = service.Handler{}
	id, err := strconv.Atoi(c.Query("id"))
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Invalid request body"})
		return
	}

	fmt.Println("req: ", id)
	// Get the product details from the catalogue
	product, err := service.GetProductDetails(id)
	if err != nil {
		c.JSON(http.StatusNotFound, gin.H{"error": "Product not found"})
		return
	}

	c.JSON(http.StatusOK, product)
}

// PlaceOrder - places an order for a product
func PlaceOrder(c *gin.Context) {

	//inorder to place an order we need to have the product id ,quantity and the category of the product

	req := models.OrderRequest{}

	if err := c.ShouldBindJSON(&req); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	fmt.Println("req: ", req)

	//check if the product quantity is greater than 0 when placing an order for a product else return an error
	if len(req.ProductID) == 0 || len(req.Quantity) <= 0 {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Invalid request body"})
		return
	}

	//validate the request
	if err := validator.Validate(req); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	var service = service.Handler{}

	// Get the product details from the catalogue
	product, err := service.PlaceOrder(req)
	if err != nil {
		c.JSON(http.StatusNotFound, gin.H{"error": "****Product not found****"})
		return
	}

	c.JSON(http.StatusOK, product)
}

// GetAllOrders - returns all the orders placed
func GetAllOrders(c *gin.Context) {

	var service = service.Handler{}

	// Get the product details from the catalogue
	orders, err := service.GetAllOrders()
	if err != nil {
		c.JSON(http.StatusNotFound, gin.H{"error": "****No orders found****"})
		return
	}

	c.JSON(http.StatusOK, orders)
}

// GetOrderDetails - returns the details of a specific order
func GetOrderDetails(c *gin.Context) {

	var service = service.Handler{}
	id, err := strconv.Atoi(c.Query("id"))
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Invalid request body"})
		return
	}

	fmt.Println("req: ", id)

	details, err := service.GetOrderDetails(id)
	if err != nil {
		c.JSON(http.StatusNotFound, gin.H{"error": "****Order not found****"})
		return
	}

	c.JSON(http.StatusOK, details)
}

// UpdateOrderStatus - updates the order status
func UpdateOrderStatus(c *gin.Context) {

	req := models.OrderUpdateRequest{}

	if err := c.ShouldBindJSON(&req); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	fmt.Println("req: ", req)

	if err := validator.Validate(req); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	var service = service.Handler{}
	updateOrder, err := service.UpdateOrderStatus(req)
	if err != nil {
		c.JSON(http.StatusNotFound, gin.H{"error": "Unable to update the order status"})
		return
	}

	c.JSON(http.StatusOK, updateOrder)
}
