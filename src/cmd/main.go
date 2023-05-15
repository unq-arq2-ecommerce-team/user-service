package main

import (
	"context"
	"github.com/unq-arq2-ecommerce-team/users-service/src/domain/action/command"
	"github.com/unq-arq2-ecommerce-team/users-service/src/domain/action/query"
	"github.com/unq-arq2-ecommerce-team/users-service/src/infrastructure/api"
	"github.com/unq-arq2-ecommerce-team/users-service/src/infrastructure/config"
	"github.com/unq-arq2-ecommerce-team/users-service/src/infrastructure/logger"
	"github.com/unq-arq2-ecommerce-team/users-service/src/infrastructure/repository/mongo"
)

func main() {
	conf := config.LoadConfig()
	baseLogger := logger.New(&logger.Config{
		ServiceName:     "users-service",
		EnvironmentName: conf.Environment,
		LogLevel:        conf.LogLevel,
		LogFormat:       logger.JsonFormat,
	})
	mongoDB := mongo.Connect(context.Background(), baseLogger, conf.MongoURI, conf.MongoDatabase)

	//repositories
	customerRepo := mongo.NewCustomerRepository(baseLogger, mongoDB, conf.MongoTimeout)
	sellerRepo := mongo.NewSellerRepository(baseLogger, mongoDB, conf.MongoTimeout)

	//customer
	findCustomerByIdQuery := query.NewFindCustomerById(customerRepo)
	createCustomerCmd := command.NewCreateCustomer(customerRepo)
	updateCustomerCmd := command.NewUpdateCustomer(customerRepo, *findCustomerByIdQuery)
	deleteCustomerCmd := command.NewDeleteCustomer(customerRepo, *findCustomerByIdQuery)

	//seller
	findSellerByIdQuery := query.NewFindSellerById(sellerRepo)
	createSellerCmd := command.NewCreateSeller(sellerRepo)
	updateSellerCmd := command.NewUpdateSeller(sellerRepo, *findSellerByIdQuery)
	deleteSellerCmd := command.NewDeleteSeller(sellerRepo, *findSellerByIdQuery)

	app := api.NewApplication(baseLogger, conf, &api.ApplicationUseCases{
		FindCustomerQuery: findCustomerByIdQuery,
		CreateCustomerCmd: createCustomerCmd,
		UpdateCustomerCmd: updateCustomerCmd,
		DeleteCustomerCmd: deleteCustomerCmd,

		FindSellerQuery: findSellerByIdQuery,
		CreateSellerCmd: createSellerCmd,
		UpdateSellerCmd: updateSellerCmd,
		DeleteSellerCmd: deleteSellerCmd,
	})
	baseLogger.Fatal(app.Run())
}
