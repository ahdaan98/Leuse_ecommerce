package handler

import (
	"net/http"
	"strconv"

	interfaces "github.com/ahdaan98/pkg/usecase/interface"
	"github.com/ahdaan98/pkg/utils/models"
	"github.com/ahdaan98/pkg/utils/response"
	"github.com/gin-gonic/gin"
)

type CategoryHandler struct {
	usecase interfaces.CategoryUseCase
}

func NewCategoryHandler(usecase interfaces.CategoryUseCase) *CategoryHandler {
	return &CategoryHandler{
		usecase: usecase,
	}
}

func (ca *CategoryHandler) AddCategory(c *gin.Context) {
	var categoryName models.AddCategory

	if err := c.ShouldBindJSON(&categoryName); err != nil {
		errRes := response.ClientResponse(http.StatusBadRequest, "error binding json format", nil, err.Error())
		c.JSON(http.StatusBadRequest, errRes)
		return
	}

	Category, err := ca.usecase.AddCategory(categoryName)
	if err != nil {
		errRes := response.ClientResponse(http.StatusBadRequest, "failed to add category", nil, err.Error())
		c.JSON(http.StatusBadRequest, errRes)
		return
	}

	successRes := response.ClientResponse(http.StatusOK, "successfully created brand...", Category, nil)
	c.JSON(http.StatusOK, successRes)
}

func (ca *CategoryHandler) EditCategory(c *gin.Context) {
	idstr := c.Query("id")
	id, err := strconv.Atoi(idstr)

	if err != nil {
		errRes := response.ClientResponse(http.StatusBadRequest, "error in id", nil, err.Error())
		c.JSON(http.StatusBadRequest, errRes)
		return
	}

	var EditCategory models.EditCategory

	if err := c.ShouldBindJSON(&EditCategory); err != nil {
		errRes := response.ClientResponse(http.StatusBadRequest, "error binding json format", nil, err.Error())
		c.JSON(http.StatusBadRequest, errRes)
		return
	}

	Category, err := ca.usecase.EditCategory(EditCategory, id)
	if err != nil {
		errRes := response.ClientResponse(http.StatusBadRequest, "failed to edit category", nil, err.Error())
		c.JSON(http.StatusBadRequest, errRes)
		return
	}

	successRes := response.ClientResponse(http.StatusOK, "successfully edited category...", Category, nil)
	c.JSON(http.StatusOK, successRes)
}

func (ca *CategoryHandler) DeleteCategory(c *gin.Context) {
	idstr := c.Query("id")
	id, err := strconv.Atoi(idstr)

	if err != nil {
		errRes := response.ClientResponse(http.StatusBadRequest, "error in id", nil, err.Error())
		c.JSON(http.StatusBadRequest, errRes)
		return
	}

	err = ca.usecase.DeleteCategory(id)
	if err != nil {
		errRes := response.ClientResponse(http.StatusBadRequest, "failed to delete category...", nil, err.Error())
		c.JSON(http.StatusBadRequest, errRes)
		return
	}

	successRes := response.ClientResponse(http.StatusOK, "successfully deleted category...", nil, nil)
	c.JSON(http.StatusOK, successRes)
}

func (ca *CategoryHandler) GetCategories(c *gin.Context) {
	categories, err := ca.usecase.ListCategories()
	if err != nil {
		errRes := response.ClientResponse(http.StatusBadRequest, "failed to retrieve all categories...", nil, err.Error())
		c.JSON(http.StatusBadRequest, errRes)
		return
	}

	successRes := response.ClientResponse(http.StatusOK, "successfully retrieved all categories...", categories, nil)
	c.JSON(http.StatusOK, successRes)
}

func (ca *CategoryHandler) FilterByCategory(c *gin.Context) {
	idstr := c.Query("id")
	id, err := strconv.Atoi(idstr)

	if err != nil {
		errRes := response.ClientResponse(http.StatusBadRequest, "error in id", nil, err.Error())
		c.JSON(http.StatusBadRequest, errRes)
		return
	}

	products, name, err := ca.usecase.FilterByCategory(id)
	if err != nil {
		errRes := response.ClientResponse(http.StatusBadRequest, "failed to get product under this category", nil, err.Error())
		c.JSON(http.StatusBadRequest, errRes)
		return
	}

	successRes := response.CategoryResponseWithProduct(http.StatusOK, id, name, products)
	c.JSON(http.StatusOK, successRes)
}
