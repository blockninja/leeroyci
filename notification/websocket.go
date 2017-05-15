package notification

import (
	"github.com/blockninja/leeroyci/database"
	"github.com/blockninja/leeroyci/websocket"
)

func sendWebsocket(job *database.Job, event string) {
	msg := websocket.NewMessage(job, event)
	websocket.Send(msg)
}
