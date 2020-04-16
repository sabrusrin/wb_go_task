//Package service http transport
//CODE GENERATED AUTOMATICALLY
//THIS FILE COULD BE EDITED BY HANDS
package httpserver

import (
	"context"
	"net/http"
	"strconv"

	"github.com/mailru/easyjson"
	"github.com/valyala/fasthttp"

	v1 "github.com/sabrusrin/wb_go_task/pkg/api/v1"
)

type errorCreator func(status int, format string, v ...interface{}) error

// GetServiceUserTransport transport interface
type GetServiceUserTransport interface {
	DecodeRequest(ctx context.Context, r *fasthttp.Request) (request v1.UserRequest, err error)
	EncodeResponse(ctx context.Context, r *fasthttp.Response, response *v1.UserResponse) (err error)
}

type getServiceUserTransport struct {
	errorCreator errorCreator
}

// DecodeRequest method for decoding requests on server side
func (t *getServiceUserTransport) DecodeRequest(ctx context.Context, r *fasthttp.Request) (request v1.UserRequest, err error) {
	id, err := strconv.ParseInt(string(r.URI().QueryArgs().Peek("id")), 10, 0)
	if err != nil {
		return v1.UserRequest{}, t.errorCreator(http.StatusBadRequest, "failed to decode JSON request: %v", err)
	}
	request.UserId = int(id)
	return
}

// EncodeResponse method for encoding response on server side
func (t *getServiceUserTransport) EncodeResponse(ctx context.Context, r *fasthttp.Response, response *v1.UserResponse) (err error) {
	r.Header.Set("Content-Type", "application/json")
	if _, err = easyjson.MarshalToWriter(response, r.BodyWriter()); err != nil {
		return t.errorCreator(http.StatusInternalServerError, "failed to encode JSON response: %s", err)
	}
	return
}

// NewGetServiceUserTransport the transport creator for http requests
func NewGetServiceUserTransport(errorCreator errorCreator) GetServiceUserTransport {
	return &getServiceUserTransport{
		errorCreator: errorCreator,
	}
}

// PutServiceOrderTransport transport interface
type PutServiceOrderTransport interface {
	DecodeRequest(ctx context.Context, r *fasthttp.Request) (request v1.OrdersRequest, err error)
	EncodeResponse(ctx context.Context, r *fasthttp.Response, response *v1.OrdersResponse) (err error)
}

type putServiceOrderTransport struct {
	errorCreator errorCreator
}

// DecodeRequest method for decoding requests on server side
func (t *putServiceOrderTransport) DecodeRequest(ctx context.Context, r *fasthttp.Request) (request v1.OrdersRequest, err error) {
	if err = request.UnmarshalJSON(r.Body()); err != nil {
		return v1.OrdersRequest{}, t.errorCreator(http.StatusBadRequest, "failed to decode JSON request: %v", err)
	}
	return
}

// EncodeResponse method for encoding response on server side
func (t *putServiceOrderTransport) EncodeResponse(ctx context.Context, r *fasthttp.Response, response *v1.OrdersResponse) (err error) {
	r.Header.Set("Content-Type", "application/json")
	if _, err = easyjson.MarshalToWriter(response, r.BodyWriter()); err != nil {
		return t.errorCreator(http.StatusInternalServerError, "failed to encode JSON response: %s", err)
	}
	return
}

// NewPutServiceOrderTransport the transport creator for http requests
func NewPutServiceOrderTransport(errorCreator errorCreator) PutServiceOrderTransport {
	return &putServiceOrderTransport{
		errorCreator: errorCreator,
	}
}

// GetUserTransport transport interface
type GetUserTransport interface {
	DecodeRequest(ctx *fasthttp.RequestCtx, r *fasthttp.Request) (request v1.UserRequest, err error)
	EncodeResponse(ctx context.Context, r *fasthttp.Response, response *v1.UserResponse) (err error)
}

type getUserTransport struct {
	errorCreator errorCreator
}

// DecodeRequest method for decoding requests on server side
func (t *getUserTransport) DecodeRequest(ctx *fasthttp.RequestCtx, r *fasthttp.Request) (request v1.UserRequest, err error) {
	request.UserId, err = strconv.Atoi(ctx.UserValue("id").(string))
	if err != nil {
		return v1.UserRequest{}, t.errorCreator(http.StatusBadRequest, "failed to decode JSON request: %v", err)
	}
	return
}

// EncodeResponse method for encoding response on server side
func (t *getUserTransport) EncodeResponse(ctx context.Context, r *fasthttp.Response, response *v1.UserResponse) (err error) {
	r.Header.Set("Content-Type", "application/json")
	if _, err = easyjson.MarshalToWriter(response, r.BodyWriter()); err != nil {
		return t.errorCreator(http.StatusInternalServerError, "failed to encode JSON response: %s", err)
	}
	return
}

// NewGetUserTransport the transport creator for http requests
func NewGetUserTransport(errorCreator errorCreator) GetUserTransport {
	return &getUserTransport{
		errorCreator: errorCreator,
	}
}

// GetOrdersTransport transport interface
type GetOrdersTransport interface {
	DecodeRequest() (err error)
	EncodeResponse(ctx context.Context, r *fasthttp.Response, response *v1.OrdersResponse) (err error)
}

type getOrdersTransport struct {
	errorCreator errorCreator
}

// DecodeRequest method for decoding requests on server side
func (t *getOrdersTransport) DecodeRequest() (err error) {
	return
}

// EncodeResponse method for encoding response on server side
func (t *getOrdersTransport) EncodeResponse(ctx context.Context, r *fasthttp.Response, response *v1.OrdersResponse) (err error) {
	r.Header.Set("Content-Type", "application/json")
	if _, err = easyjson.MarshalToWriter(response, r.BodyWriter()); err != nil {
		return t.errorCreator(http.StatusInternalServerError, "failed to encode JSON response: %s", err)
	}
	return
}

// NewGetOrdersTransport the transport creator for http requests
func NewGetOrdersTransport(errorCreator errorCreator) GetOrdersTransport {
	return &getOrdersTransport{
		errorCreator: errorCreator,
	}
}
