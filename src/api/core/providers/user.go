package providers

import (
	"context"

	"github.com/proethics/mp-conciliation/src/api/core/entities"
)

type User interface {
	Get(ctx context.Context, userName string) (*entities.User, error)
}
