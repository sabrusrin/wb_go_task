package httpclient

import (
	"context"
	v1 "github.com/sabrusrin/wb_go_task/pkg/api/v1"
	"github.com/sabrusrin/wb_go_task/pkg/service/httpserver"
	"github.com/stretchr/testify/assert"
	"log"
	"net/http"
	"testing"
	"time"

	"github.com/valyala/fasthttp"

	serv "github.com/sabrusrin/wb_go_task/pkg/httpserver"
	"github.com/sabrusrin/wb_go_task/pkg/service"
)

const (
	serverAddress = "localhost:8085"
	serverHost    = "localhost:8085"

	serverLaunchingWaitSleep = 1 * time.Second

	getServiceUserSuccess  = "GetServiceUser success test"
	putServiceOrderSuccess = "PutServiceOrder success test"
	getUserSuccess         = "GetUser success test"
	getOrdersSuccess       = "GetOrders success test"
	getServiceUserFail     = "GetServiceUser fail test"
	putServiceOrderFail    = "PutServiceOrder fail test"
	getUserFail            = "GetUser fail test"

	serviceUserError   = "error: bad user ID"
	serviceOrdersError = "error: bad orders ID"

	serviceMethodGetServiceUser  = "GetServiceUser"
	serviceMethodPutServiceOrder = "PutServiceOrder"
	serviceMethodGetUser         = "GetUser"
	serviceMethodGetOrders       = "GetOrders"

	idGood            = 15
	idBad             = -1
	maxConns          = 4096
	serverReadTimeout = 1 * time.Second
)

var (
	nilClientError error
)

func TestClient_GetServiceUserSuccess(t *testing.T) {
	initServiceUser := v1.UserRequest{UserId: idGood}
	userData := v1.User{Res: true}
	expectedResponse := v1.UserResponse{
		Error:       false,
		ErrorText:   "",
		Data:        &userData,
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
	initServiceOrder := v1.OrdersRequest{OrderId: idGood}
	orderData := v1.Order{Res: true}
	expectedResponse := v1.OrdersResponse{
		Error:       false,
		ErrorText:   "",
		Data:        &orderData,
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
	initServiceUser := v1.UserRequest{UserId: idGood}
	userData := v1.User{Res: true}
	expectedResponse := v1.UserResponse{
		Error:       false,
		ErrorText:   "",
		Data:        &userData,
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
	orderData := v1.Order{Res: true}
	expectedResponse := v1.OrdersResponse{
		Error:       false,
		ErrorText:   "",
		Data:        &orderData,
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

func TestClient_GetServiceUserFail(t *testing.T) {
	initServiceUser := v1.UserRequest{UserId: idBad}
	expectedResponse := v1.UserResponse{}

	t.Run(getServiceUserFail, func(t *testing.T) {
		serviceMock := new(service.ServiceMock)
		serviceMock.On(serviceMethodGetServiceUser, context.Background(), &initServiceUser).
			Return(expectedResponse,
				serv.NewError(http.StatusServiceUnavailable, serviceUserError)).
			Once()

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
		assert.Equal(t, serv.NewError(http.StatusServiceUnavailable, serviceUserError), err)
		assert.EqualValues(t, expectedResponse, response)
	})
}

func TestClient_PutServiceOrderFail(t *testing.T) {
	initServiceOrder := v1.OrdersRequest{OrderId: idBad}
	expectedResponse := v1.OrdersResponse{}

	t.Run(putServiceOrderFail, func(t *testing.T) {
		serviceMock := new(service.ServiceMock)
		serviceMock.On(serviceMethodPutServiceOrder, context.Background(), &initServiceOrder).
			Return(expectedResponse,
				serv.NewError(http.StatusServiceUnavailable, serviceOrdersError)).
			Once()

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
		assert.Equal(t, serv.NewError(http.StatusServiceUnavailable, serviceOrdersError), err)
		assert.EqualValues(t, expectedResponse, response)
	})
}

func TestClient_GetUserFail(t *testing.T) {
	initServiceUser := v1.UserRequest{UserId: idBad}
	expectedResponse := v1.UserResponse{}

	t.Run(getUserFail, func(t *testing.T) {
		serviceMock := new(service.ServiceMock)
		serviceMock.On(serviceMethodGetUser, context.Background(), &initServiceUser).
			Return(expectedResponse,
				serv.NewError(http.StatusServiceUnavailable, serviceUserError)).
			Once()

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
		assert.Equal(t, serv.NewError(http.StatusServiceUnavailable, serviceUserError), err)
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
	)
	router := httpserver.NewPreparedServer(svc)
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
