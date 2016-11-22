package auth

import (
	"bufio"
	"errors"
	"fmt"
	"log"
	"os"
	"strings"

	"github.com/darwinfroese/networkobserver/pkg/settings"
)

func init() {
	authDbLoc := settings.AuthenticationDBLocation + settings.AuthenticationDBName

	if _, err := os.Stat(authDbLoc); os.IsNotExist(err) {
		err := createFile(authDbLoc)

		if err != nil {
			log.Panicf("Error creating the file: %s\n%s", authDbLoc, err.Error())
		}

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

	if err != nil {
		return records, err
	}

	scanner := bufio.NewScanner(file)
	for scanner.Scan() {
		line := scanner.Text()
		parseLineToUser(line)
		records++
	}

	file.Close()
	return records, err
}

func createFile(fileLocation string) error {
	_, err := os.Create(fileLocation)

	return err
}

func parseLineToUser(line string) (bool, error) {
	data := strings.Split(line, ",")

	if len(data) != 3 {
		msg := fmt.Sprintf("Invalid database entry: %s", line)
		return false, errors.New(msg)
	}

	user := User{
		Username: data[0],
		Password: data[1],
		Salt:     data[2],
	}

	authDatabaseEntries = append(authDatabaseEntries, user)

	return true, nil
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
