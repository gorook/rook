package main

import (
	"log"
	"time"

	"github.com/rjeczalik/notify"
)

const tick = 200 * time.Millisecond
const eventBuffer = 42

type watcher struct {
	c chan notify.EventInfo
}

func newWatcher(path string, handler func(notify.EventInfo)) *watcher {
	c := make(chan notify.EventInfo, eventBuffer)
	err := notify.Watch(path, c, notify.All)
	if err != nil {
		log.Fatalf("unable to run watcher on %s: %v", path, err)
	}
	go func() {
		events := make([]notify.EventInfo, 0)
		ticker := time.NewTicker(tick)
		defer ticker.Stop()
		for {
			select {
			case event := <-c:
				events = append(events, event)
			case <-ticker.C:
				if len(events) > 0 {
					event := events[0]
					events = events[:0]
					go handler(event)
				}
			}
		}
	}()
	return &watcher{c: c}
}
