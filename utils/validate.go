package utils

import (
	"fmt"
	"net/url"
)

//CheckURL used to check if it is a valid url
func CheckURL(uri string) {
	u, err := url.ParseRequestURI(uri)

	if err != nil {
		fmt.Println(err.Error())
	}

	if u != nil {
		fmt.Println(u)
	}
}
