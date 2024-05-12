package providers

import (
	"context"

	"github.com/proethics/mp-conciliation/src/api/core/entities"
)

type Payments interface {
	Get(ctx context.Context, payment_id int64) (*entities.Payment, error)
	GetPendingPayments(ctx context.Context) ([]entities.Payment, error)
	Save(ctx context.Context, payment *entities.Payment) error
	Update(ctx context.Context, payment *entities.Payment, userName string) error
}
