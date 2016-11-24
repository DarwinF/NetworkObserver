package auth

var authDatabaseEntries []User

// dbEntryLineLen - Length of a database line
var dbEntryLineLen = 32*3 + 2

// Authenticator is an interface for base authentication methods
type Authenticator interface {
	Login(User) (bool, error)
	CreateUser(User) (bool, error)
	UpdatePassword(User, []byte) (bool, User, error)
}

// encrypter interface for different encryption methods
type encrypter interface {
	Encrypt([]byte) ([]byte, []byte, error)
	Validate([]byte, []byte, []byte) (bool, error)
}

// User stores user authentication information
type User struct {
	Username []byte
	Password []byte
	Salt     []byte
}

type baseAuthenticator struct {
	enc   encrypter
	users []User
}
