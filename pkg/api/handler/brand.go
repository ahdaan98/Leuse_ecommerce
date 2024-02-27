package handler

import (
	"net/http"
	"strconv"

	interfaces "github.com/ahdaan98/pkg/usecase/interface"
	"github.com/ahdaan98/pkg/utils/models"
	"github.com/ahdaan98/pkg/utils/response"
	"github.com/gin-gonic/gin"
)

type BrandHandler struct {
	usecase interfaces.BrandUseCase
}

func NewBrandHandler(usecase interfaces.BrandUseCase) *BrandHandler {
	return &BrandHandler{
		usecase: usecase,
	}
}

func (br *BrandHandler) AddBrand(c *gin.Context) {
	var brandName models.AddBrand

	if err := c.ShouldBindJSON(&brandName); err != nil {
		errRes := response.ClientResponse(http.StatusBadRequest, "Fields are in wrong format", nil, err.Error())
		c.JSON(http.StatusBadRequest, errRes)
		return
	}

	Brand, err := br.usecase.AddBrand(brandName)
	if err != nil {
		errRes := response.ClientResponse(http.StatusBadRequest, "failed to add brand", nil, err.Error())
		c.JSON(http.StatusBadRequest, errRes)
		return
	}

	successRes := response.ClientResponse(http.StatusOK, "successfully created brand...", Brand, nil)
	c.JSON(http.StatusOK, successRes)
}

func (br *BrandHandler) EditBrand(c *gin.Context) {
	idstr := c.Query("id")
	id, err := strconv.Atoi(idstr)

	if err != nil {
		errRes := response.ClientResponse(http.StatusBadRequest, "error in id", nil, err.Error())
		c.JSON(http.StatusBadRequest, errRes)
		return
	}

	var EditBrand models.EditBrand

	if err := c.ShouldBindJSON(&EditBrand); err != nil {
		errRes := response.ClientResponse(http.StatusBadRequest, "error binding json format", nil, err.Error())
		c.JSON(http.StatusBadRequest, errRes)
		return
	}

	Brand, err := br.usecase.EditBrand(EditBrand, id)
	if err != nil {
		errRes := response.ClientResponse(http.StatusBadRequest, "failed to edit Brand", nil, err.Error())
		c.JSON(http.StatusBadRequest, errRes)
		return
	}

	successRes := response.ClientResponse(http.StatusOK, "successfully edited Brand...", Brand, nil)
	c.JSON(http.StatusOK, successRes)
}

func (br *BrandHandler) DeleteBrand(c *gin.Context) {
	idstr := c.Query("id")
	id, err := strconv.Atoi(idstr)

	if err != nil {
		errRes := response.ClientResponse(http.StatusBadRequest, "error in id", nil, err.Error())
		c.JSON(http.StatusBadRequest, errRes)
		return
	}

	err = br.usecase.DeleteBrand(id)
	if err != nil {
		errRes := response.ClientResponse(http.StatusBadRequest, "failed to delete Brand...", nil, err.Error())
		c.JSON(http.StatusBadRequest, errRes)
		return
	}

	successRes := response.ClientResponse(http.StatusOK, "successfully deleted Brand...", nil, nil)
	c.JSON(http.StatusOK, successRes)
}

func (br *BrandHandler) GetBrands(c *gin.Context) {
	brands, err := br.usecase.ListBrands()
	if err != nil {
		errRes := response.ClientResponse(http.StatusBadRequest, "failed to retrieve all Brands...", nil, err.Error())
		c.JSON(http.StatusBadRequest, errRes)
		return
	}

	successRes := response.ClientResponse(http.StatusOK, "successfully retrieved all Brands...", brands, nil)
	c.JSON(http.StatusOK, successRes)
}

func (br *BrandHandler) FilterByBrand(c *gin.Context) {
	idstr := c.Query("id")
	id, err := strconv.Atoi(idstr)

	page, _ := strconv.Atoi(c.Query("page"))
	per_product, _ := strconv.Atoi(c.Query("per_product"))

	if err != nil {
		errRes := response.ClientResponse(http.StatusBadRequest, "error in id", nil, err.Error())
		c.JSON(http.StatusBadRequest, errRes)
		return
	}

	products, name, err := br.usecase.FilterByBrand(id, page, per_product)
	if err != nil {
		errRes := response.ClientResponse(http.StatusBadRequest, "failed to get products under this brand", nil, err.Error())
		c.JSON(http.StatusBadRequest, errRes)
		return
	}

	successRes := response.BrandResponseWithProduct(http.StatusOK, id, name, products)
	c.JSON(http.StatusOK, successRes)
}
