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

func Test_GivenDeleteSellerCmdAndSellerId_WhenDo_ThenReturnNoError(t *testing.T) {
	sellerDeleteCmd, mocks := setUpSellerDeleteCmd(t)
	ctx := context.Background()
	sellerId := int64(123)
	mocks.SellerRepo.EXPECT().FindById(ctx, sellerId).Return(&model.Seller{Id: sellerId}, nil)
	mocks.SellerRepo.EXPECT().Delete(ctx, sellerId).Return(true, nil)

	err := sellerDeleteCmd.Do(ctx, sellerId)

	assert.NoError(t, err)
}

func Test_GivenDeleteSellerCmdAndSellerIdAndSellerRepoDeleteError_WhenDo_ThenReturnThatError(t *testing.T) {
	sellerDeleteCmd, mocks := setUpSellerDeleteCmd(t)
	ctx := context.Background()
	sellerId := int64(123)
	mocks.SellerRepo.EXPECT().FindById(ctx, sellerId).Return(&model.Seller{Id: sellerId}, nil)
	mocks.SellerRepo.EXPECT().Delete(ctx, sellerId).Return(false, exception.SellerCannotDelete{Id: sellerId})

	err := sellerDeleteCmd.Do(ctx, sellerId)

	assert.ErrorIs(t, err, exception.SellerCannotDelete{Id: sellerId})
}

func Test_GivenDeleteSellerCmdAndSellerIdAndSellerRepoFindByIdError_WhenDo_ThenReturnThatError(t *testing.T) {
	sellerDeleteCmd, mocks := setUpSellerDeleteCmd(t)
	ctx := context.Background()
	sellerId := int64(123)
	mocks.SellerRepo.EXPECT().FindById(ctx, sellerId).Return(nil, exception.SellerNotFound{Id: sellerId})

	err := sellerDeleteCmd.Do(ctx, sellerId)

	assert.ErrorIs(t, err, exception.SellerNotFound{Id: sellerId})
}

func setUpSellerDeleteCmd(t *testing.T) (*DeleteSeller, *mock.InterfaceMocks) {
	mocks := mock.NewInterfaceMocks(t)
	return NewDeleteSeller(mocks.SellerRepo, *query.NewFindSellerById(mocks.SellerRepo)), mocks
}
