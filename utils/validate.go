package utils

import (
	"net/url"
)

//CheckURL used to check if it is a valid url
func CheckURL(uri string) bool {
	_, err := url.ParseRequestURI(uri)

	if err != nil {
		return false
	}

	return true
}
