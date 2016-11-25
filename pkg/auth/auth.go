package auth

import (
	"bufio"
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

	authDatabaseEntries = []User{}
	defer file.Close()

	scanner := bufio.NewScanner(file)
	for scanner.Scan() {
		line := scanner.Text()
		user, success := parseLineToUser(line)

		if success {
			authDatabaseEntries = append(authDatabaseEntries, user)
			records++
		}
	}

	return records, nil
}

func createFile(fileLocation string) error {
	file, err := os.Create(fileLocation)

	file.Close()
	return err
}

func parseLineToUser(line string) (User, bool) {
	fields := strings.Split(line, ":")

	if len(fields) != 3 {
		return User{}, false
	}

	user := User{
		Username: fields[0],
		Password: fields[1],
		Salt:     fields[2],
	}

	return user, true
}

func writeAllUsersToFile() (int, error) {
	records := 0
	file, err := os.OpenFile(authDbLoc, os.O_WRONLY, 0777)

	if err != nil {
		log.Printf("[Error] Couldn't open the database for writing. %s", err.Error())
		return records, err
	}

	defer file.Close()

	for i := range authDatabaseEntries {
		writeUserToFile(authDatabaseEntries[i], file)
		records++
	}

	return records, nil
}

func writeUserToFile(user User, file *os.File) {
	line := userToString(user)

	w := bufio.NewWriter(file)

	fmt.Fprintf(w, "%s\n", line)
	w.Flush()
}

func updateUsername(username, newUsername string) bool {
	var updated = false
	for i := range authDatabaseEntries {
		if authDatabaseEntries[i].Username == username {
			authDatabaseEntries[i].Username = newUsername
			updated = true
		}
	}

	if updated {
		writeAllUsersToFile()
	}

	return updated
}

func updatePassword(username, password, salt string) bool {
	var updated = false
	for i := range authDatabaseEntries {
		if authDatabaseEntries[i].Username == username {
			authDatabaseEntries[i].Password = password
			authDatabaseEntries[i].Salt = salt
			updated = true
		}
	}

	if updated {
		writeAllUsersToFile()
	}

	return updated
}

func userToString(user User) string {
	line := user.Username + ":" + user.Password + ":" + user.Salt

	return line
}
