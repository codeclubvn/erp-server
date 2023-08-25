package middlewares

import (
	"bytes"
	"time"

	"github.com/gin-gonic/gin"
	"go.uber.org/zap"
)

type bodyLogWriter struct {
	gin.ResponseWriter
	body *bytes.Buffer
}

func (w bodyLogWriter) Write(b []byte) (int, error) {
	w.body.Write(b)
	return w.ResponseWriter.Write(b)
}

func (w bodyLogWriter) WriteString(s string) (int, error) {
	w.body.WriteString(s)
	return w.ResponseWriter.WriteString(s)
}

func (e *GinMiddleware) Logger(zapLogger *zap.Logger) gin.HandlerFunc {
	return func(c *gin.Context) {

		start := time.Now()
		blw := &bodyLogWriter{body: bytes.NewBufferString(""), ResponseWriter: c.Writer}
		c.Writer = blw
		c.Next()

		path := c.Request.URL.Path
		raw := c.Request.URL.RawQuery

		end := time.Now()
		latency := end.Sub(start)

		clientIP := c.ClientIP()
		method := c.Request.Method
		statusCode := c.Writer.Status()

		logger := zapLogger.Info

		logger("Request",
			zap.String("Path", path),
			zap.String("Raw", raw),
			zap.String("ClientIP", clientIP),
			zap.String("Method", method),
			zap.Int("StatusCode", statusCode),
			zap.Duration("Latency", latency),
			zap.String("Response", blw.body.String()),
		)
	}
}
