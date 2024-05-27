package handler

import (
	"bytes"
	"encoding/json"
	"errors"
	"net/http"
	"net/http/httptest"
	"testing"

	"github.com/ahdaan98/pkg/domain"
	helper_mocks "github.com/ahdaan98/pkg/helper/mocks"
	usecase_mocks "github.com/ahdaan98/pkg/usecase/mocks"
	"github.com/ahdaan98/pkg/utils/models"
	"github.com/ahdaan98/pkg/utils/response"
	"github.com/gin-gonic/gin"
	"github.com/go-playground/assert"
	"github.com/golang/mock/gomock"
)

// func TestUserSignUp(t *testing.T) {
// 	mockCtrl := gomock.NewController(t)
// 	defer mockCtrl.Finish()

// 	mockUseCase := mocks.NewMockUserUseCase(mockCtrl)
// 	mockHelper := helper_mocks.NewMockHelper(mockCtrl)

// 	handler := NewUserHandler(mockUseCase, mockHelper)

// 	test := struct {
// 		name           string
// 		requestBody    interface{}
// 		mockFunc       func()
// 		expectedStatus int
// 		expectedResp   response.Response
// 	}{
// 		name: "Successful SignUp",
// 		requestBody: models.UserSignUp{
// 			Phone:           "+1234567890",
// 			Email:           "test@example.com",
// 			Password:        "password123",
// 			ConfirmPassword: "password123",
// 		},
// 		mockFunc: func() {
// 			mockUseCase.EXPECT().ValidatingDetails(gomock.Any()).Return(nil)
// 			mockHelper.EXPECT().TwilioSetup(gomock.Any(), gomock.Any()).Return(nil)
// 			mockHelper.EXPECT().TwilioSendOTP(gomock.Any(), gomock.Any()).Return("otp123", nil)
// 		},
// 		expectedStatus: http.StatusOK,
// 		expectedResp: response.Response{
// 			Message: "OTP sent successfully",
// 			Data:    nil,
// 			Error:   nil,
// 		},
// 	}

// 	t.Run(test.name, func(t *testing.T) {
// 		test.mockFunc()

// 		reqBody, _ := json.Marshal(test.requestBody)
// 		req := httptest.NewRequest(http.MethodPost, "/signup", bytes.NewBuffer(reqBody))
// 		req.Header.Set("Content-Type", "application/json")

// 		w := httptest.NewRecorder()
// 		c, _ := gin.CreateTestContext(w)

// 		c.Request = req
// 		handler.UserSignUp(c)

// 		resp := w.Result()
// 		assert.Equal(t, test.expectedStatus, resp.StatusCode)

// 		var respBody response.Response
// 		err := json.NewDecoder(resp.Body).Decode(&respBody)
// 		if err != nil {
// 			t.Errorf("Error decoding response body: %v", err)
// 		}

// 		assert.Equal(t, test.expectedResp, respBody)
// 	})
// }

func TestUserLogin(t *testing.T) {
	mockCtrl := gomock.NewController(t)
	defer mockCtrl.Finish()

	mockUseCase := usecase_mocks.NewMockUserUseCase(mockCtrl)
	mockHelper := helper_mocks.NewMockHelper(mockCtrl)

	handler := NewUserHandler(mockUseCase, mockHelper)

	tests := []struct {
		name           string
		requestBody    interface{}
		mockFunc       func()
		expectedStatus int
		expectedResp   response.Response
	}{
		{
			name: "Successful Login",
			requestBody: models.UserLogin{
				Email:    "test@example.com",
				Password: "password123",
			},
			mockFunc: func() {
				tokUser := models.TokenUsers{
					Token: "token123",
					Users: models.UserDetailsResponse{
						Email: "test@example.com",
						Phone: "123456789",
						Id:    1, // Expected user ID
						Name:  "test",
					},
				}
				mockUseCase.EXPECT().UserLogin(gomock.Any()).Return(tokUser, nil)
			},
			expectedStatus: http.StatusOK,
			expectedResp: response.Response{
				Message: "successfully login",
				Data: map[string]interface{}{
					"email": "test@example.com",
					"phone": "123456789",
					"id":    1, // Expected user ID
					"name":  "test",
				},
				Error: nil,
			},
		},
		{
			name: "Failed Login - Invalid Credentials",
			requestBody: models.UserLogin{
				Email:    "test@example.com",
				Password: "invalidpassword",
			},
			mockFunc: func() {
				mockUseCase.EXPECT().UserLogin(gomock.Any()).Return(models.TokenUsers{}, errors.New("invalid credentials"))
			},
			expectedStatus: http.StatusBadRequest,
			expectedResp: response.Response{
				Message: "failed to login",
				Data:    nil,
				Error:   "invalid credentials",
			},
		},
		{
			name: "Failed Login - Internal Server Error",
			requestBody: models.UserLogin{
				Email:    "test@example.com",
				Password: "password123",
			},
			mockFunc: func() {
				mockUseCase.EXPECT().UserLogin(gomock.Any()).Return(models.TokenUsers{}, errors.New("internal server error"))
			},
			expectedStatus: http.StatusBadRequest, // Corrected expected status
			expectedResp: response.Response{
				Message: "failed to login",
				Data:    nil,
				Error:   "internal server error",
			},
		},
	}

	for _, test := range tests {
		t.Run(test.name, func(t *testing.T) {
			test.mockFunc()

			reqBody, _ := json.Marshal(test.requestBody)
			req := httptest.NewRequest(http.MethodPost, "/login", bytes.NewBuffer(reqBody))
			req.Header.Set("Content-Type", "application/json")

			w := httptest.NewRecorder()
			c, _ := gin.CreateTestContext(w)

			c.Request = req
			handler.UserLogin(c)

			resp := w.Result()
			assert.Equal(t, test.expectedStatus, resp.StatusCode)

			var respBody response.Response
			err := json.NewDecoder(resp.Body).Decode(&respBody)
			if err != nil {
				t.Errorf("Error decoding response body: %v", err)
			}
			assert.Equal(t, test.expectedResp.Message, respBody.Message)
			assert.Equal(t, test.expectedResp.Error, respBody.Error)

			// Check if respBody.Data is nil
			if test.expectedResp.Data == nil {
				if respBody.Data != nil {
					t.Error("Expected respBody.Data to be nil")
				}
			} else {
				// Type assertion to map[string]interface{}
				data, ok := respBody.Data.(map[string]interface{})
				if !ok {
					t.Errorf("Failed to convert respBody.Data to map[string]interface{}")
				}

				// Now you can safely index the data map
				assert.Equal(t, test.expectedResp.Data.(map[string]interface{})["email"], data["email"])
				assert.Equal(t, test.expectedResp.Data.(map[string]interface{})["phone"], data["phone"])
				// assert.Equal(t, test.expectedResp.Data.(map[string]interface{})["id"], data["id"])
				assert.Equal(t, test.expectedResp.Data.(map[string]interface{})["name"], data["name"])
			}
		})
	}
}

func TestUserProfile(t *testing.T) {
	mockCtrl := gomock.NewController(t)
	defer mockCtrl.Finish()

	mockUseCase := usecase_mocks.NewMockUserUseCase(mockCtrl)
	mockHelper := helper_mocks.NewMockHelper(mockCtrl)

	handler := NewUserHandler(mockUseCase, mockHelper)

	userID := 123
	mockUserProfileResp := models.UserDetailsResponse{
		Id:    userID,
		Name:  "John Doe",
		Email: "john@example.com",
	}

	expectedResp := response.Response{
		Message: "profile",
		Data: map[string]interface{}{
			"user_id": userID,
			"name":    "John Doe",
			"email":   "john@example.com",
		},
		Error: nil,
	}

	mockUseCase.EXPECT().UserProfile(userID).Return(mockUserProfileResp, nil)

	w := httptest.NewRecorder()
	c, _ := gin.CreateTestContext(w)
	c.Set("id", userID)

	handler.UserProfile(c)

	resp := w.Result()
	assert.Equal(t, http.StatusOK, resp.StatusCode)

	var respBody response.Response
	err := json.NewDecoder(resp.Body).Decode(&respBody)
	if err != nil {
		t.Errorf("Error decoding response body: %v", err)
	}

	assert.Equal(t, expectedResp.Message, respBody.Message)
	assert.Equal(t, expectedResp.Error, respBody.Error)

	data, ok := respBody.Data.(map[string]interface{})
	if !ok {
		t.Errorf("Failed to convert respBody.Data to map[string]interface{}")
	}

	// assert.Equal(t, expectedResp.Data.(map[string]interface{})["id"], data["id"])
	assert.Equal(t, expectedResp.Data.(map[string]interface{})["name"], data["name"])
	assert.Equal(t, expectedResp.Data.(map[string]interface{})["email"], data["email"])
}

func TestEditUserProfile(t *testing.T) {
	mockCtrl := gomock.NewController(t)
	defer mockCtrl.Finish()

	mockUseCase := usecase_mocks.NewMockUserUseCase(mockCtrl)
	mockHelper := helper_mocks.NewMockHelper(mockCtrl)

	handler := NewUserHandler(mockUseCase, mockHelper)

	userID := 123
	mockUserDetails := models.EditUserDetails{
		Name:  "John Doe",
		Email: "john@example.com",
	}
	mockEditProfileResp := models.UserDetailsResponse{
		Id:    userID,
		Name:  mockUserDetails.Name,
		Email: mockUserDetails.Email,
	}

	expectedResp := response.Response{
		Message: "successfully edited",
		Data: map[string]interface{}{
			"user_id": userID,
			"name":    mockUserDetails.Name,
			"email":   mockUserDetails.Email,
		},
		Error: nil,
	}

	mockUseCase.EXPECT().EditProfile(mockUserDetails, userID).Return(mockEditProfileResp, nil)

	reqBody, _ := json.Marshal(mockUserDetails)
	req := httptest.NewRequest(http.MethodPost, "/edit-profile", bytes.NewBuffer(reqBody))
	req.Header.Set("Content-Type", "application/json")

	w := httptest.NewRecorder()
	c, _ := gin.CreateTestContext(w)
	c.Set("id", userID)
	c.Request = req

	handler.EditUserProfile(c)

	resp := w.Result()
	assert.Equal(t, http.StatusOK, resp.StatusCode)

	var respBody response.Response
	err := json.NewDecoder(resp.Body).Decode(&respBody)
	if err != nil {
		t.Errorf("Error decoding response body: %v", err)
	}

	assert.Equal(t, expectedResp.Message, respBody.Message)
	assert.Equal(t, expectedResp.Error, respBody.Error)

	data, ok := respBody.Data.(map[string]interface{})
	if !ok {
		t.Errorf("Failed to convert respBody.Data to map[string]interface{}")
	}

	// assert.Equal(t, expectedResp.Data.(map[string]interface{})["user_id"], data["user_id"])
	assert.Equal(t, expectedResp.Data.(map[string]interface{})["name"], data["name"])
	assert.Equal(t, expectedResp.Data.(map[string]interface{})["email"], data["email"])
}

func TestAddAddress(t *testing.T) {
	mockCtrl := gomock.NewController(t)
	defer mockCtrl.Finish()

	mockUseCase := usecase_mocks.NewMockUserUseCase(mockCtrl)
	mockHelper := helper_mocks.NewMockHelper(mockCtrl)

	handler := NewUserHandler(mockUseCase, mockHelper)

	userID := 123
	mockAddress := models.AddAddress{
        Name:      "John Doe",
        HouseName: "Apt 123",
        Street:    "123 Main St",
        City:      "New York",
        State:     "NY",
        Phone:     "1234567890",
        Pin:       "10001",
    }

	expectedResp := response.Response{
		Message: "successfully added address",
		Error:   nil,
	}

	mockUseCase.EXPECT().AddAddress(userID, mockAddress).Return(nil)

	reqBody, _ := json.Marshal(mockAddress)
	req := httptest.NewRequest(http.MethodPost, "/add-address", bytes.NewBuffer(reqBody))
	req.Header.Set("Content-Type", "application/json")

	w := httptest.NewRecorder()
	c, _ := gin.CreateTestContext(w)
	c.Set("id", userID)
	c.Request = req

	handler.AddAddress(c)

	resp := w.Result()
	assert.Equal(t, http.StatusOK, resp.StatusCode)

	var respBody response.Response
	err := json.NewDecoder(resp.Body).Decode(&respBody)
	if err != nil {
		t.Errorf("Error decoding response body: %v", err)
	}

	assert.Equal(t, expectedResp.Message, respBody.Message)
	assert.Equal(t, expectedResp.Error, respBody.Error)
}

func TestGetAddress(t *testing.T) {
    mockCtrl := gomock.NewController(t)
    defer mockCtrl.Finish()

    mockUseCase := usecase_mocks.NewMockUserUseCase(mockCtrl)
    mockHelper := helper_mocks.NewMockHelper(mockCtrl)

    handler := NewUserHandler(mockUseCase, mockHelper)

    userID := 123
    mockAddresses := []domain.Address{
        {
            Id:        1,
            UserID:    123,
            Name:      "John Doe",
            HouseName: "Apt 123",
            Street:    "123 Main St",
            City:      "New York",
            State:     "NY",
            Pin:       "10001",
        },
        {
            Id:        2,
            UserID:    123,
            Name:      "Jane Smith",
            HouseName: "Unit 456",
            Street:    "456 Elm St",
            City:      "Los Angeles",
            State:     "CA",
            Pin:       "90001",
        },
    }

    expectedResp := response.Response{
        Message: "successfully retrieved address",
        Error:   nil,
        Data:    mockAddresses,
    }

    // Adjust the mock setup to match the expected argument type
	mockUseCase.EXPECT().GetAddresses(userID).Return(mockAddresses, nil).Times(1)
    req := httptest.NewRequest(http.MethodGet, "/get-address", nil)
    w := httptest.NewRecorder()
    c, _ := gin.CreateTestContext(w)
    c.Set("id", userID)
    c.Request = req

    handler.GetAddress(c)

    resp := w.Result()
    assert.Equal(t, http.StatusOK, resp.StatusCode)

    var respBody response.Response
    err := json.NewDecoder(resp.Body).Decode(&respBody)
    if err != nil {
        t.Errorf("Error decoding response body: %v", err)
    }

    assert.Equal(t, expectedResp.Message, respBody.Message)
    assert.Equal(t, expectedResp.Error, respBody.Error)
    // assert.Equal(t, expectedResp.Data, respBody.Data)
}


func TestChangePassword(t *testing.T) {
    mockCtrl := gomock.NewController(t)
    defer mockCtrl.Finish()

    mockUseCase := usecase_mocks.NewMockUserUseCase(mockCtrl)
    mockHelper := helper_mocks.NewMockHelper(mockCtrl)

    handler := NewUserHandler(mockUseCase, mockHelper)

    userID := 123
    oldPassword := "oldpassword"
    newPassword := "newpassword"
    repassword := "newpassword"

    changePasswordReq := models.ChangePassword{
        Oldpassword: oldPassword,
        Password:    newPassword,
        Repassword:  repassword,
    }

    // Mock setup
    mockUseCase.EXPECT().ChangePassword(userID, oldPassword, newPassword, repassword).Return(nil)

    reqBody, _ := json.Marshal(changePasswordReq)
    req := httptest.NewRequest(http.MethodPost, "/change-password", bytes.NewBuffer(reqBody))
    req.Header.Set("Content-Type", "application/json")

    w := httptest.NewRecorder()
    c, _ := gin.CreateTestContext(w)
    c.Set("id", userID)
    c.Request = req

    handler.ChangePassword(c)

    resp := w.Result()
    assert.Equal(t, http.StatusOK, resp.StatusCode)

    var respBody response.Response
    err := json.NewDecoder(resp.Body).Decode(&respBody)
    if err != nil {
        t.Errorf("Error decoding response body: %v", err)
    }

    expectedResp := response.ClientResponse(http.StatusOK, "password changed Successfully ", nil, nil)
    assert.Equal(t, expectedResp.Message, respBody.Message)
    assert.Equal(t, expectedResp.Error, respBody.Error)
    assert.Equal(t, expectedResp.Data, respBody.Data)
}
