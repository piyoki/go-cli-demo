package main

import (
	"encoding/json"
	"io/ioutil"
)

type video struct {
	Id          string
	Title       string
	Description string
	Imageurl    string
	Url         string
}

func getVideos() (videos []video) {

	fileBytes, err := ioutil.ReadFile("./videos.json")

	if err != nil {
		panic(err)
	}

	// convert bytes array to video slices
	err = json.Unmarshal(fileBytes, &videos)

	if err != nil {
		panic(err)
	}

	return videos
}

func saveVideos(videos []video) {

	// convert video slices back to bytes array
	videoBytes, err := json.Marshal(videos)
	if err != nil {
		panic(err)
	}

	// write the bytes array back to the json file
	err = ioutil.WriteFile("./videos.json", videoBytes, 0644)
	if err != nil {
		panic(err)
	}

}
