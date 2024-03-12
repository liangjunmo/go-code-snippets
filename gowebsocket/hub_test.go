package gowebsocket

import (
	"context"
	"encoding/json"
	"errors"
	"log"
	"net/http"
	"net/url"
	"testing"
	"time"

	"github.com/gorilla/websocket"
	"github.com/stretchr/testify/require"
)

const (
	addr  = "localhost:8080"
	route = "/ws"

	RouteTestRequest  Route = "TestRequest"
	RouteTestResponse Route = "TestResponse"
)

var wsUpgrader = websocket.Upgrader{}

type TestMessage struct {
	ID      string `json:"id"`
	Route   Route  `json:"route"`
	Payload string `json:"payload"`
}

func (m TestMessage) GetID() string {
	return m.ID
}

func (m TestMessage) GetRoute() Route {
	return m.Route
}

func (m TestMessage) GetPayload() string {
	return m.Payload
}

var testMessageParser MessageParser = func(ctx context.Context, raw []byte) (Message, error) {
	var message TestMessage
	err := json.Unmarshal(raw, &message)
	if err != nil {
		return TestMessage{}, err
	}
	return message, nil
}

var testErrorReceiver ErrorReceiver = func(ctx context.Context, err error) {
	log.Println(err)
}

func TestHub(t *testing.T) {
	SetMessageParser(testMessageParser)
	SetErrorReceiver(testErrorReceiver)

	done := make(chan struct{})

	Intercept(RouteTestRequest, func(ctx context.Context, message Message, writing chan<- Message) error {
		t.Logf("%+v", message)
		writing <- TestMessage{
			ID:      time.Now().String(),
			Route:   RouteTestResponse,
			Payload: message.GetPayload() + "...",
		}
		return nil
	})

	Handle(RouteTestResponse, func(ctx context.Context, message Message, writing chan<- Message) error {
		t.Logf("%+v", message)
		done <- struct{}{}
		return nil
	})

	mux := http.NewServeMux()
	mux.HandleFunc(route, func(w http.ResponseWriter, r *http.Request) {
		conn, err := wsUpgrader.Upgrade(w, r, nil)
		require.Nil(t, err)

		hub := NewHub(conn, 10, 10)
		hub.Run(r.Context())
	})

	server := &http.Server{
		Addr:    addr,
		Handler: mux,
	}
	go func() {
		err := server.ListenAndServe()
		if !errors.Is(err, http.ErrServerClosed) {
			t.Error(err)
			return
		}
	}()

	u := url.URL{Host: addr, Scheme: "ws", Path: route}
	conn, _, err := websocket.DefaultDialer.Dial(u.String(), nil)
	require.Nil(t, err)
	defer conn.Close()

	hub := NewHub(conn, 10, 10)
	ctx := context.Background()
	go func() {
		message := TestMessage{
			ID:      time.Now().String(),
			Route:   RouteTestRequest,
			Payload: "hello world",
		}

		b, _ := json.Marshal(message)
		err = conn.WriteMessage(websocket.TextMessage, b)
		require.Nil(t, err)

		<-done
		conn.Close()
		err = server.Shutdown(ctx)
		require.Nil(t, err)
	}()
	hub.Run(ctx)
}
