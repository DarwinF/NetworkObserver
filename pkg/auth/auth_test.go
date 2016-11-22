package auth

import (
	"errors"
	"fmt"
	"os"
	"testing"
)

var testDbLoc = "../../Test/user_info.db"
var testDatabase = []User{}

func Test_FileIsCreatedIfMissing(t *testing.T) {
	fmt.Printf("Looking for auth database at: %s\n", testDbLoc)

	err := removeExistingAuthFile()

	if err != nil {
		t.Errorf(err.Error())
	}

	err = createFile(testDbLoc)

	if err != nil {
		t.Errorf("Couldn't create file %s\n%s", testDbLoc, err.Error())
	}

	if _, err := os.Stat(testDbLoc); os.IsNotExist(err) {
		t.Errorf("Cannot find created file %s\n%s", testDbLoc, err.Error())
	}

	os.Remove(testDbLoc)
}

func Test_AccountsReadFromFile(t *testing.T) {

}

func Test_AuthInformationIsWrittenToFile(t *testing.T) {

}

func Test_UsernameIsUpdatedInFile(t *testing.T) {

}

func Test_PasswordIsUpdatedInFile(t *testing.T) {

}

// Setup
func removeExistingAuthFile() error {
	if _, err := os.Stat(testDbLoc); os.IsExist(err) {
		err := os.Remove(testDbLoc)

		if err != nil {
			msg := fmt.Sprintf("Couldn't remove file %s\n%s", testDbLoc, err.Error())
			return errors.New(msg)
		}
	}

	return nil
}

func populateTestDatabase() {
	for i := 0; i < 5; i++ {
		salt, password := sha256WithSalt("Password", "")
		username := fmt.Sprintf("Test User %d", i)
		s := string(salt[:len(salt)])
		p := string(password[:len(password)])

		testDatabase = append(testDatabase, User{Username: username, Password: p, Salt: s})
	}
}

func setup() error {
	err := removeExistingAuthFile()

	if err != nil {
		return err
	}

	err = createFile(testDbLoc)

	if err != nil {
		return err
	}

	return nil
}

func setupWithUsers() error {
	err := setup()

	if err != nil {
		return err
	}

	populateTestDatabase()

}
