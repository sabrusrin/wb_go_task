package service

import (
	"context"
	"fmt"

	v1 "github.com/sabrusrin/wb_go_task/pkg/api/v1"
)

// Service ...
type Service interface {
	GetServiceUser(ctx context.Context, request *v1.UserRequest) (response v1.UserResponse, err error)
	PutServiceOrder(ctx context.Context, request *v1.OrdersRequest) (response v1.OrdersResponse, err error)
	GetUser(ctx context.Context, request *v1.UserRequest) (response v1.UserResponse, err error)
	GetOrders() (response v1.OrdersResponse)
}

type service struct {
}

func (s *service) GetServiceUser(_ context.Context, request *v1.UserRequest,
) (response v1.UserResponse, err error) {
	if request.UserId <= 0 {
		err = fmt.Errorf("error: bad user ID")
		return
	}
	data := v1.User{Res: true,}
	response.Data = &data
	return
}

func (s *service) PutServiceOrder(_ context.Context, request *v1.OrdersRequest,
) (response v1.OrdersResponse, err error) {
	if request.OrderId <= 0 {
		err = fmt.Errorf("error: bad orders ID")
		return
	}
	data := v1.Order{Res: true}
	response.Data = &data
	return
}

func (s *service) GetUser(_ context.Context, request *v1.UserRequest,
) (response v1.UserResponse, err error) {
	if request.UserId <= 0 {
		err = fmt.Errorf("error: bad user ID")
		return
	}
	data := v1.User{Res: true,}
	response.Data = &data
	return
}

func (s *service) GetOrders() (response v1.OrdersResponse) {
	data := v1.Order{Res: true}
	response.Data = &data
	return
}

// NewService ...
func NewService() Service {
	return &service{}
}
