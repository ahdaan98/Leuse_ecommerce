package interfaces

import (
	"time"

	"github.com/ahdaan98/pkg/domain"
	"github.com/ahdaan98/pkg/utils/models"
)

type AdminRepository interface {
	GetUsers()([]models.UserDetailsAtAdmin,error)
	GetUserByID(id int) (models.UserDetailsAtAdmin,error)
	UpdateBlockUserByID(k bool,id int) (error)
	GetAdminPassword(email string) (string,error)
	GetAdminByEmail(email string) (models.AdminDetailsResponse,error)
	CheckAdminExist(email string) (bool,error)

	NewPaymentMethod(string) error
	ListPaymentMethods() ([]domain.PaymentMethod, error)
	GetPaymentMethod() ([]models.PaymentMethodResponse, error)
	CheckIfPaymentMethodAlreadyExists(payment string) (bool, error)
	DeletePaymentMethod(id int) error

	TotalRevenue() (models.DashboardRevenue, error)
	DashBoardOrder() (models.DashboardOrder, error)
	AmountDetails() (models.DashboardAmount, error)
	DashBoardUserDetails() (models.DashBoardUser, error)
	DashBoardProductDetails() (models.DashBoardProduct, error)

	SalesByYear(yearInt int) ([]models.OrderDetailsAdmin, error)
	SalesByMonth(yearInt int, monthInt int) ([]models.OrderDetailsAdmin, error)
	SalesByDay(yearInt int, monthInt int, dayInt int) ([]models.OrderDetailsAdmin, error)

	CustomSalesReportByDate(startTime time.Time, endTime time.Time) (models.SalesReport, error)
}