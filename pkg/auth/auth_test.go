package auth

import "testing"

var authenticator Authenticator
var testPassword = "test password"

func Test_SaltedDifferentFromUnsalted(t *testing.T) {
	salted, _ := sha256WithSalt(testPassword, "")
	unsalted := sha256WithoutSalt(testPassword)

	if salted == unsalted {
		t.Fatal("The salted and unsalted values were the same")
	}
}

func Test_SaltingReturnsTheSameValues(t *testing.T) {
	salted1, salt := sha256WithSalt(testPassword, "")
	salted2, _ := sha256WithSalt(testPassword, string(salt[:]))

	if salted1 != salted2 {
		t.Fatal("The salted passwords were not the same")
	}
}

func Test_VerifyWorksWithSalting(t *testing.T) {
	settings := authSettings
	setupAuthenticator(&settings)

	encrypted, salt, _ := authenticator.Encrypt(testPassword)
	valid, _ := authenticator.Validate(testPassword, salt, encrypted)

	if !valid {
		t.Fatal("The passwords didn't match")
	}
}

func Test_VerifyWorksWithoutSalting(t *testing.T) {
	settings := authSettings
	settings.UseSalt = false

	setupAuthenticator(&settings)
	encrypted, salt, _ := authenticator.Encrypt(testPassword)
	valid, _ := authenticator.Validate(testPassword, salt, encrypted)

	if !valid {
		t.Fatal("The passwords didn't match")
	}
}

func setupAuthenticator(settings *Settings) {
	if settings != nil {
		authenticator, _ = NewAuthenticator(settings)
	} else {
		authenticator, _ = NewAuthenticator(nil)
	}
}
