package cron

import (
	"time"

	"github.com/gorhill/cronexpr"
)

type Schedule struct {
	Expression *cronexpr.Expression
	Next       time.Time
	Prev       time.Time
	Job
}

type Schedules []*Schedule

func (s Schedules) Len() int      { return len(s) }
func (s Schedules) Swap(i, j int) { s[i], s[j] = s[j], s[i] }
func (s Schedules) Less(i, j int) bool {
	if s[i].Next.IsZero() {
		return false
	}
	if s[j].Next.IsZero() {
		return true
	}
	return s[i].Next.Before(s[j].Next)
}

func (s Schedules) Pos(id string) int {
	for i, e := range s {
		if e.ID() == id {
			return i
		}
	}
	return -1
}
