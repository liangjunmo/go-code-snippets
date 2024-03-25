package trace

import (
	"context"

	"github.com/gin-gonic/gin"
)

func GinMiddleware() gin.HandlerFunc {
	return func(c *gin.Context) {
		ctx := c.Request.Context()
		for _, key := range traceKeys {
			ctx = context.WithValue(ctx, key, c.GetHeader(key))
		}
		if ctx.Value(traceIDKey) == "" {
			ctx = Trace(ctx)
		}
		c.Request = c.Request.WithContext(ctx)
		c.Next()
	}
}
