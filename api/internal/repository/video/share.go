package video

import (
	"context"
	"log"

	"github.com/nhan1603/CryptographicAssignment/api/internal/model"
	"github.com/nhan1603/CryptographicAssignment/api/internal/repository/dbmodel"
	"github.com/volatiletech/sqlboiler/v4/boil"

	pkgerrors "github.com/pkg/errors"
)

func (i impl) ShareVideo(ctx context.Context, sharingInput model.VideoShare) error {
	sharedVideo := dbmodel.SharedVideo{
		UserID:         sharingInput.UserID,
		YoutubeVideoID: sharingInput.VideoUrl,
		Title:          sharingInput.VideoTitle,
	}

	err := sharedVideo.Insert(ctx, i.dbConn, boil.Infer())
	if err != nil {
		log.Println(err)
		return pkgerrors.WithStack(err)
	}

	return nil
}
