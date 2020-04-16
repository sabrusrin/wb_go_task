//Package service logging wrapper
//CODE GENERATED AUTOMATICALLY
//THIS FILE COULD BE EDITED BY HANDS
package service

import (
	"context"
	"time"

	"github.com/go-kit/kit/log"
	"github.com/go-kit/kit/log/level"

	v1 "github.com/sabrusrin/wb_go_task/pkg/api/v1"
)

// loggingMiddleware wraps Service and logs request information to the provided logger
type loggingMiddleware struct {
	logger log.Logger
	svc    Service
}

func (s *loggingMiddleware) GetServiceUser(ctx context.Context, request *v1.UserRequest) (response v1.UserResponse, err error) {
	defer func(begin time.Time) {
		_ = s.wrap(err).Log(
			"method", "GetServiceUser",
			"request", request,
			"err", err,
			"elapsed", time.Since(begin),
		)
	}(time.Now())
	return s.svc.GetServiceUser(ctx, request)
}

func (s *loggingMiddleware) PutServiceOrder(ctx context.Context, request *v1.OrdersRequest) (response v1.OrdersResponse, err error) {
	defer func(begin time.Time) {
		_ = s.wrap(err).Log(
			"method", "PutServiceOrder",
			"request", request,
			"err", err,
			"elapsed", time.Since(begin),
		)
	}(time.Now())
	return s.svc.PutServiceOrder(ctx, request)
}

func (s *loggingMiddleware) GetUser(ctx context.Context, request *v1.UserRequest) (response v1.UserResponse, err error) {
	defer func(begin time.Time) {
		_ = s.wrap(err).Log(
			"method", "GetUser",
			"request", request,
			"err", err,
			"elapsed", time.Since(begin),
		)
	}(time.Now())
	return s.svc.GetUser(ctx, request)
}

func (s *loggingMiddleware) GetOrders() (response v1.OrdersResponse) {
	var err error
	defer func(begin time.Time) {
		_ = s.wrap(err).Log(
			"method", "GetOrders",

			"err", err,
			"elapsed", time.Since(begin),
		)
	}(time.Now())
	return s.svc.GetOrders()
}

func (s *loggingMiddleware) wrap(err error) log.Logger {
	lvl := level.Debug
	if err != nil {
		lvl = level.Error
	}
	return lvl(s.logger)
}

// NewLoggingMiddleware ...
func NewLoggingMiddleware(logger log.Logger, svc Service) Service {
	return &loggingMiddleware{
		logger: logger,
		svc:    svc,
	}
}
