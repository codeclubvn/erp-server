package middlewares

type GinMiddleware struct {
}

func NewMiddleware() *GinMiddleware {
	middleware := &GinMiddleware{}
	return middleware
}
