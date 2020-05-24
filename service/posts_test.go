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
	postColumns = []string{
		"id",
		"title",
		"content",
		"private",
		"created_by",
		"created_at",
		"updated_by",
		"updated_at",
	}

	testPost1 = pb.Post{
		Title:   "Test Title 1",
		Content: `Test content 1`,
	}

	testPost2 = pb.Post{
		Title:   "Test Title 2",
		Content: `Test content 2`,
	}
)

func TestCreatePostRequestSuccess(t *testing.T) {

	mock := db.Get("blog").Mock

	insertPostQueryRgx := "INSERT INTO posts"
	selectPostQueryRgx := "SELECT (.+) FROM posts"

	mock.ExpectExec(insertPostQueryRgx).
		WillReturnResult(sqlmock.NewResult(1, 1))

	mock.ExpectQuery(selectPostQueryRgx).
		WithArgs(1).
		WillReturnRows(
			sqlmock.NewRows(postColumns).
				AddRow(
					1,
					testPost1.Title,
					testPost1.Content,
					false,
					1,
					time.Time{},
					0,
					time.Time{},
				))

	svc := blogService{}
	resp, err := svc.CreatePost(context.Background(), &testPost1)
	if err != nil {
		t.Errorf("Unexpected error %s", err.Error())
	}

	if err = mock.ExpectationsWereMet(); err != nil {
		t.Errorf("There were unfulfilled expectations: %s", err)
	}

	if resp.Post.Title != testPost1.Title {
		t.Errorf("Expected name to be '%s' but saw '%s'", resp.Post.Title, resp.Post.Title)
	}

}

func TestUpdatePostRequestSuccess(t *testing.T) {

	mock := db.Get("blog").Mock

	testPost := testPost1
	testPost.Id = 1
	testPost.Title = testPost2.Title
	// testPost.Private = false

	updatePostQueryRgx := "UPDATE posts"
	selectPostQueryRgx := "SELECT (.+) FROM posts"

	mock.ExpectQuery(selectPostQueryRgx).
		WithArgs(1).
		WillReturnRows(
			sqlmock.NewRows(postColumns).
				AddRow(
					1,
					testPost1.Title,
					testPost1.Content,
					false,
					1,
					time.Time{},
					0,
					time.Time{},
				))

	mock.ExpectExec(updatePostQueryRgx).
		WithArgs(testPost.Title, testPost.Content, false, 1, 1).
		WillReturnResult(sqlmock.NewResult(1, 1))

	mock.ExpectQuery(selectPostQueryRgx).
		WithArgs(1).
		WillReturnRows(
			sqlmock.NewRows(postColumns).
				AddRow(
					1,
					testPost.Title,
					testPost.Content,
					false,
					1,
					time.Time{},
					1,
					time.Time{},
				))

	svc := blogService{}

	ctx := context.WithValue(context.Background(), middleware.UserIDKey, strconv.Itoa(int(1)))

	resp, err := svc.UpdatePost(ctx, &testPost)
	if err != nil {
		t.Errorf("Unexpected error %s", err.Error())
	}

	if err = mock.ExpectationsWereMet(); err != nil {
		t.Errorf("There were unfulfilled expectations: %s", err)
	}

	if resp.Post.Title != testPost.Title {
		t.Errorf("Expected name to be '%s' but saw '%s'", testPost.Title, resp.Post.Title)
	}

}

func TestGetPostRequestSuccess(t *testing.T) {

	mock := db.Get("blog").Mock

	selectPostQueryRgx := "SELECT (.+) FROM posts"

	mock.ExpectQuery(selectPostQueryRgx).
		WithArgs(1).
		WillReturnRows(
			sqlmock.NewRows(postColumns).
				AddRow(
					1,
					testPost1.Title,
					testPost1.Content,
					false,
					1,
					time.Time{},
					0,
					time.Time{},
				))

	svc := blogService{}

	testPost := testPost1
	testPost.Id = 1

	resp, err := svc.GetPost(context.Background(), &pb.PostRequest{Post: &testPost})
	if err != nil {
		t.Errorf("Unexpected error %s", err.Error())
	}

	if err = mock.ExpectationsWereMet(); err != nil {
		t.Errorf("There were unfulfilled expectations: %s", err)
	}

	if resp.Post.Title != testPost.Title {
		t.Errorf("Expected name to be '%s' but saw '%s'", resp.Post.Title, resp.Post.Title)
	}

}

func TestGetPostsRequestSuccess(t *testing.T) {

	mock := db.Get("blog").Mock

	selectPostQueryRgx := "SELECT (.+) FROM posts"

	mock.ExpectQuery(selectPostQueryRgx).
		WillReturnRows(
			sqlmock.NewRows(postColumns).
				AddRow(
					1,
					testAPIKey1,
					testPost1.Title,
					true,
					1,
					time.Time{},
					0,
					time.Time{},
				))

	mock.ExpectQuery(selectPostQueryRgx).
		WillReturnRows(
			sqlmock.NewRows([]string{"count"}).
				AddRow(1))

	svc := blogService{}

	resp, err := svc.GetPosts(context.Background(), &pb.PostRequest{})
	if err != nil {
		t.Errorf("Unexpected error %s", err.Error())
	}

	if err = mock.ExpectationsWereMet(); err != nil {
		t.Errorf("There were unfulfilled expectations: %s", err)
	}

	if len(resp.Posts) == 0 {
		t.Errorf("Expected results but saw zero")
	}

}

func TestDeletePostRequestSuccess(t *testing.T) {

	mock := db.Get("blog").Mock

	testPost := testPost1
	testPost.Id = 1

	selectPostQueryRgx := "SELECT (.+) FROM posts"
	updatePostQueryRgx := "UPDATE posts"
	deletePostExecRgx := "DELETE FROM posts"

	mock.ExpectQuery(selectPostQueryRgx).
		WithArgs(1).
		WillReturnRows(
			sqlmock.NewRows(postColumns).
				AddRow(
					1,
					testPost1.Title,
					testPost1.Content,
					false,
					1,
					time.Time{},
					0,
					time.Time{},
				))

	mock.ExpectExec(updatePostQueryRgx).
		WithArgs(1, 1).
		WillReturnResult(sqlmock.NewResult(1, 1))

	mock.ExpectExec(deletePostExecRgx).
		WithArgs(1).WillReturnResult(sqlmock.NewResult(1, 1))

	svc := blogService{}

	ctx := context.WithValue(context.Background(), middleware.UserIDKey, strconv.Itoa(1))

	_, err := svc.DeletePost(ctx, &testPost)
	if err != nil {
		t.Errorf("Unexpected error %s", err.Error())
	}

	if err = mock.ExpectationsWereMet(); err != nil {
		t.Errorf("There were unfulfilled expectations: %s", err)
	}

}

func TestPostNilRequests(t *testing.T) {

	svc := blogService{}

	tests := []struct {
		name        string
		requestFunc func() error
		wantErr     bool
	}{
		{
			name: "Fail on create post nil request",
			requestFunc: func() error {
				_, err := svc.CreatePost(context.Background(), nil)
				return err
			},
		},
		{
			name: "Fail on update post nil request",
			requestFunc: func() error {
				_, err := svc.UpdatePost(context.Background(), nil)
				return err
			},
		},
		{
			name: "Fail on get post nil request",
			requestFunc: func() error {
				_, err := svc.GetPost(context.Background(), nil)
				return err
			},
		},
		{
			name: "Fail on get posts nil request",
			requestFunc: func() error {
				_, err := svc.GetPosts(context.Background(), nil)
				return err
			},
		},
		{
			name: "Fail on delete post nil request",
			requestFunc: func() error {
				_, err := svc.DeletePost(context.Background(), nil)
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
