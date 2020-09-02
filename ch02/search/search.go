package search

import (
	"log"
	"sync"
)

// 이렇게 소문자로 변수/함수/구조체/메서드를 만들면 패키지 내부에서만 사용할 수 있다.
var matchers = make(map[string]Matcher)

// Run is function exact search login
func Run(searchTerm string) {
	feeds, err := RetrieveFeeds()
	if err != nil {
		log.Fatalln(err)
	}

	results := make(chan *Result)
	var waitGroup sync.WaitGroup
	waitGroup.Add(len(feeds))

	for _, feed := range feeds {
		matcher, exists := matchers[feed.Type]
		if !exists {
			matcher = matchers["default"]
		}

		go func(matcher Matcher, feed *Feed) {
			Match(matcher, feed, searchTerm, results)
			waitGroup.Done()
		}(matcher, feed)
	}

	go func() {
		waitGroup.Wait()
		close(results)
	}()

	Display(results)
}

// Register is register matcher
func Register(feedType string, matcher Matcher) {
	if _, exists := matchers[feedType]; exists {
		log.Fatalln(feedType, "검색기가 이미 등록되었습니다.")
	}

	log.Println("등록 완료", feedType, " 검색기")
	matchers[feedType] = matcher
}
