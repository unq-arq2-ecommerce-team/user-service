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

func Test_GivenUpdateSellerCmdAndSellerIdAndSellerUpdate_WhenDo_ThenReturnNoError(t *testing.T) {
	updateSellerCmd, mocks := setUpSellerUpdateCmd(t)
	ctx := context.Background()
	sellerIdToFind := int64(4)
	sellerToUpdate := &model.Seller{
		Id:    sellerIdToFind,
		Name:  "pepeGrillo",
		Email: "pepe_grillo@gmail.com",
	}
	updateSellerInfo := model.UpdateSeller{
		Name:  "fafafa",
		Email: "fafafa@mail.com",
	}
	mocks.SellerRepo.EXPECT().FindById(ctx, sellerIdToFind).Return(sellerToUpdate, nil)
	sellerToUpdate.Merge(updateSellerInfo)
	mocks.SellerRepo.EXPECT().Update(ctx, *sellerToUpdate).Return(true, nil)

	err := updateSellerCmd.Do(ctx, sellerIdToFind, updateSellerInfo)
	assert.NoError(t, err)
}

func Test_GivenUpdateSellerCmdAndSellerIdAndSellerUpdateAndAnyErrorInFindSeller_WhenDo_ThenReturnThatError(t *testing.T) {
	updateSellerCmd, mocks := setUpSellerUpdateCmd(t)
	ctx := context.Background()
	sellerIdToFind := int64(4)
	updateSellerInfo := model.UpdateSeller{
		Name:  "fafafa",
		Email: "fafafa@mail.com",
	}
	mocks.SellerRepo.EXPECT().FindById(ctx, sellerIdToFind).Return(nil, exception.SellerNotFound{Id: sellerIdToFind})
	err := updateSellerCmd.Do(ctx, sellerIdToFind, updateSellerInfo)
	assert.ErrorIs(t, err, exception.SellerNotFound{Id: sellerIdToFind})
}

func Test_GivenUpdateSellerCmdAndSellerIdAndSellerUpdateAndAnyErrorInUpdateSeller_WhenDo_ThenReturnThatError(t *testing.T) {
	updateSellerCmd, mocks := setUpSellerUpdateCmd(t)
	ctx := context.Background()
	sellerIdToFind := int64(4)
	sellerToUpdate := &model.Seller{
		Id:    sellerIdToFind,
		Name:  "pepeGrillo",
		Email: "pepe_grillo@gmail.com",
	}
	updateSellerInfo := model.UpdateSeller{
		Name:  "fafafa",
		Email: "fafafa@mail.com",
	}
	mocks.SellerRepo.EXPECT().FindById(ctx, sellerIdToFind).Return(sellerToUpdate, nil)
	sellerToUpdate.Merge(updateSellerInfo)
	mocks.SellerRepo.EXPECT().Update(ctx, *sellerToUpdate).Return(false, exception.SellerAlreadyExist{Name: updateSellerInfo.Name})

	err := updateSellerCmd.Do(ctx, sellerIdToFind, updateSellerInfo)
	assert.ErrorIs(t, err, exception.SellerAlreadyExist{Name: updateSellerInfo.Name})
}

func setUpSellerUpdateCmd(t *testing.T) (*UpdateSeller, *mock.InterfaceMocks) {
	mocks := mock.NewInterfaceMocks(t)
	return NewUpdateSeller(mocks.SellerRepo, *query.NewFindSellerById(mocks.SellerRepo)), mocks
}
