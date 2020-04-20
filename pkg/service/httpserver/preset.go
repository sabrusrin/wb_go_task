//Package service http client
//CODE GENERATED AUTOMATICALLY
//THIS FILE COULD BE EDITED BY HANDS
package httpserver

const (
	URIPrefix = "/api/v1"

	URIPathClientGetServiceUser  = URIPrefix + "/user/%d"
	URIPathClientPutServiceOrder = URIPrefix + "/orders"
	URIPathClientGetUser         = URIPrefix + "/user/%d/count"
	URIPathClientGetOrders       = URIPrefix + "/orders"

	URIPathGetServiceUser  = URIPrefix + "/user/:id"
	URIPathPutServiceOrder = URIPrefix + "/orders"
	URIPathGetUser         = URIPrefix + "/user/:id/count"
	URIPathGetOrders       = URIPrefix + "/orders"

	HTTPMethodGetServiceUser  = "GET"
	HTTPMethodPutServiceOrder = "POST"
	HTTPMethodGetUser         = "GET"
	HTTPMethodGetOrders       = "GET"
)
