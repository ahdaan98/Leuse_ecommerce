package helper

import (
	"errors"
	"fmt"
	"reflect"
	"regexp"
	"strconv"
	"time"
	"unicode"

	"github.com/360EntSecGroup-Skylar/excelize"
	cfg "github.com/ahdaan98/pkg/config"
	"github.com/ahdaan98/pkg/helper/interfaces"
	"github.com/ahdaan98/pkg/utils/models"
	"github.com/golang-jwt/jwt"
	"github.com/twilio/twilio-go"
	twilioApi "github.com/twilio/twilio-go/rest/verify/v2"
	"golang.org/x/crypto/bcrypt"
)

type helper struct {
	cfg cfg.Config
}

func NewHelper(config cfg.Config) interfaces.Helper {
	return &helper{
		cfg: config,
	}
}

var client *twilio.RestClient

type AuthCustomClaims struct {
	Id    int    `json:"id"`
	Email string `json:"email"`
	Role  string `json:"role"`
	jwt.StandardClaims
}

func (h *helper) GenerateTokenAdmin(admin models.AdminDetailsResponse) (string, error) {
	claims := &AuthCustomClaims{
		Id:    admin.ID,
		Email: admin.Email,
		Role:  "admin",
		StandardClaims: jwt.StandardClaims{
			ExpiresAt: time.Now().Add(time.Hour * 24).Unix(),
			IssuedAt:  time.Now().Unix(),
		},
	}

	token := jwt.NewWithClaims(jwt.SigningMethodHS256, claims)
	tokenString, err := token.SignedString([]byte(h.cfg.ACCESS_KEY_ADMIN))

	if err != nil {
		return "", err
	}

	return tokenString, nil
}

func (h *helper) TwilioSetup(username string, password string) {
	client = twilio.NewRestClientWithParams(twilio.ClientParams{
		Username: username,
		Password: password,
	})

}

func (h *helper) TwilioSendOTP(phone string, serviceID string) (string, error) {
	to := "+91" + phone
	params := &twilioApi.CreateVerificationParams{}
	params.SetTo(to)
	params.SetChannel("sms")

	resp, err := client.VerifyV2.CreateVerification(serviceID, params)
	if err != nil {

		return " ", err
	}

	return *resp.Sid, nil

}

func (h *helper) TwilioVerifyOTP(serviceID string, code string, phone string) error {

	params := &twilioApi.CreateVerificationCheckParams{}
	params.SetTo("+91" + phone)
	params.SetCode(code)
	resp, err := client.VerifyV2.CreateVerificationCheck(serviceID, params)

	if err != nil {
		return err
	}

	if *resp.Status == "approved" {
		return nil
	}

	return errors.New("failed to validate otp")
}

func (h *helper) GenerateTokenClients(user models.UserDetailsResponse) (string, error) {
	claims := &AuthCustomClaims{
		Id:    user.Id,
		Email: user.Email,
		Role:  "client",
		StandardClaims: jwt.StandardClaims{
			ExpiresAt: time.Now().Add(time.Hour * 24).Unix(),
			IssuedAt:  time.Now().Unix(),
		},
	}

	token := jwt.NewWithClaims(jwt.SigningMethodHS256, claims)
	tokenString, err := token.SignedString([]byte(h.cfg.ACCESS_KEY_USER))

	if err != nil {
		return "", err
	}

	return tokenString, nil
}

func (h *helper) PasswordHashing(password string) (string, error) {

	hashedPassword, err := bcrypt.GenerateFromPassword([]byte(password), 10)
	if err != nil {
		return "", errors.New("internal server error")
	}

	hash := string(hashedPassword)
	return hash, nil
}

func (h *helper) CompareHashAndPassword(a string, b string) error {
	err := bcrypt.CompareHashAndPassword([]byte(a), []byte(b))
	if err != nil {
		return err
	}
	return nil
}

func (h *helper) ValidatePhoneNumber(phone string) bool {
	pattern := `^\d{10}$`
	regex := regexp.MustCompile(pattern)
	return regex.MatchString(phone)
}

func (h *helper) ValidatePin(pin string) bool {

	match, _ := regexp.MatchString(`^\d{4}(\d{2})?$`, pin)
	return match

}

func (h *helper) ValidateDatatype(data, intOrString string) (bool, error) {

	switch intOrString {
	case "int":
		if _, err := strconv.Atoi(data); err != nil {
			return false, errors.New("data is not an integer")
		}
		return true, nil
	case "string":
		kind := reflect.TypeOf(data).Kind()
		return kind == reflect.String, nil
	default:
		return false, errors.New("data is not" + intOrString)
	}

}

func (h *helper) ValidatePassword(pass string, cfm string) error {
	if len(pass) < 8 {
		return errors.New("password must be at least 8 characters long")
	}

	if len(pass) > 20 {
		return errors.New("password can be maximum of 20 characters long")
	}

	// Regular expressions for password criteria
	containsLower := regexp.MustCompile(`[a-z]`).MatchString(pass)
	containsUpper := regexp.MustCompile(`[A-Z]`).MatchString(pass)
	containsDigit := regexp.MustCompile(`[0-9]`).MatchString(pass)
	containsSpecial := regexp.MustCompile(`[!@#$%^&*()_+\-~=[\]{}|;:'",.<>/?]`).MatchString(pass)

	if !containsLower {
		return errors.New("password must contain at least one lowercase letter")
	}

	if !containsUpper {
		return errors.New("password must contain at least one uppercase letter")
	}

	if !containsDigit {
		return errors.New("password must contain at least one digit")
	}

	if !containsSpecial {
		return errors.New("password must contain at least one special character")
	}

	if pass != cfm {
		return errors.New("passwords do not match")
	}

	return nil
}

func (h *helper) ValidateAlphabets(data string) (bool, error) {
	for _, char := range data {
		if !unicode.IsLetter(char) {
			return false, errors.New("data contains non-alphabetic characters")
		}
	}
	return true, nil
}

func (h *helper) ValidateName(name string) error {
	if len(name) < 4 {
		return errors.New("length of the username must be atleast 4 character")
	}
	if len(name) > 8 {
		return errors.New("length of the username can only be upto 8 character")
	}

	if len(name) == 0 ||name[0] <'A' || name[0]>'Z'{
		return errors.New("first letter should be capital")
	}


	pattern := "^[a-zA-Z0-9_]{4,8}$"
	regex := regexp.MustCompile(pattern)
	if !regex.MatchString(name){
		return errors.New("invalid username that")
	}

	return nil
}

func (h *helper) ValidateEmail(email string) error {
	if len(email) == 0 {
		return errors.New("email cannot be empty")
	}

	pattern := `^[a-zA-Z0-9._%+-]+@[a-zA-Z0-9.-]+\.[a-zA-Z]{2,}$`
	regex := regexp.MustCompile(pattern)
	if !regex.MatchString(email) {
		return errors.New("invalid email format")
	}

	return nil
}

func ConvertToExel(sales []models.OrderDetailsAdmin) (*excelize.File, error) {
	file := excelize.NewFile()

	// Set column headers
	headers := []string{"Item", "Total Amount Sold"}
	for col, header := range headers {
		cell := fmt.Sprintf("%s%d", string(rune('A'+col)), 1)
		file.SetCellValue("Sheet1", cell, header)
	}

	// Set sales data
	for i, sale := range sales {
		row := i + 2 // Start from row 2
		file.SetCellValue("Sheet1", fmt.Sprintf("A%d", row), sale.ProductName)
		file.SetCellValue("Sheet1", fmt.Sprintf("B%d", row), sale.TotalAmount)
	}

	return file, nil
}

func (h *helper) GetTimeFromPeriod(timePeriod string) (time.Time, time.Time) {

	endDate := time.Now()

	if timePeriod == "week" {
		startDate := endDate.AddDate(0, 0, -6)
		return startDate, endDate
	}

	if timePeriod == "month" {
		startDate := endDate.AddDate(0, -1, 0)
		return startDate, endDate
	}

	if timePeriod == "year" {
		startDate := endDate.AddDate(0, -1, 0)
		return startDate, endDate
	}

	return endDate.AddDate(0, 0, -6), endDate

}
