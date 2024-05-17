package payments

import (
	"context"
	"errors"
	"time"

	"github.com/proethics/mp-conciliation/src/api/core/entities"
	"github.com/proethics/mp-conciliation/src/api/core/entities/constants"
	errorsMessages "github.com/proethics/mp-conciliation/src/api/core/errors"
	"github.com/proethics/mp-conciliation/src/api/infrastructure/logger"
	"gorm.io/gorm"
)

type Repository struct {
	StoreClient *gorm.DB
}

func (repository Repository) Get(ctx context.Context, id int64) (*entities.Payment, error) {
	ctx = context.WithValue(ctx, logger.MpConciliationKey{}, "payments_repository")

	var payment entities.Payment
	err := repository.StoreClient.Preload("Payer").Find(&payment, id).Error

	if err != nil {
		tags := logger.Tags{"err": err.Error(), "resource": constants.Payment.String(), "payment_id": id}
		if errors.Is(err, gorm.ErrRecordNotFound) {
			logger.Error(ctx, errorsMessages.ErrorRecordNotFound.GetMessage(), tags)
			return nil, errorsMessages.NewRepositoryError(errorsMessages.ErrorRecordNotFound.GetMessage())
		}

		logger.Error(ctx, errorsMessages.ErrorGettingResource.GetMessage(), tags)
		return nil, errorsMessages.NewRepositoryError(errorsMessages.ErrorGettingResource.GetMessage())
	}

	return &payment, nil
}

func (repository Repository) GetPaymentsByStatus(ctx context.Context, status constants.ConcilliedStatus) ([]entities.Payment, error) {
	ctx = context.WithValue(ctx, logger.MpConciliationKey{}, "payments_repository")

	var payments []entities.Payment
	err := repository.StoreClient.Where(&entities.Payment{
		ConcilliedStatus: status,
	}).Preload("Payer").Find(&payments).Error

	if err != nil {
		tags := logger.Tags{"err": err.Error(), "resource": constants.Payment.String()}
		logger.Error(ctx, errorsMessages.ErrorGettingResource.GetMessage(), tags)
		return nil, errorsMessages.NewRepositoryError(errorsMessages.ErrorGettingResource.GetMessage())
	}

	return payments, nil
}

func (repository Repository) Save(ctx context.Context, payment *entities.Payment) error {
	ctx = context.WithValue(ctx, logger.MpConciliationKey{}, "payments_repository")
	var err error

	db := repository.StoreClient.Begin()
	defer func() {
		if r := recover(); r != nil {
			logger.Error(ctx, errorsMessages.ErrorRecoverFunction.GetMessageWithParams(
				errorsMessages.Parameters{"resource": constants.Payment.String()}), logger.Tags{})
			db.Rollback()
		}
	}()

	if err := db.Error; err != nil {
		tags := logger.Tags{"resource": constants.Payment.String()}
		logger.Error(ctx, errorsMessages.ErrorBeginTransaction.GetMessage(), tags)
		return errorsMessages.NewRepositoryError(errorsMessages.ErrorBeginTransaction.GetMessage())
	}

	payment.ConcilliedStatus = constants.ConcilliedPending
	err = db.Create(payment).Error

	if err != nil {
		db.Rollback()
		tags := logger.Tags{"err": err.Error(), "resource": constants.Payment.String()}
		logger.Error(ctx, errorsMessages.ErrorCreatingResource.GetMessage(), tags)
		return errorsMessages.NewRepositoryError(errorsMessages.ErrorCreatingResource.GetMessage())
	}

	db.Commit()
	if err := db.Error; err != nil {
		errorMessage := errorsMessages.ErrorCommitTransaction.GetMessageWithParams(
			errorsMessages.Parameters{"resource": constants.Payment.String()})
		logger.Error(ctx, errorMessage, logger.Tags{})
		return errorsMessages.NewRepositoryError(errorMessage)
	}

	return nil
}

func (repository Repository) Update(ctx context.Context, payment *entities.Payment, userName string) error {
	ctx = context.WithValue(ctx, logger.MpConciliationKey{}, "payments_repository")
	var err error

	db := repository.StoreClient.Begin()
	defer func() {
		if r := recover(); r != nil {
			logger.Error(ctx, errorsMessages.ErrorRecoverFunction.GetMessageWithParams(
				errorsMessages.Parameters{"resource": constants.Payment.String()}), logger.Tags{})
			db.Rollback()
		}
	}()

	if err := db.Error; err != nil {
		tags := logger.Tags{"resource": constants.Payment.String()}
		logger.Error(ctx, errorsMessages.ErrorBeginTransaction.GetMessage(), tags)
		return errorsMessages.NewRepositoryError(errorsMessages.ErrorBeginTransaction.GetMessage())
	}

	now := time.Now()
	err = db.Model(&payment).Updates(entities.Payment{
		ConcilliedStatus: constants.Concillied,
		ConcilliedDate:   &now,
		ConcilliedUser:   &userName,
	}).Error

	if err != nil {
		db.Rollback()
		tags := logger.Tags{"err": err.Error(), "resource": constants.Payment.String()}
		logger.Error(ctx, errorsMessages.ErrorUpdatingResource.GetMessage(), tags)
		return errorsMessages.NewRepositoryError(errorsMessages.ErrorCreatingResource.GetMessage())
	}

	db.Commit()
	if err := db.Error; err != nil {
		errorMessage := errorsMessages.ErrorCommitTransaction.GetMessageWithParams(
			errorsMessages.Parameters{"resource": constants.Payment.String()})
		logger.Error(ctx, errorMessage, logger.Tags{})
		return errorsMessages.NewRepositoryError(errorMessage)
	}

	return nil
}
