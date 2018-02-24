package cron

import "errors"

var (
	ErrJobIDNotExists = errors.New("job id not exits")
	ErrJobIDExists    = errors.New("job id exists")
)
