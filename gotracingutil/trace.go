package gotracingutil

import (
	"context"
)

type TracingIDGenerator func() string

var tracingIDGenerator TracingIDGenerator

func SetTracingIDGenerator(generator TracingIDGenerator) {
	tracingIDGenerator = generator
}

func Trace(ctx context.Context) context.Context {
	return context.WithValue(ctx, tracingIDKey, tracingIDGenerator())
}
