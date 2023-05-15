package command

import (
	"context"
	"github.com/stretchr/testify/assert"
	"github.com/unq-arq2-ecommerce-team/users-service/src/domain/mock"
	"github.com/unq-arq2-ecommerce-team/users-service/src/domain/model"
	"github.com/unq-arq2-ecommerce-team/users-service/src/domain/model/exception"
	"testing"
)

func Test_GivenCreateCustomerCmdAndNewCustomer_WhenDo_ThenReturnNoErrorAndANewId(t *testing.T) {
	createCustomerCmd, mocks := setUpCustomerCreateCmd(t)
	ctx := context.Background()
	customer := model.Customer{
		Firstname: "pepe",
		Lastname:  "grillo",
		Email:     "pepegrillo@mail.com",
	}
	newCustomerId := int64(15)
	mocks.CustomerRepo.EXPECT().Create(ctx, customer).Return(newCustomerId, nil)

	resCustomerId, err := createCustomerCmd.Do(ctx, customer)

	assert.Equal(t, newCustomerId, resCustomerId)
	assert.NoError(t, err)
}

func Test_GivenCreateCustomerCmdAndNewCustomerAndCustomerRepoCreateError_WhenDo_ThenReturnThatError(t *testing.T) {
	createCustomerCmd, mocks := setUpCustomerCreateCmd(t)
	ctx := context.Background()
	customer := model.Customer{
		Firstname: "pepe",
		Lastname:  "grillo",
		Email:     "pepegrillo@mail.com",
	}
	mocks.CustomerRepo.EXPECT().Create(ctx, customer).Return(int64(0), exception.CustomerAlreadyExist{Email: customer.Email})

	resCustomerId, err := createCustomerCmd.Do(ctx, customer)

	assert.Equal(t, int64(0), resCustomerId)
	assert.ErrorIs(t, err, exception.CustomerAlreadyExist{Email: customer.Email})
}

func setUpCustomerCreateCmd(t *testing.T) (*CreateCustomer, *mock.InterfaceMocks) {
	mocks := mock.NewInterfaceMocks(t)
	return NewCreateCustomer(mocks.CustomerRepo), mocks
}
