package auth

import (
	"log"
	"os"

	"bytes"

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

	// TODO: Refactor this into something better
	buffer := make([]byte, 32*3+2)
	for n, err := file.Read(buffer); n > 0 && err != nil; {
		parseLineToUser(buffer)
		records++
	}

	return records, nil
}

func createFile(fileLocation string) error {
	_, err := os.Create(fileLocation)

	return err
}

func parseLineToUser(line []byte) (bool, error) {
	user := User{
		Username: line[:32],
		Password: line[33:65],
		Salt:     line[66:],
	}

	authDatabaseEntries = append(authDatabaseEntries, user)

	return true, nil
}

func writeAllUsersToFile() (int, error) {
	records := 0
	file, err := os.OpenFile(authDbLoc, os.O_WRONLY, 0777)

	if err != nil {
		return records, err
	}

	defer file.Close()

	for i := range authDatabaseEntries {
		success, err := writeUserToFile(authDatabaseEntries[i])

		if success {
			records++
		} else {
			log.Printf("[Error] Couldn't write entry to database. %s", err.Error())
		}
	}

	return records, nil
}

func writeUserToFile(user User) (bool, error) {
	line := makeByteSliceFromUser(user)
	file, err := os.Open(authDbLoc)

	if err != nil {
		log.Printf("Couldn't open the database for writing.\n%s", err.Error())
		return false, err
	}

	_, err = file.Write(line)

	if err != nil {
		log.Printf("Couldn't write the record to the database.\n%s", err.Error())
		return false, err
	}

	return true, nil
}

func updateUserInFile(user User, offset int64) (bool, error) {
	line := makeByteSliceFromUser(user)
	file, err := os.Open(authDbLoc)

	if err != nil {
		log.Printf("Couldn't open the database for updating.\n%s", err.Error())
		return false, err
	}

	_, err = file.WriteAt(line, offset)

	if err != nil {
		log.Printf("Couldn't update the database at offest: %d\n%s", offset, err.Error())
		return false, err
	}

	return true, nil
}

func updateUsername(username, newUsername []byte) bool {
	var updated = false
	for i := range authDatabaseEntries {
		if bytes.Equal(authDatabaseEntries[i].Username, username) {
			authDatabaseEntries[i].Username = newUsername
			updateUserInFile(authDatabaseEntries[i], int64(dbEntryLineLen*i))
			updated = true
		}
	}

	written, err := writeAllUsersToFile()

	if err != nil {
		log.Printf("[ERROR] Couldn't write updated users to file. %d users written.\n%s", written, err.Error())
		updated = false
	}

	return updated
}

func updatePassword(username, password, salt []byte) bool {
	for i := range authDatabaseEntries {
		if bytes.Equal(authDatabaseEntries[i].Username, username) {
			authDatabaseEntries[i].Password = password
			authDatabaseEntries[i].Salt = salt
			return true
		}
	}

	return false
}

func makeByteSliceFromUser(user User) []byte {
	line := make([]byte, dbEntryLineLen)

	line = append(line, getUsername(user.Username)...)

	line = append(line, []byte(",")...)
	line = append(line, user.Password...)
	line = append(line, []byte(",")...)
	line = append(line, user.Salt...)

	return line
}

func getUsername(username []byte) []byte {
	if len(username) == usernameMaxLen {
		return username
	}

	newUsername := make([]byte, usernameMaxLen)

	newUsername = append(newUsername, username...)

	for i := len(username); i < usernameMaxLen; i++ {
		newUsername[i] = '\x00'
	}

	return newUsername
}
