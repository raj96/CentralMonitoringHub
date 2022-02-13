package wss

import (
	"github.com/gorilla/websocket"
)

var WSS_FUNCS map[string]func(*websocket.Conn) error

func InitWssFuncs() {
	WSS_FUNCS = make(map[string]func(*websocket.Conn) error)

	WSS_FUNCS["get-all-data"] = getAllData
}

func getAllData(conn *websocket.Conn) error {

	return nil
}
