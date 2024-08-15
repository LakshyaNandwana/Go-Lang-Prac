package auth

type AuthError struct {
	User   string
	Reason string
}

func (e *AuthError) Error() string {
	return e.Reason
}

type User struct {
	Username string
}

func Login(user, passwd string) (*User, error) {
	const correctUsername = "admin"
	const correctPassword = "password"

	if user == correctUsername && passwd == correctPassword {
		return &User{Username: user}, nil
	}

	//If the credentials are incorrect, return AuthError
	err := &AuthError{
		User:   user,
		Reason: "invalid username or password",
	}
	return nil, err
}
