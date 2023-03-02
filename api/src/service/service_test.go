package service

import (
	"fmt"
	"hdfc-assignment/src/models"
	"testing"

	"github.com/stretchr/testify/assert"
)

// GetAllProducts - unit test for GetAllProducts
func TestGetAllProducts(t *testing.T) {
	var service = Handler{}
	catalogue := service.GetAllProducts()
	if catalogue == nil {
		t.Errorf("Expected catalogue to be not nil")
	}
}

// GetProductDetails - unit test for GetProductDetails
func TestGetProductDetails(t *testing.T) {
	var service = Handler{}

	//positive test case
	id := 1
	product, err := service.GetProductDetails(id)
	if err != nil {
		t.Errorf("Expected product to be not nil")
	}
	if product == nil {
		t.Errorf("Expected product to be not nil")
	}

	//negative test case
	id = 100
	product, err = service.GetProductDetails(id)
	if err == nil {
		t.Errorf("Expected product to be nil")
	}
	if product != nil {
		t.Errorf("Expected product to be nil")
	}
}

// PlaceOrder - unit test for PlaceOrder
func TestPlaceOrder(t *testing.T) {
	var service = Handler{}

	//positive test case
	req := []*models.OrderReq{
		{
			ProductId: 1,
			Quantity:  1,
		},
		{
			ProductId: 4,
			Quantity:  1,
		},
		{
			ProductId: 7,
			Quantity:  1,
		},
	}
	res, err := service.PlaceOrder(req)
	if err != nil {
		t.Errorf("Expected err to be nil")
	}
	if res == "" {
		t.Errorf("Expected res to be not nil")
	}

	//negative test case
	req = []*models.OrderReq{
		{
			ProductId: 1,
			Quantity:  100,
		},
		{
			ProductId: 2,
			Quantity:  100,
		},
	}
	res, err = service.PlaceOrder(req)
	if err == nil {
		t.Errorf("Expected err to be not nil")
	}
	if res != "" {
		t.Errorf("Expected res to be nil")
	}

	//negative test case for quantity 0
	req = []*models.OrderReq{
		{
			ProductId: 1,
			Quantity:  0,
		},
	}
	res, err = service.PlaceOrder(req)
	if err == nil {
		t.Errorf("Expected err to be not nil")
	}
	if res != "" {
		t.Errorf("Expected res to be nil")
	}
}

// PlaceOrder - unit test for PlaceOrder for found= false
func TestPlaceOrderFoundFalse(t *testing.T) {
	var service = Handler{}

	//positive test case
	req := []*models.OrderReq{
		{
			ProductId: 1,
			Quantity:  1,
		},
		{
			ProductId: 4,
			Quantity:  1,
		},
		{
			ProductId: 7,
			Quantity:  1,
		},
	}
	res, err := service.PlaceOrder(req)
	if err != nil {
		t.Errorf("Expected err to be nil")
	}
	if res == "" {
		t.Errorf("Expected res to be not nil")
	}

	//negative test case
	req = []*models.OrderReq{
		{
			ProductId: 1,
			Quantity:  100,
		},
		{
			ProductId: 2,
			Quantity:  100,
		},
	}

	//found := false
	res, err = service.PlaceOrder(req)
	if err == nil {
		t.Errorf("Expected err to be not nil")
	}
	if res != "" {
		t.Errorf("Expected res to be nil")
	}

	//negative test case for quantity 0
	req = []*models.OrderReq{
		{
			ProductId: 0,
			Quantity:  0,
		},
	}
	res, err = service.PlaceOrder(req)
	if err == nil {
		t.Errorf("Expected err to be not nil")
	}
	if res != "" {
		t.Errorf("Expected res to be nil")
	}
}

// GetAllOrders - unit test for GetAllOrders
func TestGetAllOrders(t *testing.T) {
	var service = Handler{}

	var err error
	orders := []*models.Order{
		{
			ID: 1,
			ProductDetails: []*models.OrderReq{
				{
					ProductId: 1,
					Quantity:  1,
				},
			},
			OrderStatus: "Placed",
		},
	}

	fmt.Println("orders: ", orders)

	orders, err = service.GetAllOrders()
	if err != nil {
		t.Errorf("Expected err to be nil")
	}
	fmt.Println("orders: ", orders)

	//negative test case
	orders = nil
	orders, err = service.GetAllOrders()
	if err != nil {
		t.Errorf("Expected err to be not nil")
	}
	fmt.Println("orders: ", orders)
}

// GetOrderDetailsByID - unit test for GetOrderDetailsByID
func TestGetOrderDetailsByID(t *testing.T) {
	var service = Handler{}

	orders = []*models.Order{
		{
			ID: 1,
			ProductDetails: []*models.OrderReq{
				{
					ProductId: 1,
					Quantity:  1,
				},
			},
			OrderStatus: "Placed",
		},
	}

	//positive test case
	id := 1
	order, err := service.GetOrderDetailsByID(id)
	assert.Nil(t, err)
	assert.NotNil(t, order)

	//negative test case
	id = 100
	order, err = service.GetOrderDetailsByID(id)
	assert.NotNil(t, err)
	assert.Nil(t, order)
}

// UpdateOrderStatus - unit test for UpdateOrderStatus
func TestUpdateOrderStatus(t *testing.T) {
	var service = Handler{}

	orders = []*models.Order{
		{
			ID: 1,
			ProductDetails: []*models.OrderReq{
				{
					ProductId: 1,
					Quantity:  1,
				},
			},
			OrderStatus: "Placed",
		},
	}

	//positive test case
	req := models.OrderUpdateRequest{
		OrderID:     1,
		OrderStatus: "Dispatched",
	}
	res, err := service.UpdateOrderStatus(req)
	assert.Nil(t, err)
	assert.NotNil(t, res)

	//negative test case
	req = models.OrderUpdateRequest{
		OrderID:     100,
		OrderStatus: "Dispatched",
	}
	res, err = service.UpdateOrderStatus(req)
	assert.NotNil(t, err)
	assert.Equal(t, "", res)
}
