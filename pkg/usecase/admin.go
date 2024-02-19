package usecase

import (
	"errors"
	"fmt"
	"strconv"
	"time"

	"github.com/ahdaan98/pkg/domain"
	helper "github.com/ahdaan98/pkg/helper/interfaces"
	interfaces "github.com/ahdaan98/pkg/repository/interface"
	usecase "github.com/ahdaan98/pkg/usecase/interface"
	"github.com/ahdaan98/pkg/utils/models"
	"github.com/jung-kurt/gofpdf"
)

type AdminUseCase struct {
	repo   interfaces.AdminRepository
	helper helper.Helper
}

func NewAdminUseCase(repo interfaces.AdminRepository,helper helper.Helper) usecase.AdminUseCase {
	return &AdminUseCase{
		repo: repo,
		helper: helper,
	}
}

func (ad *AdminUseCase) AdminLogin(admin models.AdminLogin) (domain.TokenAdmin, error) {
	Exist, err := ad.repo.CheckAdminExist(admin.Email)
	if err != nil {
		return domain.TokenAdmin{}, err
	}

	if !Exist {
		return domain.TokenAdmin{}, errors.New("admin does not exist")
	}

	resp, err := ad.repo.GetAdminByEmail(admin.Email)
	if err != nil {
		return domain.TokenAdmin{}, err
	}

	pass, err := ad.repo.GetAdminPassword(resp.Email)
	if err != nil {
		return domain.TokenAdmin{}, err
	}

	if admin.Password != pass {
		return domain.TokenAdmin{}, errors.New("incorrect password")
	}

	token, err := ad.helper.GenerateTokenAdmin(resp)
	if err != nil {
		return domain.TokenAdmin{}, err
	}

	return domain.TokenAdmin{
		Admin: resp,
		Token: token,
	}, nil
}

func (ad *AdminUseCase) GetUsers() ([]models.UserDetailsAtAdmin, error) {
	users, err := ad.repo.GetUsers()
	if err != nil {
		return []models.UserDetailsAtAdmin{}, err
	}

	return users, err
}

func (ad *AdminUseCase) GetUserByID(id int) (models.UserDetailsAtAdmin, error) {
	user, err := ad.repo.GetUserByID(id)
	if err != nil {
		return models.UserDetailsAtAdmin{}, err
	}

	return user, nil
}

func (ad *AdminUseCase) Blockuser(id int) error {
	err := ad.repo.UpdateBlockUserByID(true, id)
	if err != nil {
		return err
	}
	return nil
}

func (ad *AdminUseCase) UnBlockUser(id int) error {
	err := ad.repo.UpdateBlockUserByID(false, id)
	if err != nil {
		return err
	}
	return nil
}


func (i *AdminUseCase) NewPaymentMethod(id string) error {

	// parsedID, err := strconv.Atoi(id)
	// if err != nil || parsedID <= 0 {
	// 	return errors.New("invalid id")
	// }

	exists, err := i.repo.CheckIfPaymentMethodAlreadyExists(id)
	if err != nil {
		return err
	}

	if exists {
		return errors.New("payment method already exists")
	}

	err = i.repo.NewPaymentMethod(id)
	if err != nil {
		return err
	}

	return nil

}

func (a *AdminUseCase) ListPaymentMethods() ([]domain.PaymentMethod, error) {

	categories, err := a.repo.ListPaymentMethods()
	if err != nil {
		return []domain.PaymentMethod{}, err
	}
	return categories, nil

}

func (a *AdminUseCase) DeletePaymentMethod(id int) error {

	if id <= 0 {
		return errors.New("invalid page number")
	}

	err := a.repo.DeletePaymentMethod(id)
	if err != nil {
		return err
	}
	return nil

}

func (ad *AdminUseCase) DashBoard() (models.CompleteAdminDashboard, error) {
	userDetails, err := ad.repo.DashBoardUserDetails()
	if err != nil {
		return models.CompleteAdminDashboard{}, err
	}
	productDetails, err := ad.repo.DashBoardProductDetails()
	if err != nil {
		return models.CompleteAdminDashboard{}, err
	}
	orderDetails, err := ad.repo.DashBoardOrder()
	if err != nil {
		return models.CompleteAdminDashboard{}, err
	}
	totalRevenue, err := ad.repo.TotalRevenue()
	if err != nil {
		return models.CompleteAdminDashboard{}, err
	}
	amountDetails, err := ad.repo.AmountDetails()
	if err != nil {
		return models.CompleteAdminDashboard{}, err
	}
	return models.CompleteAdminDashboard{
		DashboardUser:    userDetails,
		DashboardProduct: productDetails,
		DashboardOrder:   orderDetails,
		DashboardRevenue: totalRevenue,
		DashboardAmount:  amountDetails,
	}, nil
}
func (ad *AdminUseCase) SalesByDate(dayInt int, monthInt int, yearInt int) ([]models.OrderDetailsAdmin, error) {

	if dayInt == 0 && monthInt == 0 && yearInt == 0 {
		return []models.OrderDetailsAdmin{}, errors.New("must enter a value for day, month, and year")
	}

	if dayInt < 0 || monthInt < 0 || yearInt < 0 {
		return []models.OrderDetailsAdmin{}, errors.New("no such values are allowded")
	}

	if yearInt >= 2020 {
		if monthInt == 0 && dayInt == 0 {

			body, err := ad.repo.SalesByYear(yearInt)
			if err != nil {
				return []models.OrderDetailsAdmin{}, err
			}
			fmt.Println("body at usecase year", body)
			return body, nil
		} else if monthInt > 0 && monthInt <= 12 && dayInt == 0 {

			body, err := ad.repo.SalesByMonth(yearInt, monthInt)
			if err != nil {
				return []models.OrderDetailsAdmin{}, err
			}
			fmt.Println("body at usecase month", body)
			return body, nil
		} else if monthInt > 0 && monthInt <= 12 && dayInt > 0 && dayInt <= 31 {

			body, err := ad.repo.SalesByDay(yearInt, monthInt, dayInt)
			if err != nil {
				return []models.OrderDetailsAdmin{}, err
			}
			fmt.Println("body at usecase day", body)
			return body, nil
		}
	}

	return []models.OrderDetailsAdmin{}, errors.New("invalid date parameters")
}

func (ad *AdminUseCase) PrintSalesReport(sales []models.OrderDetailsAdmin) (*gofpdf.Fpdf, error) {

	pdf := gofpdf.New("P", "mm", "A4", "")

	pdf.AddPage()
	pdf.SetFont("Arial", "B", 22)
	pdf.SetTextColor(31, 73, 125)
	pdf.CellFormat(0, 20, "Total Sales Report", "0", 1, "C", false, 0, "")

	pdf.SetFont("Arial", "", 16)
	pdf.SetTextColor(0, 0, 0)

	for _, item := range sales {
		pdf.CellFormat(0, 10, "Product: "+item.ProductName, "0", 1, "L", false, 0, "")
		amount := strconv.FormatFloat(item.TotalAmount, 'f', 2, 64)
		pdf.CellFormat(0, 10, "Amount Sold: $"+amount, "0", 1, "L", false, 0, "")
		pdf.Ln(5)
	}

	pdf.SetFont("Arial", "I", 12)
	pdf.SetTextColor(150, 150, 150)

	pdf.Cell(0, 10, "Generated by Leuse India Pvt Ltd. - "+time.Now().Format("2006-01-02 15:04:05"))
 
	return pdf, nil
}

func (ad *AdminUseCase) CustomSalesReportByDate(startDate, endDate time.Time) (models.SalesReport, error) {
	orders, err := ad.repo.CustomSalesReportByDate(startDate, endDate)
	if err != nil {
		return models.SalesReport{}, errors.New("report fetching failed")
	}
	return orders, nil
}