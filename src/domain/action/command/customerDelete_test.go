package command

import (
	"context"
	"github.com/stretchr/testify/assert"
	"github.com/unq-arq2-ecommerce-team/users-service/src/domain/action/query"
	"github.com/unq-arq2-ecommerce-team/users-service/src/domain/mock"
	"github.com/unq-arq2-ecommerce-team/users-service/src/domain/model"
	"github.com/unq-arq2-ecommerce-team/users-service/src/domain/model/exception"
	"testing"
)

func Test_GivenDeleteCustomerCmdAndCustomerId_WhenDo_ThenReturnNoError(t *testing.T) {
	customerDeleteCmd, mocks := setUpCustomerDeleteCmd(t)
	ctx := context.Background()
	customerId := int64(123)
	mocks.CustomerRepo.EXPECT().FindById(ctx, customerId).Return(&model.Customer{Id: customerId}, nil)
	mocks.CustomerRepo.EXPECT().Delete(ctx, customerId).Return(true, nil)

	err := customerDeleteCmd.Do(ctx, customerId)

	assert.NoError(t, err)
}

func Test_GivenDeleteCustomerCmdAndCustomerIdAndCustomerRepoDeleteError_WhenDo_ThenReturnThatError(t *testing.T) {
	customerDeleteCmd, mocks := setUpCustomerDeleteCmd(t)
	ctx := context.Background()
	customerId := int64(123)
	mocks.CustomerRepo.EXPECT().FindById(ctx, customerId).Return(&model.Customer{Id: customerId}, nil)
	mocks.CustomerRepo.EXPECT().Delete(ctx, customerId).Return(false, exception.CustomerCannotDelete{Id: customerId})

	err := customerDeleteCmd.Do(ctx, customerId)

	assert.ErrorIs(t, err, exception.CustomerCannotDelete{Id: customerId})
}

func Test_GivenDeleteCustomerCmdAndCustomerIdAndCustomerRepoFindByIdError_WhenDo_ThenReturnThatError(t *testing.T) {
	customerDeleteCmd, mocks := setUpCustomerDeleteCmd(t)
	ctx := context.Background()
	customerId := int64(123)
	mocks.CustomerRepo.EXPECT().FindById(ctx, customerId).Return(nil, exception.CustomerNotFound{Id: customerId})

	err := customerDeleteCmd.Do(ctx, customerId)

	assert.ErrorIs(t, err, exception.CustomerNotFound{Id: customerId})
}

func setUpCustomerDeleteCmd(t *testing.T) (*DeleteCustomer, *mock.InterfaceMocks) {
	mocks := mock.NewInterfaceMocks(t)
	return NewDeleteCustomer(mocks.CustomerRepo, *query.NewFindCustomerById(mocks.CustomerRepo)), mocks
}
