package cron

type Job interface {
	Run()
	ID() string
	Cronexpr() string
}
