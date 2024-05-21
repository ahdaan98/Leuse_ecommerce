package usecase

import (
	"testing"

	"github.com/ahdaan98/pkg/repository/mocks"
	"github.com/ahdaan98/pkg/utils/models"
	"github.com/golang/mock/gomock"
	"github.com/stretchr/testify/assert"
)

func TestUserProfile(t *testing.T) {
	ctrl := gomock.NewController(t)
	defer ctrl.Finish()

	// Create a mock repository
	mockRepo := mocks.NewMockUserRepository(ctrl)

	// Create a new instance of UserUseCase with the mock repository
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

	// Mock the behavior of GetUserDetails method
	mockRepo.EXPECT().GetUserDetails(1).Return(expectedResponse, nil)

	// Call the method under test
	resp, err := userUC.UserProfile(1)

	// Assert that no error is returned
	assert.NoError(t, err)

	// Assert that the response matches the expected response
	assert.Equal(t, expectedResponse, resp)
}