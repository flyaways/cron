package cron

import "time"

func New() *Daemon {
	return NewWithLocation(time.Now().Location())
}

func NewWithLocation(location *time.Location) *Daemon {
	return &Daemon{
		schedules: nil,
		add:       make(chan *Schedule),
		del:       make(chan *Schedule),
		update:    make(chan *Schedule),
		stop:      make(chan struct{}),
		snapshot:  make(chan Schedules),
		running:   false,
		now:       time.Now().In(location),
		location:  location,
	}
}
