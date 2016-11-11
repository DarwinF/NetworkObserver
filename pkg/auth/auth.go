// Package auth provides functionality for authenticating
// users to the system.
//
// Default Configuration Values
// * Password Encryption Method: sha256
// * Salt length in bits: 64
package auth

import (
	"crypto/sha256"
	"math/rand"

	"fmt"
)

// Settings is a struct of all the cofiguration options
// used for configuring the auth package
type Settings struct {
	EncryptionMethod string
	SaltLength       int
	UseSalt          bool
}

// Authenticator object for authentication
type Authenticator interface {
	Encrypt(string) (string, string, error)
	Validate(string, string, string) (bool, error)
}

type authAdapter struct {
}

var authSettings Settings

func init() {
	authSettings = Settings{
		EncryptionMethod: "sha256",
		SaltLength:       64,
		UseSalt:          true,
	}
}

// NewAuthenticator returns a new authenticator that can be used for authenticating
func NewAuthenticator(settings *Settings) (Authenticator, error) {
	a := authAdapter{}

	if settings != nil {
		authSettings = *settings
	}

	return &a, nil
}

func (a *authAdapter) Encrypt(v string) (eValue, salt string, err error) {
	if authSettings.UseSalt {
		e, s := sha256WithSalt(v, "")
		eValue = string(e[:])
		salt = string(s[:])
		return
	}

	e := sha256WithoutSalt(v)
	eValue = string(e[:])

	return
}

func (a *authAdapter) Validate(input, salt, password string) (valid bool, err error) {
	var encryptedString string
	valid = false

	if authSettings.UseSalt {
		e, _ := sha256WithSalt(input, salt)
		encryptedString = string(e[:])
	} else {
		e := sha256WithoutSalt(input)
		encryptedString = string(e[:])
	}

	if encryptedString == password {
		valid = true
	}

	return
}

/*
// Read the stored usernames and password hashes from the file
func CheckCredentials(uname string, pword [32]byte) bool {
	valid := false
	file, _ := os.Open(pwloc)
	defer file.Close()

	pws := string(pword[:])
	scanner := bufio.NewScanner(file)

	for scanner.Scan() {
		str := scanner.Text()
		line := strings.Split(str, ":")
		pw := line[1]

		if line[0] == uname && pw == pws {
			valid = true
			break
		} else if line[0] == uname && pw == pws {
			break
		}
	}

	return valid
}

// Create a cookie with a life of one day
func SetSessionID(w http.ResponseWriter) {
	var value int

	value = rgen.Intn(max-min) + min
	for used(value) {
		value = rgen.Intn(max-min) + min
	}

	expiration := time.Now().Add(time.Duration(24) * time.Hour)
	cookie := http.Cookie{Name: cookieName, Value: strconv.Itoa(value), Expires: expiration}
	http.SetCookie(w, &cookie)

	writeCookieValue(value)
}

// Check if there is a valid cookie stored
func CheckSessionID(r *http.Request) bool {
	cookie, _ := r.Cookie(cookieName)

	if cookie != nil {
		return checkID(cookie.Value)
	} else {
		return false
	}
}

func RemoveCookie(w http.ResponseWriter, r *http.Request) {
	cookie, _ := r.Cookie(cookieName)

	if cookie != nil {
		cookie.MaxAge = -1
		http.SetCookie(w, cookie)
	}
}

// Check if the ID is in the database of cookies
func checkID(value string) bool {
	valid := false

	file, _ := os.Open(cloc)
	defer file.Close()

	scanner := bufio.NewScanner(file)

	for scanner.Scan() {
		str := scanner.Text()

		if str == value {
			valid = true
			break
		}
	}

	return valid
}

func UsernameInUse(uname string) bool {
	used := false

	file, _ := os.Open(pwloc)
	defer file.Close()

	scanner := bufio.NewScanner(file)

	for scanner.Scan() {
		str := scanner.Text()
		line := strings.Split(str, ":")

		if line[0] == uname {
			used = true
			break
		}
	}

	return used
}

func SavePassword(uname string, pword [32]byte) {
	file, _ := os.OpenFile(pwloc, os.O_APPEND|os.O_WRONLY, 0600)
	defer file.Close()

	w := bufio.NewWriter(file)
	defer w.Flush()

	w.WriteString(uname + ":" + string(pword[:]) + "\n")
}

func writeCookieValue(value int) {
	file, _ := os.OpenFile(cloc, os.O_APPEND|os.O_WRONLY, 0600)
	defer file.Close()

	w := bufio.NewWriter(file)
	defer w.Flush()

	w.WriteString(strconv.Itoa(value) + "\n")
	cookieValues = append(cookieValues, value)
}

func used(value int) bool {
	for _, v := range cookieValues {
		if v == value {
			return true
		}
	}

	return false
}
*/
func sha256WithoutSalt(value string) [sha256.Size]byte {
	encrypted := sha256.Sum256([]byte(value))

	return encrypted
}

func sha256WithSalt(value, saltValue string) ([sha256.Size]byte, []byte) {
	salt := make([]byte, authSettings.SaltLength)

	if saltValue != "" {
		salt = []byte(saltValue)
	}
	n, err := rand.Read(salt)

	if err != nil {
		fmt.Println("There was an error generating a salt: ", err)
		return [sha256.Size]byte{}, nil
	}

	if n != authSettings.SaltLength {
		fmt.Printf("Only %d characters were read.\n", n)
		return [sha256.Size]byte{}, nil
	}

	saltedVal := append([]byte(value), salt...)
	encrypted := sha256.Sum256(saltedVal)

	return encrypted, salt
}
