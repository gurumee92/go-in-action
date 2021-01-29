package search

import "log"

// Result is model of result
type Result struct {
	Field   string
	Content string
}

// Matcher is interface feed, and searchTerm
type Matcher interface {
	Search(feed *Feed, searchTerm string) ([]*Result, error)
}

// Match is amtch function
func Match(matcher Matcher, feed *Feed, searchTerm string, results chan<- *Result) {
	searchResults, err := matcher.Search(feed, searchTerm)
	if err != nil {
		log.Println(err)
		return
	}

	for _, result := range searchResults {
		results <- result
	}
}

// Display is fucntion display result
func Display(results chan *Result) {
	for result := range results {
		log.Printf("%s:\n%s\n\n", result.Field, result.Content)
	}
}
