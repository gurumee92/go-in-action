package search

import (
	"encoding/json"
	"os"
)

const dataFile = "go-in-action/ch02/data/data.json"

// Feed is model of feed
type Feed struct {
	Name string `json:"site"`
	URI  string `json:"link"`
	Type string `json:"type"`
}

// RetrieveFeeds is retrieve feeds to data/data.json
func RetrieveFeeds() ([]*Feed, error) {
	file, err := os.Open(dataFile)
	if err != nil {
		return nil, err
	}

	// defer는 함수 컨텍스트가 끝나는 동시에 실행된다.
	defer file.Close()
	var feeds []*Feed
	err = json.NewDecoder(file).Decode(&feeds)
	return feeds, err
}
