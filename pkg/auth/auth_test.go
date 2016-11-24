package auth

import (
	"errors"
	"fmt"
	"log"
	"os"
	"testing"
)

var testUserCount = 5

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

	//cleanupTests()
}

func Test_UsernameIsUpdatedInFile(t *testing.T) {
	oldUser := []byte("Test User 3")
	newUser := []byte("New Test User 3")

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

	cleanupTests()
}

func Test_PasswordIsUpdatedInFile(t *testing.T) {

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
	password := []byte("Password")
	for i := 0; i < testUserCount; i++ {
		salt, pass := sha256WithSalt(password, nil)
		username := []byte(fmt.Sprintf("Test User %d", i))

		authDatabaseEntries = append(authDatabaseEntries, User{Username: username, Password: pass[:], Salt: salt[:]})
	}
}

func printCurrentMemoryDatabase() {
	for i := 0; i < len(authDatabaseEntries); i++ {
		fmt.Printf("\tUser: %s\n", authDatabaseEntries[i].Username)
	}
}
