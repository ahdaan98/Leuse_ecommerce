package usecase

import (
	"errors"
	"testing"

	"github.com/ahdaan98/pkg/config"

	helper_mocks "github.com/ahdaan98/pkg/helper/mocks"
	repo_mocks "github.com/ahdaan98/pkg/repository/mocks"
	"github.com/ahdaan98/pkg/utils/models"
	"github.com/golang/mock/gomock"
	"github.com/stretchr/testify/assert"
)

func TestUserProfile(t *testing.T) {
	ctrl := gomock.NewController(t)
	defer ctrl.Finish()

	// Create a mock repository
	mockRepo := repo_mocks.NewMockUserRepository(ctrl)

	userUC := UserUseCase{
		repo: mockRepo,
	}

	// Define the expected response
	expectedResponse := models.UserDetailsResponse{
		Id:    1,
		Name:  "John Doe",
		Email: "john@example.com",
		Phone: "1234567890",
	}

	mockRepo.EXPECT().GetUserDetails(1).Return(expectedResponse, nil)

	resp, err := userUC.UserProfile(1)
	assert.NoError(t, err)
	assert.Equal(t, expectedResponse, resp)
}

func TestValidatingDetails(t *testing.T) {
	mockCtrl := gomock.NewController(t)
	defer mockCtrl.Finish()

	mockHelper := helper_mocks.NewMockHelper(mockCtrl)
	mockUserRepo := repo_mocks.NewMockUserRepository(mockCtrl)
	mockInventoryRepo := repo_mocks.NewMockInventoryRepository(mockCtrl)
	cfg := config.Config{
		DBUrl:              "dummy_db_url",
		AUTHTOKEN:          "dummy_auth_token",
		ACCOUNTSID:         "dummy_account_sid",
		SERVICESID:         "dummy_services_sid",
		ACCESS_KEY_ADMIN:   "dummy_access_key_admin",
		ACCESS_KEY_USER:    "dummy_access_key_user",
		KEY_ID_FOR_PAY:     "dummy_key_id_for_pay",
		SECRET_KEY_FOR_PAY: "dummy_secret_key_for_pay",
		PORT:               "dummy_port",
	}
	userUseCase := NewUserUseCase(mockUserRepo, mockHelper, cfg, mockInventoryRepo)

	tests := []struct {
		name    string
		input   models.UserSignUp
		stub    func(*helper_mocks.MockHelper, *repo_mocks.MockUserRepository, models.UserSignUp)
		wantErr error
	}{
		{
			name: "success",
			input: models.UserSignUp{
				Name:            "John Doe",
				Email:           "john.doe@example.com",
				Phone:           "+1234567890",
				Password:        "password123",
				ConfirmPassword: "password123",
			},
			stub: func(helper *helper_mocks.MockHelper, userRepo *repo_mocks.MockUserRepository, user models.UserSignUp) {
				helper.EXPECT().ValidateName(user.Name).Return(nil).Times(1)
				helper.EXPECT().ValidateEmail(user.Email).Return(nil).Times(1)
				userRepo.EXPECT().CheckUserExist(user.Email).Return(false, nil).Times(1)
				helper.EXPECT().ValidatePhoneNumber(user.Phone).Return(true).Times(1)
				helper.EXPECT().ValidatePassword(user.Password, user.ConfirmPassword).Return(nil).Times(1)
			},
			wantErr: nil,
		},
		{
			name: "invalid name",
			input: models.UserSignUp{
				Name:            "",
				Email:           "john.doe@example.com",
				Phone:           "+1234567890",
				Password:        "password123",
				ConfirmPassword: "password123",
			},
			stub: func(helper *helper_mocks.MockHelper, userRepo *repo_mocks.MockUserRepository, user models.UserSignUp) {
				helper.EXPECT().ValidateName(user.Name).Return(errors.New("invalid name")).Times(1)
			},
			wantErr: errors.New("invalid name"),
		},
		{
			name: "email already exists",
			input: models.UserSignUp{
				Name:            "John Doe",
				Email:           "john.doe@example.com",
				Phone:           "+1234567890",
				Password:        "password123",
				ConfirmPassword: "password123",
			},
			stub: func(helper *helper_mocks.MockHelper, userRepo *repo_mocks.MockUserRepository, user models.UserSignUp) {
				helper.EXPECT().ValidateName(user.Name).Return(nil).Times(1)
				helper.EXPECT().ValidateEmail(user.Email).Return(nil).Times(1)
				userRepo.EXPECT().CheckUserExist(user.Email).Return(true, nil).Times(1)
			},
			wantErr: errors.New("user with this email already exists"),
		},
		{
			name: "invalid phone number",
			input: models.UserSignUp{
				Name:            "John Doe",
				Email:           "john.doe@example.com",
				Phone:           "invalid_phone",
				Password:        "password123",
				ConfirmPassword: "password123",
			},
			stub: func(helper *helper_mocks.MockHelper, userRepo *repo_mocks.MockUserRepository, user models.UserSignUp) {
				helper.EXPECT().ValidateName(user.Name).Return(nil).Times(1)
				helper.EXPECT().ValidateEmail(user.Email).Return(nil).Times(1)
				userRepo.EXPECT().CheckUserExist(user.Email).Return(false, nil).Times(1)
				helper.EXPECT().ValidatePhoneNumber(user.Phone).Return(false).Times(1)
			},
			wantErr: errors.New("invalid phone number"),
		},
		{
			name: "passwords do not match",
			input: models.UserSignUp{
				Name:            "John Doe",
				Email:           "john.doe@example.com",
				Phone:           "+1234567890",
				Password:        "password123",
				ConfirmPassword: "different_password",
			},
			stub: func(helper *helper_mocks.MockHelper, userRepo *repo_mocks.MockUserRepository, user models.UserSignUp) {
				helper.EXPECT().ValidateName(user.Name).Return(nil).Times(1)
				helper.EXPECT().ValidateEmail(user.Email).Return(nil).Times(1)
				userRepo.EXPECT().CheckUserExist(user.Email).Return(false, nil).Times(1)
				helper.EXPECT().ValidatePhoneNumber(user.Phone).Return(true).Times(1)
				helper.EXPECT().ValidatePassword(user.Password, user.ConfirmPassword).Return(errors.New("passwords do not match")).Times(1)
			},
			wantErr: errors.New("passwords do not match"),
		},
	}

	for _, tc := range tests {
		t.Run(tc.name, func(t *testing.T) {
			tc.stub(mockHelper, mockUserRepo, tc.input)
			err := userUseCase.ValidatingDetails(tc.input)
			assert.Equal(t, tc.wantErr, err)
		})
	}
}

func TestUserSignUp(t *testing.T) {
	mockCtrl := gomock.NewController(t)
	defer mockCtrl.Finish()

	mockRepo := repo_mocks.NewMockUserRepository(mockCtrl)
	mockHelper := helper_mocks.NewMockHelper(mockCtrl)
	mockInventoryRepo := repo_mocks.NewMockInventoryRepository(mockCtrl)

	cfg := config.Config{
		DBUrl:              "dummy_db_url",
		AUTHTOKEN:          "dummy_auth_token",
		ACCOUNTSID:         "dummy_account_sid",
		SERVICESID:         "dummy_services_sid",
		ACCESS_KEY_ADMIN:   "dummy_access_key_admin",
		ACCESS_KEY_USER:    "dummy_access_key_user",
		KEY_ID_FOR_PAY:     "dummy_key_id_for_pay",
		SECRET_KEY_FOR_PAY: "dummy_secret_key_for_pay",
		PORT:               "dummy_port",
	}

	userUC := NewUserUseCase(mockRepo, mockHelper, cfg, mockInventoryRepo)

	tests := []struct {
		name      string
		input     models.UserSignUp
		mockSetup func(*repo_mocks.MockUserRepository, *helper_mocks.MockHelper, models.UserSignUp)
		want      models.TokenUsers
		wantErr   error
	}{
		{
			name: "successful sign up",
			input: models.UserSignUp{
				Name:            "John Doe",
				Email:           "john.doe@example.com",
				Phone:           "+1234567890",
				Password:        "password123",
				ConfirmPassword: "password123",
			},
			mockSetup: func(repo *repo_mocks.MockUserRepository, helper *helper_mocks.MockHelper, user models.UserSignUp) {
				repo.EXPECT().CreateUser(user).Return(models.UserDetailsResponse{}, nil).Times(1)
				helper.EXPECT().GenerateTokenClients(gomock.Any()).Return("dummy_token", nil).Times(1)
			},
			want: models.TokenUsers{
				Users: models.UserDetailsResponse{},
				Token: "dummy_token",
			},
			wantErr: nil,
		},
		{
			name: "create user error",
			input: models.UserSignUp{
				Name:            "John Doe",
				Email:           "john.doe@example.com",
				Phone:           "+1234567890",
				Password:        "password123",
				ConfirmPassword: "password123",
			},
			mockSetup: func(repo *repo_mocks.MockUserRepository, helper *helper_mocks.MockHelper, user models.UserSignUp) {
				repo.EXPECT().CreateUser(user).Return(models.UserDetailsResponse{}, errors.New("create user error")).Times(1)
			},
			want:    models.TokenUsers{},
			wantErr: errors.New("create user error"),
		},
		{
			name: "generate token error",
			input: models.UserSignUp{
				Name:            "John Doe",
				Email:           "john.doe@example.com",
				Phone:           "+1234567890",
				Password:        "password123",
				ConfirmPassword: "password123",
			},
			mockSetup: func(repo *repo_mocks.MockUserRepository, helper *helper_mocks.MockHelper, user models.UserSignUp) {
				repo.EXPECT().CreateUser(user).Return(models.UserDetailsResponse{}, nil).Times(1)
				helper.EXPECT().GenerateTokenClients(gomock.Any()).Return("", errors.New("generate token error")).Times(1)
			},
			want:    models.TokenUsers{},
			wantErr: errors.New("generate token error"),
		},
	}

	for _, tc := range tests {
		t.Run(tc.name, func(t *testing.T) {
			tc.mockSetup(mockRepo, mockHelper, tc.input)
			got, err := userUC.UserSignUp(tc.input)
			assert.Equal(t, tc.wantErr, err)
			assert.Equal(t, tc.want, got)
		})
	}
}

func TestUserLogin(t *testing.T) {
	mockCtrl := gomock.NewController(t)
	defer mockCtrl.Finish()

	mockRepo := repo_mocks.NewMockUserRepository(mockCtrl)
	mockHelper := helper_mocks.NewMockHelper(mockCtrl)
	mockInventoryRepo := repo_mocks.NewMockInventoryRepository(mockCtrl)

	cfg := config.Config{
		DBUrl:              "dummy_db_url",
		AUTHTOKEN:          "dummy_auth_token",
		ACCOUNTSID:         "dummy_account_sid",
		SERVICESID:         "dummy_services_sid",
		ACCESS_KEY_ADMIN:   "dummy_access_key_admin",
		ACCESS_KEY_USER:    "dummy_access_key_user",
		KEY_ID_FOR_PAY:     "dummy_key_id_for_pay",
		SECRET_KEY_FOR_PAY: "dummy_secret_key_for_pay",
		PORT:               "dummy_port",
	}

	userUC := NewUserUseCase(mockRepo, mockHelper, cfg, mockInventoryRepo)

	tests := []struct {
		name      string
		input     models.UserLogin
		mockSetup func(*repo_mocks.MockUserRepository, *helper_mocks.MockHelper, models.UserLogin)
		want      models.TokenUsers
		wantErr   error
	}{
		{
			name: "successful login",
			input: models.UserLogin{
				Email:    "john.doe@example.com",
				Password: "password123",
			},
			mockSetup: func(repo *repo_mocks.MockUserRepository, helper *helper_mocks.MockHelper, user models.UserLogin) {
				repo.EXPECT().CheckUserExist(user.Email).Return(true, nil).Times(1)
				repo.EXPECT().GetUserByEmail(user.Email).Return(models.UserDetailsResponse{}, nil).Times(1)
				repo.EXPECT().GetUserPassword(user.Email).Return("password123", nil).Times(1)
				repo.EXPECT().CheckBlockStatus(user.Email).Return(false, nil).Times(1)
				helper.EXPECT().GenerateTokenClients(gomock.Any()).Return("dummy_token", nil).Times(1)
			},
			want: models.TokenUsers{
				Users: models.UserDetailsResponse{},
				Token: "dummy_token",
			},
			wantErr: nil,
		},
		{
			name: "user does not exist",
			input: models.UserLogin{
				Email:    "john.doe@example.com",
				Password: "password123",
			},
			mockSetup: func(repo *repo_mocks.MockUserRepository, helper *helper_mocks.MockHelper, user models.UserLogin) {
				repo.EXPECT().CheckUserExist(user.Email).Return(false, nil).Times(1)
			},
			want:    models.TokenUsers{},
			wantErr: errors.New("user does not exist"),
		},
		{
			name: "incorrect password",
			input: models.UserLogin{
				Email:    "john.doe@example.com",
				Password: "password123",
			},
			mockSetup: func(repo *repo_mocks.MockUserRepository, helper *helper_mocks.MockHelper, user models.UserLogin) {
				repo.EXPECT().CheckUserExist(user.Email).Return(true, nil).Times(1)
				repo.EXPECT().GetUserByEmail(user.Email).Return(models.UserDetailsResponse{}, nil).Times(1)
				repo.EXPECT().GetUserPassword(user.Email).Return("wrong_password", nil).Times(1)
			},
			want:    models.TokenUsers{},
			wantErr: errors.New("incorrect password"),
		},
		{
			name: "user blocked by admin",
			input: models.UserLogin{
				Email:    "john.doe@example.com",
				Password: "password123",
			},
			mockSetup: func(repo *repo_mocks.MockUserRepository, helper *helper_mocks.MockHelper, user models.UserLogin) {
				repo.EXPECT().CheckUserExist(user.Email).Return(true, nil).Times(1)
				repo.EXPECT().GetUserByEmail(user.Email).Return(models.UserDetailsResponse{}, nil).Times(1)
				repo.EXPECT().GetUserPassword(user.Email).Return("password123", nil).Times(1)
				repo.EXPECT().CheckBlockStatus(user.Email).Return(true, nil).Times(1)
			},
			want:    models.TokenUsers{},
			wantErr: errors.New("you are blocked by admin"),
		},
	}

	for _, tc := range tests {
		t.Run(tc.name, func(t *testing.T) {
			tc.mockSetup(mockRepo, mockHelper, tc.input)
			got, err := userUC.UserLogin(tc.input)
			assert.Equal(t, tc.wantErr, err)
			assert.Equal(t, tc.want, got)
		})
	}
}

func TestEditProfile(t *testing.T) {
	mockCtrl := gomock.NewController(t)
	defer mockCtrl.Finish()

	mockRepo := repo_mocks.NewMockUserRepository(mockCtrl)
	mockHelper := helper_mocks.NewMockHelper(mockCtrl)
	mockInventoryRepo := repo_mocks.NewMockInventoryRepository(mockCtrl)

	cfg := config.Config{
		DBUrl:              "dummy_db_url",
		AUTHTOKEN:          "dummy_auth_token",
		ACCOUNTSID:         "dummy_account_sid",
		SERVICESID:         "dummy_services_sid",
		ACCESS_KEY_ADMIN:   "dummy_access_key_admin",
		ACCESS_KEY_USER:    "dummy_access_key_user",
		KEY_ID_FOR_PAY:     "dummy_key_id_for_pay",
		SECRET_KEY_FOR_PAY: "dummy_secret_key_for_pay",
		PORT:               "dummy_port",
	}

	userUC := NewUserUseCase(mockRepo, mockHelper, cfg, mockInventoryRepo)

	tests := []struct {
		name      string
		input     models.EditUserDetails
		id        int
		mockSetup func(*repo_mocks.MockUserRepository, models.EditUserDetails, int)
		want      models.UserDetailsResponse
		wantErr   error
	}{
		{
			name: "successful edit",
			input: models.EditUserDetails{
				Name:  "John Doe",
				Email: "john.doe@example.com",
				Phone: "+1234567890",
			},
			id: 1,
			mockSetup: func(repo *repo_mocks.MockUserRepository, details models.EditUserDetails, id int) {
				repo.EXPECT().EditDetails(details, id).Return(models.UserDetailsResponse{}, nil).Times(1)
			},
			want:    models.UserDetailsResponse{},
			wantErr: nil,
		},
		{
			name: "edit error",
			input: models.EditUserDetails{
				Name:  "John Doe",
				Email: "john.doe@example.com",
				Phone: "+1234567890",
			},
			id: 1,
			mockSetup: func(repo *repo_mocks.MockUserRepository, details models.EditUserDetails, id int) {
				repo.EXPECT().EditDetails(details, id).Return(models.UserDetailsResponse{}, errors.New("edit error")).Times(1)
			},
			want:    models.UserDetailsResponse{},
			wantErr: errors.New("edit error"),
		},
	}

	for _, tc := range tests {
		t.Run(tc.name, func(t *testing.T) {
			tc.mockSetup(mockRepo, tc.input, tc.id)
			got, err := userUC.EditProfile(tc.input, tc.id)
			assert.Equal(t, tc.wantErr, err)
			assert.Equal(t, tc.want, got)
		})
	}
}

func TestChangePassword(t *testing.T) {
	mockCtrl := gomock.NewController(t)
	defer mockCtrl.Finish()

	mockRepo := repo_mocks.NewMockUserRepository(mockCtrl)
	mockHelper := helper_mocks.NewMockHelper(mockCtrl)
	mockInventoryRepo := repo_mocks.NewMockInventoryRepository(mockCtrl)

	cfg := config.Config{
		DBUrl:              "dummy_db_url",
		AUTHTOKEN:          "dummy_auth_token",
		ACCOUNTSID:         "dummy_account_sid",
		SERVICESID:         "dummy_services_sid",
		ACCESS_KEY_ADMIN:   "dummy_access_key_admin",
		ACCESS_KEY_USER:    "dummy_access_key_user",
		KEY_ID_FOR_PAY:     "dummy_key_id_for_pay",
		SECRET_KEY_FOR_PAY: "dummy_secret_key_for_pay",
		PORT:               "dummy_port",
	}

	userUC := NewUserUseCase(mockRepo, mockHelper, cfg, mockInventoryRepo)

	tests := []struct {
		name       string
		id         int
		old        string
		password   string
		repassword string
		mockSetup  func(*repo_mocks.MockUserRepository, *helper_mocks.MockHelper, int, string, string, string)
		wantErr    error
	}{
		{
			name:       "successful password change",
			id:         1,
			old:        "old_password",
			password:   "new_password",
			repassword: "new_password",
			mockSetup: func(repo *repo_mocks.MockUserRepository, helper *helper_mocks.MockHelper, id int, old, password, repassword string) {
				repo.EXPECT().GetUserDetails(id).Return(models.UserDetailsResponse{}, nil).Times(1)
				repo.EXPECT().GetUserPassword(gomock.Any()).Return("old_password", nil).Times(1)
				helper.EXPECT().ValidatePassword(gomock.Any(), gomock.Any()).Return(nil).Times(1)
				repo.EXPECT().ChangePassword(id, password).Return(nil).Times(1)
			},
			wantErr: nil,
		},
		{
			name:    "invalid id",
			id:      0,
			wantErr: errors.New("invalid id"),
		},
		{
			name: "user details retrieval error",
			id:   1,
			mockSetup: func(repo *repo_mocks.MockUserRepository, helper *helper_mocks.MockHelper, id int, old, password, repassword string) {
				repo.EXPECT().GetUserDetails(id).Return(models.UserDetailsResponse{}, errors.New("user details retrieval error")).Times(1)
			},
			wantErr: errors.New("user details retrieval error"),
		},
		{
			name:    "incorrect old password",
			id:      1,
			old:     "old_password",
			wantErr: errors.New("please enter correct password"),
			mockSetup: func(repo *repo_mocks.MockUserRepository, helper *helper_mocks.MockHelper, id int, old, password, repassword string) {
				repo.EXPECT().GetUserDetails(id).Return(models.UserDetailsResponse{}, nil).Times(1)
				repo.EXPECT().GetUserPassword(gomock.Any()).Return("different_password", nil).Times(1)
			},
		},
		// {
		// 	name:       "passwords do not match",
		// 	id:         1,
		// 	old:        "old_password",
		// 	password:   "new_password",
		// 	repassword: "different_password",
		// 	wantErr:    errors.New("passwords do not match"),
		// 	mockSetup: func(repo *repo_mocks.MockUserRepository, helper *helper_mocks.MockHelper, id int, old, password, repassword string) {
		// 		repo.EXPECT().GetUserDetails(id).Return(models.UserDetailsResponse{}, nil).Times(1)
		// 		helper.EXPECT().ValidatePassword(password, repassword).Return(errors.New("passwords do not match")).Times(1)
		// 	},
		// },
		{
			name:       "password validation error",
			id:         1,
			old:        "old_password",
			password:   "invalid_password",
			repassword: "invalid_password",
			mockSetup: func(repo *repo_mocks.MockUserRepository, helper *helper_mocks.MockHelper, id int, old, password, repassword string) {
				repo.EXPECT().GetUserDetails(id).Return(models.UserDetailsResponse{}, nil).Times(1)
				repo.EXPECT().GetUserPassword(gomock.Any()).Return("old_password", nil).Times(1)
				helper.EXPECT().ValidatePassword(gomock.Any(), gomock.Any()).Return(errors.New("password validation error")).Times(1)
			},
			wantErr: errors.New("password validation error"),
		},
	}

	for _, tc := range tests {
		t.Run(tc.name, func(t *testing.T) {
			if tc.mockSetup != nil {
				tc.mockSetup(mockRepo, mockHelper, tc.id, tc.old, tc.password, tc.repassword)
			}
			err := userUC.ChangePassword(tc.id, tc.old, tc.password, tc.repassword)
			assert.Equal(t, tc.wantErr, err)
		})
	}
}
