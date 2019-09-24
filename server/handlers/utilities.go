package handlers

import "net/http"

func newCookie(name, value string) *http.Cookie {
	return &http.Cookie{Name: name, Value: value}
}
