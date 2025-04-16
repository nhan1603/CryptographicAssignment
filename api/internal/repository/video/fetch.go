package video

import (
	"context"

	"github.com/nhan1603/ReminoAssignment/api/internal/model"
	"github.com/nhan1603/ReminoAssignment/api/internal/repository/dbmodel"
	pkgerrors "github.com/pkg/errors"
	"github.com/volatiletech/sqlboiler/v4/queries/qm"
)

func (i impl) RetrieveSharedVideo(ctx context.Context) (model.ListSharedVideo, error) {
	sharedVideos, err := dbmodel.SharedVideos(
		qm.OrderBy(dbmodel.SharedVideoColumns.SharedAt+" DESC"),
		qm.Load(dbmodel.SharedVideoRels.User),
	).All(ctx, i.dbConn)
	if err != nil {
		return nil, pkgerrors.WithStack(err)
	}

	var result model.ListSharedVideo
	for _, sharedVideo := range sharedVideos {
		if sharedVideo.R == nil || sharedVideo.R.User == nil {
			return nil, pkgerrors.WithStack(ErrNoUser)
		}
		user := sharedVideo.R.User
		result = append(result, model.SharedVideo{
			VideoUrl:    sharedVideo.YoutubeVideoID,
			VideoTitle:  sharedVideo.Title,
			SharerEmail: user.Email,
		})
	}

	return result, nil
}
