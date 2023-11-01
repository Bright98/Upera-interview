package api

import (
	"github.com/gin-gonic/gin"
	"go.mongodb.org/mongo-driver/bson"
	"net/http"
	"revision/domain"
)

type RestHandler struct {
	ServiceInterface domain.ServiceInterface
}

func NewRestApi(serviceInterface domain.ServiceInterface) *RestHandler {
	return &RestHandler{ServiceInterface: serviceInterface}
}

func (h *RestHandler) GetRevisionByProductIDAndNo(c *gin.Context) {
	productID := c.Param("product-id")
	versionNo := c.Param("version-no")

	revision, resErr := h.ServiceInterface.GetRevisionByProductIDAndNoService(productID, versionNo)
	if resErr != nil {
		c.JSON(http.StatusInternalServerError, bson.M{"data": nil, "error": resErr, "message": ""})
		return
	}
	c.JSON(http.StatusOK, bson.M{"data": revision, "error": nil, "message": ""})
}
func (h *RestHandler) GetAllRevisionsOfOneProduct(c *gin.Context) {
	productID := c.Param("product-id")
	skip, limit := domain.GetSkipLimitFromQuery(c)

	revision, resErr := h.ServiceInterface.GetAllRevisionsOfOneProductService(skip, limit, productID)
	if resErr != nil {
		c.JSON(http.StatusInternalServerError, bson.M{"data": nil, "error": resErr, "message": ""})
		return
	}
	c.JSON(http.StatusOK, bson.M{"data": revision, "error": nil, "message": ""})
}
