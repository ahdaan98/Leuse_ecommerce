package repository

import (
	"fmt"
	"time"

	"github.com/ahdaan98/pkg/domain"
	interfaces "github.com/ahdaan98/pkg/repository/interface"
	"github.com/ahdaan98/pkg/utils/models"
	"gorm.io/gorm"
)

type adminRepository struct {
	DB *gorm.DB
}

func NewAdminRepository(DB *gorm.DB) interfaces.AdminRepository {
	return &adminRepository{
		DB: DB,
	}
}

func (ad *adminRepository) GetUsers() ([]models.UserDetailsAtAdmin, error) {
	query:=`
	SELECT id, name, email, phone, blocked AS block_status FROM users
	`

	var users []models.UserDetailsAtAdmin
	if err:=ad.DB.Raw(query).Scan(&users).Error; err!=nil{
		return []models.UserDetailsAtAdmin{},err
	}

	return users,nil
}

func (ad *adminRepository) GetUserByID(id int) (models.UserDetailsAtAdmin,error){
	query:=`
	SELECT id, name, email, phone, blocked AS block_status FROM users WHERE id = ?
	`

	var user models.UserDetailsAtAdmin
	if err:=ad.DB.Raw(query,id).Scan(&user).Error; err!=nil{
		return models.UserDetailsAtAdmin{},err
	}

	return user,nil
}

func (ad *adminRepository) UpdateBlockUserByID(k bool,id int) (error) {
	query:=`
	UPDATE users SET blocked = ? WHERE id = ?
	`

	if err:=ad.DB.Exec(query,k,id).Error; err!=nil{
		return err
	}

	return nil
}

func (ad *adminRepository) GetAdminPassword(email string) (string,error){
	var password string
	query := `
	SELECT password FROM admins 
	WHERE email = ?
	`

	if err := ad.DB.Raw(query, email).Scan(&password).Error; err != nil {
		return "", err
	}

	return password, nil
}

func (ad *adminRepository)CheckAdminExist(email string) (bool,error){
	var count int
	query := `
	SELECT COUNT(*) FROM admins
	WHERE email = ?
	`

	if err := ad.DB.Raw(query, email).Scan(&count).Error; err != nil {
		return false, err
	}

	return count > 0, nil
}

func (ad *adminRepository) GetAdminByEmail(email string) (models.AdminDetailsResponse,error) {
	var resp models.AdminDetailsResponse
	query := `
	SELECT id, name, email FROM admins
	WHERE email = ?
	`

	if err := ad.DB.Raw(query, email).Scan(&resp).Error; err != nil {
		return models.AdminDetailsResponse{}, err
	}

	return resp, nil
}

func (i *adminRepository) NewPaymentMethod(pay string) error {

	if err := i.DB.Exec("insert into payment_methods(payment_name)values($1)", pay).Error; err != nil {
		return err
	}

	return nil

}

func (a *adminRepository) ListPaymentMethods() ([]domain.PaymentMethod, error) {
	var model []domain.PaymentMethod
	err := a.DB.Raw("SELECT * FROM payment_methods where is_deleted = false").Scan(&model).Error
	if err != nil {
		return []domain.PaymentMethod{}, err
	}

	return model, nil
}

func (a *adminRepository) CheckIfPaymentMethodAlreadyExists(payment string) (bool, error) {
	var count int64
	err := a.DB.Raw("SELECT COUNT(*) FROM payment_methods WHERE payment_name = $1 and is_deleted = false", payment).Scan(&count).Error
	if err != nil {
		return false, err
	}

	return count > 0, nil
}

func (a *adminRepository) DeletePaymentMethod(id int) error {
	err := a.DB.Exec("UPDATE payment_methods SET is_deleted = true WHERE id = $1 ", id).Error
	if err != nil {
		return err
	}

	return nil
}

func (a *adminRepository) GetPaymentMethod() ([]models.PaymentMethodResponse, error) {
	var model []models.PaymentMethodResponse
	err := a.DB.Raw("SELECT * FROM payment_methods").Scan(&model).Error
	if err != nil {
		return []models.PaymentMethodResponse{}, err
	}

	return model, nil
}

func (a *adminRepository) TotalRevenue() (models.DashboardRevenue, error) {
	var revenueDetails models.DashboardRevenue

	startTime := time.Now().AddDate(0, 0, -1)

	err := a.DB.Raw("select coalesce(sum(final_price),0) from orders where payment_status = 'PAID'  and created_at >= ?", startTime).Scan(&revenueDetails.TodayRevenue).Error
	if err != nil {
		return models.DashboardRevenue{}, nil
	}

	startTime = time.Now().AddDate(0, -1, 1).UTC()
	err = a.DB.Raw("select coalesce(sum(final_price),0) from orders where payment_status = 'PAID'  and created_at >= ?", startTime).Scan(&revenueDetails.MonthRevenue).Error
	if err != nil {
		return models.DashboardRevenue{}, nil
	}
	startTime = time.Now().AddDate(-1, 1, 1).UTC()
	err = a.DB.Raw("select coalesce(sum(final_price),0) from orders where payment_status = 'PAID'  and created_at >= ?", startTime).Scan(&revenueDetails.YearRevenue).Error
	if err != nil {
		return models.DashboardRevenue{}, nil
	}

	return revenueDetails, nil
}

func (ad *adminRepository) DashBoardOrder() (models.DashboardOrder, error) {

	var orderDetails models.DashboardOrder
	err := ad.DB.Raw("select count(*) from orders where payment_status = 'PAID'").Scan(&orderDetails.CompletedOrder).Error
	if err != nil {
		return models.DashboardOrder{}, nil
	}

	err = ad.DB.Raw("select count(*) from orders where order_status = 'PENDING' or order_status = 'PROCESSING'").Scan(&orderDetails.PendingOrder).Error
	if err != nil {
		return models.DashboardOrder{}, nil
	}

	err = ad.DB.Raw("select count(*) from orders where order_status = 'CANCELED'").Scan(&orderDetails.CancelledOrder).Error
	if err != nil {
		return models.DashboardOrder{}, nil
	}

	err = ad.DB.Raw("select count(*) from orders").Scan(&orderDetails.TotalOrder).Error
	if err != nil {
		return models.DashboardOrder{}, nil
	}

	err = ad.DB.Raw("select sum(quantity) from order_items").Scan(&orderDetails.TotalOrderItem).Error
	if err != nil {
		return models.DashboardOrder{}, nil
	}

	return orderDetails, nil

}

func (ad *adminRepository) AmountDetails() (models.DashboardAmount, error) {

	var amountDetails models.DashboardAmount
	err := ad.DB.Raw("select coalesce(sum(final_price),0) from orders where payment_status = 'PAID' ").Scan(&amountDetails.CreditedAmount).Error
	if err != nil {
		return models.DashboardAmount{}, nil
	}

	err = ad.DB.Raw("select coalesce(sum(final_price),0) from orders where payment_status = 'NOT PAID' and order_status = 'PROCESSING' or order_status = 'PENDING' or order_status = 'ORDER PLACED' ").Scan(&amountDetails.PendingAmount).Error
	if err != nil {
		return models.DashboardAmount{}, nil
	}

	return amountDetails, nil

}
func (ad *adminRepository) DashBoardUserDetails() (models.DashBoardUser, error) {
	var userDetails models.DashBoardUser
	err := ad.DB.Raw("SELECT COUNT(*) FROM users").Scan(&userDetails.TotalUsers).Error
	if err != nil {
		return models.DashBoardUser{}, nil
	}
	err = ad.DB.Raw("SELECT COUNT(*)  FROM users WHERE blocked=true").Scan(&userDetails.BlockedUser).Error
	if err != nil {
		return models.DashBoardUser{}, nil
	}
	return userDetails, nil
}

func (ad *adminRepository) DashBoardProductDetails() (models.DashBoardProduct, error) {
	var productDetails models.DashBoardProduct
	err := ad.DB.Raw("SELECT COUNT(*) FROM inventories").Scan(&productDetails.TotalProducts).Error
	if err != nil {
		return models.DashBoardProduct{}, nil
	}
	err = ad.DB.Raw("SELECT COUNT(*) FROM inventories WHERE stock<=0").Scan(&productDetails.OutofStockProduct).Error
	if err != nil {
		return models.DashBoardProduct{}, nil
	}
	return productDetails, nil
}

func (ad *adminRepository) SalesByYear(yearInt int) ([]models.OrderDetailsAdmin, error) {
	var orderDetails []models.OrderDetailsAdmin

	query := `SELECT i.product_name, SUM(oi.total_price) AS total_amount
              FROM orders o
              JOIN order_items oi ON o.id = oi.order_id
              JOIN inventories i ON oi.inventory_id = i.id
              WHERE o.payment_status = 'PAID'
                AND EXTRACT(YEAR FROM o.created_at) = ?
              GROUP BY i.product_name`

	if err := ad.DB.Raw(query, yearInt).Scan(&orderDetails).Error; err != nil {
		return []models.OrderDetailsAdmin{}, err
	}

	fmt.Println("body at repo year", orderDetails)

	return orderDetails, nil
}

func (ad *adminRepository) SalesByMonth(yearInt int, monthInt int) ([]models.OrderDetailsAdmin, error) {
	var orderDetails []models.OrderDetailsAdmin

	query := `SELECT i.product_name, SUM(oi.total_price) AS total_amount
              FROM orders o
              JOIN order_items oi ON o.id = oi.order_id
              JOIN inventories i ON oi.inventory_id = i.id
              WHERE o.payment_status = 'PAID'
			  AND EXTRACT(YEAR FROM o.created_at) = ?
			  AND EXTRACT(MONTH FROM o.created_at) = ?
              GROUP BY i.product_name`

	if err := ad.DB.Raw(query, yearInt, monthInt).Scan(&orderDetails).Error; err != nil {
		return []models.OrderDetailsAdmin{}, err
	}

	fmt.Println("body at repo month", orderDetails)

	return orderDetails, nil
}

func (ad *adminRepository) SalesByDay(yearInt int, monthInt int, dayInt int) ([]models.OrderDetailsAdmin, error) {
	var orderDetails []models.OrderDetailsAdmin

	query := `SELECT i.product_name, SUM(oi.total_price) AS total_amount
              FROM orders o
              JOIN order_items oi ON o.id = oi.order_id
              JOIN inventories i ON oi.inventory_id = i.id
              WHERE o.payment_status = 'PAID'
			  AND EXTRACT(YEAR FROM o.created_at) = ?
			  AND EXTRACT(MONTH FROM o.created_at) = ?
                AND EXTRACT(DAY FROM o.created_at) = ?
              GROUP BY i.product_name`

	if err := ad.DB.Raw(query, yearInt, monthInt, dayInt).Scan(&orderDetails).Error; err != nil {
		return []models.OrderDetailsAdmin{}, err
	}

	fmt.Println("body at repo day", orderDetails)

	return orderDetails, nil
}

func (ad *adminRepository) CustomSalesReportByDate(startTime time.Time, endTime time.Time) (models.SalesReport, error) {
	var salesReport models.SalesReport
	result := ad.DB.Raw("SELECT COALESCE(SUM(final_price),0) FROM orders WHERE payment_status='PAID' AND created_at >= ? AND created_at <= ?", startTime, endTime).Scan(&salesReport.TotalSales)
	if result.Error != nil {
		return models.SalesReport{}, result.Error
	}
	result = ad.DB.Raw("SELECT COUNT(*) FROM orders").Scan(&salesReport.TotalOrders)
	if result.Error != nil {
		return models.SalesReport{}, result.Error
	}
	result = ad.DB.Raw("SELECT COUNT(*) FROM orders WHERE payment_status = 'PAID' and created_at >= ? AND created_at <= ?", startTime, endTime).Scan(&salesReport.CompletedOrders)
	if result.Error != nil {
		return models.SalesReport{}, result.Error
	}
	result = ad.DB.Raw("SELECT COUNT(*) FROM orders WHERE order_status = 'PENDING'  AND created_at >= ? AND created_at<=?", startTime, endTime).Scan(&salesReport.PendingOrders)
	if result.Error != nil {
		return models.SalesReport{}, result.Error
	}
	var productID int
	result = ad.DB.Raw("SELECT inventory_id FROM order_items GROUP BY inventory_id order by SUM(quantity) DESC LIMIT 1").Scan(&productID)
	if result.Error != nil {
		return models.SalesReport{}, result.Error
	}
	result = ad.DB.Raw("SELECT product_name FROM inventories WHERE id = ?", productID).Scan(&salesReport.TrendingProduct)
	if result.Error != nil {
		return models.SalesReport{}, result.Error
	}
	return salesReport, nil
}