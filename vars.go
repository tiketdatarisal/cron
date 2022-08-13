package cron

import (
	"errors"
	"sync"
)

var (
	c     *Cron
	mutex sync.Mutex

	ErrDuplicateJobInstance = errors.New("duplicate job instance")
	ErrJobInstanceNotFound  = errors.New("job not found")
)
