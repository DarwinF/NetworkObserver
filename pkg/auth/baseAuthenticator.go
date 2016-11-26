package auth

import (
	"errors"
	"fmt"
	"log"
)

// NewBaseAuthenticator - Returns an authenticator
func NewBaseAuthenticator() Authenticator {
	settings := Settings{
		EncryptionMethod: "sha256",
		SaltLength:       saltMaxLen,
		UseSalt:          true,
	}

	shaEncrypter, err := newShaEncrypter(&settings)

	if err != nil {
		log.Printf("[ERROR] Couldn't create the encryptor for the base authenticator.\n%s", err.Error())
	}

	authenticator := baseAuthenticator{enc: shaEncrypter}

	return &authenticator
}

func (adapter *baseAuthenticator) Login(username, password string) (bool, error) {
	user, inuse := checkIfUsernameInDatabase(username)

	if !inuse {
		errMsg := fmt.Sprintf("The username %s is not in the database.", username)
		err := errors.New(errMsg)
		return false, err
	}

	valid := adapter.enc.Validate(password, user.Salt, user.Password)

	if !valid {
		err := errors.New("The password entered was incorrect.")
		return false, err
	}

	return true, nil
}

func (adapter *baseAuthenticator) CreateUser(username, password string) (bool, error) {
	_, inuse := checkIfUsernameInDatabase(username)

	if inuse {
		errMsg := fmt.Sprintf("The username %s is already in use.", username)
		return false, errors.New(errMsg)
	}

	encryptedPass, salt := adapter.enc.Encrypt(password)

	user := user{
		Username: username,
		Password: encryptedPass,
		Salt:     salt,
	}

	authDatabaseEntries = append(authDatabaseEntries, user)

	return true, nil
}

func (adapter *baseAuthenticator) UpdatePassword(username, oldPassword, newPassword string) (bool, error) {
	user, inuse := checkIfUsernameInDatabase(username)

	if !inuse {
		errMsg := fmt.Sprintf("There is no account with the username %s in the database.", username)
		return false, errors.New(errMsg)
	}

	correctPassword := adapter.enc.Validate(oldPassword, user.Salt, user.Password)

	if !correctPassword {
		return false, errors.New("The password entered was incorrect.")
	}

	encryptedPass, salt := adapter.enc.Encrypt(newPassword)
	updated := updatePassword(username, encryptedPass, salt)

	if !updated {
		errMsg := fmt.Sprintf("There was an error updated the password for the user %s in the database.", username)
		return false, errors.New(errMsg)
	}

	return true, nil
}

func (adapter *baseAuthenticator) UpdateUsername(oldUsername, newUsername string) (bool, error) {
	_, inuse := checkIfUsernameInDatabase(oldUsername)

	if !inuse {
		errMsg := fmt.Sprintf("There is no account with the username %s in the database.", oldUsername)
		return false, errors.New(errMsg)
	}

	_, available := checkIfUsernameInDatabase(newUsername)

	if available {
		errMsg := fmt.Sprintf("The username %s is not available.", newUsername)
		return false, errors.New(errMsg)
	}

	updated := updateUsername(oldUsername, newUsername)

	if !updated {
		err := errors.New("Unable to update the username in the database.")
		return false, err
	}

	return true, nil
}

func checkIfUsernameInDatabase(username string) (user, bool) {
	user := user{}
	found := false
	for i := range authDatabaseEntries {
		if username == authDatabaseEntries[i].Username {
			user = authDatabaseEntries[i]
			found = true
		}
	}

	return user, found
}
