package domain

type Inventory struct {
	ID          uint     `json:"id" gorm:"primaryKey"`
	ProductName string   `json:"product_name" gorm:"not null"`
	BrandID     uint     `json:"brand_id" gorm:"not null"`
	Brand       Brand    `json:"brand" gorm:"foreignKey:BrandID;constraint:OnDelete:CASCADE"`
	CategoryID  uint     `json:"category_id" gorm:"not null"`
	Category    Category `json:"category" gorm:"foreignKey:CategoryID;constraint:OnDelete:CASCADE"`
	Stock       int      `json:"stock" gorm:"not null"`
	Price       float64  `json:"price" gorm:"not null"`
}

type Category struct {
	ID           uint   `json:"id" gorm:"primaryKey"`
	CategoryName string `json:"category_name" gorm:"not null"`
}

type Brand struct {
	ID        uint   `json:"id" gorm:"primaryKey"`
	BrandName string `json:"brand_name" gorm:"not null"`
}

type Image struct {
    InventoryID uint      `json:"inventory_id"`
    Inventory   Inventory `gorm:"foreignKey:InventoryID"`
    Image       string    `json:"image" gorm:"not null"`
}