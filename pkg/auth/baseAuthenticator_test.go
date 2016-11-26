package auth

import "testing"
import "log"

func Test_BaseAuthenticatorLoginIsSuccessful(t *testing.T) {
	log.Printf("Testing logging in with the base authenticator.")
	authenticator := testSetupWithDefaultUser()

	loggedIn, err := authenticator.Login(defaultUsername, defaultPassword)

	if !loggedIn {
		t.Errorf("Couldn't log into the account.\n%s", err.Error())
	}
}

func Test_BaseAuthenticatorLoginFailsWithIncorrectPassword(t *testing.T) {
	log.Printf("Testing logging in with an incorrect password fails.")
	authenticator := testSetupWithDefaultUser()

	loggedIn, err := authenticator.Login(defaultUsername, "incorrect password")

	if loggedIn {
		t.Errorf("Logged into the account with the wrong password.\n%s", err.Error())
	}
}

func Test_BaseAuthenticatorLoginFailsWithUnusedUsername(t *testing.T) {
	log.Printf("Testing logging in with an unused username fails.")
	authenticator := testSetupWithDefaultUser()

	loggedIn, err := authenticator.Login("incorrect user", defaultPassword)

	if loggedIn {
		t.Errorf("Logged into an account that didn't exist.\n%s", err.Error())
	}
}

func Test_BaseAuthenticatorCreateAccountIsSuccessful(t *testing.T) {
	log.Printf("Testing creating an account with the base authenticator.")
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

func Test_BaseAuthenticatorUpdatesUsernameSuccessfully(t *testing.T) {
	log.Printf("Testing updating to a username that isn't in use.")
	newUsername := "New User"

	authenticator := testSetupWithDefaultUser()

	updated, err := authenticator.UpdateUsername(defaultUsername, newUsername)

	if !updated {
		t.Errorf("Unable to update the user %s in the database.\n", err.Error())
	}

	found := searchForUser(newUsername)

	if !found {
		t.Errorf("The user %s was not found.", newUsername)
	}
}

func Test_BaseAuthenticatorCannotUpdateUsernameToInUseUsername(t *testing.T) {
	log.Printf("Testing updating to a username that is in use.")
	inUseName := "User in use"

	authenticator := testSetupWithDefaultUser()
	created, err := authenticator.CreateUser(inUseName, defaultPassword)

	if !created {
		t.Errorf("Couldn't create the user %s.\n%s", inUseName, err.Error())
	}

	updated, _ := authenticator.UpdateUsername(defaultUsername, inUseName)

	if updated {
		t.Errorf("Updated the username to a username that is in use.")
	}
}

func Test_BaseAuthenticatorUpdatePassword(t *testing.T) {
	log.Printf("Testing updating a password with the base authenticator.")
}

func testSetup() Authenticator {
	authDatabaseEntries = []user{}

	authenticator := NewBaseAuthenticator()
	return authenticator
}

func testSetupWithDefaultUser() Authenticator {
	authenticator := testSetup()

	authenticator.CreateUser(defaultUsername, defaultPassword)

	return authenticator
}

func searchForUser(username string) bool {
	found := false
	for i := range authDatabaseEntries {
		if username == authDatabaseEntries[i].Username {
			log.Printf("found user %s", authDatabaseEntries[i].Username)
			found = true
		}
	}

	return found
}
