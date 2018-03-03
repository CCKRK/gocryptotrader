package binance

import (
	"net/http"
	"testing"

	"github.com/gorilla/websocket"
)

func TestWebsocketSubscribe(t *testing.T) {
	websocketSubcribe := Binance{}
	var Dialer websocket.Dialer
	var err error
	params := make(map[string]string)
	params["pair"] = "BTCUSD"

	websocketSubcribe.WebsocketConn, _, err = Dialer.Dial(binanceWebsocket, http.Header{})
	if err != nil {
		t.Errorf("Test Failed - Binance Dialer error: %s", err)
	}
	err = websocketSubcribe.WebsocketSubscribe("ticker", params)
	if err != nil {
		t.Errorf("Test Failed - Binance WebsocketSubscribe() error: %s", err)
	}

	err = websocketSubcribe.WebsocketConn.Close()
	if err != nil {
		t.Errorf("Test Failed - Binance websocketConn.Close() error: %s", err)
	}
}
