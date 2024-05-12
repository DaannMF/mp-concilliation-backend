package process_payments_news

import (
	"time"

	"github.com/proethics/mp-conciliation/src/api/core/entities/constants"
)

type Request struct {
	ID          string              `json:"id"`
	LiveMode    bool                `json:"live_mode"`
	Type        constants.Event     `json:"type"`
	DateCreated time.Time           `json:"date_created"`
	UserID      int64               `json:"user_id"`
	APIVersion  string              `json:"api_version"`
	Action      constants.EventType `json:"action"`
	Data        struct {
		PaymentID string `json:"id"`
	} `json:"data"`
}
