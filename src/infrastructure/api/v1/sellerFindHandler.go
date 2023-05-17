package v1

import (
	"github.com/gin-gonic/gin"
	"github.com/unq-arq2-ecommerce-team/users-service/src/domain/action/query"
	"github.com/unq-arq2-ecommerce-team/users-service/src/domain/model"
	"github.com/unq-arq2-ecommerce-team/users-service/src/domain/model/exception"
	"net/http"
)

// FindSellerHandler
// @Summary      Endpoint find seller
// @Description  find seller
// @Param sellerId path int true "Seller ID" minimum(1)
// @Tags         Seller
// @Produce json
// @Success 200 {object} model.Seller
// @Success 400 {object} dto.ErrorMessage
// @Failure 404 {object} dto.ErrorMessage
// @Router       /api/v1/seller/{sellerId} [get]
func FindSellerHandler(log model.Logger, findSellerByIdQuery *query.FindSellerById) gin.HandlerFunc {
	return func(c *gin.Context) {
		id, err := parsePathParamPositiveIntId(c, "sellerId")
		if err != nil {
			log.WithFields(model.LoggerFields{"error": err}).Error("invalid path param")
			writeJsonErrorMessageWithNoDesc(c, http.StatusBadRequest, err)
			return
		}
		seller, err := findSellerByIdQuery.Do(c.Request.Context(), id)
		if err != nil {
			switch err.(type) {
			case exception.SellerNotFound:
				writeJsonErrorMessageWithNoDesc(c, http.StatusNotFound, err)
			default:
				defaultInternalServerError(log, c, "uncaught error when find seller", err)
			}
			return
		}
		c.JSON(http.StatusOK, seller)
	}
}
