package notification

import (
	"github.com/nii236/leeroyci/database"
	"github.com/nii236/leeroyci/websocket"
)

func sendWebsocket(job *database.Job, event string) {
	msg := websocket.NewMessage(job, event)
	websocket.Send(msg)
}
