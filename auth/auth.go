package auth

import (
	"net/http"
	"time"
)

func CheckCredentials(uname, pword string) bool {
	return false
}

func SetSessionID(w http.ResponseWriter) {
	// One day long cookie
	expiration := time.Now().Add(time.Duration(24)*time.Hour)
	cookie := http.Cookie{Name:"darwin", Value:"01", Expires:expiration}
	http.SetCookie(w, &cookie)
}

func CheckSessionID(r *http.Request) bool {
	cookie,_ := r.Cookie("darwin")

	if cookie != nil {
		return checkID(cookie.Value)
	} else {
		return false
	}
}

func checkID(value string) bool {
	return true
}