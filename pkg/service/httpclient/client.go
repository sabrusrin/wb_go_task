//Package service http client
//CODE GENERATED AUTOMATICALLY
//THIS FILE COULD BE EDITED BY HANDS
package httpclient

import (
	"context"

	"github.com/valyala/fasthttp"

	v1 "github.com/sabrusrin/wb_go_task/pkg/api/v1"
)

var (
	GetServiceUser  = option{}
	PutServiceOrder = option{}
	GetUser         = option{}
	GetOrders       = option{}
)

type option struct{}

// Option ...
type Option interface {
	Prepare(ctx context.Context, r *fasthttp.Request)
}

// Service implements Service interface
type Service interface {
	GetServiceUser(ctx context.Context, request *v1.UserRequest) (response v1.UserResponse, err error)
	PutServiceOrder(ctx context.Context, request *v1.OrdersRequest) (response v1.OrdersResponse, err error)
	GetUser(ctx context.Context, request *v1.UserRequest) (response v1.UserResponse, err error)
	GetOrders(ctx context.Context, ) (response v1.OrdersResponse, err error)
}

type client struct {
	cli *fasthttp.HostClient

	transportGetServiceUser  GetServiceUserClientTransport
	transportPutServiceOrder PutServiceOrderClientTransport
	transportGetUser         GetUserClientTransport
	transportGetOrders       GetOrdersClientTransport
	options                  map[interface{}]Option
}

// GetServiceUser ...
func (s *client) GetServiceUser(ctx context.Context, request *v1.UserRequest) (response v1.UserResponse, err error) {
	req, res := fasthttp.AcquireRequest(), fasthttp.AcquireResponse()
	defer func() {
		fasthttp.ReleaseRequest(req)
		fasthttp.ReleaseResponse(res)
	}()
	if opt, ok := s.options[GetServiceUser]; ok {
		opt.Prepare(ctx, req)
	}
	if err = s.transportGetServiceUser.EncodeRequest(ctx, req, request); err != nil {
		return
	}
	err = s.cli.Do(req, res)
	if err != nil {
		return
	}
	return s.transportGetServiceUser.DecodeResponse(ctx, res)
}

// PutServiceOrder ...
func (s *client) PutServiceOrder(ctx context.Context, request *v1.OrdersRequest) (response v1.OrdersResponse, err error) {
	req, res := fasthttp.AcquireRequest(), fasthttp.AcquireResponse()
	defer func() {
		fasthttp.ReleaseRequest(req)
		fasthttp.ReleaseResponse(res)
	}()
	if opt, ok := s.options[PutServiceOrder]; ok {
		opt.Prepare(ctx, req)
	}
	if err = s.transportPutServiceOrder.EncodeRequest(ctx, req, request); err != nil {
		return
	}
	err = s.cli.Do(req, res)
	if err != nil {
		return
	}
	return s.transportPutServiceOrder.DecodeResponse(ctx, res)
}

// GetUser ...
func (s *client) GetUser(ctx context.Context, request *v1.UserRequest) (response v1.UserResponse, err error) {
	req, res := fasthttp.AcquireRequest(), fasthttp.AcquireResponse()
	defer func() {
		fasthttp.ReleaseRequest(req)
		fasthttp.ReleaseResponse(res)
	}()
	if opt, ok := s.options[GetUser]; ok {
		opt.Prepare(ctx, req)
	}
	if err = s.transportGetUser.EncodeRequest(ctx, req, request); err != nil {
		return
	}
	err = s.cli.Do(req, res)
	if err != nil {
		return
	}
	return s.transportGetUser.DecodeResponse(ctx, res)
}

// GetOrders ...
func (s *client) GetOrders(ctx context.Context, ) (response v1.OrdersResponse, err error) {
	req, res := fasthttp.AcquireRequest(), fasthttp.AcquireResponse()
	defer func() {
		fasthttp.ReleaseRequest(req)
		fasthttp.ReleaseResponse(res)
	}()
	if opt, ok := s.options[GetOrders]; ok {
		opt.Prepare(ctx, req)
	}
	if err = s.transportGetOrders.EncodeRequest(ctx, req, ); err != nil {
		return
	}
	err = s.cli.Do(req, res)
	if err != nil {
		return
	}
	return s.transportGetOrders.DecodeResponse(ctx, res)
}

// NewClient the client creator
func NewClient(
	cli *fasthttp.HostClient,

	transportGetServiceUser GetServiceUserClientTransport,
	transportPutServiceOrder PutServiceOrderClientTransport,
	transportGetUser GetUserClientTransport,
	transportGetOrders GetOrdersClientTransport,
	options map[interface{}]Option,
) Service {
	return &client{
		cli: cli,

		transportGetServiceUser:  transportGetServiceUser,
		transportPutServiceOrder: transportPutServiceOrder,
		transportGetUser:         transportGetUser,
		transportGetOrders:       transportGetOrders,
		options:                  options,
	}
}

// NewPreparedClient create and set up http client
func NewPreparedClient(
	serverURL string,
	serverHost string,
	maxConns int,
	options map[interface{}]Option,
	errorProcessor errorProcessor,
	errorCreator errorCreator,

	uriPathGetServiceUser string,
	uriPathPutServiceOrder string,
	uriPathGetUser string,
	uriPathGetOrders string,

	httpMethodGetServiceUser string,
	httpMethodPutServiceOrder string,
	httpMethodGetUser string,
	httpMethodGetOrders string,
) Service {

	transportGetServiceUser := NewGetServiceUserClientTransport(
		errorProcessor,
		errorCreator,
		serverURL+uriPathGetServiceUser,
		httpMethodGetServiceUser,
	)

	transportPutServiceOrder := NewPutServiceOrderClientTransport(
		errorProcessor,
		errorCreator,
		serverURL+uriPathPutServiceOrder,
		httpMethodPutServiceOrder,
	)

	transportGetUser := NewGetUserClientTransport(
		errorProcessor,
		errorCreator,
		serverURL+uriPathGetUser,
		httpMethodGetUser,
	)

	transportGetOrders := NewGetOrdersClientTransport(
		errorProcessor,
		errorCreator,
		serverURL+uriPathGetOrders,
		httpMethodGetOrders,
	)

	return NewClient(
		&fasthttp.HostClient{
			Addr:     serverHost,
			MaxConns: maxConns,
		},

		transportGetServiceUser,
		transportPutServiceOrder,
		transportGetUser,
		transportGetOrders,
		options,
	)
}
