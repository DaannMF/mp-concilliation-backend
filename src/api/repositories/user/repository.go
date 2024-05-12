package user

import (
	"context"
	"errors"

	"github.com/proethics/mp-conciliation/src/api/core/entities"
	"github.com/proethics/mp-conciliation/src/api/core/entities/constants"
	errorsMessages "github.com/proethics/mp-conciliation/src/api/core/errors"
	"github.com/proethics/mp-conciliation/src/api/infrastructure/logger"
	"gorm.io/gorm"
)

type Repository struct {
	StoreClient *gorm.DB
}

func (repository Repository) Get(ctx context.Context, userName string) (*entities.User, error) {
	ctx = context.WithValue(ctx, logger.MpConciliationKey{}, "user_repository")

	var user entities.User
	err := repository.StoreClient.Where(&entities.User{
		UserName: userName,
	}).First(&user).Error

	if err != nil {
		tags := logger.Tags{"err": err.Error(), "resource": constants.User.String(), "user_name": userName}
		if errors.Is(err, gorm.ErrRecordNotFound) {
			logger.Error(ctx, errorsMessages.ErrorRecordNotFound.GetMessage(), tags)
			return nil, errorsMessages.NewRepositoryError(errorsMessages.ErrorRecordNotFound.GetMessage())
		}

		logger.Error(ctx, errorsMessages.ErrorGettingResource.GetMessage(), tags)
		return nil, errorsMessages.NewRepositoryError(errorsMessages.ErrorGettingResource.GetMessage())
	}

	return &user, nil
}
