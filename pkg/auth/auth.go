// Package auth provides functionality for authenticating
// users to the system.
//
// Default Configuration Values
// * Password Encryption Method: sha256
// * Salt length in bits: 64
package auth

import (
	"crypto/sha256"
	"math/rand"

	"fmt"
)

// Settings is a struct of all the cofiguration options
// used for configuring the auth package
type Settings struct {
	EncryptionMethod string
	SaltLength       int
	UseSalt          bool
}

// Authenticator object for authentication
type Authenticator interface {
	Encrypt(string) (string, string, error)
	Validate(string, string, string) (bool, error)
}

type authAdapter struct {
}

var authSettings Settings

func init() {
	authSettings = Settings{
		EncryptionMethod: "sha256",
		SaltLength:       64,
		UseSalt:          true,
	}
}

// NewAuthenticator returns a new authenticator that can be used for authenticating
func NewAuthenticator(settings *Settings) (Authenticator, error) {
	a := authAdapter{}

	if settings != nil {
		authSettings = *settings
	}

	return &a, nil
}

func (a *authAdapter) Encrypt(v string) (eValue, salt string, err error) {
	if authSettings.UseSalt {
		e, s := sha256WithSalt(v, "")
		eValue = string(e[:])
		salt = string(s[:])
		return
	}

	e := sha256WithoutSalt(v)
	eValue = string(e[:])

	return
}

func (a *authAdapter) Validate(input, salt, password string) (valid bool, err error) {
	var encryptedString string
	valid = false

	if authSettings.UseSalt {
		e, _ := sha256WithSalt(input, salt)
		encryptedString = string(e[:])
	} else {
		e := sha256WithoutSalt(input)
		encryptedString = string(e[:])
	}

	if encryptedString == password {
		valid = true
	}

	return
}

func sha256WithoutSalt(value string) [sha256.Size]byte {
	encrypted := sha256.Sum256([]byte(value))

	return encrypted
}

func sha256WithSalt(value, saltValue string) ([sha256.Size]byte, []byte) {
	salt := make([]byte, authSettings.SaltLength)

	if saltValue != "" {
		salt = []byte(saltValue)
	} else {
		n, err := rand.Read(salt)

		if err != nil {
			fmt.Println("There was an error generating a salt: ", err)
			return [sha256.Size]byte{}, nil
		}

		if n != authSettings.SaltLength {
			fmt.Printf("Only %d characters were read.\n", n)
			return [sha256.Size]byte{}, nil
		}
	}

	saltedVal := append([]byte(value), salt...)
	encrypted := sha256.Sum256(saltedVal)

	return encrypted, salt
}
