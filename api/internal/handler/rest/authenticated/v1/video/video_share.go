package video

import (
	"encoding/json"
	"net/http"

	"github.com/nhan1603/CryptographicAssignment/api/internal/appconfig/httpserver"
	"github.com/nhan1603/CryptographicAssignment/api/internal/appconfig/iam"
	"github.com/nhan1603/CryptographicAssignment/api/internal/controller/videos"
)

func (h Handler) ShareVideo() http.HandlerFunc {
	type request struct {
		VideoUrl   string `json:"videoUrl"`
		VideoTitle string `json:"videoTitle"`
	}

	return httpserver.HandlerErr(func(w http.ResponseWriter, r *http.Request) error {
		var req request
		if err := json.NewDecoder(r.Body).Decode(&req); err != nil {
			http.Error(w, "Invalid request body", http.StatusBadRequest)
			return err
		}

		ctx := r.Context()

		userProfile := iam.UserProfileFromContext(ctx)

		err := h.videoCtrl.ShareVideo(ctx, req.VideoUrl, req.VideoTitle, userProfile.Email)
		if err != nil {
			switch err {
			case videos.ErrInvalidVideoUrl:
				http.Error(w, "Invalid YouTube video URL", http.StatusBadRequest)
			case videos.ErrUserNotFound:
				http.Error(w, "Target user does not exist", http.StatusBadRequest)
			default:
				http.Error(w, "Failed to share video", http.StatusInternalServerError)
			}
			return err
		}

		h.videoCtrl.Push(ctx, userProfile.Email)

		w.WriteHeader(http.StatusOK)
		return nil
	})
}
