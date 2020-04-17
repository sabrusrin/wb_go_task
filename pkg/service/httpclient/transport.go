//Package service http client
//CODE GENERATED AUTOMATICALLY
//THIS FILE COULD BE EDITED BY HANDS
package httpclient

import (
	"bufio"
	"context"
	"encoding/json"
	"fmt"
	"net/http"
	"strconv"

	"github.com/valyala/fasthttp"

	v1 "github.com/sabrusrin/wb_go_task/pkg/api/v1"
)

type errorCreator func(status int, format string, v ...interface{}) error

type errorProcessor interface {
	Encode(ctx context.Context, r *fasthttp.Response, err error)
	Decode(r *fasthttp.Response) error
}

// GetServiceUserClientTransport transport interface
type GetServiceUserClientTransport interface {
	EncodeRequest(ctx context.Context, r *fasthttp.Request, request *v1.UserRequest) (err error)
	DecodeResponse(ctx context.Context, r *fasthttp.Response) (response v1.UserResponse, err error)
}

type getServiceUserClientTransport struct {
	errorProcessor errorProcessor
	errorCreator   errorCreator
	pathTemplate   string
	method         string
}

// EncodeRequest method for encoding requests on client side
func (t *getServiceUserClientTransport) EncodeRequest(ctx context.Context, r *fasthttp.Request, request *v1.UserRequest) (err error) {
	r.Header.SetMethod(t.method)
	r.SetRequestURI(t.pathTemplate)
	r.Header.Set("Content-Type", "application/json")
	r.URI().QueryArgs().Add("id", strconv.FormatInt(int64(request.UserId), 10))
	return
}

// DecodeResponse method for decoding response on client side
func (t *getServiceUserClientTransport) DecodeResponse(ctx context.Context, r *fasthttp.Response) (response v1.UserResponse, err error) {
	if r.StatusCode() != http.StatusOK {
		err = t.errorProcessor.Decode(r)
		return
	}
	err = response.UnmarshalJSON(r.Body())
	return
}

// NewGetServiceUserClientTransport the transport creator for http requests
func NewGetServiceUserClientTransport(
	errorProcessor errorProcessor,
	errorCreator errorCreator,
	pathTemplate string,
	method string,
) GetServiceUserClientTransport {
	return &getServiceUserClientTransport{
		errorProcessor: errorProcessor,
		errorCreator:   errorCreator,
		pathTemplate:   pathTemplate,
		method:         method,
	}
}

// PutServiceOrderClientTransport transport interface
type PutServiceOrderClientTransport interface {
	EncodeRequest(ctx context.Context, r *fasthttp.Request, request *v1.OrdersRequest) (err error)
	DecodeResponse(ctx context.Context, r *fasthttp.Response) (response v1.OrdersResponse, err error)
}

type putServiceOrderClientTransport struct {
	errorProcessor errorProcessor
	errorCreator   errorCreator
	pathTemplate   string
	method         string
}

// EncodeRequest method for encoding requests on client side
func (t *putServiceOrderClientTransport) EncodeRequest(ctx context.Context, r *fasthttp.Request, request *v1.OrdersRequest) (err error) {
	r.Header.SetMethod(t.method)
	r.SetRequestURI(t.pathTemplate)
	r.Header.Set("Content-Type", "application/json")
	r.SetBodyStreamWriter(func(w *bufio.Writer) {
		if err = json.NewEncoder(w).Encode(request); err != nil {
			return
		}
	})
	return
}

// DecodeResponse method for decoding response on client side
func (t *putServiceOrderClientTransport) DecodeResponse(ctx context.Context, r *fasthttp.Response) (response v1.OrdersResponse, err error) {
	if r.StatusCode() != http.StatusOK {
		err = t.errorProcessor.Decode(r)
		return
	}
	err = response.UnmarshalJSON(r.Body())
	return
}

// NewPutServiceOrderClientTransport the transport creator for http requests
func NewPutServiceOrderClientTransport(
	errorProcessor errorProcessor,
	errorCreator errorCreator,
	pathTemplate string,
	method string,
) PutServiceOrderClientTransport {
	return &putServiceOrderClientTransport{
		errorProcessor: errorProcessor,
		errorCreator:   errorCreator,
		pathTemplate:   pathTemplate,
		method:         method,
	}
}

// GetUserClientTransport transport interface
type GetUserClientTransport interface {
	EncodeRequest(ctx context.Context, r *fasthttp.Request, request *v1.UserRequest) (err error)
	DecodeResponse(ctx context.Context, r *fasthttp.Response) (response v1.UserResponse, err error)
}

type getUserClientTransport struct {
	errorProcessor errorProcessor
	errorCreator   errorCreator
	pathTemplate   string
	method         string
}

// EncodeRequest method for encoding requests on client side
func (t *getUserClientTransport) EncodeRequest(ctx context.Context, r *fasthttp.Request, request *v1.UserRequest) (err error) {
	r.Header.SetMethod(t.method)
	r.SetRequestURI(fmt.Sprintf(t.pathTemplate, request.UserId))
	r.Header.Set("Content-Type", "application/json")
	r.SetBodyStreamWriter(func(w *bufio.Writer) {
		if err = json.NewEncoder(w).Encode(request); err != nil {
			return
		}
	})
	return
}

// DecodeResponse method for decoding response on client side
func (t *getUserClientTransport) DecodeResponse(ctx context.Context, r *fasthttp.Response) (response v1.UserResponse, err error) {
	if r.StatusCode() != http.StatusOK {
		err = t.errorProcessor.Decode(r)
		return
	}
	err = response.UnmarshalJSON(r.Body())
	return
}

// NewGetUserClientTransport the transport creator for http requests
func NewGetUserClientTransport(
	errorProcessor errorProcessor,
	errorCreator errorCreator,
	pathTemplate string,
	method string,
) GetUserClientTransport {
	return &getUserClientTransport{
		errorProcessor: errorProcessor,
		errorCreator:   errorCreator,
		pathTemplate:   pathTemplate,
		method:         method,
	}
}

// GetOrdersClientTransport transport interface
type GetOrdersClientTransport interface {
	EncodeRequest() (err error)
	DecodeResponse(r *fasthttp.Response) (response v1.OrdersResponse, err error)
}

type getOrdersClientTransport struct {
	errorProcessor errorProcessor
	errorCreator   errorCreator
	pathTemplate   string
	method         string
}

// EncodeRequest method for encoding requests on client side
func (t *getOrdersClientTransport) EncodeRequest() (err error) {
	return
}

// DecodeResponse method for decoding response on client side
func (t *getOrdersClientTransport) DecodeResponse(r *fasthttp.Response) (response v1.OrdersResponse, err error) {
	if r.StatusCode() != http.StatusOK {
		err = t.errorProcessor.Decode(r)
		return
	}
	err = response.UnmarshalJSON(r.Body())
	return
}

// NewGetOrdersClientTransport the transport creator for http requests
func NewGetOrdersClientTransport(
	errorProcessor errorProcessor,
	errorCreator errorCreator,
	pathTemplate string,
	method string,
) GetOrdersClientTransport {
	return &getOrdersClientTransport{
		errorProcessor: errorProcessor,
		errorCreator:   errorCreator,
		pathTemplate:   pathTemplate,
		method:         method,
	}
}
