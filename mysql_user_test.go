package blog

import (
	"fmt"
	"testing"
)

func TestCreateUserFail(t *testing.T) {

	tests := []struct {
		name     string
		testFunc func() error
	}{
		{
			name: "on user is nil",
			testFunc: func() error {

				err := Store.CreateUser(nil)
				if err != ErrCreateUserIsNil {
					return fmt.Errorf("Expected error %s but saw %s", ErrCreateUserIsNil.Error(), err.Error())
				}

				return nil
			},
		},
		{
			name: "on user id is present",
			testFunc: func() error {

				err := Store.CreateUser(&User{ID: 1})
				if err != ErrCreateUserIsNil {
					return fmt.Errorf("Expected error %s but saw %s", ErrCreateUserIsNil.Error(), err.Error())
				}

				return nil
			},
		},
		{
			name: "on user name is empty",
			testFunc: func() error {

				err := Store.CreateUser(&User{})
				if err != ErrCreateUserIsNil {
					return fmt.Errorf("Expected error %s but saw %s", ErrCreateUserIsNil.Error(), err.Error())
				}

				return nil
			},
		},
	}

	for _, tt := range tests {
		tt := tt
		t.Run(tt.name, func(t *testing.T) {
			t.Parallel()

			err := tt.testFunc()
			if err != nil {
				t.Errorf("Unexpected error for '%s' but saw %s", tt.name, err.Error())
			}

		})
	}
}

func TestUpdateUserFail(t *testing.T) {

	tests := []struct {
		name     string
		testFunc func() error
	}{
		{
			name: "on user is nil",
			testFunc: func() error {

				err := Store.UpdateUser(nil)
				if err != ErrUpdateUserIsNil {
					return fmt.Errorf("Expected error %s but saw %s", ErrCreateUserIsNil.Error(), err.Error())
				}

				return nil
			},
		},
		{
			name: "on user id is 0",
			testFunc: func() error {

				err := Store.UpdateUser(&User{})
				if err != ErrUpdateUserInvalidID {
					return fmt.Errorf("Expected error %s but saw %s", ErrCreateUserIsNil.Error(), err.Error())
				}

				return nil
			},
		},
		{
			name: "on user name is empty",
			testFunc: func() error {

				err := Store.UpdateUser(&User{ID: 1})
				if err != ErrUpdateUserNameMissing {
					return fmt.Errorf("Expected error %s but saw %s", ErrCreateUserIsNil.Error(), err.Error())
				}

				return nil
			},
		},
		{
			name: "on user update by id is 0",
			testFunc: func() error {

				err := Store.UpdateUser(&User{ID: 1, Name: "Bobby Tables"})
				if err != ErrUpdateUserUpdatedByMissing {
					return fmt.Errorf("Expected error %s but saw %s", ErrCreateUserIsNil.Error(), err.Error())
				}

				return nil
			},
		},
	}

	for _, tt := range tests {
		tt := tt
		t.Run(tt.name, func(t *testing.T) {
			t.Parallel()

			err := tt.testFunc()
			if err != nil {
				t.Errorf("Unexpected error for '%s' but saw %s", tt.name, err.Error())
			}

		})
	}
}
