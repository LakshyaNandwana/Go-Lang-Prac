package auth_test

import (
	"auth/auth"
	"testing"
)

//SIMPLE PASSING TEST CASE

// func TestLogin(t *testing.T) {
// 	_, err := auth.Login("admin", "password")
// 	if err != nil {
// 		t.Fatalf("can't loging: %#v", err)
// 	}
// }

//TABLE DRIVEN TEST CASES

func TestLogin(t *testing.T) {
	type testCase struct {
		name         string
		user         string
		passwd       string
		expectError  bool
		expectedUser *auth.User
		expectedErr  *auth.AuthError
	}

	testCases := []testCase{
		{
			name:         "Successful login with correct credentials",
			user:         "admin",
			passwd:       "password",
			expectError:  false,
			expectedUser: &auth.User{Username: "admin"},
		},
		{
			name:        "Failed login with incorrect password",
			user:        "admin",
			passwd:      "incorrect",
			expectError: true,
			expectedErr: &auth.AuthError{User: "admin", Reason: "invalid username or password"},
		},
		{
			name:        "Failed login with incorrect username",
			user:        "wronguser",
			passwd:      "password",
			expectError: true,
			expectedErr: &auth.AuthError{User: "wronguser", Reason: "invalid username or password"},
		},
		{
			name:        "Failed login with both incorrect username and password",
			user:        "wronguser",
			passwd:      "incorrect",
			expectError: true,
			expectedErr: &auth.AuthError{User: "wronguser", Reason: "invalid username or password"},
		},
	}

	//Iterating over the test cases
	for _, tc := range testCases {
		t.Run(tc.name, func(t *testing.T) {
			user, err := auth.Login(tc.user, tc.passwd)

			if tc.expectError {
				if err == nil {
					t.Fatalf("expected error, got successful login with user: %#v", user)
				}

				authErr, ok := err.(*auth.AuthError)
				if !ok {
					t.Fatalf("expected AuthError, got %#v", err)
				}

				if authErr.User != tc.expectedErr.User || authErr.Reason != tc.expectedErr.Reason {
					t.Fatalf("expected error %v, got %v", tc.expectedErr, authErr)
				}
			} else {
				if err != nil {
					t.Fatalf("expected successful login, got error: %#v", err)
				}

				if user == nil || user.Username != tc.expectedUser.Username {
					t.Fatalf("expected user %v, got %v", tc.expectedUser, user)
				}
			}
		})
	}
}
