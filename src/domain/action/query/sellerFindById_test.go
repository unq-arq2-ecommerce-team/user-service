package query

import (
	"context"
	"fmt"
	"github.com/stretchr/testify/assert"
	"github.com/unq-arq2-ecommerce-team/users-service/src/domain/mock"
	"github.com/unq-arq2-ecommerce-team/users-service/src/domain/model"
	"github.com/unq-arq2-ecommerce-team/users-service/src/domain/model/exception"
	"testing"
)

func Test_GivenFindSellerByIdQueryAndSellerThatExistWithSomeId_WhenDoWithSameId_ThenReturnSellerWithSameIdAndNoError(t *testing.T) {
	findSellerById, mocks := setUpFindSellerById(t)
	ctx := context.Background()
	sellerIdToFind := int64(4)
	sellerToFind := &model.Seller{
		Id: sellerIdToFind,
	}
	mocks.SellerRepo.EXPECT().FindById(ctx, sellerIdToFind).Return(sellerToFind, nil)
	seller, err := findSellerById.Do(ctx, sellerIdToFind)

	assert.Equal(t, sellerToFind, seller)
	assert.NoError(t, err)
}

func Test_GivenFindSellerByIdQuery_WhenDoWithId_ThenReturnNoSellerAndAnUnexpectedError(t *testing.T) {
	findSellerById, mocks := setUpFindSellerById(t)
	ctx := context.Background()
	sellerIdToFind := int64(4)
	errMsg := "unexpected error"
	mocks.SellerRepo.EXPECT().FindById(ctx, sellerIdToFind).Return(nil, fmt.Errorf(errMsg))

	seller, err := findSellerById.Do(ctx, sellerIdToFind)

	assert.Nil(t, seller)
	assert.EqualError(t, err, errMsg)
}

func Test_GivenFindSellerByIdQuery_WhenDoWithIdThatNotExists_ThenReturnNoSellerAndSellerNotFoundErrorWithThatId(t *testing.T) {
	findSellerById, mocks := setUpFindSellerById(t)
	ctx := context.Background()
	sellerIdToFind := int64(999)
	mocks.SellerRepo.EXPECT().FindById(ctx, sellerIdToFind).Return(nil, exception.SellerNotFound{Id: sellerIdToFind})

	seller, err := findSellerById.Do(ctx, sellerIdToFind)

	assert.Nil(t, seller)
	assert.ErrorIs(t, err, exception.SellerNotFound{Id: sellerIdToFind})
}

func setUpFindSellerById(t *testing.T) (*FindSellerById, *mock.InterfaceMocks) {
	mocks := mock.NewInterfaceMocks(t)
	return NewFindSellerById(mocks.SellerRepo), mocks
}
