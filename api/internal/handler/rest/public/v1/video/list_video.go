package video

import (
	"net/http"

	"github.com/nhan1603/CryptographicAssignment/api/internal/appconfig/httpserver"
	"github.com/nhan1603/CryptographicAssignment/api/internal/model"
)

type ListVideoResponse struct {
	Videos model.ListSharedVideo `json:"videos"`
}

// Web errors
var (
	webInternalSerror = &httpserver.Error{Status: http.StatusInternalServerError, Code: "internal_error", Desc: "Something went wrong"}
)

func (h Handler) ListVideo() http.HandlerFunc {
	return httpserver.HandlerErr(func(w http.ResponseWriter, r *http.Request) error {

		videoData, err := h.videoCtrl.GetSharedVideo(r.Context())
		if err != nil {
			return webInternalSerror
		}

		httpserver.RespondJSON(w, ListVideoResponse{
			Videos: videoData,
		})

		return nil
	})
}
