package videos

import (
	"context"

	"github.com/gorilla/websocket"
	"github.com/nhan1603/CryptographicAssignment/api/internal/model"
	"github.com/nhan1603/CryptographicAssignment/api/internal/repository"
)

// Controller represents the specification of this pkg
type Controller interface {
	ShareVideo(ctx context.Context, videoUrl, videoTitle, sharerEmail string) error
	GetSharedVideo(ctx context.Context) (model.ListSharedVideo, error)
	// Push send a response message to ws
	Push(ctx context.Context, message string)
	// BroadCastResponse create a ws for broadcast response
	BroadCastResponse(ctx context.Context, ws *websocket.Conn)
}

// New initializes a new Controller instance and returns it
func New(repo repository.Registry,
	clients map[*websocket.Conn]bool,
	broadcast chan model.NewVideoMessage) Controller {
	return impl{
		repo:      repo,
		clients:   clients,
		broadcast: broadcast,
	}
}

type impl struct {
	repo      repository.Registry
	clients   map[*websocket.Conn]bool
	broadcast chan model.NewVideoMessage
}
