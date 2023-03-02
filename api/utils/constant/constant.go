package constant

import (
	"encoding/json"
	"fmt"
)

var (
	PREMIUM    = "Premium"
	DISPATCHED = "Dispatched"
	PLACED     = "Placed"

	ErrProductNotFound      = fmt.Errorf("product not found")
	ErrQuantityExceedsLimit = fmt.Errorf("requested quantity exceeds the maximum limit")
	ErrInsufficientQuantity = fmt.Errorf("product quantity not available")
	ErrOrderNotFound        = fmt.Errorf("order not found")

	ORDERPLACED     = "Ordered placed sucessfully"
	ORDERUPDATED    = "Order status updated successfully"
	ORDERNOTUPDATED = "Unable to update the order status"
	ORDERNOTFOUND   = "Order not found"
	QUANTITYEXCEEDS = "requested quantity exceeds available quantity: %d > %d"
	INVALIDREQUEST  = "Invalid request body"
)

// unmarshalData - unmarshals the data
func UnmarshalData(data string, v interface{}) error {
	err := json.Unmarshal([]byte(data), v)
	if err != nil {
		return err
	}
	return nil
}
