package handler

import (
	"net/http"
	"strconv"

	"github.com/ahdaan98/pkg/config"
	helper "github.com/ahdaan98/pkg/helper/interfaces"
	interfaces "github.com/ahdaan98/pkg/usecase/interface"
	"github.com/ahdaan98/pkg/utils/models"
	"github.com/ahdaan98/pkg/utils/response"
	"github.com/gin-gonic/gin"
)

type UserHandler struct {
	userUseCase interfaces.UserUseCase
	helper  helper.Helper
}

func NewUserHandler(usecase interfaces.UserUseCase, helper helper.Helper) *UserHandler {
	return &UserHandler{
		userUseCase: usecase,
		helper: helper,
	}
}
var userDetails = make(map[string]string)

func (u *UserHandler) UserSignUp(c *gin.Context) {
    var user models.UserSignUp

    // Bind the JSON request body to the user model
    if err := c.ShouldBindJSON(&user); err != nil {
        errRes := response.ClientResponse(http.StatusBadRequest, "fields are in wrong format", nil, err.Error())
        c.JSON(http.StatusBadRequest, errRes)
        return
    }

    // Check if the user already exists by email
    err := u.userUseCase.ValidatingDetails(user)
	if err!=nil{
		errRes := response.ClientResponse(http.StatusBadRequest, "failed to signup", nil, err.Error())
        c.JSON(http.StatusBadRequest, errRes)
        return
	}

    // Store user information in the map
    userDetails["phone"] = user.Phone
    userDetails["email"] = user.Email
    userDetails["password"] = user.Password
    userDetails["cpass"] = user.ConfirmPassword

    // Load Twilio configuration from environment variables
    cfg, _ := config.LoadEnvVariables()

    // Initialize Twilio client
    u.helper.TwilioSetup(cfg.ACCOUNTSID, cfg.AUTHTOKEN)

    // Retrieve phone number from the map
    phone := userDetails["phone"]

    // Generate OTP and send to user's phone
    _, err = u.helper.TwilioSendOTP(phone, cfg.SERVICESID)
    if err != nil {
        errRes := response.ClientResponse(http.StatusInternalServerError, "failed to send OTP", nil, err.Error())
        c.JSON(http.StatusInternalServerError, errRes)
        return
    }

    // Send success response
    successRes := response.ClientResponse(http.StatusOK, "OTP sent successfully", nil, nil)
    c.JSON(http.StatusOK, successRes)
}

func (u *UserHandler) VerifyOTP(c *gin.Context) {
    // Retrieve phone number from the map
    phone := userDetails["phone"]
    email := userDetails["email"]
    pass := userDetails["password"]
    cpass := userDetails["cpass"]

    // Retrieve OTP code from request body
    var req struct {
        Code string `json:"code"`
    }
    if err := c.ShouldBindJSON(&req); err != nil {
        errRes := response.ClientResponse(http.StatusBadRequest, "fields are in wrong format", nil, err.Error())
        c.JSON(http.StatusBadRequest, errRes)
        return
    }

    cfg, _ := config.LoadEnvVariables()

    u.helper.TwilioSetup(cfg.ACCOUNTSID, cfg.AUTHTOKEN)

    // Verify OTP
    err := u.helper.TwilioVerifyOTP(cfg.SERVICESID, req.Code, phone)
    if err != nil {
        errRes := response.ClientResponse(http.StatusUnauthorized, "failed to verify OTP", nil, err.Error())
        c.JSON(http.StatusUnauthorized, errRes)
        return
    }

    // Call UserSignUp method to sign up the user
    tokUser, err := u.userUseCase.UserSignUp(models.UserSignUp{Email: email, Phone: phone, Password: pass, ConfirmPassword: cpass})
    if err != nil {
        errRes := response.ClientResponse(http.StatusInternalServerError, "failed to sign up user", nil, err.Error())
        c.JSON(http.StatusInternalServerError, errRes)
        return
    }

    // Proceed with any additional logic after signing up the user
    // For example, you can send a success response or perform any other action

    successRes := response.ClientResponse(http.StatusOK, "user signed up successfully", tokUser.Users, nil)
    c.JSON(http.StatusOK, successRes)
}

func (u *UserHandler) UserLogin(c *gin.Context) {
	var user models.UserLogin

	if err := c.ShouldBindJSON(&user); err != nil {
		errRes := response.ClientResponse(http.StatusBadRequest, "field are in wrong format", nil, err.Error())
		c.JSON(http.StatusBadRequest, errRes)
		return
	}

	tokUser, err := u.userUseCase.UserLogin(user)
	if err != nil {
		errRes := response.ClientResponse(http.StatusBadRequest, "failed to login", nil, err.Error())
		c.JSON(http.StatusBadRequest, errRes)
		return
	}

	c.SetSameSite(http.SameSiteLaxMode)
	c.SetCookie("client", tokUser.Token, 3600*24*30, "/", "", false, true)

	successRes := response.ClientResponse(http.StatusOK, "successfully login", tokUser.Users, nil)
	c.JSON(http.StatusOK, successRes)
}

func (u *UserHandler) UserProfile(c *gin.Context) {
	id, _ := c.Get("id")
	k := id.(int)

	resp, err := u.userUseCase.UserProfile(k)
	if err != nil {
		errRes := response.ClientResponse(http.StatusBadRequest, "failed to get profile", nil, err)
		c.JSON(http.StatusBadRequest, errRes)
		return
	}

	successRes := response.ClientResponse(http.StatusOK, "profile", resp, nil)
	c.JSON(http.StatusOK, successRes)
}

func (u *UserHandler) EditUserProfile(c *gin.Context) {
	id, _ := c.Get("id")
	k := id.(int)

	var user models.EditUserDetails

	if err := c.ShouldBindJSON(&user); err != nil {
		errRes := response.ClientResponse(http.StatusBadRequest, "fields are in wrong format", nil, err)
		c.JSON(http.StatusBadRequest, errRes)
		return
	}

	resp, err := u.userUseCase.EditProfile(user, k)
	if err != nil {
		errRes := response.ClientResponse(http.StatusBadRequest, "failed to edit", nil, err)
		c.JSON(http.StatusBadRequest, errRes)
	}

	successRes := response.ClientResponse(http.StatusOK, "successfully edited", resp, nil)
	c.JSON(http.StatusOK, successRes)
}

func (u *UserHandler) AddAddress(c *gin.Context) {
	id, _ := c.Get("id")
	k := id.(int)

	var address models.AddAddress

	if err := c.ShouldBindJSON(&address); err != nil {
		errRes := response.ClientResponse(http.StatusBadRequest, "fields are in wrong format", nil, err)
		c.JSON(http.StatusBadRequest, errRes)
		return
	}

	if err := u.userUseCase.AddAddress(k, address); err != nil {
		errRes := response.ClientResponse(http.StatusBadRequest, "failed to add address", nil, err)
		c.JSON(http.StatusBadRequest, errRes)
		return
	}

	successRes := response.ClientResponse(http.StatusOK, "successfully added address", nil, nil)
	c.JSON(http.StatusOK, successRes)
}

func (u *UserHandler) GetAddress(c *gin.Context) {
	id, _ := c.Get("id")
	k := id.(int)

	addressess,err:=u.userUseCase.GetAddresses(k)
	if err!=nil{
		errRes := response.ClientResponse(http.StatusBadRequest, "failed to get address", nil, err)
		c.JSON(http.StatusBadRequest, errRes)
		return
	}

	successRes := response.ClientResponse(http.StatusOK, "successfully retrieved address", addressess, nil)
	c.JSON(http.StatusOK, successRes)
}

func (i *UserHandler) GetCart(c *gin.Context) {
	idString, _ := c.Get("id")
	id, _ := idString.(int)

	products, err := i.userUseCase.GetCart(id)
	if err != nil {
		errorRes := response.ClientResponse(http.StatusBadRequest, "could not retrieve cart", nil, err.Error())
		c.JSON(http.StatusBadRequest, errorRes)
		return
	}
	successRes := response.ClientResponse(http.StatusOK, "Successfully got all products in cart", products, nil)
	c.JSON(http.StatusOK, successRes)
}

func (i *UserHandler) RemoveFromCart(c *gin.Context) {

	idString, _ := c.Get("id")
	id, _ := idString.(int)

	InventoryID, err := strconv.Atoi(c.Query("inventory_id"))
	if err != nil {
		errorRes := response.ClientResponse(http.StatusBadRequest, "check parameters properly", nil, err.Error())
		c.JSON(http.StatusBadRequest, errorRes)
		return
	}

	if err := i.userUseCase.RemoveFromCart(id, InventoryID); err != nil {
		errorRes := response.ClientResponse(http.StatusBadRequest, "could not remove from cart", nil, err.Error())
		c.JSON(http.StatusBadRequest, errorRes)
		return
	}

	successRes := response.ClientResponse(http.StatusOK, "Successfully Removed product from cart", nil, nil)
	c.JSON(http.StatusOK, successRes)
}

func (i *UserHandler) UpdateQuantity(c *gin.Context) {
	idString, _ := c.Get("id")
	id, _ := idString.(int)

	
	inv, err := strconv.Atoi(c.Query("inventory"))
	if err != nil {
		errorRes := response.ClientResponse(http.StatusBadRequest, "check parameters properly", nil, err.Error())
		c.JSON(http.StatusBadRequest, errorRes)
		return
	}

	qty, err := strconv.Atoi(c.Query("quantity"))
	if err != nil {
		errorRes := response.ClientResponse(http.StatusBadRequest, "check parameters properly", nil, err.Error())
		c.JSON(http.StatusBadRequest, errorRes)
		return
	}

	if err := i.userUseCase.UpdateQuantity(id, inv, qty); err != nil {
		errorRes := response.ClientResponse(http.StatusBadRequest, "could not Add the quantity", nil, err.Error())
		c.JSON(http.StatusBadRequest, errorRes)
		return
	}

	successRes := response.ClientResponse(http.StatusOK, "Successfully added quantity", nil, nil)
	c.JSON(http.StatusOK, successRes)
}

func (i *UserHandler) ChangePassword(c *gin.Context) {

	idString, _ := c.Get("id")
	id, _ := idString.(int)

	var ChangePassword models.ChangePassword
	if err := c.BindJSON(&ChangePassword); err != nil {
		errorRes := response.ClientResponse(http.StatusBadRequest, "fields provided are in wrong format", nil, err.Error())
		c.JSON(http.StatusBadRequest, errorRes)
		return
	}

	if err := i.userUseCase.ChangePassword(id, ChangePassword.Oldpassword, ChangePassword.Password, ChangePassword.Repassword); err != nil {
		errorRes := response.ClientResponse(http.StatusBadRequest, "Could not change the password", nil, err.Error())
		c.JSON(http.StatusBadRequest, errorRes)
		return
	}

	successRes := response.ClientResponse(http.StatusOK, "password changed Successfully ", nil, nil)
	c.JSON(http.StatusOK, successRes)

}