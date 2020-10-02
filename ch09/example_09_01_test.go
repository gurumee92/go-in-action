package ch09

import (
	"net/http"
	"testing"
)

const checkMark = "\u2713"
const ballotX = "\u2717"

func TestDownload(t *testing.T) {
	// given
	params := []struct {
		url        string
		statusCode int
	}{
		{
			url:        "https://api.hnpwa.com/v0/news/1.json",
			statusCode: 200,
		},
		{
			url:        "https://api.hnpwa.com/v1/news/1.json",
			statusCode: 404,
		},
		{
			url:        "https://api.hnpwa.com/v0/news/100.json",
			statusCode: 500,
		},
	}

	t.Log("컨텐츠 다운로드 시작")
	for _, param := range params {
		t.Logf("\tURL \"%s\" check status code \"%d\"", param.url, param.statusCode)

		resp, err := http.Get(param.url)

		if err != nil {
			t.Fatal("\t\t HTTP GET Check", ballotX, err)
		}

		t.Log("\t\t HTTP GET Check", checkMark)
		defer resp.Body.Close()

		// then
		if resp.StatusCode == param.statusCode {
			t.Logf("\t\t Status Code Check \"%d\": \"%v\"", param.statusCode, checkMark)
		} else {
			t.Errorf("\t\t Status Code Check \"%d\": \"%v\" \"%v\"", param.statusCode, ballotX, resp.StatusCode)
		}
	}

}
