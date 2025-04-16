package video

import "github.com/nhan1603/CryptographicAssignment/api/internal/controller/videos"

// Handler is the web handler for this pkg
type Handler struct {
	videoCtrl videos.Controller
}

// New instantiates a new Handler and returns it
func New(videoCtrl videos.Controller) Handler {
	return Handler{videoCtrl: videoCtrl}
}
