package main

import (
	//"fmt"
	"log"
	//"io"
	//"os"
	//"github.com/zergon321/reisen"
	youtube "github.com/knadh/go-get-youtube/youtube"
)

func main() {
	// get the video object (with metdata)
	video, err := youtube.Get("FTl0tl9BGdc")
  if err != nil {
	log.Fatal(err)
  }

	// download the video and write to file
	option := &youtube.Option{
		Rename: true,  // rename file using video title
		Resume: true,  // resume cancelled download
		Mp3:    true,  // extract audio to MP3
	}
	video.Download(0, "video.mp4", option)
}
