package api

import (
	"fmt"
	"github.com/gin-gonic/gin"
	swaggerFiles "github.com/swaggo/files"
	ginSwagger "github.com/swaggo/gin-swagger"
	swaggerDocs "github.com/unq-arq2-ecommerce-team/users-service/docs"
	"github.com/unq-arq2-ecommerce-team/users-service/src/domain/action/command"
	"github.com/unq-arq2-ecommerce-team/users-service/src/domain/action/query"
	"github.com/unq-arq2-ecommerce-team/users-service/src/domain/model"
	"github.com/unq-arq2-ecommerce-team/users-service/src/infrastructure/api/middleware"
	"github.com/unq-arq2-ecommerce-team/users-service/src/infrastructure/api/v1"
	"github.com/unq-arq2-ecommerce-team/users-service/src/infrastructure/config"
	"io"
	"net/http"
)

// Application
// @title users-service API
// @version 1.0
// @description api for tp arq2
// @contact.name API SUPPORT
// @contact.url http://www.swagger.io/support
// @contact.email support@swagger.io
// @license.name MIT
// @license.url https://opensource.org/licenses/MIT
// @BasePath /
// @query.collection.format multi
type Application interface {
	Run() error
}

type application struct {
	logger model.Logger
	config config.Config
	*ApplicationUseCases
}

type ApplicationUseCases struct {
	//customer
	CreateCustomerCmd *command.CreateCustomer
	UpdateCustomerCmd *command.UpdateCustomer
	DeleteCustomerCmd *command.DeleteCustomer
	FindCustomerQuery *query.FindCustomerById
	//seller
	CreateSellerCmd *command.CreateSeller
	UpdateSellerCmd *command.UpdateSeller
	DeleteSellerCmd *command.DeleteSeller
	FindSellerQuery *query.FindSellerById
}

func NewApplication(l model.Logger, conf config.Config, applicationUseCases *ApplicationUseCases) Application {
	return &application{
		logger:              l,
		config:              conf,
		ApplicationUseCases: applicationUseCases,
	}
}

func (app *application) Run() error {
	swaggerDocs.SwaggerInfo.Host = fmt.Sprintf("localhost:%v", app.config.Port)

	gin.SetMode(gin.ReleaseMode)
	gin.DefaultWriter = io.Discard

	router := gin.Default()
	router.GET("/", HealthCheck)

	rv1 := router.Group("/api/v1")
	rv1.Use(middleware.TracingRequestId())
	{
		rv1Customer := rv1.Group("/customer")
		rv1Customer.POST("", v1.CreateCustomerHandler(app.logger, app.CreateCustomerCmd))
		rv1Customer.GET("/:customerId", v1.FindCustomerHandler(app.logger, app.FindCustomerQuery))
		rv1Customer.DELETE("/:customerId", v1.DeleteCustomerHandler(app.logger, app.DeleteCustomerCmd))
		rv1Customer.PUT("/:customerId", v1.UpdateCustomerHandler(app.logger, app.UpdateCustomerCmd))
	}
	{
		rv1Seller := rv1.Group("/seller")
		rv1Seller.POST("", v1.CreateSellerHandler(app.logger, app.CreateSellerCmd))
		rv1Seller.GET("/:sellerId", v1.FindSellerHandler(app.logger, app.FindSellerQuery))
		rv1Seller.DELETE("/:sellerId", v1.DeleteSellerHandler(app.logger, app.DeleteSellerCmd))
		rv1Seller.PUT("/:sellerId", v1.UpdateSellerHandler(app.logger, app.UpdateSellerCmd))
	}

	router.GET("/docs/*any", ginSwagger.WrapHandler(swaggerFiles.Handler))

	app.logger.Infof("running http server on port %d", app.config.Port)
	return router.Run(fmt.Sprintf(":%d", app.config.Port))
}

// HealthCheck
// @Summary Show the status of server.
// @Description get the status of server.
// @Tags Health check
// @Accept */*
// @Produce json
// @Success 200 {object} HealthCheckRes
// @Router / [get]
func HealthCheck(c *gin.Context) {
	c.JSON(http.StatusOK, HealthCheckRes{Data: "Server is up and running"})
}

type HealthCheckRes struct {
	Data string `json:"data"`
}
