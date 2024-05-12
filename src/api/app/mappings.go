package app

import (
	"github.com/gin-gonic/gin"
	"github.com/proethics/mp-conciliation/src/api/infrastructure/dependencies"
)

func configureMappings(router *gin.Engine, handlers *dependencies.HandlerContainer) {
	configureApiMappings(router, handlers)
	configureConsumersMappings(router, handlers)
}

func configureApiMappings(router *gin.Engine, handlers *dependencies.HandlerContainer) {
	groupInternal := router.Group("/internal")

	groupConcilliation := groupInternal.Group("/concilliation")
	groupConcilliation.GET("search", handlers.AuthMiddleWare.Handle, handlers.Search.Handle)
	groupConcilliation.PUT(":payment_id/save", handlers.AuthMiddleWare.Handle, handlers.Concilliation.Handle)

	groupAuth := groupInternal.Group("/auth")
	groupAuth.POST("login", handlers.Login.Handle)
}

func configureConsumersMappings(router *gin.Engine, handlers *dependencies.HandlerContainer) {
	groupConsumers := router.Group("/consumers")
	groupConsumers.POST("payments/news", handlers.ProcessPaymentsNews.Handle)
}
