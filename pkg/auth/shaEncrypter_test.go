package auth

import (
	"log"
	"testing"
)

var enc encrypter

func Test_SaltingIsRandom(t *testing.T) {
	log.Println("Testing that salting is random.")

	settings := shaSettings
	setupEncrypter(&settings)

	encrypted1, salt1 := enc.Encrypt(defaultPassword)
	encrypted2, salt2 := enc.Encrypt(defaultPassword)

	if encrypted1 == encrypted2 {
		t.Fatal("The salted passwords are the same")
	}

	if salt1 == salt2 {
		t.Fatal("The salt values are the same")
	}
}

func Test_VerifyWorksWithSalting(t *testing.T) {
	log.Println("Testing that using the same salt will return the same encrypted value.")
	settings := shaSettings
	setupEncrypter(&settings)

	encrypted, salt := enc.Encrypt(defaultPassword)
	valid := enc.Validate(defaultPassword, salt, encrypted)

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
