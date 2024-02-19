package repository

import (
	"errors"

	"github.com/ahdaan98/pkg/domain"
	interfaces "github.com/ahdaan98/pkg/repository/interface"
	"github.com/ahdaan98/pkg/utils/models"
	"gorm.io/gorm"
)

type userDatabase struct {
	DB *gorm.DB
}

func NewUserRepository(DB *gorm.DB) interfaces.UserRepository {
	return &userDatabase{
		DB: DB,
	}
}

func (i *userDatabase) CreateUser(user models.UserSignUp) (models.UserDetailsResponse, error) {
	var resp models.UserDetailsResponse
	query := `
	INSERT INTO users (name,email,password,phone) VALUES (?,?,?,?) RETURNING id, name, email, phone
	`

	if err := i.DB.Raw(query, user.Name, user.Email, user.Password, user.Phone).Scan(&resp).Error; err != nil {
		return models.UserDetailsResponse{}, err
	}

	// create wallet
	err := i.DB.Exec("INSERT INTO wallets (user_id, amount) VALUES (?, ?)", resp.Id, 0).Error
	if err != nil {
		return models.UserDetailsResponse{}, err
	}

	return resp, nil
}

func (i *userDatabase) CheckUserExist(email string) (bool, error) {
	var count int
	query := `
	SELECT COUNT(*) FROM users
	WHERE email = ?
	`

	if err := i.DB.Raw(query, email).Scan(&count).Error; err != nil {
		return false, err
	}

	return count > 0, nil
}

func (i *userDatabase) CheckUserExistByID(id int) (bool, error) {
	var count int
	query := `
	SELECT COUNT(*) FROM users
	WHERE id = ?
	`

	if err := i.DB.Raw(query, id).Scan(&count).Error; err != nil {
		return false, err
	}

	return count > 0, nil
}

func (i *userDatabase) CheckBlockStatus(email string) (bool, error) {
	var Blocked bool
	query := `
	SELECT blocked FROM users WHERE email = ?
	`

	if err := i.DB.Raw(query, email).Scan(&Blocked).Error; err != nil {
		return false, err
	}

	return Blocked, nil
}

func (i *userDatabase) GetUserDetails(id int) (models.UserDetailsResponse, error) {
	var resp models.UserDetailsResponse

	query := `
	SELECT id, name, email, phone FROM users
	WHERE id = ?
	`

	if err := i.DB.Raw(query, id).Scan(&resp).Error; err != nil {
		return models.UserDetailsResponse{}, err
	}

	return resp, nil
}

func (i *userDatabase) EditDetails(details models.EditUserDetails, id int) (models.UserDetailsResponse, error) {
	var resp models.UserDetailsResponse

	query := `
	UPDATE users SET name = ?, email = ?, phone = ? WHERE id = ?
	`
	if err := i.DB.Exec(query, details.Name, details.Email, details.Phone, id).Error; err != nil {
		return models.UserDetailsResponse{}, err
	}

	if err := i.DB.Raw("SELECT id, name, email, phone FROM users WHERE id = ?", id).Scan(&resp).Error; err != nil {
		return models.UserDetailsResponse{}, err
	}

	return resp, nil
}

func (i *userDatabase) GetUserByEmail(email string) (models.UserDetailsResponse, error) {
	var resp models.UserDetailsResponse
	query := `
	SELECT id, name, email, phone FROM users
	WHERE email = ?
	`

	if err := i.DB.Raw(query, email).Scan(&resp).Error; err != nil {
		return models.UserDetailsResponse{}, err
	}

	return resp, nil
}

func (i *userDatabase) GetUserPassword(email string) (string, error) {
	var password string
	query := `
	SELECT password FROM users 
	WHERE email = ?
	`

	if err := i.DB.Raw(query, email).Scan(&password).Error; err != nil {
		return "", err
	}

	return password, nil
}

func (i *userDatabase) AddAddress(userID int, address models.AddAddress) error{
	err := i.DB.Exec(`
		INSERT INTO addresses (user_id, name, house_name, street, city, state, phone, pin)
		VALUES ($1, $2, $3, $4, $5, $6, $7, $8 )`,
		userID, address.Name, address.HouseName, address.Street, address.City, address.State, address.Phone, address.Pin).Error
	if err != nil {
		return errors.New("could not add address")
	}

	return nil
}

func (i *userDatabase) GetAddresses(id int) ([]domain.Address, error){

	var addresses []domain.Address

	if err := i.DB.Raw("select * from addresses where user_id=?", id).Scan(&addresses).Error; err != nil {
		return []domain.Address{}, errors.New("error in getting addresses")
	}

	return addresses, nil

}

func (i *userDatabase) CheckIfFirstAddress(id int) bool{
	var count int

	if err := i.DB.Raw("select count(*) from addresses where user_id=$1", id).Scan(&count).Error; err != nil {
		return false
	}
	
	// if count is greater than 0 that means the address already exist
	return count > 0
}

func (ad *userDatabase) GetCartID(id int) (int, error) {

	var cart_id int

	if err := ad.DB.Raw("select id from carts where user_id=?", id).Scan(&cart_id).Error; err != nil {
		return 0, err
	}

	return cart_id, nil

}

func (ad *userDatabase) GetProductsInCart(cart_id int) ([]int, error) {

	var cart_products []int

	if err := ad.DB.Raw("select inventory_id from line_items where cart_id=?", cart_id).Scan(&cart_products).Error; err != nil {
		return []int{}, err
	}

	return cart_products, nil

}

func (ad *userDatabase) FindProductNames(inventory_id int) (string, error) {

	var product_name string

	if err := ad.DB.Raw("select product_name from inventories where id=?", inventory_id).Scan(&product_name).Error; err != nil {
		return "", err
	}

	return product_name, nil

}

func (ad *userDatabase) FindCartQuantity(cart_id, inventory_id int) (int, error) {

	var quantity int

	if err := ad.DB.Raw("select quantity from line_items where cart_id=$1 and inventory_id=$2", cart_id, inventory_id).Scan(&quantity).Error; err != nil {
		return 0, err
	}

	return quantity, nil

}

func (ad *userDatabase) FindPrice(inventory_id int) (float64, error) {

	var price float64

	if err := ad.DB.Raw("select price from inventories where id=?", inventory_id).Scan(&price).Error; err != nil {
		return 0, err
	}

	return price, nil

}

func (ad *userDatabase) FindCategory(inventoryID int) (int, error) {

	var categoryID int

	if err := ad.DB.Raw("SELECT category_id FROM inventories WHERE id = ?", inventoryID).Scan(&categoryID).Error; err != nil {
		return 0, err
	}

	return categoryID, nil
}

func (i *userDatabase) FindStock(id int) (int, error) {
	var stock int
	err := i.DB.Raw("SELECT stock FROM inventories WHERE id = ?", id).Scan(&stock).Error
	if err != nil {
		return 0, err
	}

	return stock, nil
}

func (ad *userDatabase) RemoveFromCart(cart, inventory int) error {

	if err := ad.DB.Exec(`DELETE FROM line_items WHERE cart_id = $1 AND inventory_id = $2`, cart, inventory).Error; err != nil {
		return err
	}

	return nil

}
func (ad *userDatabase) UpdateQuantity(id, invID, qty int) error {
		query := `
        UPDATE line_items
        SET quantity = $1
        WHERE cart_id = $2 AND inventory_id = $3
        `

		result := ad.DB.Exec(query, qty, id, invID)
		if result.Error != nil {
			return result.Error
		}
	

	return nil
}

func (i *userDatabase) ChangePassword(id int, password string) error {

	err := i.DB.Exec("UPDATE users SET password=$1 WHERE id=$2", password, id).Error
	if err != nil {
		return err
	}

	return nil

}