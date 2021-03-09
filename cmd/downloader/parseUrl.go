package downloader

import (
	"fmt"
	"net/url"
	"regexp"
)

//TODO: Working on extracting video id from the url
//https://gobyexample.com/url-parsing
//https://gobyexample.com/regular-expressions

//GetVideoIDFromURL get video url from url link
func GetVideoIDFromURL(videoURL string) string {

	url, err := url.Parse(videoURL)

	var videoID string

	if err != nil {
		panic(err)
	}

	urlPattern := regexp.MustCompile("([^\\?]*)/i")

	if len(url.Host) > 0 && url.Host == "www.youtube.com" {

		foundStatements := urlPattern.FindAllString(url.String(), -1)
		fmt.Println(foundStatements)

		if len(foundStatements) <= 0 {
			panic("The video id is not defined")
		}

	} else {
		panic("Invalid video url")
	}

	return videoID
}
