package service

import (
	"context"
	v1 "github.com/sabrusrin/wb_go_task/pkg/api/v1"
	"github.com/stretchr/testify/assert"
	"log"
	"net/http"
	"testing"
	"time"

	"github.com/valyala/fasthttp"

	serv "github.com/sabrusrin/wb_go_task/pkg/httpserver"
	"github.com/sabrusrin/wb_go_task/pkg/service/httpclient"
	"github.com/sabrusrin/wb_go_task/pkg/service/httpserver"
)

const (
	uriPathUserQuery   = "/api/v1/user"
	uriPathUser   = "/api/v1/user/{id}/count"
	uriPathOrders = "/api/v1/orders"

	serverAddress  = "localhost:38812"
	serverHost  = "localhost:38812"
	//getServiceUserSuccessAddress  = "localhost:38812"
	//getServiceUserFailAddress     = "localhost:38813"
	//putServiceOrderSuccessAddress = "localhost:38814"
	//putServiceOrderFailAddress    = "localhost:38815"
	//getUserSuccessAddress         = "localhost:38816"
	//getUserFailAddress            = "localhost:38817"
	//getOrdersSuccessAddress       = "localhost:38818"
	//getOrdersFailAddress          = "localhost:38819"

	maxConns          = 4096
	serverReadTimeout = 1 * time.Second
)

var (
	nilClientError error
)

func TestClient_GetServiceUserSuccess(t *testing.T) {
	initServiceUser := v1.UserRequest{UserId:15}
	expectedResponse := v1.UserResponse{
		Error:       false,
		ErrorText:   "",
		Data:        &v1.User{Res:true},
		CustomError: nil,
	}
	t.Run("Success serviceUser init", func(t *testing.T) {
		serviceMock := new(ServiceMock)
		serviceMock.On("GetServiceUser", context.Background(), &initServiceUser).
			Return(expectedResponse, nilClientError).Once()
		server, client := makeServer(
			serverAddress,
			serverHost,
			serviceMock,
			)
		defer func() {
			err := server.Shutdown()
			if err != nil {
				log.Printf("server shut down err: %v", err)
			}
		}()
		response, err := client.GetServiceUser(context.Background(), &initServiceUser)
		assert.NoError(t, err, "unexpected error:", err)
		assert.EqualValues(t, expectedResponse, response)
	})
}

func makeServer(
	serverAddress string,
	serverHost string,
	svc Service,
) (server *fasthttp.Server, client Service) {
	errorProcessor := serv.NewErrorProcessor(http.StatusInternalServerError, "internal error")
	client = httpclient.NewPreparedClient(
		"http://"+serverAddress,
		serverHost,
		maxConns,
		errorProcessor,
		serv.NewError,
		uriPathUser,
		uriPathOrders,
		uriPathUser,
		uriPathOrders,
		http.MethodGet,
		http.MethodPost,
		http.MethodGet,
		http.MethodGet,
		)
	router := serv.MakeFastHTTPRouter(
		[]*serv.HandlerSettings{
			{
				Path:   uriPathUserQuery,
				Method: http.MethodGet,
				Handler: httpserver.NewGetServiceUserServer(
					httpserver.NewGetServiceUserTransport(serv.NewError),
					svc,
					errorProcessor,
				),
			},
			{
				Path:   uriPathOrders,
				Method: http.MethodPost,
				Handler: httpserver.NewPutServiceOrderServer(
					httpserver.NewPutServiceOrderTransport(serv.NewError),
					svc,
					errorProcessor,
				),
			},
			{
				Path:   uriPathUser,
				Method: http.MethodGet,
				Handler: httpserver.NewGetUserServer(
					httpserver.NewGetUserTransport(serv.NewError),
					svc,
					errorProcessor,
				),
			},
			{
				Path:   uriPathOrders,
				Method: http.MethodGet,
				Handler: httpserver.NewGetOrdersServer(
					httpserver.NewGetOrdersTransport(serv.NewError),
					svc,
					errorProcessor,
				),
			},
		},
	)
	server = &fasthttp.Server{
		Handler:     router.Handler,
		ReadTimeout: serverReadTimeout,
	}
	go func() {
		err := server.ListenAndServe(serverAddress)
		if err != nil {
			log.Printf("server shut down err: %v", err)
		}
	}()
	return
}
