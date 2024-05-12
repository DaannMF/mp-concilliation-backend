package providers

import (
	"context"

	"github.com/proethics/mp-conciliation/src/api/core/entities"
)

type MercadoPago interface {
	GetPayment(ctx context.Context, payment_id int64) (*entities.Payment, error)
}
