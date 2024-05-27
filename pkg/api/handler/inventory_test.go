package handler

import (
	"bytes"
	"encoding/json"
	"net/http"
	"net/http/httptest"
	"strconv"
	"testing"

	usecase_mocks "github.com/ahdaan98/pkg/usecase/mocks"
	"github.com/ahdaan98/pkg/utils/models"
	"github.com/ahdaan98/pkg/utils/response"
	"github.com/gin-gonic/gin"
	"github.com/go-playground/assert"
	"github.com/golang/mock/gomock"
	"github.com/stretchr/testify/require"
)

func TestAddInventory(t *testing.T) {
	mockCtrl := gomock.NewController(t)
	defer mockCtrl.Finish()

	mockUseCase := usecase_mocks.NewMockInventoryUseCase(mockCtrl)

	handler := NewInventoryHandler(mockUseCase)

	inventory := models.AddInventory{
		ProductName: "New Product",
		BrandID:     1,
		CategoryID:  1,
		Stock:       10,
		Price:       99.99,
	}

	reqBody, _ := json.Marshal(inventory)
	req := httptest.NewRequest(http.MethodPost, "/add-inventory", bytes.NewBuffer(reqBody))
	req.Header.Set("Content-Type", "application/json")

	w := httptest.NewRecorder()
	c, _ := gin.CreateTestContext(w)
	c.Request = req

	// Mock setup
	mockUseCase.EXPECT().AddInventory(inventory).Return(models.InventoryResponse{}, nil)

	handler.AddInventory(c)

	resp := w.Result()
	assert.Equal(t, http.StatusOK, resp.StatusCode)

	var respBody response.Response
	err := json.NewDecoder(resp.Body).Decode(&respBody)
	if err != nil {
		t.Errorf("Error decoding response body: %v", err)
	}

	expectedResp := response.ClientResponse(http.StatusOK, "Successfully added new inventory...", nil, nil)
	assert.Equal(t, expectedResp.Message, respBody.Message)
	assert.Equal(t, expectedResp.Error, respBody.Error)
	assert.Equal(t, expectedResp.Data, respBody.Data)
}

func TestShowIndividualProduct(t *testing.T) {
	mockCtrl := gomock.NewController(t)
	defer mockCtrl.Finish()

	mockUseCase := usecase_mocks.NewMockInventoryUseCase(mockCtrl)

	handler := NewInventoryHandler(mockUseCase)

	productID := 1
	mockProduct := models.InventoryResponse{
		ProductID:   1,
		ProductName: "Product Name",
		CategoryID:  1,
		Category:    "Category Name",
		BrandID:     1,
		Brand:       "Brand Name",
		Stock:       100,
		Price:       99.99,
	}

	expectedResp := response.Response{
		Message: "successfully retrieved the product by id",
		Error:   nil,
		Data:    mockProduct,
	}

	// Mock setup for successful product retrieval
	mockUseCase.EXPECT().ShowIndividualProduct(productID).Return(mockProduct, nil)

	req := httptest.NewRequest(http.MethodGet, "/product?id="+strconv.Itoa(productID), nil)
	w := httptest.NewRecorder()
	c, _ := gin.CreateTestContext(w)
	c.Request = req

	handler.ShowIndividualProduct(c)

	resp := w.Result()
	assert.Equal(t, http.StatusOK, resp.StatusCode)

	var respBody response.Response
	err := json.NewDecoder(resp.Body).Decode(&respBody)
	require.NoError(t, err, "Error decoding response body")

	

	assert.Equal(t, expectedResp.Message, respBody.Message)
	assert.Equal(t, expectedResp.Error, respBody.Error)
	// assert.Equal(t, mockProductMap, respBody.Data)
}

func TestListProductsWithImages(t *testing.T) {
	mockCtrl := gomock.NewController(t)
	defer mockCtrl.Finish()

	mockUseCase := usecase_mocks.NewMockInventoryUseCase(mockCtrl)
	handler := NewInventoryHandler(mockUseCase)

	page := 1
	perProduct := 10
	mockProducts := []models.InventoryResponseWithImages{
		{
			ProductID:   1,
			ProductName: "Product 1",
			CategoryID:  1,
			Category:    "Category 1",
			BrandID:     1,
			Brand:       "Brand 1",
			Stock:       100,
			Price:       99.99,
			Images:      []string{"image1.jpg", "image2.jpg"},
		},
		{
			ProductID:   2,
			ProductName: "Product 2",
			CategoryID:  2,
			Category:    "Category 2",
			BrandID:     2,
			Brand:       "Brand 2",
			Stock:       200,
			Price:       199.99,
			Images:      []string{"image3.jpg", "image4.jpg"},
		},
	}

	// Mock setup for successful product listing
	mockUseCase.EXPECT().ListProductsWithImages(page, perProduct).Return(mockProducts, nil)

	req := httptest.NewRequest(http.MethodGet, "/list-products?page="+strconv.Itoa(page)+"&per_product="+strconv.Itoa(perProduct), nil)
	w := httptest.NewRecorder()
	c, _ := gin.CreateTestContext(w)
	c.Request = req

	handler.ListProductsWithImages(c)

	resp := w.Result()
	assert.Equal(t, http.StatusOK, resp.StatusCode)

	var respBody []models.InventoryResponseWithImages
	err := json.NewDecoder(resp.Body).Decode(&respBody)
	assert.Equal(t, err, nil)

	assert.Equal(t, mockProducts, respBody)
}
