package gotracing

import (
	"context"
	"errors"
	"net/http"
	"testing"

	"github.com/gin-gonic/gin"
	"github.com/levigross/grequests"
	"github.com/stretchr/testify/require"
)

func TestGinMiddleware(t *testing.T) {
	resetTracingKeys()

	key := "key"
	value := "value"
	key2 := "key2"
	value2 := "value2"

	SetTracingIDKey(key)
	SetTracingIDGenerator(func() string {
		return value
	})
	AppendTracingKeys([]string{key2})

	gin.SetMode(gin.ReleaseMode)

	router := gin.New()
	router.Use(GinMiddleware())
	router.GET("/", func(c *gin.Context) {
		ctx := c.Request.Context()
		require.Equal(t, value, ctx.Value(key))
		require.Equal(t, value2, ctx.Value(key2))
	})

	server := &http.Server{
		Addr:    ":8000",
		Handler: router,
	}

	go func() {
		err := server.ListenAndServe()
		if !errors.Is(err, http.ErrServerClosed) {
			t.Error(err)
			return
		}
	}()

	_, err := grequests.Get("http://127.0.0.1:8000/", &grequests.RequestOptions{
		Headers: map[string]string{
			key2: value2,
		},
	})
	require.Nil(t, err)

	err = server.Shutdown(context.Background())
	require.Nil(t, err)
}
