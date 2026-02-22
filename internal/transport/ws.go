package transport

import (
	"context"
	"fmt"
	"net/http"
	"net/url"

	"github.com/gorilla/websocket"
)

type WSConnection struct {
	conn *websocket.Conn
}

func (t *Transport) DialWebSocket(
	ctx context.Context,
	path string,
	query url.Values,
) (*WSConnection, error) {

	parsed, err := url.Parse(t.BaseURL)
	if err != nil {
		return nil, err
	}

	switch parsed.Scheme {
	case "https":
		parsed.Scheme = "wss"
	case "http":
		parsed.Scheme = "ws"
	}

	parsed.Path = path
	parsed.RawQuery = query.Encode()

	header := http.Header{}
	header.Set("Api-Subscription-Key", t.APIKey)

	dialer := websocket.DefaultDialer
	// debug statement with the full URL being dialed
	fmt.Printf("Dialing WebSocket URL: %s\n", parsed.String())
	conn, _, err := dialer.DialContext(ctx, parsed.String(), header)
	if err != nil {
		return nil, err
	}

	return &WSConnection{conn: conn}, nil
}

func (w *WSConnection) ReadMessage() (int, []byte, error) {
	mtype, data, err := w.conn.ReadMessage()
	return mtype, data, err
}

func (w *WSConnection) WriteJSON(v any) error {
	return w.conn.WriteJSON(v)
}

func (w *WSConnection) Close() error {
	return w.conn.Close()
}
