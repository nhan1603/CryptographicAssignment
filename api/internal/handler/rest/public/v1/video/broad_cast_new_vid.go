package video

import (
	"log"
	"net/http"

	"github.com/gorilla/websocket"
	"github.com/nhan1603/CryptographicAssignment/api/internal/appconfig/httpserver"
)

// BroadCastResponse broadcasts email of new video to client
func (h Handler) BroadCastResponse() http.HandlerFunc {
	return httpserver.HandlerErr(func(w http.ResponseWriter, r *http.Request) error {
		log.Println("[BroadCastResponse] START processing requests")
		var upgrader = websocket.Upgrader{
			ReadBufferSize:  1024,
			WriteBufferSize: 1024,
			CheckOrigin:     func(r *http.Request) bool { return true },
		}

		// init websocket
		ws, err := upgrader.Upgrade(w, r, nil)
		if err != nil {
			log.Fatal(err)
		}

		h.videoCtrl.BroadCastResponse(r.Context(), ws)
		return nil
	})
}
