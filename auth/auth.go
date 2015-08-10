//--------------------------------------------
// auth/auth.go
//
// Handles the verification of credentials and
// sessionIDs as well as creation of cookies.
//--------------------------------------------

package auth

import (
	"bufio"
	"math/rand"
	"net/http"
	"os"
	"strconv"
	"strings"
	"time"
)

var cookieName = "NetObAuth"

var rgen *rand.Rand
var max = 999999
var min = 111111

func init() {
	src := rand.NewSource(time.Now().UnixNano())
	rgen = rand.New(src)
}

// Read the stored usernames and password hashes from the file
func CheckCredentials(uname string, pword [32]byte) bool {
	valid := false
	file, _ := os.Open(".password")
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
	value := rgen.Intn(max-min) + min

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

// Check if the ID is in the database of cookies
func checkID(value string) bool {
	valid := false

	file, _ := os.Open(".cookies")
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

	file, _ := os.Open(".password")
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
	file, _ := os.OpenFile(".password", os.O_APPEND|os.O_WRONLY, 0600)
	defer file.Close()

	w := bufio.NewWriter(file)
	defer w.Flush()

	w.WriteString(uname + ":" + string(pword[:]) + "\n")
}

func writeCookieValue(value int) {
	file, _ := os.OpenFile(".cookies", os.O_APPEND|os.O_WRONLY, 0600)
	defer file.Close()

	w := bufio.NewWriter(file)
	defer w.Flush()

	w.WriteString(strconv.Itoa(value) + "\n")
}
