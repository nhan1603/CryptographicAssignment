package videos

import "errors"

var (
	// ErrUserNotFound means user not found
	ErrUserNotFound    = errors.New("user not found")
	ErrInvalidVideoUrl = errors.New("invalid video url")
)
