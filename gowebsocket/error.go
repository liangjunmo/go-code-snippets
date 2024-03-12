package gowebsocket

import (
	"context"
)

type ErrorReceiver func(ctx context.Context, err error)

var errorReceiver ErrorReceiver

func SetErrorReceiver(receiver ErrorReceiver) {
	errorReceiver = receiver
}
