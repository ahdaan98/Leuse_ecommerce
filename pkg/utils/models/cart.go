package models

type GetCart struct {
	ProductID   int     `json:"product_id"`
	ProductName string  `json:"product_name"`
	BrandID     uint    `json:"brand_id"`
	Brand       string  `json:"brand"`
	CategoryID  uint    `json:"category_id"`
	Category    string  `json:"category"`
	Quantity    int     `json:"quantity"`
	Price       int     `json:"price"`
	Total       float64 `json:"total_price"`
}

// check out
type CheckOut struct {
	CartID        int
	Addresses     []Address
	Products      []GetCart
	PaymentMethod []PaymentMethodResponse
}

type AddToCart struct {
	UserID      int `json:"user_id"`
	InventoryID int `json:"inventory_id"`
}