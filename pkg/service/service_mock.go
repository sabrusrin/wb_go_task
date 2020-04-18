package service

import (
	"context"

	"github.com/stretchr/testify/mock"

	v1 "github.com/sabrusrin/wb_go_task/pkg/api/v1"
)

type ServiceMock struct {
	mock.Mock
}

func (s *ServiceMock) GetServiceUser(
	ctx context.Context, request *v1.UserRequest,
) (response v1.UserResponse, err error) {
	args := s.Called(context.Background(), request)
	if a, ok := args.Get(0).(v1.UserResponse); ok {
		return a, args.Error(1)
	}
	return response, args.Error(1)
}

func (s *ServiceMock) PutServiceOrder(
	ctx context.Context, request *v1.OrdersRequest,
) (response v1.OrdersResponse, err error) {
	args := s.Called(context.Background(), request)
	if a, ok := args.Get(0).(v1.OrdersResponse); ok {
		return a, args.Error(1)
	}
	return response, args.Error(1)
}

func (s *ServiceMock) GetUser(
	ctx context.Context, request *v1.UserRequest,
) (response v1.UserResponse, err error) {
	args := s.Called(context.Background(), request)
	if a, ok := args.Get(0).(v1.UserResponse); ok {
		return a, args.Error(1)
	}
	return response, args.Error(1)
}

func (s *ServiceMock) GetOrders() (response v1.OrdersResponse) {
	args := s.Called()
	if a, ok := args.Get(0).(v1.OrdersResponse); ok {
		return a
	}
	return
}
