package interfaces

import (
	"time"

	"github.com/ahdaan98/pkg/domain"
	"github.com/ahdaan98/pkg/utils/models"
	"github.com/jung-kurt/gofpdf"
)

type AdminUseCase interface {
	AdminLogin(admin models.AdminLogin) (domain.TokenAdmin, error)
	GetUsers() ([]models.UserDetailsAtAdmin, error)
	GetUserByID(id int) (models.UserDetailsAtAdmin, error)
	Blockuser(id int) error
	UnBlockUser(id int) error

	NewPaymentMethod(string) error
	ListPaymentMethods() ([]domain.PaymentMethod, error)
	DeletePaymentMethod(id int) error

	DashBoard() (models.CompleteAdminDashboard, error)
	SalesByDate(dayInt int, monthInt int, yearInt int) ([]models.OrderDetailsAdmin, error)
	PrintSalesReport(sales []models.OrderDetailsAdmin) (*gofpdf.Fpdf, error)
	CustomSalesReportByDate(startDate, endDate time.Time) (models.SalesReport, error)
}