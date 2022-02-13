package wss

import (
	"log"
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/gorilla/websocket"
)

var wssUpgrader = websocket.Upgrader{
	ReadBufferSize:  1024,
	WriteBufferSize: 1024,
	CheckOrigin:     func(r *http.Request) bool { return true },
}

func handleWssConnection(conn *websocket.Conn) {
	dontExit := true
	conn.SetCloseHandler(func(code int, text string) error {
		dontExit = false
		log.Println("CloseHandler", code, text)
		conn.Close()

		return nil
	})
	for dontExit {
		_, msgBytes, err := conn.ReadMessage()
		if err != nil {
			log.Println("Error occurred while handling WSS", err)
			dontExit = false
			if err := conn.Close(); err != nil {
				panic(err)
			}
		} else {
			processWssMessage(msgBytes, conn)
		}
	}
}

func processWssMessage(_msg []byte, conn *websocket.Conn) {
	msg := string(_msg)
	if WSS_FUNCS[msg] != nil {
		err := WSS_FUNCS[msg](conn)
		if err != nil {
			log.Println("Error occurred, while processing ", msg, err)
		}
	} else {
		log.Println("Unknown command", msg, "from", conn.RemoteAddr())
	}
}

func acceptWs(c *gin.Context) {
	if c.IsWebsocket() {
		conn, err := wssUpgrader.Upgrade(c.Writer, c.Request, nil)
		if err != nil {
			log.Println("WSS connection failed", err)
			return
		} else {
			go handleWssConnection(conn)
		}
	}
}

func IntializeWss(router *gin.Engine) {
	router.GET("/ws", acceptWs)
}
