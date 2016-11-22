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

var authDbLoc string

func init() {
	authDbLoc = settings.AuthenticationDBLocation + settings.AuthenticationDBName

	if _, err := os.Stat(authDbLoc); os.IsNotExist(err) {
		err := createFile(authDbLoc)

		if err != nil {
			log.Panicf("Error creating the file: %s\n%s", authDbLoc, err.Error())
		}

		return
	}

	records, err := readFile(authDatabaseEntries, authDbLoc)

	if err != nil {
		log.Panicf("Error reading users from file: %s\n%d Records were read in.", err.Error(), records)
	}
}

func readFile(users []User, fileLocation string) (int, error) {
	var records = 0
	file, err := os.Open(fileLocation)

	if err != nil {
		return records, err
	}

	defer file.Close()

	scanner := bufio.NewScanner(file)
	for scanner.Scan() {
		line := scanner.Text()
		parseLineToUser(users, line)
		records++
	}

	return records, nil
}

func createFile(fileLocation string) error {
	_, err := os.Create(fileLocation)

	return err
}

func parseLineToUser(users []User, line string) (bool, error) {
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

	users = append(users, user)

	return true, nil
}

func writeAllUsersToFile(users []User) (int, error) {
	records := 0
	file, err := os.Open(authDbLoc)

	if err != nil {
		return records, err
	}

	defer file.Close()

	for i := range users {
		line := users[i].Username + "," + users[i].Password + "," + users[i].Salt
		_, err := file.WriteString(line)

		if err != nil {
			return records, err
		}

		records++
	}

	return records, nil
}

func writeUserToFile(username, password, salt string) (bool, error) {
	line := username + "," + password + "," + salt
	file, err := os.Open(authDbLoc)

	if err != nil {
		log.Panicf("Couldn't open the database for writing.\n%s", err.Error())
	}

	_, err = file.WriteString(line)

	if err != nil {
		log.Panicf("Couldn't write the record to the database.\n%s", err.Error())
	}

	return true, nil
}

func updateUsername(users []User, username, newUsername string) bool {
	for i := range users {
		if users[i].Username == username {
			users[i].Username = newUsername
			return true
		}
	}

	return false
}

func updatePassword(users []User, username, password, salt string) bool {
	for i := range users {
		if users[i].Username == username {
			users[i].Password = password
			users[i].Salt = salt
			return true
		}
	}

	return false
}
