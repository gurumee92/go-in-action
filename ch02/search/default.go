package search

type deafultMatcher struct{}

func init() {
	var matcher deafultMatcher
	Register("default", matcher)
}

func (m deafultMatcher) Search(feed *Feed, searchTerm string) ([]*Result, error) {
	return nil, nil
}
