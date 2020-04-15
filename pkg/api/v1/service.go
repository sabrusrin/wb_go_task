package v1

// Order data
type Order struct {
	Res bool
}

// OrdersRequest ...
type OrdersRequest struct {
	OrderId int
}

// OrdersResponse ...
type OrdersResponse struct {
	Error       bool
	ErrorText   string
	Data        *Order
	CustomError map[string]string
}

// User data
type User struct {
	Res bool
}

// UserRequest
type UserRequest struct {
	UserId int
}

// UserResponse ...
type UserResponse struct {
	Error       bool
	ErrorText   string
	Data        *User
	CustomError map[string]string
}
