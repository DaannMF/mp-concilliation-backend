package providers

import (
	"context"

	"github.com/proethics/mp-conciliation/src/api/core/entities"
	"github.com/proethics/mp-conciliation/src/api/core/entities/constants"
)

type Payments interface {
	Get(ctx context.Context, payment_id int64) (*entities.Payment, error)
	GetPaymentsByStatus(ctx context.Context, status constants.ConcilliedStatus) ([]entities.Payment, error)
	Save(ctx context.Context, payment *entities.Payment) error
	Update(ctx context.Context, payment *entities.Payment, userName string) error
}
