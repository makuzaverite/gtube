package cmd

import (
	"fmt"
	"regexp"

	"github.com/makuzaverite/gtube/utils"
)

//DownloadDefaultVideo  a  video when url is specified from youtube
func DownloadDefaultVideo(url string) {
	isValidURL := utils.CheckURL(url)

	if !isValidURL {
		fmt.Printf("\nInvalid url %s\n", url)
	}

}

func getVideoID(videoID string) string {
	checkid := regexp.MustCompile("v=+")
	return checkid.FindString(videoID)
}
