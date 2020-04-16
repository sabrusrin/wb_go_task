//Package service http server
//CODE GENERATED AUTOMATICALLY
//THIS FILE COULD BE EDITED BY HANDS
package httpserver

import (
	"context"

	"github.com/valyala/fasthttp"

	v1 "github.com/sabrusrin/wb_go_task/pkg/api/v1"
)

type errorProcessor interface {
	Encode(ctx context.Context, r *fasthttp.Response, err error)
}

type service interface {
	GetServiceUser(ctx context.Context, request *v1.UserRequest) (response v1.UserResponse, err error)
	PutServiceOrder(ctx context.Context, request *v1.OrdersRequest) (response v1.OrdersResponse, err error)
	GetUser(ctx context.Context, request *v1.UserRequest) (response v1.UserResponse, err error)
	GetOrders() (response v1.OrdersResponse)
}

type getServiceUserServer struct {
	transport      GetServiceUserTransport
	service        service
	errorProcessor errorProcessor
}

// ServeHTTP implements http.Handler.
func (s *getServiceUserServer) ServeHTTP(ctx *fasthttp.RequestCtx) {
	request, err := s.transport.DecodeRequest(ctx, &ctx.Request)
	if err != nil {
		s.errorProcessor.Encode(ctx, &ctx.Response, err)
		return
	}

	response, err := s.service.GetServiceUser(ctx, &request)
	if err != nil {
		s.errorProcessor.Encode(ctx, &ctx.Response, err)
		return
	}

	if err := s.transport.EncodeResponse(ctx, &ctx.Response, &response); err != nil {
		s.errorProcessor.Encode(ctx, &ctx.Response, err)
		return
	}
}

// NewGetServiceUserServer the server creator
func NewGetServiceUserServer(transport GetServiceUserTransport, service service, errorProcessor errorProcessor) fasthttp.RequestHandler {
	ls := getServiceUserServer{
		transport:      transport,
		service:        service,
		errorProcessor: errorProcessor,
	}
	return ls.ServeHTTP
}

type putServiceOrderServer struct {
	transport      PutServiceOrderTransport
	service        service
	errorProcessor errorProcessor
}

// ServeHTTP implements http.Handler.
func (s *putServiceOrderServer) ServeHTTP(ctx *fasthttp.RequestCtx) {
	request, err := s.transport.DecodeRequest(ctx, &ctx.Request)
	if err != nil {
		s.errorProcessor.Encode(ctx, &ctx.Response, err)
		return
	}

	response, err := s.service.PutServiceOrder(ctx, &request)
	if err != nil {
		s.errorProcessor.Encode(ctx, &ctx.Response, err)
		return
	}

	if err := s.transport.EncodeResponse(ctx, &ctx.Response, &response); err != nil {
		s.errorProcessor.Encode(ctx, &ctx.Response, err)
		return
	}
}

// NewPutServiceOrderServer the server creator
func NewPutServiceOrderServer(transport PutServiceOrderTransport, service service, errorProcessor errorProcessor) fasthttp.RequestHandler {
	ls := putServiceOrderServer{
		transport:      transport,
		service:        service,
		errorProcessor: errorProcessor,
	}
	return ls.ServeHTTP
}

type getUserServer struct {
	transport      GetUserTransport
	service        service
	errorProcessor errorProcessor
}

// ServeHTTP implements http.Handler.
func (s *getUserServer) ServeHTTP(ctx *fasthttp.RequestCtx) {
	request, err := s.transport.DecodeRequest(ctx, &ctx.Request)
	if err != nil {
		s.errorProcessor.Encode(ctx, &ctx.Response, err)
		return
	}

	response, err := s.service.GetUser(ctx, &request)
	if err != nil {
		s.errorProcessor.Encode(ctx, &ctx.Response, err)
		return
	}

	if err := s.transport.EncodeResponse(ctx, &ctx.Response, &response); err != nil {
		s.errorProcessor.Encode(ctx, &ctx.Response, err)
		return
	}
}

// NewGetUserServer the server creator
func NewGetUserServer(transport GetUserTransport, service service, errorProcessor errorProcessor) fasthttp.RequestHandler {
	ls := getUserServer{
		transport:      transport,
		service:        service,
		errorProcessor: errorProcessor,
	}
	return ls.ServeHTTP
}

type getOrdersServer struct {
	transport      GetOrdersTransport
	service        service
	errorProcessor errorProcessor
}

// ServeHTTP implements http.Handler.
func (s *getOrdersServer) ServeHTTP(ctx *fasthttp.RequestCtx) {
	err := s.transport.DecodeRequest()
	if err != nil {
		s.errorProcessor.Encode(ctx, &ctx.Response, err)
		return
	}

	response := s.service.GetOrders()

	if err := s.transport.EncodeResponse(ctx, &ctx.Response, &response); err != nil {
		s.errorProcessor.Encode(ctx, &ctx.Response, err)
		return
	}
}

// NewGetOrdersServer the server creator
func NewGetOrdersServer(transport GetOrdersTransport, service service, errorProcessor errorProcessor) fasthttp.RequestHandler {
	ls := getOrdersServer{
		transport:      transport,
		service:        service,
		errorProcessor: errorProcessor,
	}
	return ls.ServeHTTP
}
