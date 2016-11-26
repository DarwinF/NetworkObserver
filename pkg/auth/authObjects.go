package auth

var authDatabaseEntries []user

var saltMaxLen = 64

// Authenticator is an interface for base authentication methods
type Authenticator interface {
	Login(string, string) (bool, error)
	CreateUser(string, string) (bool, error)
	UpdatePassword(string, string, string) (bool, error)
	UpdateUsername(string, string) (bool, error)
}

// encrypter interface for different encryption methods
type encrypter interface {
	Encrypt(string) (string, string)
	Validate(string, string, string) bool
}

// User stores user authentication information
type user struct {
	Username string
	Password string
	Salt     string
}

type baseAuthenticator struct {
	enc encrypter
}
