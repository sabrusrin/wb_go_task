//Package service instrumenting wrapper
//CODE GENERATED AUTOMATICALLY
//THIS FILE COULD BE EDITED BY HANDS
package service

import (
	"context"
	"strconv"
	"time"

	"github.com/go-kit/kit/metrics"

	v1 "github.com/sabrusrin/wb_go_task/pkg/api/v1"
)

// instrumentingMiddleware wraps Service and enables request metrics
type instrumentingMiddleware struct {
	reqCount    metrics.Counter
	reqDuration metrics.Histogram
	svc         Service
}

func (s *instrumentingMiddleware) GetServiceUser(ctx context.Context, request *v1.UserRequest) (response v1.UserResponse, err error) {
	defer s.recordMetrics("GetServiceUser", time.Now(), err)
	return s.svc.GetServiceUser(ctx, request)
}

func (s *instrumentingMiddleware) PutServiceOrder(ctx context.Context, request *v1.OrdersRequest) (response v1.OrdersResponse, err error) {
	defer s.recordMetrics("PutServiceOrder", time.Now(), err)
	return s.svc.PutServiceOrder(ctx, request)
}

func (s *instrumentingMiddleware) GetUser(ctx context.Context, request *v1.UserRequest) (response v1.UserResponse, err error) {
	defer s.recordMetrics("GetUser", time.Now(), err)
	return s.svc.GetUser(ctx, request)
}

func (s *instrumentingMiddleware) GetOrders() (response v1.OrdersResponse) {
	var err error
	defer s.recordMetrics("GetOrders", time.Now(), err)
	return s.svc.GetOrders()
}

func (s *instrumentingMiddleware) recordMetrics(method string, startTime time.Time, err error) {
	labels := []string{
		"method", method,
		"error", strconv.FormatBool(err != nil),
	}
	s.reqCount.With(labels...).Add(1)
	s.reqDuration.With(labels...).Observe(time.Since(startTime).Seconds())
}

// NewInstrumentingMiddleware ...
func NewInstrumentingMiddleware(reqCount metrics.Counter, reqDuration metrics.Histogram, svc Service) Service {
	return &instrumentingMiddleware{
		reqCount:    reqCount,
		reqDuration: reqDuration,
		svc:         svc,
	}
}
