package interfaces

import "github.com/ahdaan98/pkg/utils/models"

type Helper interface {
	GenerateTokenAdmin(admin models.AdminDetailsResponse) (string, error)
	TwilioSetup(username string, password string)
	TwilioSendOTP(phone string, serviceID string) (string, error)
	TwilioVerifyOTP(serviceID string, code string, phone string) error
	GenerateTokenClients(user models.UserDetailsResponse) (string, error)
	PasswordHashing(string) (string, error)
	ValidateAlphabets(data string) (bool, error)
	CompareHashAndPassword(a string, b string) error
	ValidatePin(pin string) bool
	ValidatePhoneNumber(phone string) bool
	ValidateDatatype(data, intOrString string) (bool, error)
	ValidatePassword(pass string,cfm string) error
	ValidateName(name string) error
	ValidateEmail(email string) error
}
