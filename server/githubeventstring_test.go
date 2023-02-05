package server

import "testing"

func TestGithubEventString_ToString(t *testing.T) {
	githubEventString_ToString_WrongEvent(t)
	testGithubEventString_ToString_CommitCommentToString(t)
}

func githubEventString_ToString_WrongEvent(t *testing.T) {
	var event = GithubEventString("test")
	var convertedEventToString = event.ToString()

	if convertedEventToString != "" {
		t.Errorf(`Expected empty string, got: %s`, convertedEventToString)
	}
}

func testGithubEventString_ToString_CommitCommentToString(t *testing.T) {
	var event = GithubEventString(COMMITCOMMENT)
	var convertedEventToString = event.ToString()

	if convertedEventToString != COMMITCOMMENT_BEAUTIFY {
		t.Errorf(`Expected empty string, got: %s`, convertedEventToString)
	}
}
