package trace

import (
	"context"
)

type TraceIDGenerator func() string

var traceIDGenerator TraceIDGenerator

func SetTraceIDGenerator(generator TraceIDGenerator) {
	traceIDGenerator = generator
}

func Trace(ctx context.Context) context.Context {
	return context.WithValue(ctx, traceIDKey, traceIDGenerator())
}
