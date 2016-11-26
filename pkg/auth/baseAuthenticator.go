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
	return false, nil
}

func (adapter *baseAuthenticator) UpdateUsername(oldUsername, newUsername string) (bool, error) {
	return false, nil
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
