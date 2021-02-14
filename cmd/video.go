package cmd

import (
	"fmt"
	"regexp"

	"github.com/makuzaverite/gtube/utils"
)

//DownloadDefaultVideo  a  video when url is specified from youtube
func DownloadDefaultVideo(url string) {
	fmt.Printf("Donwloaing video from youtube with url %s\n", url)
	utils.CheckURL(url)
}

func getVideoID(videoID string) string {
	checkid := regexp.MustCompile("v=+")
	return checkid.FindString(videoID)
}
