package gowebsocket

import (
	"context"
	"fmt"
)

type Route string

type Handler func(ctx context.Context, message Message, writing chan<- Message) error

var (
	handlers     = make(map[Route]Handler)
	interceptors = make(map[Route]Handler)
)

func Handle(route Route, handler Handler) {
	if _, ok := handlers[route]; ok {
		panic(fmt.Sprintf("handler already exists: %s", route))
	}
	handlers[route] = handler
}

func Intercept(route Route, interceptor Handler) {
	if _, ok := interceptors[route]; ok {
		panic(fmt.Sprintf("interceptor already exists: %s", route))
	}
	interceptors[route] = interceptor
}
