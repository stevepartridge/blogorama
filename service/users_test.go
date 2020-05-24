package service

import (
	"context"
	"strconv"
	"testing"
	"time"

	"github.com/DATA-DOG/go-sqlmock"
	"github.com/stevepartridge/blogorama/pkg/middleware"
	pb "github.com/stevepartridge/blogorama/protos/go"
	"github.com/stevepartridge/db"
)

var (
	userColumns = []string{
		"id",
		"api_key",
		"name",
		"active",
		"created_by",
		"created_at",
		"updated_by",
		"updated_at",
	}

	testUser1 = pb.User{
		Name: "Bobby Tables",
	}
	testUser2 = pb.User{
		Name: "Little Bobby Tables",
	}
	testAPIKey1 = "c2fcf476-9d46-11ea-9862-dca904888369"
)

func TestCreateUserRequestSuccess(t *testing.T) {

	mock := db.Get("blog").Mock

	insertUserQueryRgx := "INSERT INTO users"
	selectUserQueryRgx := "SELECT (.+) FROM users"
	selectAPIKeyQueryRgx := "SELECT api_key FROM users"

	mock.ExpectExec(insertUserQueryRgx).
		WillReturnResult(sqlmock.NewResult(1, 1))

	mock.ExpectQuery(selectUserQueryRgx).
		WithArgs(1).
		WillReturnRows(
			sqlmock.NewRows(userColumns).
				AddRow(
					1,
					testAPIKey1,
					testUser1.Name,
					true,
					1,
					time.Time{},
					0,
					time.Time{},
				))

	mock.ExpectQuery(selectAPIKeyQueryRgx).
		WithArgs(1).
		WillReturnRows(
			sqlmock.NewRows([]string{"api_key"}).
				AddRow(
					testAPIKey1,
				))

	svc := blogService{}
	resp, err := svc.CreateUser(context.Background(), &testUser1)
	if err != nil {
		t.Errorf("Unexpected error %s", err.Error())
	}

	if err = mock.ExpectationsWereMet(); err != nil {
		t.Errorf("There were unfulfilled expectations: %s", err)
	}

	if resp.User.Name != testUser1.Name {
		t.Errorf("Expected name to be '%s' but saw '%s'", resp.User.Name, resp.User.Name)
	}

	if resp.User.ApiKey != testAPIKey1 {
		t.Errorf("Expected API Key to be '%s' but saw '%s'", testAPIKey1, resp.User.ApiKey)
	}

}

func TestUpdateUserRequestSuccess(t *testing.T) {

	mock := db.Get("blog").Mock

	testUser := testUser1
	testUser.Id = 1
	testUser.Name = testUser2.Name
	testUser.Active = true

	updateUserQueryRgx := "UPDATE users"
	selectUserQueryRgx := "SELECT (.+) FROM users"

	mock.ExpectQuery(selectUserQueryRgx).
		WithArgs(1).
		WillReturnRows(
			sqlmock.NewRows(userColumns).
				AddRow(
					testUser.Id,
					testAPIKey1,
					testUser1.Name,
					true,
					1,
					time.Time{},
					0,
					time.Time{},
				))

	mock.ExpectExec(updateUserQueryRgx).
		WithArgs(testUser.Name, true, 1, 1).
		WillReturnResult(sqlmock.NewResult(1, 1))

	mock.ExpectQuery(selectUserQueryRgx).
		WithArgs(1).
		WillReturnRows(
			sqlmock.NewRows(userColumns).
				AddRow(
					testUser.Id,
					testAPIKey1,
					testUser.Name,
					true,
					testUser.Id,
					time.Time{},
					testUser.Id,
					time.Time{},
				))

	svc := blogService{}

	ctx := context.WithValue(context.Background(), middleware.UserIDKey, strconv.Itoa(int(testUser.Id)))

	resp, err := svc.UpdateUser(ctx, &testUser)
	if err != nil {
		t.Errorf("Unexpected error %s", err.Error())
	}

	if err = mock.ExpectationsWereMet(); err != nil {
		t.Errorf("There were unfulfilled expectations: %s", err)
	}

	if resp.User.Name != testUser.Name {
		t.Errorf("Expected name to be '%s' but saw '%s'", resp.User.Name, resp.User.Name)
	}

}

func TestGetUserRequestSuccess(t *testing.T) {

	mock := db.Get("blog").Mock

	selectUserQueryRgx := "SELECT (.+) FROM users"

	mock.ExpectQuery(selectUserQueryRgx).
		WithArgs(1).
		WillReturnRows(
			sqlmock.NewRows(userColumns).
				AddRow(
					1,
					testAPIKey1,
					testUser1.Name,
					true,
					1,
					time.Time{},
					0,
					time.Time{},
				))

	svc := blogService{}

	testUser := testUser1
	testUser.Id = 1

	resp, err := svc.GetUser(context.Background(), &pb.UserRequest{User: &testUser})
	if err != nil {
		t.Errorf("Unexpected error %s", err.Error())
	}

	if err = mock.ExpectationsWereMet(); err != nil {
		t.Errorf("There were unfulfilled expectations: %s", err)
	}

	if resp.User.Name != testUser.Name {
		t.Errorf("Expected name to be '%s' but saw '%s'", resp.User.Name, resp.User.Name)
	}

}

func TestGetUsersRequestSuccess(t *testing.T) {

	mock := db.Get("blog").Mock

	selectUserQueryRgx := "SELECT (.+) FROM users"

	mock.ExpectQuery(selectUserQueryRgx).
		WillReturnRows(
			sqlmock.NewRows(userColumns).
				AddRow(
					1,
					testAPIKey1,
					testUser1.Name,
					true,
					1,
					time.Time{},
					0,
					time.Time{},
				))

	mock.ExpectQuery(selectUserQueryRgx).
		WillReturnRows(
			sqlmock.NewRows([]string{"count"}).
				AddRow(1))

	svc := blogService{}

	resp, err := svc.GetUsers(context.Background(), &pb.UserRequest{})
	if err != nil {
		t.Errorf("Unexpected error %s", err.Error())
	}

	if err = mock.ExpectationsWereMet(); err != nil {
		t.Errorf("There were unfulfilled expectations: %s", err)
	}

	if len(resp.Users) == 0 {
		t.Errorf("Expected results but saw zero")
	}

}

func TestDeleteUserRequestSuccess(t *testing.T) {

	mock := db.Get("blog").Mock

	testUser := testUser1
	testUser.Id = 1

	updateUserQueryRgx := "UPDATE users"
	deleteUserExecRgx := "DELETE FROM users"

	mock.ExpectExec(updateUserQueryRgx).
		WithArgs(1, 1).
		WillReturnResult(sqlmock.NewResult(1, 1))

	mock.ExpectExec(deleteUserExecRgx).
		WithArgs(1).WillReturnResult(sqlmock.NewResult(1, 1))

	svc := blogService{}

	ctx := context.WithValue(context.Background(), middleware.UserIDKey, strconv.Itoa(1))

	_, err := svc.DeleteUser(ctx, &testUser)
	if err != nil {
		t.Errorf("Unexpected error %s", err.Error())
	}

	if err = mock.ExpectationsWereMet(); err != nil {
		t.Errorf("There were unfulfilled expectations: %s", err)
	}

}

func TestUserNilRequests(t *testing.T) {

	svc := blogService{}

	tests := []struct {
		name        string
		requestFunc func() error
		wantErr     bool
	}{
		{
			name: "Fail on create user nil request",
			requestFunc: func() error {
				_, err := svc.CreateUser(context.Background(), nil)
				return err
			},
		},
		{
			name: "Fail on update user nil request",
			requestFunc: func() error {
				_, err := svc.UpdateUser(context.Background(), nil)
				return err
			},
		},
		{
			name: "Fail on get user nil request",
			requestFunc: func() error {
				_, err := svc.GetUser(context.Background(), nil)
				return err
			},
		},
		{
			name: "Fail on get users nil request",
			requestFunc: func() error {
				_, err := svc.GetUsers(context.Background(), nil)
				return err
			},
		},
		{
			name: "Fail on delete user nil request",
			requestFunc: func() error {
				_, err := svc.DeleteUser(context.Background(), nil)
				return err
			},
		},
	}

	for _, tt := range tests {
		tt := tt
		t.Run(tt.name, func(t *testing.T) {
			t.Parallel()

			err := tt.requestFunc()
			if err == nil {
				t.Errorf("Unexpected error for '%s' but didn't see one", tt.name)
			}

		})
	}

}
