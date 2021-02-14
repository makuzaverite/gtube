package cmd

import (
	"fmt"
	"regexp"
)

//DownloadDefaultVideo  a  video when url is specified from youtube
func DownloadDefaultVideo(url string) {
	fmt.Printf("Donwloaing video from youtube with url %s\n", url)
	videoID := getVideoID(url)

	fmt.Println(videoID)
}

func getVideoID(videoID string) string {
	checkid := regexp.MustCompile("v=+")
	return checkid.FindString(videoID)
}
