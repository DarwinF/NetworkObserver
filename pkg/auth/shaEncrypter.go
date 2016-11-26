package auth

import (
	"crypto/sha256"
	"log"
	"math/rand"

	"encoding/base64"
	"fmt"
)

// Settings is a struct of all the cofiguration options
// used for configuring the auth package
type Settings struct {
	EncryptionMethod string
	SaltLength       int
	UseSalt          bool
}

type shaAdapter struct {
}

var shaSettings Settings

func init() {
	shaSettings = Settings{
		EncryptionMethod: "sha256",
		SaltLength:       saltMaxLen,
		UseSalt:          true,
	}
}

// newShaEncrypter returns a new authenticator that can be used for authenticating
func newShaEncrypter(settings *Settings) (encrypter, error) {
	a := shaAdapter{}

	if settings != nil {
		shaSettings = *settings
	}

	return &a, nil
}

func (a *shaAdapter) Encrypt(input string) (eValue, salt string) {
	e, s := sha256WithSalt([]byte(input), nil)

	eStr := base64.StdEncoding.EncodeToString(e[:])
	sStr := base64.StdEncoding.EncodeToString(s)

	return eStr, sStr
}

// Validate - Takes in a value (input) and encrypts it with the salt value
// and returns whether the encrypted input is the same as the password value
// taken in
func (a *shaAdapter) Validate(input, salt, password string) (valid bool) {
	var encryptedString string
	valid = false

	decoded, err := base64.StdEncoding.DecodeString(salt)

	if err != nil {
		log.Printf("[ERROR] There was an error decoding the string %s\n", err.Error())
		return
	}
	e, _ := sha256WithSalt([]byte(input), decoded)
	encryptedString = base64.StdEncoding.EncodeToString(e[:])

	if encryptedString == password {
		valid = true
	}

	return
}

func sha256WithSalt(value, saltValue []byte) ([sha256.Size]byte, []byte) {
	salt := make([]byte, shaSettings.SaltLength)

	if saltValue != nil {
		salt = []byte(saltValue)
	} else {
		n, err := rand.Read(salt)

		if err != nil {
			fmt.Println("There was an error generating a salt: ", err)
			return [sha256.Size]byte{}, nil
		}

		if n != shaSettings.SaltLength {
			fmt.Printf("Only %d characters were read.\n", n)
			return [sha256.Size]byte{}, nil
		}
	}

	saltedVal := append([]byte(value), salt...)
	encrypted := sha256.Sum256(saltedVal)

	return encrypted, salt
}
