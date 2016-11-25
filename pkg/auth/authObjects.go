package auth

var authDatabaseEntries []User

var saltMaxLen = 64

// Authenticator is an interface for base authentication methods
type Authenticator interface {
	Login(User) (bool, error)
	CreateUser(User) (bool, error)
	UpdatePassword(User, []byte) (bool, User, error)
}

// encrypter interface for different encryption methods
type encrypter interface {
	Encrypt(string) (string, string)
	Validate(string, string, string) bool
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
