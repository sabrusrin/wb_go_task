package v1

// Order data
type Order struct {
	Res bool `json:"res"`
}

// OrdersRequest ...
type OrdersRequest struct {
	OrderId int `json:"orderId"`
}

// OrdersResponse ...
type OrdersResponse struct {
	Error       bool              `json:"error"`
	ErrorText   string            `json:"errorText"`
	Data        *Order            `json:"data"`
	CustomError map[string]string `json:"additionalErrors"`
}

// User data
type User struct {
	Res bool `json:"res"`
}

// UserRequest
type UserRequest struct {
	UserId int `json:"userId"`
}

// UserResponse ...
type UserResponse struct {
	Error       bool              `json:"error"`
	ErrorText   string            `json:"errorText"`
	Data        *User             `json:"data"`
	CustomError map[string]string `json:"additionalErrors"`
}
