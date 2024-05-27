package usecase

import (
	"errors"
	"testing"

	repo_mocks "github.com/ahdaan98/pkg/repository/mocks"
	"github.com/ahdaan98/pkg/utils/models"
	"github.com/golang/mock/gomock"
	"github.com/stretchr/testify/assert"
)

func TestAddInventory(t *testing.T) {
	mockCtrl := gomock.NewController(t)
	defer mockCtrl.Finish()

	mockRepo := repo_mocks.NewMockInventoryRepository(mockCtrl)

	uc := NewInventoryUseCase(mockRepo)

	tests := []struct {
		name        string
		inventory   models.AddInventory
		mockFunc    func()
		expectedErr error
	}{
		{
			name: "Valid Inventory",
			inventory: models.AddInventory{
				Stock:      10,
				BrandID:    1,
				CategoryID: 1,
				Price:      100,
			},
			mockFunc: func() {
				mockRepo.EXPECT().CheckInventoryExist(gomock.Any()).Return(false, nil)
				mockRepo.EXPECT().AddInventory(gomock.Any()).Return(models.InventoryResponse{}, nil)
			},
			expectedErr: nil,
		},
		{
			name: "Invalid Inventory - Negative Stock",
			inventory: models.AddInventory{
				Stock:      -10,
				BrandID:    1,
				CategoryID: 1,
				Price:      100,
			},
			mockFunc: func() {},
			expectedErr: errors.New("check values properly, id cannot be negative or zero"),
		},
	}

	for _, tc := range tests {
		t.Run(tc.name, func(t *testing.T) {
			tc.mockFunc()
			result, err := uc.AddInventory(tc.inventory)
			assert.Equal(t, tc.expectedErr, err)
			assert.Equal(t, models.InventoryResponse{}, result)
		})
	}
}

func TestListProducts(t *testing.T) {
	mockCtrl := gomock.NewController(t)
	defer mockCtrl.Finish()

	mockRepo := repo_mocks.NewMockInventoryRepository(mockCtrl)

	uc := NewInventoryUseCase(mockRepo)

	test := struct {
		name         string
		page         int
		perProduct   int
		mockFunc     func()
		expectedResp []models.InventoryResponse
		expectedErr  error
	}{
		name:       "Valid Page and PerProduct",
		page:       1,
		perProduct: 10,
		mockFunc: func() {
			mockRepo.EXPECT().ListProducts(1, 10).Return([]models.InventoryResponse{}, nil)
		},
		expectedResp: []models.InventoryResponse{},
		expectedErr:  nil,
	}

	t.Run(test.name, func(t *testing.T) {
		test.mockFunc()
		resp, err := uc.ListProducts(test.page, test.perProduct)
		assert.Equal(t, test.expectedResp, resp)
		assert.Equal(t, test.expectedErr, err)
	})
}

func TestEditInventory(t *testing.T) {
    mockCtrl := gomock.NewController(t)
    defer mockCtrl.Finish()

    mockRepo := repo_mocks.NewMockInventoryRepository(mockCtrl)
    uc := NewInventoryUseCase(mockRepo)

    tests := []struct {
        name      string
        inventory models.EditInventory
        id        int
        mockFunc  func()
        wantErr   bool
    }{
        {
            name: "Valid Inventory Edit",
            inventory: models.EditInventory{
                BrandID:    1,
                CategoryID: 1,
                Price:      100,
                ProductName: "EditedProduct",
            },
            id: 1,
            mockFunc: func() {
                mockRepo.EXPECT().CheckInventoryExist(gomock.Any()).Return(false, nil)
                mockRepo.EXPECT().CheckInventoryExistByID(1).Return(true, nil)
                mockRepo.EXPECT().EditInventory(gomock.Any(), 1).Return(models.InventoryResponse{}, nil)
            },
            wantErr: false,
        },
        {
            name: "Invalid Inventory Edit - Negative ID",
            inventory: models.EditInventory{
                BrandID:    1,
                CategoryID: 1,
                Price:      100,
                ProductName: "EditedProduct",
            },
            id: -1,
            mockFunc: func() {},
            wantErr:  true,
        },
    }

    for _, tc := range tests {
        t.Run(tc.name, func(t *testing.T) {
            tc.mockFunc()
            _, err := uc.EditInventory(tc.inventory, tc.id)
            if (err != nil) != tc.wantErr {
                t.Errorf("EditInventory() error = %v, wantErr %v", err, tc.wantErr)
                return
            }
        })
    }
}

func TestUpdateInventory(t *testing.T) {
    mockCtrl := gomock.NewController(t)
    defer mockCtrl.Finish()

    mockRepo := repo_mocks.NewMockInventoryRepository(mockCtrl)
    uc := NewInventoryUseCase(mockRepo)

    tests := []struct {
        name      string
        inventory models.UpdateInventory
        id        int
        mockFunc  func()
        wantErr   bool
    }{
        {
            name: "Valid Inventory Update",
            inventory: models.UpdateInventory{
                Stock: 20,
            },
            id: 1,
            mockFunc: func() {
                mockRepo.EXPECT().CheckInventoryExistByID(1).Return(true, nil)
                mockRepo.EXPECT().UpdateInventory(gomock.Any(), 1).Return(models.InventoryResponse{}, nil)
            },
            wantErr: false,
        },
        {
            name: "Invalid Inventory Update - Negative ID",
            inventory: models.UpdateInventory{
                Stock: 20,
            },
            id: -1,
            mockFunc: func() {},
            wantErr:  true,
        },
        {
            name: "Invalid Inventory Update - Negative Stock",
            inventory: models.UpdateInventory{
                Stock: -20,
            },
            id: 1,
            mockFunc: func() {},
            wantErr:  true,
        },
        {
            name: "Invalid Inventory Update - Non-existent Product",
            inventory: models.UpdateInventory{
                Stock: 20,
            },
            id: 1,
            mockFunc: func() {
                mockRepo.EXPECT().CheckInventoryExistByID(1).Return(false, nil)
            },
            wantErr:  true,
        },
    }

    for _, tc := range tests {
        t.Run(tc.name, func(t *testing.T) {
            tc.mockFunc()
            _, err := uc.UpdateInventory(tc.inventory, tc.id)
            if (err != nil) != tc.wantErr {
                t.Errorf("UpdateInventory() error = %v, wantErr %v", err, tc.wantErr)
                return
            }
        })
    }
}

func TestShowIndividualProduct(t *testing.T) {
    mockCtrl := gomock.NewController(t)
    defer mockCtrl.Finish()

    mockRepo := repo_mocks.NewMockInventoryRepository(mockCtrl)
    uc := NewInventoryUseCase(mockRepo)

    tests := []struct {
        name     string
        productID int
        mockFunc func()
        wantErr  bool
    }{
        {
            name:     "Valid Product ID",
            productID: 1,
            mockFunc: func() {
                mockRepo.EXPECT().CheckInventoryExistByID(1).Return(true, nil)
                mockRepo.EXPECT().ShowIndividualProduct(1).Return(models.InventoryResponse{}, nil)
            },
            wantErr: false,
        },
        {
            name:     "Non-existent Product",
            productID: 2,
            mockFunc: func() {
                mockRepo.EXPECT().CheckInventoryExistByID(2).Return(false, nil)
            },
            wantErr:  true,
        },
        {
            name:     "Repository Error",
            productID: 3,
            mockFunc: func() {
                mockRepo.EXPECT().CheckInventoryExistByID(3).Return(true, nil)
                mockRepo.EXPECT().ShowIndividualProduct(3).Return(models.InventoryResponse{}, errors.New("repository error"))
            },
            wantErr:  true,
        },
    }

    for _, tc := range tests {
        t.Run(tc.name, func(t *testing.T) {
            tc.mockFunc()
            _, err := uc.ShowIndividualProduct(tc.productID)
            if (err != nil) != tc.wantErr {
                t.Errorf("ShowIndividualProduct() error = %v, wantErr %v", err, tc.wantErr)
                return
            }
        })
    }
}

func TestCheckStock(t *testing.T) {
    mockCtrl := gomock.NewController(t)
    defer mockCtrl.Finish()

    mockRepo := repo_mocks.NewMockInventoryRepository(mockCtrl)
    uc := NewInventoryUseCase(mockRepo)

    tests := []struct {
        name      string
        productID int
        mockFunc  func()
        wantErr   bool
    }{
        {
            name:      "Valid Product ID",
            productID: 1,
            mockFunc: func() {
                mockRepo.EXPECT().CheckInventoryExistByID(1).Return(true, nil)
                mockRepo.EXPECT().CheckStock(1).Return(models.CheckStockResponse{}, nil)
            },
            wantErr:   false,
        },
        {
            name:      "Invalid Product ID",
            productID: -1,
            mockFunc:  func() {},
            wantErr:   true,
        },
        {
            name:      "Non-existent Product",
            productID: 2,
            mockFunc: func() {
                mockRepo.EXPECT().CheckInventoryExistByID(2).Return(false, nil)
            },
            wantErr:   true,
        },
        {
            name:      "Repository Error",
            productID: 3,
            mockFunc: func() {
                mockRepo.EXPECT().CheckInventoryExistByID(3).Return(true, nil)
                mockRepo.EXPECT().CheckStock(3).Return(models.CheckStockResponse{}, errors.New("repository error"))
            },
            wantErr:   true,
        },
    }

    for _, tc := range tests {
        t.Run(tc.name, func(t *testing.T) {
            tc.mockFunc()
            _, err := uc.CheckStock(tc.productID)
            if (err != nil) != tc.wantErr {
                t.Errorf("CheckStock() error = %v, wantErr %v", err, tc.wantErr)
                return
            }
        })
    }
}

func TestAddImage(t *testing.T) {
    mockCtrl := gomock.NewController(t)
    defer mockCtrl.Finish()

    mockRepo := repo_mocks.NewMockInventoryRepository(mockCtrl)
    uc := NewInventoryUseCase(mockRepo)

    tests := []struct {
        name      string
        id        int
        image     string
        mockFunc  func()
        wantErr   bool
    }{
        {
            name:     "Valid Parameters",
            id:       1,
            image:    "example.jpg",
            mockFunc: func() {
                mockRepo.EXPECT().UploadImage(1, "example.jpg").Return(nil)
            },
            wantErr:  false,
        },
        {
            name:     "Repository Error",
            id:       2,
            image:    "example.jpg",
            mockFunc: func() {
                mockRepo.EXPECT().UploadImage(2, "example.jpg").Return(errors.New("repository error"))
            },
            wantErr:  true,
        },
    }

    for _, tc := range tests {
        t.Run(tc.name, func(t *testing.T) {
            tc.mockFunc()
            err := uc.AddImage(tc.id, tc.image)
            if (err != nil) != tc.wantErr {
                t.Errorf("AddImage() error = %v, wantErr %v", err, tc.wantErr)
                return
            }
        })
    }
}
