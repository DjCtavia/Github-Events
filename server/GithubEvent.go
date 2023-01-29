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

func (eventName GithubEvent) ToString() string {
	switch eventName {
	case COMMITCOMMENT:
		return "Commit comment"
	case CREATE:
		return "Create"
	case DELETE:
		return "Delete"
	case FORK:
		return "Fork"
	case GOLLUM:
		return "Gollum"
	case ISSUECOMMENT:
		return "Issue comment"
	case ISSUES:
		return "Issues"
	case MEMBER:
		return "Member"
	case PUBLIC:
		return "Public"
	case PULLREQUEST:
		return "Pull request"
	case PULLREQUESTREVIEW:
		return "Pull request review"
	case PULLREQUESTREVIEWCOMMENT:
		return "Pull request review comment"
	case PULLREQUESTREVIEWTHREAD:
		return "Pull request review thread"
	case PUSH:
		return "Push"
	case RELEASE:
		return "Release"
	case SPONSORSHIP:
		return "Sponsorship"
	case WATCH:
		return "Watch"
	}
	return ""
}
