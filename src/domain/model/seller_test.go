package model

import (
	"github.com/stretchr/testify/assert"
	"testing"
)

func Test_Seller_String(t *testing.T) {
	seller1 := Seller{
		Name:  "pepe",
		Email: "pepegrillo@mail.com",
	}

	seller2 := Seller{
		Id:    2,
		Name:  "sarasa",
		Email: "sarasa@mail.com",
	}
	assert.Equal(t, `[Seller]{"id":0,"name":"pepe","email":"pepegrillo@mail.com"}`, seller1.String())
	assert.Equal(t, `[Seller]{"id":2,"name":"sarasa","email":"sarasa@mail.com"}`, seller2.String())
}

func Test_Seller_Merge(t *testing.T) {
	seller := Seller{
		Id:    5,
		Name:  "pepe",
		Email: "pepegrillo@mail.com",
	}
	seller.Merge(UpdateSeller{
		Name:  "apple station",
		Email: "apple@asd.com",
	})

	resultSeller := Seller{
		Id:    5,
		Name:  "apple station",
		Email: "apple@asd.com",
	}
	assert.Equal(t, resultSeller, seller)
}
