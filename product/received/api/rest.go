package api

import (
	"encoding/json"
	"github.com/gin-gonic/gin"
	"go.mongodb.org/mongo-driver/bson"
	"io"
	"net/http"
	"product/domain"
)

type RestHandler struct {
	ServiceInterface domain.ServiceInterface
}

func NewRestApi(serviceInterface domain.ServiceInterface) *RestHandler {
	return &RestHandler{ServiceInterface: serviceInterface}
}

func (h RestHandler) InsertProduct(c *gin.Context) {
	//get request body
	body, err := io.ReadAll(c.Request.Body)
	if err != nil {
		_err := domain.SetError(domain.InvalidationErr, err.Error())
		c.JSON(http.StatusBadRequest, bson.M{"data": nil, "error": _err, "message": ""})
		return
	}

	//request body validation
	product := &domain.Products{}
	err = json.Unmarshal(body, product)
	if err != nil {
		_err := domain.SetError(domain.InvalidationErr, err.Error())
		c.JSON(http.StatusBadRequest, bson.M{"data": nil, "error": _err, "message": ""})
		return
	}

	insertedID, resErr := h.ServiceInterface.InsertProductService(product)
	if resErr != nil {
		c.JSON(http.StatusInternalServerError, bson.M{"data": nil, "error": resErr, "message": ""})
		return
	}
	c.JSON(http.StatusCreated, bson.M{"data": insertedID, "error": nil, "message": "inserted"})
}
func (h RestHandler) UpdateProduct(c *gin.Context) {
	//get request body
	body, err := io.ReadAll(c.Request.Body)
	if err != nil {
		_err := domain.SetError(domain.InvalidationErr, err.Error())
		c.JSON(http.StatusBadRequest, bson.M{"data": nil, "error": _err, "message": ""})
		return
	}

	//request body validation
	attributes := &domain.ProductAttributes{}
	err = json.Unmarshal(body, attributes)
	if err != nil {
		_err := domain.SetError(domain.InvalidationErr, err.Error())
		c.JSON(http.StatusBadRequest, bson.M{"data": nil, "error": _err, "message": ""})
		return
	}

	productID := c.Param("product-id")

	resErr := h.ServiceInterface.UpdateProductService(productID, attributes)
	if resErr != nil {
		c.JSON(http.StatusInternalServerError, bson.M{"data": nil, "error": resErr, "message": ""})
		return
	}
	c.JSON(http.StatusOK, bson.M{"data": nil, "error": nil, "message": "updated"})
}
func (h RestHandler) GetProductByID(c *gin.Context) {
	productID := c.Param("product-id")

	product, resErr := h.ServiceInterface.GetProductByIDService(productID)
	if resErr != nil {
		c.JSON(http.StatusInternalServerError, bson.M{"data": nil, "error": resErr, "message": ""})
		return
	}
	c.JSON(http.StatusOK, bson.M{"data": product, "error": nil, "message": ""})
}
func (h RestHandler) GetAllProducts(c *gin.Context) {
	skip, limit := domain.GetSkipLimitFromQuery(c)

	product, resErr := h.ServiceInterface.GetAllProductsService(skip, limit)
	if resErr != nil {
		c.JSON(http.StatusInternalServerError, bson.M{"data": nil, "error": resErr, "message": ""})
		return
	}
	c.JSON(http.StatusOK, bson.M{"data": product, "error": nil, "message": ""})
}
