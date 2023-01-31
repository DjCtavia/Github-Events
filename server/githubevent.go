package server

type GithubEventStringify interface {
	ToString() string
}

type GithubEvent string

const (
	COMMITCOMMENT            = "commitcomment"
	CREATE                   = "create"
	DELETE                   = "delete"
	FORK                     = "fork"
	GOLLUM                   = "gollum"
	ISSUECOMMENT             = "issuecomment"
	ISSUES                   = "issues"
	MEMBER                   = "member"
	PUBLIC                   = "public"
	PULLREQUEST              = "pullrequest"
	PULLREQUESTREVIEW        = "pullrequestreview"
	PULLREQUESTREVIEWCOMMENT = "pullrequestreviewcomment"
	PULLREQUESTREVIEWTHREAD  = "pullrequestreviewthread"
	PUSH                     = "push"
	RELEASE                  = "release"
	SPONSORSHIP              = "sponsorship"
	WATCH                    = "watch"
)

const (
	COMMITCOMMENT_BEAUTIFY            = "Commit comment"
	CREATE_BEAUTIFY                   = "Create"
	DELETE_BEAUTIFY                   = "Delete"
	FORK_BEAUTIFY                     = "Fork"
	GOLLUM_BEAUTIFY                   = "Gollum"
	ISSUECOMMENT_BEAUTIFY             = "Issue comment"
	ISSUES_BEAUTIFY                   = "Issues"
	MEMBER_BEAUTIFY                   = "Member"
	PUBLIC_BEAUTIFY                   = "Public"
	PULLREQUEST_BEAUTIFY              = "Pull request"
	PULLREQUESTREVIEW_BEAUTIFY        = "Pull request review"
	PULLREQUESTREVIEWCOMMENT_BEAUTIFY = "Pull request review comment"
	PULLREQUESTREVIEWTHREAD_BEAUTIFY  = "Pull request review thread"
	PUSH_BEAUTIFY                     = "Push"
	RELEASE_BEAUTIFY                  = "Release"
	SPONSORSHIP_BEAUTIFY              = "Sponsorship"
	WATCH_BEAUTIFY                    = "Watch"
)

func (eventName GithubEvent) ToString() string {
	switch eventName {
	case COMMITCOMMENT:
		return COMMITCOMMENT_BEAUTIFY
	case CREATE:
		return CREATE_BEAUTIFY
	case DELETE:
		return DELETE_BEAUTIFY
	case FORK:
		return FORK_BEAUTIFY
	case GOLLUM:
		return GOLLUM_BEAUTIFY
	case ISSUECOMMENT:
		return ISSUECOMMENT_BEAUTIFY
	case ISSUES:
		return ISSUES_BEAUTIFY
	case MEMBER:
		return MEMBER_BEAUTIFY
	case PUBLIC:
		return PUBLIC_BEAUTIFY
	case PULLREQUEST:
		return PULLREQUEST_BEAUTIFY
	case PULLREQUESTREVIEW:
		return PULLREQUESTREVIEW_BEAUTIFY
	case PULLREQUESTREVIEWCOMMENT:
		return PULLREQUESTREVIEWCOMMENT_BEAUTIFY
	case PULLREQUESTREVIEWTHREAD:
		return PULLREQUESTREVIEWTHREAD_BEAUTIFY
	case PUSH:
		return PUSH_BEAUTIFY
	case RELEASE:
		return RELEASE_BEAUTIFY
	case SPONSORSHIP:
		return SPONSORSHIP_BEAUTIFY
	case WATCH:
		return WATCH_BEAUTIFY
	}
	return ""
}
