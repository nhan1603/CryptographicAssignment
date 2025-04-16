package videos

import (
	"context"

	"github.com/nhan1603/ReminoAssignment/api/internal/model"
)

func (i impl) GetSharedVideo(ctx context.Context) (model.ListSharedVideo, error) {
	return i.repo.Video().RetrieveSharedVideo(ctx)
}
