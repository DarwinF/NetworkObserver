package auth

import (
	"errors"
	"fmt"
	"log"
	"os"
	"testing"
)

var testUserCount = 5
var defaultPassword = "Password"

func Test_FileIsCreatedIfMissing(t *testing.T) {
	log.Printf("Looking for auth database at: %s\n", authDbLoc)

	err := removeExistingAuthFile()

	if err != nil {
		t.Errorf(err.Error())
	}

	err = createFile(authDbLoc)

	if err != nil {
		t.Errorf("Couldn't create file %s\n%s", authDbLoc, err.Error())
	}

	if _, err := os.Stat(authDbLoc); os.IsNotExist(err) {
		t.Errorf("Cannot find created file %s\n%s", authDbLoc, err.Error())
	}

	os.Remove(authDbLoc)
}

func Test_AccountsReadFromFile(t *testing.T) {
	log.Println("Testing reading accounts from file.")

	err := setupTest()

	if err != nil {
		t.Errorf("An error occured setting up the test: %s", err.Error())
	}

	usersRead, err := readFile(authDbLoc)

	if err != nil {
		t.Errorf("An error occured reading users from the file: %s", err.Error())
	}
	if usersRead != testUserCount {
		t.Errorf("Failed to read %d users from file. Total users read: %d", testUserCount, usersRead)
	}

	cleanupTests()
}

func Test_UsernameIsUpdatedInFile(t *testing.T) {
	oldUser := "Test User 3"
	newUser := "New Test User 3"

	log.Println("Testing that username changes are saved to file.")

	err := setupTest()

	if err != nil {
		t.Errorf("An error occured setting up the test: %s", err.Error())
	}

	nameUpdated := updateUsername(oldUser, newUser)

	if !nameUpdated {
		t.Errorf("An error occured updating the name of the user.")
	}

	usersRead, err := readFile(authDbLoc)

	if err != nil {
		t.Errorf("An error occured reading users from the file: %s", err.Error())
	}
	if usersRead != testUserCount {
		t.Errorf("Failed to read %d users from file. Total users read: %d", testUserCount, usersRead)
	}

	found := false
	for i := range authDatabaseEntries {
		if authDatabaseEntries[i].Username == newUser {
			found = true
		}
	}

	if !found {
		t.Errorf("Failed to update the username in file.")
	}

	cleanupTests()
}

func Test_PasswordIsUpdatedInFile(t *testing.T) {
	log.Println("Testing that password changes are saved to file.")

	enc, _ := newShaEncrypter(&shaSettings)
	user := "Test User 4"
	newPassword := "new passw0rd"

	err := setupTest()

	if err != nil {
		t.Errorf("An error occured setting up the test: %s", err.Error())
	}

	oldPass := authDatabaseEntries[3].Password
	oldSalt := authDatabaseEntries[3].Salt
	pass, salt := enc.Encrypt(newPassword)
	passwordUpdated := updatePassword(user, pass, salt)

	if !passwordUpdated {
		t.Errorf("Failed to update the users password")
	}

	usersRead, err := readFile(authDbLoc)

	if err != nil {
		t.Errorf("An error occured reading the users from the file: %s", err.Error())
	}
	if usersRead != testUserCount {
		t.Errorf("Failed to read %d users from file. Total users read: %d", testUserCount, usersRead)
	}

	updated := false
	for i := range authDatabaseEntries {
		if authDatabaseEntries[i].Username == user {
			p := authDatabaseEntries[i].Password
			s := authDatabaseEntries[i].Salt

			if p == pass && p != oldPass && s == salt && s != oldSalt {
				updated = true
			}
		}
	}

	if !updated {
		t.Errorf("Failed to update the password in file")
	}

	cleanupTests()
}

// Setup functions
func removeExistingAuthFile() error {
	if _, err := os.Stat(authDbLoc); os.IsExist(err) {
		err := os.Remove(authDbLoc)

		if err != nil {
			msg := fmt.Sprintf("Couldn't remove file %s\n%s", authDbLoc, err.Error())
			return errors.New(msg)
		}
	}

	return nil
}

func setupTest() error {
	err := createFile(authDbLoc)

	if err != nil {
		log.Printf("[ERROR] Couldn't create the file %d\n", authDbLoc)
		return err
	}

	populateDatabase()

	usersWritten, err := writeAllUsersToFile()

	if usersWritten != testUserCount {
		log.Printf("Failed to write %d users to file. Total users written: %d", testUserCount, usersWritten)
	}

	return err
}

func cleanupTests() {
	authDatabaseEntries = []User{}
	os.Remove(authDbLoc)
}

func populateDatabase() {
	encrypter, _ := newShaEncrypter(&shaSettings)

	for i := 0; i < testUserCount; i++ {
		pass, salt := encrypter.Encrypt(defaultPassword)
		username := fmt.Sprintf("Test User %d", i)

		authDatabaseEntries = append(authDatabaseEntries, User{Username: username, Password: pass, Salt: salt})
	}
}

func printCurrentMemoryDatabase() {
	for i := 0; i < len(authDatabaseEntries); i++ {
		fmt.Printf("\tUser: %s\n", authDatabaseEntries[i].Username)
	}
}
