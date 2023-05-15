package v1

import (
	"github.com/gin-gonic/gin"
	"github.com/unq-arq2-ecommerce-team/users-service/src/domain/action/query"
	"github.com/unq-arq2-ecommerce-team/users-service/src/domain/model"
	"github.com/unq-arq2-ecommerce-team/users-service/src/domain/model/exception"
	"github.com/unq-arq2-ecommerce-team/users-service/src/infrastructure/logger"
	"net/http"
)

// FindCustomerHandler
// @Summary      Endpoint find customer
// @Description  find customer
// @Param customerId path int true "Customer ID" minimum(1)
// @Tags         Customer
// @Produce json
// @Success 200 {object} model.Customer
// @Success 400 {object} dto.ErrorMessage
// @Failure 404 {object} dto.ErrorMessage
// @Router       /api/v1/customer/{customerId} [get]
func FindCustomerHandler(log model.Logger, findCustomerQuery *query.FindCustomerById) gin.HandlerFunc {
	return func(c *gin.Context) {
		id, err := parsePathParamPositiveIntId(c, "customerId")
		if err != nil {
			log.WithFields(logger.Fields{"error": err}).Error("invalid path param")
			writeJsonErrorMessageWithNoDesc(c, http.StatusBadRequest, err)
			return
		}
		customer, err := findCustomerQuery.Do(c.Request.Context(), id)
		if err != nil {
			switch err.(type) {
			case exception.CustomerNotFound:
				writeJsonErrorMessageWithNoDesc(c, http.StatusNotFound, err)
			default:
				defaultInternalServerError(log, c, "uncaught error when find customer", err)
			}
			return
		}
		c.JSON(http.StatusOK, customer)
	}
}
