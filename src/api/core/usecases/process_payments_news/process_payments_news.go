/*
Package process_payments_news implements golden gate consumer for audit table.
*/
package process_payments_news

import (
	"context"

	"github.com/proethics/mp-conciliation/src/api/core/providers"
	"github.com/proethics/mp-conciliation/src/api/infrastructure/logger"
)

type UseCase interface {
	Execute(ctx context.Context, payment_id int64) error
}

type Implementation struct {
	MercadoPagoProvider providers.MercadoPago
	PaymentsProvider    providers.Payments
}

func (useCase *Implementation) Execute(ctx context.Context, payment_id int64) error {
	ctx = context.WithValue(ctx, logger.MpConciliationKey{}, "process_payments_news_use_case")

	payment, err := useCase.MercadoPagoProvider.GetPayment(ctx, payment_id)
	if err != nil {
		return err
	}

	return useCase.PaymentsProvider.Save(ctx, payment)
}
