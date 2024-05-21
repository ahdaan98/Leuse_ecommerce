package repository

import (
	"errors"
	"regexp"
	"testing"

	"github.com/DATA-DOG/go-sqlmock"
	"github.com/ahdaan98/pkg/utils/models"
	"github.com/go-playground/assert/v2"
	"gorm.io/driver/postgres"
	"gorm.io/gorm"
)

func CheckUserExist(t *testing.T) {
	tests := []struct {
		name string
		args string
		stub func(sqlmock.Sqlmock)
		want bool
	}{
		{
			name: "user exist",
			args: "randomrabbit@gmail.com",
			stub: func(mocksql sqlmock.Sqlmock) {
				mocksql.ExpectQuery(regexp.QuoteMeta("SELECT COUNT(*) FROM users WHERE email = $1")).WithArgs("randomrabbit@gmail.com").WillReturnRows(sqlmock.NewRows([]string{"count"}).AddRow("1"))
			},
			want: true,
		},
		{
			name: "user does not exist",
			args: "randomrabbit@gmail.com",
			stub: func(mocksql sqlmock.Sqlmock) {
				mocksql.ExpectQuery(regexp.QuoteMeta("SELECT COUNT(*) FROM users WHERE email = $1")).WithArgs("randomrabbit@gmail.com").WillReturnRows(sqlmock.NewRows([]string{"count"}).AddRow("0"))
			},
			want: false,
		},
	}

	for _, tc := range tests {
		t.Run(tc.name, func(t *testing.T) {
			db, mock, err := sqlmock.New()
			if err != nil {
				t.Fatalf("Error creating mock DB: %v", err)
			}
			defer db.Close()

			DB, _ := gorm.Open(postgres.New(postgres.Config{Conn: db}), &gorm.Config{})

			tc.stub(mock)

			userRepository := NewUserRepository(DB)
			result, _ := userRepository.CheckUserExist(tc.args)

			assert.Equal(t, tc.want, result)
		})
	}
}

func TestCreateUser(t *testing.T) {
	tests := []struct {
		name string
		args *models.UserSignUp
		stub func(mocksql sqlmock.Sqlmock)
		want models.UserDetailsResponse
		wantErr error
	}{
		// Test case for successful creation
		{
			name: "Successfully created User",
			args: &models.UserSignUp{
				Name:     "Random Rabbit",
				Email:    "randomrabbit@gmail.com",
				Phone:    "+91000000000",
				Password: "EncryptedPassoword",
			},
			stub: func(mocksql sqlmock.Sqlmock) {
				mocksql.ExpectQuery(regexp.QuoteMeta("INSERT INTO users (name, email, password, phone) VALUES ($1,$2,$3,$4) RETURNING id, name, email, phone")).
					WithArgs("Random Rabbit", "randomrabbit@gmail.com", "EncryptedPassoword", "+91000000000").
					WillReturnRows(sqlmock.NewRows([]string{"id", "name", "email", "phone"}).AddRow(1, "Random Rabbit", "randomrabbit@gmail.com", "+91000000000"))
			},
			want: models.UserDetailsResponse{
				Id:     1,
				Name:   "Random Rabbit",
				Email:  "randomrabbit@gmail.com",
				Phone:  "+91000000000",
			},
			wantErr: nil,
		},
		{
			name: "Database error",
			args: &models.UserSignUp{
				Name:     "Error Rabbit",
				Email:    "errorrabbit@gmail.com",
				Phone:    "+91999999999",
				Password: "EncryptedPassoword",
			},
			stub: func(mocksql sqlmock.Sqlmock) {
				mocksql.ExpectQuery(regexp.QuoteMeta("INSERT INTO users (name, email, password, phone) VALUES ($1,$2,$3,$4) RETURNING id, name, email, phone")).
					WillReturnError(errors.New("data mismatching cant store in data base"))
			},
			wantErr: errors.New("data mismatching cant store in data base"),
		},
	}

	for _,tc := range tests {
		t.Run(tc.name, func(t *testing.T){
			db, mock, err := sqlmock.New()
			if err != nil {
				t.Fatalf("Error creating mock DB: %v", err)
			}
			defer db.Close()

			DB, _ := gorm.Open(postgres.New(postgres.Config{Conn: db}), &gorm.Config{})

			tc.stub(mock)

			userRepository := NewUserRepository(DB)
			result, err := userRepository.CreateUser(*tc.args)

			assert.Equal(t,tc.want,result)
			assert.Equal(t,err,tc.wantErr)
		})
	}
}
