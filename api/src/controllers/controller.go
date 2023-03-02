// Package controllers - This package contains the controllers for the qed request
package controllers

import (
	"net/http"
	"strconv"

	"hdfc-assignment/src/models"
	"hdfc-assignment/src/service"
	"hdfc-assignment/utils/constant"

	"github.com/gin-gonic/gin"
)

// GetAllProducts - returns all the products in the catalogue
func GetAllProducts(c *gin.Context) {

	var service = service.Handler{}
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
	product, err := service.GetProductDetails(id)
	if err != nil {
		c.JSON(http.StatusNotFound, gin.H{"error": "Product not found"})
		return
	}
	c.JSON(http.StatusOK, product)
}

// PlaceOrder - places an order for a product
func PlaceOrder(c *gin.Context) {

	req := []*models.OrderReq{}
	res := make(map[string]interface{})

	if err := c.ShouldBindJSON(&req); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}
	var service = service.Handler{}
	product, err := service.PlaceOrder(req)
	if err != nil {
		res["error"] = err.Error()
		c.JSON(http.StatusBadRequest, res)
		return
	}
	res["data"] = product
	c.JSON(http.StatusOK, res)
}

// GetAllOrders - returns all the orders placed
func GetAllOrders(c *gin.Context) {
	var service = service.Handler{}
	orders, err := service.GetAllOrders()
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": constant.ORDERNOTFOUND})
		return
	}
	c.JSON(http.StatusOK, orders)
}

// GetOrderDetails - returns the details of a specific order
func GetOrderDetails(c *gin.Context) {

	var service = service.Handler{}
	id, err := strconv.Atoi(c.Query("id"))
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": constant.INVALIDREQUEST})
		return
	}
	details, err := service.GetOrderDetailsByID(id)
	if err != nil {
		c.JSON(http.StatusNotFound, gin.H{"error": constant.ORDERNOTFOUND})
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
	var service = service.Handler{}
	updateOrder, err := service.UpdateOrderStatus(req)
	if err != nil {
		c.JSON(http.StatusNotFound, gin.H{"error": constant.ORDERNOTUPDATED})
		return
	}
	c.JSON(http.StatusOK, gin.H{"data": updateOrder})
}
