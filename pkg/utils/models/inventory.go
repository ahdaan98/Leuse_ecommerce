package models

// Add Details

type AddInventory struct {
	ProductName string  `json:"product_name"`
	BrandID     uint    `json:"brand_id"`
	CategoryID  uint    `json:"category_id"`
	Stock       int     `json:"stock"`
	Price       float64 `json:"price"`
}

type AddCategory struct {
	CategoryName string `json:"category_name"`
}

type AddBrand struct {
	BrandName string `json:"brand_name"`
}

type Order struct {
	AddressID       int `json:"address_id"`
	PaymentMethodID int `json:"payment_id"`
	CouponID        int `json:"coupon_id"`
}

// Edit Details

type EditInventory struct {
	ProductName string  `json:"product_name"`
	CategoryID  uint    `json:"category_id"`
	BrandID     uint    `json:"brand_id"`
	Price       float64 `json:"price"`
}

type EditCategory struct {
	CategoryID   uint   `json:"category_id"`
	CategoryName string `json:"category_name"`
}

type EditBrand struct {
	BrandID   uint   `json:"brand_id"`
	BrandName string `json:"brand_name"`
}

// Update

type UpdateInventory struct {
	Stock int `json:"stock"`
}

// Response

type ReturingInventories struct {
	ProductID   uint
	ProductName string
	BrandID     uint
	CategoryID  uint
	Stock       int
	Price       float64
}

type InventoryResponse struct {
	ProductID   uint64   `json:"id"`
	ProductName string   `json:"product_name"`
	CategoryID  uint     `json:"category_id"`
	Category    string   `json:"category"`
	BrandID     uint     `json:"brand_id"`
	Brand       string   `json:"brand"`
	Stock       int      `json:"stock"`
	Price       float64  `json:"price"`
}

type InventoryResponseWithImages struct {
    ProductID   uint64  `json:"id"`
    ProductName string  `json:"product_name"`
    CategoryID  uint    `json:"category_id"`
    Category    string  `json:"category"`
    BrandID     uint    `json:"brand_id"`
    Brand       string  `json:"brand"`
    Stock       int     `json:"stock"`
    Price       float64 `json:"price"`
    Images      []string `json:"images"`
}

type CheckStockResponse struct {
	ProductName string `json:"product_name"`
	Stock       int    `json:"stock"`
}

type FilterByCategoryResponse struct {
	ProductID   uint64  `json:"id"`
	ProductName string  `json:"product_name"`
	BrandID     uint    `json:"brand_id"`
	Brand       string  `json:"brand"`
	Stock       int     `json:"stock"`
	Price       float64 `json:"price"`
}

type FilterByBrandResponse struct {
	ProductID   uint64  `json:"id"`
	ProductName string  `json:"product_name"`
	CategoryID  uint    `json:"category_id"`
	Category    string  `json:"category"`
	Stock       int     `json:"stock"`
	Price       float64 `json:"price"`
}
