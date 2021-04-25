package main

import (
	"flag"
	"fmt"
	"os"
)

func main() {

	// 'videos get' subcommand
	getCmd := flag.NewFlagSet("get", flag.ExitOnError)

	// inputs for `videos get` command
	getAll := getCmd.Bool("all", false, "Get all videos") // default is set to false
	getID := getCmd.String("id", "", "YouTube video ID")

	// 'videos add' subcommand
	addCmd := flag.NewFlagSet("add", flag.ExitOnError)

	addID := addCmd.String("id", "", "YouTube video ID")

	// inputs for `videos add` command
	addTitle := addCmd.String("title", "", "YouTube video Title")
	addUrl := addCmd.String("url", "", "YouTube video URL")
	addImageUrl := addCmd.String("imageurl", "", "YouTube video Image URL")
	addDesc := addCmd.String("desc", "", "YouTube video description")

	// Validation
	if len(os.Args) < 2 {
		fmt.Println("expected 'get' or 'add' subcommands")
		os.Exit(1)
	}

	//look at the 2nd argument's value
	switch os.Args[1] {
	case "get": // if its the 'get' command
		//hande get here
		HandleGet(getCmd, getAll, getID)
	case "add": // if its the 'add' command
		//hande add here
		HandleAdd(addCmd, addID, addTitle, addUrl, addImageUrl, addDesc)
	default: // if we don't understand the input
	}

}

func HandleGet(getCmd *flag.FlagSet, all *bool, id *string) {
	// parse everything after the videos <subcommand>
	getCmd.Parse(os.Args[2:])
	// input validation
	if *all == false && *id == "" {
		fmt.Printf("id is required or specify --all for all videos \n")
		getCmd.PrintDefaults()
		os.Exit(1)
	}

	// handle the scenario if user passing --all flag
	if *all {
		//return all videos
		videos := getVideos()

		fmt.Printf("ID \t Title \t URL \t ImageURL \t Description \n")
		for _, video := range videos {
			fmt.Printf("%v \t %v \t %v \t %v \t %v \n", video.Id, video.Title, video.Url, video.Imageurl, video.Description)
		}

		// stops the program
		return
	}

	// handle when user is searching for a video by ID
	if *id != "" {
		videos := getVideos()
		id := *id
		for _, video := range videos {
			if id == video.Id {
				fmt.Printf("ID \t Title \t URL \t ImageURL \t Description \n")
				fmt.Printf("%v \t %v \t %v \t %v \t %v \n", video.Id, video.Title, video.Url, video.Imageurl, video.Description)

			}
		}
	}
}

func ValidateVideo(addCmd *flag.FlagSet, id *string, title *string, url *string, imageUrl *string, description *string) {
	if *id == "" || *title == "" || *url == "" || *imageUrl == "" || *description == "" {
		fmt.Printf("all fields are required for adding a video \n")
		addCmd.PrintDefaults()
		os.Exit(1)
	}
}

func HandleAdd(addCmd *flag.FlagSet, id *string, title *string, url *string, imageUrl *string, description *string) {
	ValidateVideo(addCmd, id, title, url, imageUrl, description)

	// define a video struct with the CLI input
	video := video{
		Id:          *id,
		Title:       *title,
		Description: *description,
		Imageurl:    *imageUrl,
		Url:         *url,
	}

	// get the existing videos
	videos := getVideos()

	// append the input video to the list
	videos = append(videos, video)

	// save the new updated video list
	saveVideos(videos)
}
