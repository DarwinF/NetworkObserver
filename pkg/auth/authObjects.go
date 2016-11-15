package auth

// Authenticator is an interface for base authentication methods
type Authenticator interface {
	Login(User) (bool, error)
	CreateUser(User) (bool, error)
	UpdatePassword(User, string) (bool, User, error)
}

// encrypter interface for different encryption methods
type encrypter interface {
	Encrypt(string) (string, string, error)
	Validate(string, string, string) (bool, error)
}

// User stores user authentication information
type User struct {
	Username string
	Password string
	Salt     string
}

type baseAuthenticator struct {
	enc   encrypter
	users []User
}
