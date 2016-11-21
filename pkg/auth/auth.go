package auth

import (
	"os"

	"bufio"

	"log"

	"github.com/darwinfroese/networkobserver/pkg/settings"
)

func init() {
	authDbLoc := settings.AuthenticationDBLocation + settings.AuthenticationDBName

	if _, err := os.Stat(authDbLoc); os.IsNotExist(err) {
		return
	}

	records, err := readFile(authDbLoc)

	if err != nil {
		log.Panicf("Error reading users from file: %s\n%d Records were read in.", err.Error(), records)
	}
}

func readFile(fileLocation string) (int, error) {
	var records = 0
	file, err := os.Open(fileLocation)
	defer file.Close()

	if err != nil {
		return records, err
	}

	scanner := bufio.NewScanner(file)
	for scanner.Scan() {
		line := scanner.Text()
		parseLineToUser(line)
		records++
	}

	if err := scanner.Err(); err != nil {
		return records, err
	}

	return records, nil
}

func createFile(fileLocation string) (bool, error) {
	return false, nil
}

func parseLineToUser(line string) (bool, error) {

	return false, nil
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
