package payment

import (
	"context"

	"github.com/nhan1603/CryptographicAssignment/api/internal/model"
	"github.com/nhan1603/CryptographicAssignment/api/internal/repository/dbmodel"
	"github.com/volatiletech/null/v8"
	"github.com/volatiletech/sqlboiler/v4/boil"

	pkgerrors "github.com/pkg/errors"
)

func (r impl) Create(ctx context.Context, payment model.PayPalTransaction) error {
	paypalTrasaction := dbmodel.PaypalTransaction{
		OrderID:             int(payment.OrderID),
		PaypalTransactionID: payment.PayPalTransactionID,
		PaymentStatus:       payment.PaymentStatus,
		PaymentAmount:       payment.PaymentAmount,
		PayerEmail:          null.StringFrom(payment.PayerEmail),
	}
	err := paypalTrasaction.Insert(ctx, r.dbConn, boil.Infer())

	if err != nil {
		pkgerrors.WithStack(err)
	}
	return nil
}
