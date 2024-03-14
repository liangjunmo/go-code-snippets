package gowebsocket

import (
	"context"
	"encoding/json"
	"fmt"
	"time"

	"github.com/gorilla/websocket"
)

type Hub struct {
	conn *websocket.Conn

	canceled chan struct{}

	reading chan Message
	writing chan Message

	// https://github.com/gorilla/websocket/blob/666c197fc9157896b57515c3a3326c3f8c8319fe/examples/chat/client.go
	pingPeriod     time.Duration
	readWait       time.Duration // pongWait
	writeWait      time.Duration
	maxMessageSize int64
}

func NewHub(conn *websocket.Conn, readingCapacity uint32, writingCapacity uint32) *Hub {
	if messageParser == nil {
		panic(fmt.Sprintf("message parser is nil"))
	}
	if errorReceiver == nil {
		panic(fmt.Sprintf("error receiver is nil"))
	}
	if len(handlers) == 0 && len(interceptors) == 0 {
		panic(fmt.Sprintf("handlers and interceptors are both empty"))
	}
	return &Hub{
		conn: conn,

		canceled: make(chan struct{}),

		reading: make(chan Message, readingCapacity),
		writing: make(chan Message, writingCapacity),

		pingPeriod:     time.Second * 60 * 9 / 10,
		readWait:       time.Second * 60,
		writeWait:      time.Second * 10,
		maxMessageSize: 512,
	}
}

func (hub *Hub) SetPingPeriod(period time.Duration) {
	hub.pingPeriod = period
}

func (hub *Hub) SetReadWait(wait time.Duration) {
	hub.readWait = wait
}

func (hub *Hub) SetWriteWait(wait time.Duration) {
	hub.writeWait = wait
}

func (hub *Hub) SetMaxMessageSize(size int64) {
	hub.maxMessageSize = size
}

func (hub *Hub) Run(ctx context.Context) {
	ctx, cancel := context.WithCancel(ctx)

	go hub.read(ctx)
	go hub.write(ctx)
	go hub.handle(ctx)

	<-hub.canceled
	cancel()
	<-hub.canceled
	<-hub.canceled

	hub.conn.Close()
	close(hub.canceled)
	close(hub.reading)
	close(hub.writing)
}

func (hub *Hub) read(ctx context.Context) {
	var (
		err     error
		message Message
	)

	defer func() {
		hub.canceled <- struct{}{}
	}()

	hub.conn.SetReadLimit(hub.maxMessageSize)

	err = hub.conn.SetReadDeadline(time.Now().Add(hub.readWait))
	if err != nil {
		err = fmt.Errorf("set read deadline: %w", err)
		errorReceiver(ctx, err)
		return
	}

	hub.conn.SetPongHandler(func(string) error {
		err := hub.conn.SetReadDeadline(time.Now().Add(hub.readWait))
		if err != nil {
			err = fmt.Errorf("set read deadline: %w", err)
			errorReceiver(ctx, err)
			return err
		}
		return nil
	})

	for {
		select {
		case <-ctx.Done():
			return
		default:
			_, b, e := hub.conn.ReadMessage()
			if e != nil {
				err = fmt.Errorf("read message: %w", e)
				errorReceiver(ctx, err)
				return
			}

			message, e = messageParser(ctx, b)
			if e != nil {
				errorReceiver(ctx, fmt.Errorf("parse message: %w", e))
				continue
			}

			if interceptor, ok := interceptors[message.GetRoute()]; ok {
				e = interceptor(ctx, message, hub.writing)
				if e != nil {
					errorReceiver(ctx, fmt.Errorf("call interceptor %s: %w", message.GetRoute(), e))
					continue
				}
				continue
			}

			hub.reading <- message
		}
	}
}

func (hub *Hub) write(ctx context.Context) {
	var err error

	defer func() {
		if err != nil {
			err = hub.conn.WriteMessage(websocket.CloseMessage, []byte{})
			if err != nil {
				errorReceiver(ctx, fmt.Errorf("write message: %w", err))
			}
		}
		hub.canceled <- struct{}{}
	}()

	ticker := time.NewTicker(hub.pingPeriod)
	for {
		select {
		case <-ctx.Done():
			return
		case <-ticker.C:
			e := hub.conn.SetWriteDeadline(time.Now().Add(hub.writeWait))
			if e != nil {
				err = fmt.Errorf("set write deadline: %w", e)
				errorReceiver(ctx, err)
				return
			}

			e = hub.conn.WriteMessage(websocket.PingMessage, nil)
			if e != nil {
				err = fmt.Errorf("write message: %w", e)
				errorReceiver(ctx, err)
				return
			}
		case message, ok := <-hub.writing:
			if !ok {
				err = fmt.Errorf("read writing chan failed")
				errorReceiver(ctx, err)
				return
			}

			e := hub.conn.SetWriteDeadline(time.Now().Add(hub.writeWait))
			if e != nil {
				err = fmt.Errorf("set write deadline: %w", e)
				errorReceiver(ctx, err)
				return
			}

			b, e := json.Marshal(message)
			if e != nil {
				errorReceiver(ctx, fmt.Errorf("marshal message: %w", e))
				continue
			}

			e = hub.conn.WriteMessage(websocket.TextMessage, b)
			if e != nil {
				err = fmt.Errorf("write message: %w", e)
				errorReceiver(ctx, err)
				return
			}
		}
	}
}

func (hub *Hub) handle(ctx context.Context) {
	var err error

	defer func() {
		hub.canceled <- struct{}{}
	}()

	for {
		select {
		case <-ctx.Done():
			return
		case message, ok := <-hub.reading:
			if !ok {
				err = fmt.Errorf("read reading chan failed")
				errorReceiver(ctx, err)
				return
			}

			handler, ok := handlers[message.GetRoute()]
			if !ok {
				errorReceiver(ctx, fmt.Errorf("handler not found: %s", message.GetRoute()))
				continue
			}

			e := handler(ctx, message, hub.writing)
			if e != nil {
				errorReceiver(ctx, fmt.Errorf("call handler %s: %w", message.GetRoute(), e))
				continue
			}
		}
	}
}
