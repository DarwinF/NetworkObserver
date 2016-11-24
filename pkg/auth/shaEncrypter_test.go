package auth

import (
	"bytes"
	"testing"
)

var enc encrypter
var testPassword = []byte("test password")

func Test_SaltingIsRandom(t *testing.T) {
	salted1, salt1 := sha256WithSalt(testPassword, nil)
	salted2, salt2 := sha256WithSalt(testPassword, nil)

	if bytes.Equal(salted1[:], salted2[:]) {
		t.Fatal("The salted passwords are the same")
	}

	if bytes.Equal(salt1[:], salt2[:]) {
		t.Fatal("The salt values are the same")
	}
}

func Test_SaltingReturnsTheSameValues(t *testing.T) {
	salted1, salt := sha256WithSalt(testPassword, nil)
	salted2, _ := sha256WithSalt(testPassword, salt[:])

	if !bytes.Equal(salted1[:], salted2[:]) {
		t.Fatal("The salted passwords were not the same")
	}
}

func Test_VerifyWorksWithSalting(t *testing.T) {
	settings := shaSettings
	setupEncrypter(&settings)

	encrypted, salt, _ := enc.Encrypt(testPassword)
	valid, _ := enc.Validate(testPassword, salt, encrypted)

	if !valid {
		t.Fatal("The passwords didn't match")
	}
}

func setupEncrypter(settings *Settings) {
	if settings != nil {
		enc, _ = newShaEncrypter(settings)
	} else {
		enc, _ = newShaEncrypter(nil)
	}
}
