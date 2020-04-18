package httpclient

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
	"github.com/sabrusrin/wb_go_task/pkg/service"
	"github.com/sabrusrin/wb_go_task/pkg/service/httpserver"
)

const (
	serverAddress = "localhost:8085"
	serverHost    = "localhost:8085"

	serverLaunchingWaitSleep = 1 * time.Second

	getServiceUserSuccess  = "GetServiceUser success test"
	putServiceOrderSuccess = "PutServiceOrder success test"
	getUserSuccess         = "GetUser success test"
	getOrdersSuccess       = "GetOrders success test"

	serviceMethodGetServiceUser  = "GetServiceUser"
	serviceMethodPutServiceOrder = "PutServiceOrder"
	serviceMethodGetUser         = "GetUser"
	serviceMethodGetOrders       = "GetOrders"

	maxConns          = 4096
	serverReadTimeout = 1 * time.Second
)

var (
	nilClientError error
)

func TestClient_GetServiceUserSuccess(t *testing.T) {
	initServiceUser := v1.UserRequest{UserId: 15}
	expectedResponse := v1.UserResponse{
		Error:       false,
		ErrorText:   "",
		Data:        nil,
		CustomError: nil,
	}
	t.Run(getServiceUserSuccess, func(t *testing.T) {
		serviceMock := new(service.ServiceMock)
		serviceMock.On(serviceMethodGetServiceUser, context.Background(), &initServiceUser).
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
		time.Sleep(serverLaunchingWaitSleep)
		response, err := client.GetServiceUser(context.Background(), &initServiceUser)
		assert.NoError(t, err, "unexpected error:", err)
		assert.EqualValues(t, expectedResponse, response)
	})
}

func TestClient_PutServiceOrderSuccess(t *testing.T) {
	initServiceOrder := v1.OrdersRequest{OrderId: 0}
	expectedResponse := v1.OrdersResponse{
		Error:       false,
		ErrorText:   "",
		Data:        nil,
		CustomError: nil,
	}
	t.Run(putServiceOrderSuccess, func(t *testing.T) {
		serviceMock := new(service.ServiceMock)
		serviceMock.On(serviceMethodPutServiceOrder, context.Background(), &initServiceOrder).
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
		time.Sleep(serverLaunchingWaitSleep)
		response, err := client.PutServiceOrder(context.Background(), &initServiceOrder)
		assert.NoError(t, err, "unexpected error:", err)
		assert.EqualValues(t, expectedResponse, response)
	})
}

func TestClient_GetUserSuccess(t *testing.T) {
	initServiceUser := v1.UserRequest{UserId: 17}
	expectedResponse := v1.UserResponse{
		Error:       false,
		ErrorText:   "",
		Data:        nil,
		CustomError: nil,
	}
	t.Run(getUserSuccess, func(t *testing.T) {
		serviceMock := new(service.ServiceMock)
		serviceMock.On(serviceMethodGetUser, context.Background(), &initServiceUser).
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
		time.Sleep(serverLaunchingWaitSleep)
		response, err := client.GetUser(context.Background(), &initServiceUser)
		assert.NoError(t, err, "unexpected error:", err)
		assert.EqualValues(t, expectedResponse, response)
	})
}

func TestClient_GetOrdersSuccess(t *testing.T) {
	expectedResponse := v1.OrdersResponse{
		Error:       false,
		ErrorText:   "",
		Data:        nil,
		CustomError: nil,
	}
	t.Run(getOrdersSuccess, func(t *testing.T) {
		serviceMock := new(service.ServiceMock)
		serviceMock.On(serviceMethodGetOrders).
			Return(expectedResponse).Once()
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
		time.Sleep(serverLaunchingWaitSleep)
		response := client.GetOrders()
		assert.EqualValues(t, expectedResponse, response)
	})
}

func makeServer(
	serverAddress string,
	serverHost string,
	svc Service,
) (server *fasthttp.Server, client Service) {
	errorProcessor := serv.NewErrorProcessor(http.StatusInternalServerError, "internal error")
	client = NewPreparedClient(
		"http://"+serverAddress,
		serverHost,
		maxConns,
		errorProcessor,
		serv.NewError,
		httpserver.URIPathClientGetServiceUser,
		httpserver.URIPathClientPutServiceOrder,
		httpserver.URIPathClientGetUser,
		httpserver.URIPathClientGetOrders,
		http.MethodGet,
		http.MethodPost,
		http.MethodGet,
		http.MethodGet,
	)
	router := serv.MakeFastHTTPRouter(
		[]*serv.HandlerSettings{
			{
				Path:   httpserver.URIPathGetServiceUser,
				Method: http.MethodGet,
				Handler: httpserver.NewGetServiceUserServer(
					httpserver.NewGetServiceUserTransport(serv.NewError),
					svc,
					errorProcessor,
				),
			},
			{
				Path:   httpserver.URIPathPutServiceOrder,
				Method: http.MethodPost,
				Handler: httpserver.NewPutServiceOrderServer(
					httpserver.NewPutServiceOrderTransport(serv.NewError),
					svc,
					errorProcessor,
				),
			},
			{
				Path:   httpserver.URIPathGetUser,
				Method: http.MethodGet,
				Handler: httpserver.NewGetUserServer(
					httpserver.NewGetUserTransport(serv.NewError),
					svc,
					errorProcessor,
				),
			},
			{
				Path:   httpserver.URIPathGetOrders,
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
