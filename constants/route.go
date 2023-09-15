package constants

import "net/http"

var PublicRoutes = map[string]string{
	"/v1/api/auth/login":    http.MethodPost,
	"/v1/api/auth/register": http.MethodPost,
	"/v1/api/health/":       http.MethodGet,
	"/v1/api/erp/store/":    http.MethodPost,
}
