package constants

import "net/http"

var PublicRoutes = map[string]string{
	"/v1/handler/auth/login":    http.MethodPost,
	"/v1/handler/auth/register": http.MethodPost,
	"/v1/handler/health/":       http.MethodGet,
	"/v1/handler/erp/store/":    http.MethodPost,
	"/v1/handler/erp/customer/": http.MethodPost,
}
