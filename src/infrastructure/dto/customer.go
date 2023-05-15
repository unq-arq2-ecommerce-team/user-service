package dto

import "github.com/unq-arq2-ecommerce-team/users-service/src/domain/model"

type CustomerCreateReq struct {
	Firstname string `json:"firstname" binding:"required"`
	Lastname  string `json:"lastname" binding:"required"`
	Email     string `json:"email" binding:"required,email"`
}

func (req *CustomerCreateReq) MapToModel() model.Customer {
	return model.Customer{
		Firstname: req.Firstname,
		Lastname:  req.Lastname,
		Email:     req.Email,
	}
}
