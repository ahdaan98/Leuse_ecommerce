package handler

import (
	// "fmt"
	"net/http"
	"strconv"
	"time"

	"github.com/ahdaan98/pkg/helper"
	interfaces "github.com/ahdaan98/pkg/usecase/interface"
	"github.com/ahdaan98/pkg/utils/models"
	"github.com/ahdaan98/pkg/utils/response"
	"github.com/gin-gonic/gin"
)

type AdminHandler struct {
	usecase interfaces.AdminUseCase
}

func NewAdminHandler(usecase interfaces.AdminUseCase) *AdminHandler {
	return &AdminHandler{
		usecase: usecase,
	}
}

func (ad *AdminHandler) AdminLogin(c *gin.Context) {
	var adminDetails models.AdminLogin
	if err := c.ShouldBindJSON(&adminDetails); err != nil {
		errRes := response.ClientResponse(http.StatusBadRequest, "details not in correct format", nil, err.Error())
		c.JSON(http.StatusBadRequest, errRes)
		return
	}

	admin, err := ad.usecase.AdminLogin(adminDetails)
	if err != nil {
		errRes := response.ClientResponse(http.StatusBadRequest, "cannot authenticate user", nil, err.Error())
		c.JSON(http.StatusBadRequest, errRes)
		return
	}

	c.SetSameSite(http.SameSiteLaxMode)
	c.SetCookie("admin", admin.Token, 3600*24*30, "/", "", false, true)
	// c.Set("Refresh", admin.RefreshToken)

	successRes := response.ClientResponse(http.StatusOK, "Admin authenticated successfully", admin.Admin, nil)
	c.JSON(http.StatusOK, successRes)

}

func (ad *AdminHandler) GetUserByID(c *gin.Context) {
	id, _ := strconv.Atoi(c.Query("id"))

	user, err := ad.usecase.GetUserByID(id)
	if err != nil {
		errRes := response.ClientResponse(http.StatusBadRequest, "failed to login", nil, err)
		c.JSON(http.StatusBadRequest, errRes)
		return
	}

	successRes := response.ClientResponse(http.StatusOK, "successfully get the user", user, nil)
	c.JSON(http.StatusOK, successRes)
}

func (ad *AdminHandler) GetUsers(c *gin.Context) {
	users, err := ad.usecase.GetUsers()
	if err != nil {
		errRes := response.ClientResponse(http.StatusBadRequest, "failed to retrieve users", nil, err)
		c.JSON(http.StatusBadRequest, errRes)
		return
	}

	successRes := response.ClientResponse(http.StatusOK, "successfully login", users, nil)
	c.JSON(http.StatusOK, successRes)
}

func (ad *AdminHandler) BlockUser(c *gin.Context) {
	id, _ := strconv.Atoi(c.Query("id"))
	err:=ad.usecase.Blockuser(id)
	if err != nil {
		errRes := response.ClientResponse(http.StatusBadRequest, "failed to block", nil, err)
		c.JSON(http.StatusBadRequest, errRes)
		return
	}

	successRes := response.ClientResponse(http.StatusOK, "successfully blocked", nil, nil)
	c.JSON(http.StatusOK, successRes)
}

func (ad *AdminHandler) UnBlockUser(c *gin.Context) {
	id, _ := strconv.Atoi(c.Query("id"))
	err:=ad.usecase.UnBlockUser(id)
	if err != nil {
		errRes := response.ClientResponse(http.StatusBadRequest, "failed to block", nil, err)
		c.JSON(http.StatusBadRequest, errRes)
		return
	}

	successRes := response.ClientResponse(http.StatusOK, "successfully blocked", nil, nil)
	c.JSON(http.StatusOK, successRes)
}

func (i *AdminHandler) NewPaymentMethod(c *gin.Context) {

	var method models.NewPaymentMethod
	if err := c.BindJSON(&method); err != nil {
		errorRes := response.ClientResponse(http.StatusBadRequest, "fields provided are in wrong format", nil, err.Error())
		c.JSON(http.StatusBadRequest, errorRes)
		return
	}

	err := i.usecase.NewPaymentMethod(method.PaymentMethod)
	if err != nil {
		errorRes := response.ClientResponse(http.StatusBadRequest, "Could not add the payment method", nil, err.Error())
		c.JSON(http.StatusBadRequest, errorRes)
		return
	}

	successRes := response.ClientResponse(http.StatusOK, "Successfully added Payment Method", nil, nil)
	c.JSON(http.StatusOK, successRes)

}

func (a *AdminHandler) ListPaymentMethods(c *gin.Context) {

	categories, err := a.usecase.ListPaymentMethods()
	if err != nil {
		errorRes := response.ClientResponse(http.StatusInternalServerError, "fields provided are in wrong format", nil, err.Error())
		c.JSON(http.StatusBadRequest, errorRes)
		return
	}

	successRes := response.ClientResponse(http.StatusOK, "Successfully got all payment methods", categories, nil)
	c.JSON(http.StatusOK, successRes)

}

func (a *AdminHandler) DeletePaymentMethod(c *gin.Context) {

	id, err := strconv.Atoi(c.Query("id"))
	if err != nil {
		errorRes := response.ClientResponse(http.StatusBadRequest, "fields provided are in wrong format", nil, err.Error())
		c.JSON(http.StatusBadRequest, errorRes)
		return
	}

	err = a.usecase.DeletePaymentMethod(id)
	if err != nil {
		errorRes := response.ClientResponse(http.StatusInternalServerError, "error in deleting data", nil, err.Error())
		c.JSON(http.StatusBadRequest, errorRes)
		return
	}

	successRes := response.ClientResponse(http.StatusOK, "Successfully deleted the payment method", nil, nil)
	c.JSON(http.StatusOK, successRes)

}

func (a *AdminHandler) DashBoard(c *gin.Context) {
	dashBoard, err := a.usecase.DashBoard()
	if err != nil {
		errRes := response.ClientResponse(http.StatusBadRequest, "error in getting dashboard details", nil, err.Error())
		c.JSON(http.StatusBadRequest, errRes)
		return
	}

	sucessRes := response.ClientResponse(http.StatusOK, "succesfully recevied all records", dashBoard, nil)
	c.JSON(http.StatusOK, sucessRes)
}

func (a *AdminHandler) SalesByDate(c *gin.Context) {
	year := c.Query("year")
	yearInt, err := strconv.Atoi(year)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "error in getting year"})
		return
	}

	month := c.Query("month")
	monthInt, err := strconv.Atoi(month)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "error in getting month"})
		return
	}

	day := c.Query("day")
	dayInt, err := strconv.Atoi(day)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "error in getting day"})
		return
	}

	body, err := a.usecase.SalesByDate(dayInt, monthInt, yearInt)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "error in getting sales details"})
		return
	}

	download := c.Query("download")
	if download == "pdf" {
		pdf, err := a.usecase.PrintSalesReport(body)
		if err != nil {
			c.JSON(http.StatusBadGateway, gin.H{"error": "error in printing sales report"})
			return
		}
		err = pdf.OutputFileAndClose("SalesReport.pdf")
		if err != nil {
			c.JSON(http.StatusBadGateway, gin.H{"error": "error in printing sales report"})
			return
		}
		c.File("SalesReport.pdf")
		return
	}

	excel, err := helper.ConvertToExel(body)
	if err != nil {
		c.JSON(http.StatusBadGateway, gin.H{"error": "error in printing sales report"})
		return
	}

	fileName := "sales_report.xlsx"
	filePath := "pkg/tmp/" + fileName // Set your desired directory for temporary file storage

	if err := excel.SaveAs(filePath); err != nil {
		c.JSON(http.StatusBadGateway, gin.H{"error": "error in serving the sales report"})
		return
}

	// Send HTML response with download button
	htmlData := gin.H{
		"FileName": fileName,
	}
	c.HTML(http.StatusOK, "salesReport.html", htmlData)
}

/*
func DownloadExcel(c *gin.Context) {
	// Set response headers for file download
	c.Header("Content-Disposition", "attachment; filename=sales_report.xlsx")
	c.Header("Content-Type", "application/vnd.openxmlformats-officedocument.spreadsheetml.sheet")

	// Serve the Excel file for download
	c.File("pkg/tmp/sales_report.xlsx")
}

*/
func DownloadExcel(c *gin.Context) {
	fileName := "sales_report.xlsx"
	filePath := "pkg/tmp/" + fileName // Set your desired directory for temporary file storage

	// Set response headers for file download
	c.Header("Content-Disposition", "attachment; filename="+fileName)
	c.Header("Content-Type", "application/vnd.openxmlformats-officedocument.spreadsheetml.sheet")

	// Serve the file for download
	c.File(filePath)
}

/*
func (a *AdminHandler) SalesByDate(c *gin.Context) {
	year := c.Query("year")
	yearInt, err := strconv.Atoi(year)

	if err != nil {
		errRes := response.ClientResponse(http.StatusBadRequest, "error in getting year", nil, err.Error())
		c.JSON(http.StatusBadRequest, errRes)
		return
	}

	month := c.Query("month")
	monthInt, err := strconv.Atoi(month)

	if err != nil {
		errRes := response.ClientResponse(http.StatusBadRequest, "error in getting month", nil, err.Error())
		c.JSON(http.StatusBadRequest, errRes)
		return
	}

	day := c.Query("day")
	dayInt, err := strconv.Atoi(day)

	if err != nil {
		errRes := response.ClientResponse(http.StatusBadRequest, "error in getting day", nil, err.Error())
		c.JSON(http.StatusBadRequest, errRes)
		return
	}

	body, err := a.usecase.SalesByDate(dayInt, monthInt, yearInt)

	fmt.Println("body handler", dayInt)
	fmt.Println("body handler", monthInt)
	fmt.Println("body handler", yearInt)

	fmt.Println("body ", body)

	if err != nil {
		errRes := response.ClientResponse(http.StatusBadRequest, "error in getting sales details", nil, err.Error())
		c.JSON(http.StatusBadRequest, errRes)
		return
	}

	download := c.Query("download")
	if download == "pdf" {
		pdf, err := a.usecase.PrintSalesReport(body)
		if err != nil {
			errRes := response.ClientResponse(http.StatusBadGateway, "error in printing sales report", nil, err)
			c.JSON(http.StatusBadRequest, errRes)
			return
		}
		c.Header("Content-Disposition", "attachment;filename=totalsalesreport.pdf")

		pdfFilePath := "salesReport/totalsalesreport.pdf"

		err = pdf.OutputFileAndClose(pdfFilePath)
		if err != nil {
			errRes := response.ClientResponse(http.StatusBadGateway, "error in printing sales report", nil, err)
			c.JSON(http.StatusBadRequest, errRes)
			return
		}

		c.Header("Content-Disposition", "attachment; filename=total_sales_report.pdf")
		c.Header("Content-Type", "application/pdf")

		c.File(pdfFilePath)

		c.Header("Content-Type", "application/pdf")

		err = pdf.Output(c.Writer)
		if err != nil {
			errRes := response.ClientResponse(http.StatusBadGateway, "error in printing sales report", nil, err)
			c.JSON(http.StatusBadRequest, errRes)
			return
		}
	} else {
		fmt.Println("body ", body)
		excel, err := helper.ConvertToExel(body)
		if err != nil {
			errRes := response.ClientResponse(http.StatusBadGateway, "error in printing sales report", nil, err)
			c.JSON(http.StatusBadRequest, errRes)
			return
		}

		fileName := "sales_report.xlsx"

		c.Header("Content-Disposition", "attachment; filename="+fileName)
		c.Header("Content-Type", "application/vnd.openxmlformats-officedocument.spreadsheetml.sheet")

		if err := excel.Write(c.Writer); err != nil {
			errRes := response.ClientResponse(http.StatusBadGateway, "Error in serving the sales report", nil, err)
			c.JSON(http.StatusBadRequest, errRes)
			return
		}
	}

	succesRes := response.ClientResponse(http.StatusOK, "success", body, nil)
	c.JSON(http.StatusOK, succesRes)
}
*/

func (ad *AdminHandler) CustomSalesReport(c *gin.Context) {
	startDateStr := c.Query("start")
	endDateStr := c.Query("end")
	if startDateStr == "" || endDateStr == "" {
		err := response.ClientResponse(http.StatusBadRequest, "start or end date is empty", nil, "Empty date string")
		c.JSON(http.StatusBadRequest, err)
		return
	}
	startDate, err := time.Parse("02-01-2006", startDateStr)
	if err != nil {
		err := response.ClientResponse(http.StatusBadRequest, "start date conversion failed", nil, err.Error())
		c.JSON(http.StatusBadRequest, err)
		return
	}
	endDate, err := time.Parse("02-01-2006", endDateStr)
	if err != nil {
		err := response.ClientResponse(http.StatusBadRequest, "end date conversion failed", nil, err.Error())
		c.JSON(http.StatusBadRequest, err)
		return
	}

	if startDate.After(endDate) {
		err := response.ClientResponse(http.StatusBadRequest, "start date is after end date", nil, "Invalid date range")
		c.JSON(http.StatusBadRequest, err)
		return
	}

	report, err := ad.usecase.CustomSalesReportByDate(startDate, endDate)
	if err != nil {
		errorRes := response.ClientResponse(http.StatusInternalServerError, "sales report could not be retrieved", nil, err.Error())
		c.JSON(http.StatusInternalServerError, errorRes)
		return
	}

	success := response.ClientResponse(http.StatusOK, "custom report retrieved successfully", report, nil)
	c.JSON(http.StatusOK, success)
}
