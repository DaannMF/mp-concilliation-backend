package mercado_pago

import (
	"context"
	"encoding/json"
	"fmt"
	"io"
	"net/http"
	"os"

	"github.com/proethics/mp-conciliation/src/api/core/entities"
	"github.com/proethics/mp-conciliation/src/api/core/entities/constants"
	"github.com/proethics/mp-conciliation/src/api/core/errors"
	"github.com/proethics/mp-conciliation/src/api/infrastructure/logger"
)

type Repository struct {
	PaymentsURL string
}

func (repository Repository) GetPayment(ctx context.Context, paymentID int64) (*entities.Payment, error) {
	ctx = context.WithValue(ctx, logger.MpConciliationKey{}, "mercado_pago_repository")
	apiUrl := fmt.Sprintf("%s/%d", repository.PaymentsURL, paymentID)

	request, _ := http.NewRequest(http.MethodGet, apiUrl, nil)
	request.Header.Set("Content-Type", "application/json; charset=utf-8")
	request.Header.Set("Authorization", "Bearer "+os.Getenv("MP_ACCESS_TOKEN"))

	client := &http.Client{}
	restResponse, err := client.Do(request)

	if err != nil {
		return nil, err
	}

	if restResponse.StatusCode == http.StatusNotFound {
		tags := logger.Tags{"payment_id": paymentID, "resource_type": constants.Payment}
		logger.Error(ctx, errors.ErrorRecordNotFound.GetMessage(), tags)
		return nil, errors.NewRepositoryError(errors.ErrorRecordNotFound.GetMessage())
	}

	bytes, err := io.ReadAll(restResponse.Body)
	if err != nil {
		return nil, err
	}

	var response response
	err = json.Unmarshal(bytes, &response)
	if err != nil {
		tags := logger.Tags{"err": err.Error()}
		logger.Error(ctx, errors.ErrorBindingRequest.GetMessage(), tags)
		return nil, err
	}

	restResponse.Body.Close()
	entity := response.GetEntity()

	return &entity, nil
}
