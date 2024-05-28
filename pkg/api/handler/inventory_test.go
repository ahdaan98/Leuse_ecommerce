package handler

import (
	"bytes"
	"encoding/json"
	"fmt"
	"net/http"
	"net/http/httptest"
	"strconv"
	"testing"

	"github.com/ahdaan98/pkg/domain"
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
func TestAddCategory(t *testing.T) {
    mockCtrl := gomock.NewController(t)
    defer mockCtrl.Finish()

    mockUseCase := usecase_mocks.NewMockCategoryUseCase(mockCtrl)
    handler := NewCategoryHandler(mockUseCase)

    type testCase struct {
        description     string
        input           interface{}
        mockReturn      domain.Category
        mockError       error
        expectedStatus  int
        expectedMessage string
        expectedError   string
        expectedData    map[string]interface{}
    }

    testCases := []testCase{
        {
            description: "Successful Add Category",
            input: models.AddCategory{
                CategoryName: "New Category",
            },
            mockReturn: domain.Category{
                ID:           1,
                CategoryName: "New Category",
            },
            mockError:       nil,
            expectedStatus:  http.StatusOK,
            expectedMessage: "successfully created category...",
            expectedError:   "",
            expectedData: map[string]interface{}{
                "id":           float64(1),
                "category_name": "New Category",
            },
        },
        {
            description:     "Failed to add category",
            input:           `{"invalid_json":}`, // invalid JSON
            expectedStatus:  http.StatusBadRequest,
            expectedMessage: "error binding json format",
            expectedError:   "invalid character '}' looking for beginning of value",
            expectedData:    nil,
        },
    }

    for _, tc := range testCases {
        t.Run(tc.description, func(t *testing.T) {
            var reqBody []byte
            var err error

            switch input := tc.input.(type) {
            case models.AddCategory:
                reqBody, err = json.Marshal(input)
                if err != nil {
                    t.Fatalf("failed to marshal input: %v", err)
                }
            case string:
                reqBody = []byte(input)
            }

            req := httptest.NewRequest(http.MethodPost, "/add-category", bytes.NewBuffer(reqBody))
            req.Header.Set("Content-Type", "application/json")

            w := httptest.NewRecorder()
            c, _ := gin.CreateTestContext(w)
            c.Request = req

            if tc.mockError == nil {
                mockUseCase.EXPECT().AddCategory(gomock.Any()).Return(tc.mockReturn, tc.mockError).AnyTimes()
            } else {
                mockUseCase.EXPECT().AddCategory(gomock.Any()).Return(tc.mockReturn, tc.mockError).AnyTimes()
            }

            handler.AddCategory(c)

            resp := w.Result()
            // assert.Equal(t, tc.expectedStatus, resp.StatusCode)

            var respBody response.Response
            err = json.NewDecoder(resp.Body).Decode(&respBody)
            if err != nil {
                t.Fatalf("failed to decode response body: %v", err)
            }

            assert.Equal(t, tc.expectedMessage, respBody.Message)

            // assert.Equal(t, tc.expectedError, respBody.Error)

            assert.Equal(t, tc.expectedData, respBody.Data)
        })
    }
}

func TestEditCategory(t *testing.T) {
    mockCtrl := gomock.NewController(t)
    defer mockCtrl.Finish()

    mockCategoryUseCase := usecase_mocks.NewMockCategoryUseCase(mockCtrl)
    handler := NewCategoryHandler(mockCategoryUseCase)

    type testCase struct {
        description      string
        id               string
        editCategory     models.EditCategory
        mockReturn       domain.Category
        mockError        error
        expectedStatus   int
        expectedMessage  string
        expectedData     interface{}
        expectedErrorMsg string
    }

    testCases := []testCase{
        {
            description: "Successful Edit Category",
            id:          "123",
            editCategory: models.EditCategory{
                CategoryName: "Edited Category",
            },
            mockReturn: domain.Category{
                ID:           123,
                CategoryName: "Edited Category",
            },
            mockError:       nil,
            expectedStatus:  http.StatusOK,
            expectedMessage: "successfully edited category...",
            expectedData: map[string]interface{}{
                "id":             float64(123),
                "category_name":  "Edited Category",
            },
            expectedErrorMsg: "",
        },
        // {
        //     description:     "Invalid ID",
        //     id:              "invalid_id",
        //     expectedStatus:  http.StatusBadRequest,
        //     expectedMessage: "error in id",
        //     expectedData:    nil,
        //     expectedErrorMsg: "strconv.Atoi: parsing \"invalid_id\": invalid syntax",
        // },
    }

    for _, tc := range testCases {
        t.Run(tc.description, func(t *testing.T) {
            reqBody, err := json.Marshal(tc.editCategory)
            if err != nil {
                t.Fatalf("failed to marshal input: %v", err)
            }

            req := httptest.NewRequest(http.MethodPost, fmt.Sprintf("/edit-category?id=%s", tc.id), bytes.NewBuffer(reqBody))
            req.Header.Set("Content-Type", "application/json")

            w := httptest.NewRecorder()
            c, _ := gin.CreateTestContext(w)
            c.Request = req

            if tc.mockError == nil {
                mockCategoryUseCase.EXPECT().EditCategory(gomock.Any(), gomock.Any()).Return(tc.mockReturn, tc.mockError).AnyTimes()
            } else {
                mockCategoryUseCase.EXPECT().EditCategory(gomock.Any(), gomock.Any()).Return(tc.mockReturn, tc.mockError).AnyTimes()
            }

            handler.EditCategory(c)

            resp := w.Result()
            assert.Equal(t, tc.expectedStatus, resp.StatusCode)

            var respBody response.Response
            err = json.NewDecoder(resp.Body).Decode(&respBody)
            if err != nil {
                t.Fatalf("failed to decode response body: %v", err)
            }

            assert.Equal(t, tc.expectedMessage, respBody.Message)
            assert.Equal(t, tc.expectedData, respBody.Data)

            if tc.expectedErrorMsg != "" {
                err, ok := respBody.Error.(error)
                if !ok {
                    t.Errorf("expected error type %T, got %T", error(nil), respBody.Error)
                } else {
                    assert.Equal(t, tc.expectedErrorMsg, err.Error())
                }
            } else if respBody.Error != nil {
                t.Errorf("expected no error, got %v", respBody.Error)
            }
        })
    }
}