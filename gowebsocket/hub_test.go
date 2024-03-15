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

	routeTestRequest  Route = "TestRequest"
	routeTestResponse Route = "TestResponse"
)

type testMessage struct {
	ID      string `json:"id"`
	Route   Route  `json:"route"`
	Payload string `json:"payload"`
}

func (m testMessage) GetID() string {
	return m.ID
}

func (m testMessage) GetRoute() Route {
	return m.Route
}

func (m testMessage) GetPayload() string {
	return m.Payload
}

var testMessageParser MessageParser = func(ctx context.Context, raw []byte) (Message, error) {
	var message testMessage
	err := json.Unmarshal(raw, &message)
	if err != nil {
		return testMessage{}, err
	}
	return message, nil
}

var testErrorReceiver ErrorReceiver = func(ctx context.Context, err error) {
	log.Println(err)
}

var wsUpgrader = websocket.Upgrader{}

func TestHub(t *testing.T) {
	SetMessageParser(testMessageParser)
	SetErrorReceiver(testErrorReceiver)

	done := make(chan struct{})

	Intercept(routeTestRequest, func(ctx context.Context, message Message, writing chan<- Message) error {
		t.Logf("%+v", message)
		writing <- testMessage{
			ID:      time.Now().String(),
			Route:   routeTestResponse,
			Payload: message.GetPayload() + "...",
		}
		return nil
	})

	Handle(routeTestResponse, func(ctx context.Context, message Message, writing chan<- Message) error {
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

	ctx := context.Background()
	hub := NewHub(conn, 10, 10)
	go func() {
		message := testMessage{
			ID:      time.Now().String(),
			Route:   routeTestRequest,
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
