package main

import (
	"encoding/json"
	"fmt"
	"log"
	"net/http"
)

type Post struct {
	Id            int    `json:"id"`
	Title         string `json:"title"`
	Points        int    `json:"points"`
	User          string `json:"user"`
	Time          int    `json:"time"`
	TimeAgo       string `json:"time_ago"`
	CommentsCount int    `json:"comments_count"`
	Type          string `json:"type"`
	Url           string `json:"url"`
	Domain        string `json:"domain"`
}

func main() {
	uri := "https://api.hnpwa.com/v0/news/1.json"
	resp, err := http.Get(uri)

	if err != nil {
		log.Fatalln("Error: ", err)
	}

	defer resp.Body.Close()

	var posts []Post
	err = json.NewDecoder(resp.Body).Decode(&posts)

	if err != nil {
		log.Fatalln("Error: ", err)
	}

	pretty, err := json.MarshalIndent(posts, "", "    ")

	if err != nil {
		log.Fatalln("Error: ", err)
	}

	fmt.Println(string(pretty))
}
