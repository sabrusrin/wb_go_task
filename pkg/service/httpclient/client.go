//Package service http client
//CODE GENERATED AUTOMATICALLY
//THIS FILE COULD BE EDITED BY HANDS
package httpclient

import (
	"context"
	"github.com/valyala/fasthttp"

	v1 "github.com/sabrusrin/wb_go_task/pkg/api/v1"
)

// Service implements Service interface
type Service interface {
	GetServiceUser(ctx context.Context, request *v1.UserRequest) (response v1.UserResponse, err error)
	PutServiceOrder(ctx context.Context, request *v1.OrdersRequest) (response v1.OrdersResponse, err error)
	GetUser(ctx context.Context, request *v1.UserRequest) (response v1.UserResponse, err error)
	GetOrders() (response v1.OrdersResponse)
}

type client struct {
	cli *fasthttp.HostClient

	transportGetServiceUser  GetServiceUserClientTransport
	transportPutServiceOrder PutServiceOrderClientTransport
	transportGetUser         GetUserClientTransport
	transportGetOrders       GetOrdersClientTransport
}

// GetServiceUser ...
func (s *client) GetServiceUser(ctx context.Context, request *v1.UserRequest) (response v1.UserResponse, err error) {
	req, res := fasthttp.AcquireRequest(), fasthttp.AcquireResponse()
	defer func() {
		fasthttp.ReleaseRequest(req)
		fasthttp.ReleaseResponse(res)
	}()
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
func (s *client) GetOrders() (response v1.OrdersResponse) {
	var err error
	req, res := fasthttp.AcquireRequest(), fasthttp.AcquireResponse()
	defer func() {
		fasthttp.ReleaseRequest(req)
		fasthttp.ReleaseResponse(res)
	}()
	if err = s.transportGetOrders.EncodeRequest(); err != nil {
		return
	}
	err = s.cli.Do(req, res)
	if err != nil {
		return
	}
	response, err = s.transportGetOrders.DecodeResponse(res)
	return
}

// NewClient the client creator
func NewClient(
	cli *fasthttp.HostClient,

	transportGetServiceUser GetServiceUserClientTransport,
	transportPutServiceOrder PutServiceOrderClientTransport,
	transportGetUser GetUserClientTransport,
	transportGetOrders GetOrdersClientTransport,
) Service {
	return &client{
		cli: cli,

		transportGetServiceUser:  transportGetServiceUser,
		transportPutServiceOrder: transportPutServiceOrder,
		transportGetUser:         transportGetUser,
		transportGetOrders:       transportGetOrders,
	}
}

// NewPreparedClient create and set up http client
func NewPreparedClient(
	serverURL string,
	serverHost string,
	maxConns int,
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
	)
}
