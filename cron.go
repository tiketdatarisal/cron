package cron

import (
	"fmt"
	"github.com/robfig/cron/v3"
)

type Cron struct {
	cron   *cron.Cron
	jobMap map[string]cron.EntryID
}

// NewCron return a new cron.
func NewCron() *Cron {
	c := cron.New()
	c.Start()
	return &Cron{cron: c, jobMap: map[string]cron.EntryID{}}
}

// Close will close cron.
func (c *Cron) Close() {
	if c.cron != nil {
		c.cron.Stop()
	}
}

// Engine return cron engine.
func (c *Cron) Engine() *cron.Cron { return c.cron }

// AddUniqueJob add unique job to cron.
func (c *Cron) AddUniqueJob(name, spec string, cmd func()) error {
	const msg = "could not add new job to cron: %w"

	if _, exists := c.jobMap[name]; exists {
		return fmt.Errorf(msg, ErrDuplicateJobInstance)
	}

	id, err := c.cron.AddFunc(spec, cmd)
	if err != nil {
		return fmt.Errorf(msg, err)
	}

	c.jobMap[name] = id
	return nil
}

// TerminateJob terminate job.
func (c *Cron) TerminateJob(name string) error {
	const msg = "could not terminate job: %w"

	if _, exists := c.jobMap[name]; !exists {
		return fmt.Errorf(msg, ErrJobInstanceNotFound)
	}

	c.cron.Remove(c.jobMap[name])
	delete(c.jobMap, name)
	return nil
}

// UpdateUniqueJob update unique job.
func (c *Cron) UpdateUniqueJob(name, spec string, cmd func()) error {
	if err := c.TerminateJob(name); err != nil {
		return err
	}

	return c.AddUniqueJob(name, spec, cmd)
}
