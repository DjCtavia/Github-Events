package server

import "testing"

func TestGithubEvent_ToString(t *testing.T) {
	githubEvent_ToString_WrongEvent(t)
	testGithubEvent_ToString_CommitCommentToString(t)
}

func githubEvent_ToString_WrongEvent(t *testing.T) {
	var event = GithubEvent("test")
	var convertedEventToString = event.ToString()

	if convertedEventToString != "" {
		t.Errorf(`Expected empty string, got: %s`, convertedEventToString)
	}
}

func testGithubEvent_ToString_CommitCommentToString(t *testing.T) {
	var event = GithubEvent(COMMITCOMMENT)
	var convertedEventToString = event.ToString()

	if convertedEventToString != COMMITCOMMENT_BEAUTIFY {
		t.Errorf(`Expected empty string, got: %s`, convertedEventToString)
	}
}
