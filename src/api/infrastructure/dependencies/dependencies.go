/*
Package dependencies implements service dependency injection within the application.
*/
package dependencies

import (
	"os"

	"github.com/proethics/mp-conciliation/src/api/config/database"
	"github.com/proethics/mp-conciliation/src/api/core/errors"
	"github.com/proethics/mp-conciliation/src/api/core/usecases/concilliation"
	"github.com/proethics/mp-conciliation/src/api/core/usecases/login"
	"github.com/proethics/mp-conciliation/src/api/core/usecases/process_payments_news"
	"github.com/proethics/mp-conciliation/src/api/entrypoints"
	"github.com/proethics/mp-conciliation/src/api/entrypoints/handlers/api"
	"github.com/proethics/mp-conciliation/src/api/entrypoints/handlers/consumer"
	"github.com/proethics/mp-conciliation/src/api/entrypoints/handlers/middlewares"
	"github.com/proethics/mp-conciliation/src/api/repositories/mercado_pago"
	"github.com/proethics/mp-conciliation/src/api/repositories/payments"
	"github.com/proethics/mp-conciliation/src/api/repositories/user"
)

type HandlerContainer struct {
	Search              entrypoints.Handler
	ProcessPaymentsNews entrypoints.Handler
	Concilliation       entrypoints.Handler
	Login               entrypoints.Handler
	AuthMiddleWare      entrypoints.Handler
}

type StartConnection struct {
	StoreConnection database.Connection
}

func (connections StartConnection) Start() *HandlerContainer {
	// Database
	storeClient, err := connections.StoreConnection.Connect()
	if err != nil {
		panic(errors.NewDependencyError(errors.ErrorDataBaseConnection.GetMessageWithParams(errors.Parameters{"cause": err.Error()})))
	}

	// Repositories
	mercadoPagoProvider := mercado_pago.Repository{
		PaymentsURL: os.Getenv("MERCADO_PAGO_PAYMENTS_URL"),
	}

	paymentsProvider := payments.Repository{
		StoreClient: storeClient,
	}

	userProvider := user.Repository{
		StoreClient: storeClient,
	}

	// UseCases
	processPaymentsNewsUseCase := process_payments_news.Implementation{
		MercadoPagoProvider: &mercadoPagoProvider,
		PaymentsProvider:    paymentsProvider,
	}

	concilliationUseCase := concilliation.Implementation{
		PaymentsProvider: paymentsProvider,
	}

	loginUseCase := login.Implementation{
		UserProvider: userProvider,
	}

	// Handlers
	handlers := HandlerContainer{}

	handlers.ProcessPaymentsNews = &consumer.PaymentsNews{
		ProcessPaymentsNewsUseCase: &processPaymentsNewsUseCase,
	}

	handlers.Search = &api.Search{
		PaymentsProvider: paymentsProvider,
	}

	handlers.Concilliation = &api.Concilliation{
		ConcilliationUseCase: &concilliationUseCase,
	}

	handlers.Login = &api.Login{
		LoginUseCase: &loginUseCase,
	}

	handlers.AuthMiddleWare = &middlewares.Auth{
		StoreClient: storeClient,
	}

	return &handlers
}
