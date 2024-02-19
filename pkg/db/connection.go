package db

import (
	"errors"

	"github.com/ahdaan98/pkg/config"
	"github.com/ahdaan98/pkg/domain"
	"gorm.io/driver/postgres"
	"gorm.io/gorm"
)

func ConnectDB(cfg config.Config) (*gorm.DB, error) {

	if cfg.DBUrl == "" {
		return nil, errors.New("empty database url")
	}

	DB, err := gorm.Open(postgres.Open(cfg.DBUrl), &gorm.Config{SkipDefaultTransaction: true})
	if err != nil {
		return nil, errors.New("failed to connect database")
	}

	if err := DB.AutoMigrate(&domain.Inventory{}); err != nil {
		return DB, err
	}
	if err := DB.AutoMigrate(&domain.Category{}); err != nil {
		return DB, err
	}
	if err := DB.AutoMigrate(&domain.Brand{}); err != nil {
		return DB, err

	}
	if err := DB.AutoMigrate(&domain.User{}); err != nil {
		return DB, err

	}
	if err := DB.AutoMigrate(&domain.Admin{}); err != nil {
		return DB, err

	}
	if err := DB.AutoMigrate(&domain.Address{}); err != nil {
		return DB, err

	}
	if err := DB.AutoMigrate(&domain.Cart{}); err != nil {
		return DB, err

	}
	if err := DB.AutoMigrate(domain.LineItems{}); err != nil {
		return DB, err

	}

	if err := DB.AutoMigrate(domain.Order{}); err != nil {
		return DB, err
	}
	if err := DB.AutoMigrate(domain.OrderItem{}); err != nil {
		return DB, err
	}
	if err := DB.AutoMigrate(domain.PaymentMethod{}); err != nil {
		return DB, err
	}
	if err := DB.AutoMigrate(domain.Payment{}); err != nil {
		return DB, err
	}

	if err := DB.AutoMigrate(domain.Wallet{}); err != nil {
		return DB, err
	}

	if err := DB.AutoMigrate(domain.Coupon{}); err != nil {
		return DB, err
	}

	if err := DB.AutoMigrate(domain.Image{}); err != nil {
		return DB, err
	}

	if err := DB.AutoMigrate(domain.OrderItemInv{}); err != nil {
		return DB, err
	}

	CheckAndCreateAdmin(DB)

	return DB, nil
}

func Migration(db *gorm.DB) error {
	if err := db.AutoMigrate(&domain.Inventory{}); err != nil {
		return err
	}
	if err := db.AutoMigrate(&domain.Category{}); err != nil {
		return err
	}
	if err := db.AutoMigrate(&domain.Brand{}); err != nil {
		return err
	}
	if err := db.AutoMigrate(&domain.User{}); err != nil {
		return err
	}
	if err := db.AutoMigrate(&domain.Admin{}); err != nil {
		return err
	}
	if err := db.AutoMigrate(&domain.Address{}); err != nil {
		return err
	}
	if err := db.AutoMigrate(&domain.Cart{}); err != nil {
		return err
	}
	if err := db.AutoMigrate(domain.LineItems{}); err != nil {
		return err
	}

	return nil
}

// Create admin
func CheckAndCreateAdmin(db *gorm.DB) {
	var count int64
	db.Model(&domain.Admin{}).Count(&count)
	if count == 0 {
		password := "2024"
		// hashedPassword, err := bcrypt.GenerateFromPassword([]byte(password), bcrypt.DefaultCost)
		// if err != nil {
		// 	return
		// }
		admin := domain.Admin{
			ID:       1,
			Name:     "leuse",
			Email:    "leuse@gmail.com",
			Password: password,
		}
		db.Create(&admin)
	}
}
