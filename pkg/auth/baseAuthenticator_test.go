package auth

import "testing"
import "log"

func Test_BaseAuthenticatorLogin(t *testing.T) {
	log.Printf("Testing logging in with the base authenticator")
}

func Test_BaseAuthenticatorCreateAccount(t *testing.T) {
	log.Printf("Testing creating an account with the base authenticator")
	authenticator := testSetup()

	created, err := authenticator.CreateUser(defaultUsername, defaultPassword)

	if !created {
		t.Errorf("The account wasn't be created.\n%s", err.Error())
	}
}

func Test_BaseAuthenticatorDoesntAllowTheSameUsername(t *testing.T) {
	log.Printf("Testing cerating an account using an in-use username.")
	authenticator := testSetup()

	created, err := authenticator.CreateUser(defaultUsername, defaultPassword)

	if !created {
		t.Errorf("The account wasn't be created.\n%s", err.Error())
	}

	created, err = authenticator.CreateUser(defaultUsername, defaultPassword)

	if created {
		t.Errorf("An account was created using the same username.")
	}
}

func Test_BaseAuthenticatorUpdatePassword(t *testing.T) {
	log.Printf("Testing updating a password with the base authenticator")
}

func testSetup() Authenticator {
	authDatabaseEntries = []user{}

	authenticator := NewBaseAuthenticator()
	return authenticator
}
