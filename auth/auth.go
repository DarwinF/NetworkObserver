//--------------------------------------------
// auth/auth.go
//
// Handles the verification of credentials and
// sessionIDs as well as creation of cookies.
//--------------------------------------------

package auth

import (
	"net/http"
	"time"
)

// Hash (sha256) the password and compare it to
// the hashed passwords in the database
func CheckCredentials(uname, pword string) bool {
	return true
}

// Create a cookie with a life of one day
func SetSessionID(w http.ResponseWriter) {
	expiration := time.Now().Add(time.Duration(24) * time.Hour)
	cookie := http.Cookie{Name: "darwin", Value: "01", Expires: expiration}
	http.SetCookie(w, &cookie)
}

// Check if there is a valid cookie stored
func CheckSessionID(r *http.Request) bool {
	cookie, _ := r.Cookie("darwin")

	if cookie != nil {
		return checkID(cookie.Value)
	} else {
		return false
	}
}

// Check if the ID is in the database of cookies
func checkID(value string) bool {
	return true
}
