package ch09

import (
	"fmt"
	"net/http"
	"net/http/httptest"
	"testing"
)

var json = `
	[
		{
			id: 24661126,
			title: "Flatpak: A security nightmare – two years later",
			points: 90,
			user: "krimeo",
			time: 1601634959,
			time_ago: "an hour ago",
			comments_count: 50,
			type: "link",
			url: "https://www.flatkill.org/2020/",
			domain: "flatkill.org"
		},
	]
`

func mockServer() *httptest.Server {
	f := func(w http.ResponseWriter, r *http.Request) {
		w.WriteHeader(200)
		w.Header().Set("Content-Type", "application/json")
		fmt.Fprintln(w, json)
	}
	return httptest.NewServer(http.HandlerFunc(f))
}

func TestMockDownload(t *testing.T) {
	// given
	const checkMark = "\u2713"
	const ballotX = "\u2717"
	server := mockServer()
	statusCode := 200

	t.Log("컨텐츠 다운로드 시작")
	t.Logf("\tURL \"%s\" check status code \"%d\"", server.URL, statusCode)
	// when
	resp, err := http.Get(server.URL)

	if err != nil {
		t.Fatal("\t\t HTTP GET Check", ballotX, err)
	}

	t.Log("\t\t HTTP GET Check", checkMark)
	defer resp.Body.Close()

	// then
	if resp.StatusCode == statusCode {
		t.Logf("\t\t Status Code Check \"%d\": \"%v\"", statusCode, checkMark)
	} else {
		t.Errorf("\t\t Status Code Check \"%d\": \"%v\" \"%v\"", statusCode, ballotX, resp.StatusCode)
	}
}
