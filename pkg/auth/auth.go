package auth

import (
	"os"

	"github.com/darwinfroese/networkobserver/pkg/settings"
)

var authDb *os.File

func init() {
	authDbLoc := settings.AuthenticationDBLocation + settings.AuthenticationDBName

	if _, err := os.Stat(authDbLoc); os.IsNotExist(err) {
		file, err := createFile(authDbLoc)

		if err != nil {
			panic(err.Error())
		}

		authDb = file
		return
	}

	file, err := readFile(authDbLoc)

	if err != nil {
		panic(err.Error())
	}

	authDb = file
}

func readFile(fileLocation string) (*os.File, error) {
	file, err := os.Open(fileLocation)

	return file, err
}

func createFile(fileLocation string) (*os.File, error) {
	file, err := os.Create(fileLocation)

	return file, err
}

func writeUserToFile(username, password, salt string) (bool, error) {
	return false, nil
}

func updateUsernameInFile(username, newUsername string) (bool, error) {
	return false, nil
}

func updatePasswordInFile(username, password, salt string) (bool, error) {
	return false, nil
}
