package user

import "testing"

func TestAccount_NewUser(t *testing.T) {

	type testCase struct {
		testName    string
		email       string
		firstName   string
		lastName    string
		expectedErr error
	}

	testCases := []testCase{
		{
			testName:    "missing email",
			email:       "",
			firstName:   "A",
			lastName:    "B",
			expectedErr: ErrEmailInvalid,
		},
		{
			testName:    "missing first name",
			email:       "A",
			firstName:   "",
			lastName:    "B",
			expectedErr: ErrFirstNameInvalid,
		},
		{
			testName:    "missing last name",
			email:       "A",
			firstName:   "B",
			lastName:    "",
			expectedErr: ErrLastNameInvalid,
		},
	}

	for _, tc := range testCases {
		t.Run(tc.testName, func(t *testing.T) {

			_, err := NewUser(tc.email, tc.firstName, tc.lastName)
			if err != tc.expectedErr {
				t.Errorf("Expected error %v, got %v", tc.expectedErr, err)
			}
		})
	}
}
