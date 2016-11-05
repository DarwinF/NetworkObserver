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

	"github.com/darwinfroese/networkobserver/pkg/settings"
)

var cookieName string = "NetObAuth"

var loc string = settings.AppLocation
var pwloc string = loc + "/.password"
var cloc string = loc + "/.cookies"

var cookieValues []int

var rgen *rand.Rand
var max = 999999
var min = 111111

func init() {
	src := rand.NewSource(time.Now().UnixNano())
	rgen = rand.New(src)

	cookieValues = make([]int, 0)
}

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
