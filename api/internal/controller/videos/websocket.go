package videos

import (
	"context"
	"log"

	"github.com/gorilla/websocket"
	"github.com/nhan1603/CryptographicAssignment/api/internal/model"
)

// BroadCastResponse broadcasts alert to client
func (i impl) BroadCastResponse(_ context.Context, ws *websocket.Conn) {
	i.clients[ws] = true
	go func(i impl) {
		for {
			val := <-i.broadcast
			// send to every client that is currently connected
			for client := range i.clients {
				err := client.WriteJSON(val)
				if err != nil {
					log.Printf("Websocket error: %s", err)
					client.Close()
					delete(i.clients, client)
				}
			}
		}
	}(i)
}

// Push pushes message to ws
func (i impl) Push(ctx context.Context, email string) {
	go func() {
		message := model.NewVideoMessage{
			SharerEmail: email,
		}
		i.broadcast <- message
	}()
}
