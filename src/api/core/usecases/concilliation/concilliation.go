package concilliation

import (
	"context"
	"errors"

	"github.com/proethics/mp-conciliation/src/api/core/entities"
	"github.com/proethics/mp-conciliation/src/api/core/entities/constants"
	messages "github.com/proethics/mp-conciliation/src/api/core/errors"
	"github.com/proethics/mp-conciliation/src/api/core/providers"
	"github.com/proethics/mp-conciliation/src/api/infrastructure/logger"
)

type UseCase interface {
	Execute(ctx context.Context, payment_id int64, user entities.User) error
}

type Implementation struct {
	PaymentsProvider providers.Payments
}

func (useCase *Implementation) Execute(ctx context.Context, payment_id int64, user entities.User) error {
	ctx = context.WithValue(ctx, logger.MpConciliationKey{}, "concilliation_use_case")

	payment, err := useCase.PaymentsProvider.Get(ctx, payment_id)
	if err != nil {
		return err
	}

	if payment.ConcilliedStatus != constants.ConcilliedPending && !user.UserRole.IsAdmin() {
		tags := logger.Tags{"payment_id": payment_id, "user_name": user.UserName}
		logger.Error(ctx, messages.ErrorPaymentAlreadyConcillied.GetMessage(), tags)
		return errors.New(messages.ErrorPaymentAlreadyConcillied.GetMessage())
	}

	return useCase.PaymentsProvider.Update(ctx, payment, user.UserName)
}
