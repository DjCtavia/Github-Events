package observer

import (
	"djctavia/server"
	"sync"
)

type Subject interface {
	register(observer Observer)
	deregister(observer Observer)
	notify(event *server.GithubEvent)
}

type GithubSubject struct {
	Observers []Observer
	mu        sync.Mutex
}

var GlobalSubject *GithubSubject = &GithubSubject{}

func (gs *GithubSubject) Register(observer Observer) {
	gs.mu.Lock()
	defer gs.mu.Unlock()

	gs.Observers = append(gs.Observers, observer)
}

func (gs *GithubSubject) Deregister(observer Observer) {
	gs.mu.Lock()
	defer gs.mu.Unlock()

	for i, o := range gs.Observers {
		if o == observer {
			gs.Observers = append(gs.Observers[:i], gs.Observers[i+1:]...)
			return
		}
	}
}

func (gs *GithubSubject) Notify(event *server.GithubEvent) {
	gs.mu.Lock()
	defer gs.mu.Unlock()

	for _, observer := range gs.Observers {
		observer.update(event)
	}
}
