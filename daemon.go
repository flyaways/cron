package cron

import (
	"log"
	"runtime"
	"sort"
	"time"

	"github.com/gorhill/cronexpr"
)

type Daemon struct {
	schedules Schedules
	snapshot  chan Schedules

	running bool
	stop    chan struct{}

	add    chan *Schedule
	del    chan *Schedule
	update chan *Schedule

	now      time.Time
	location *time.Location
}

func (c *Daemon) Add(j Job) error {
	i := c.schedules.Pos(j.ID())
	if i != -1 {
		return ErrJobIDExists
	}

	expression, err := cronexpr.Parse(j.Cronexpr())
	if err != nil {
		return err
	}

	s := &Schedule{
		Expression: expression,
		Job:        j,
		Next:       expression.Next(c.now),
	}

	if !c.running {
		c.schedules = append(c.schedules, s)
		return nil
	}

	c.add <- s

	return nil
}

func (c *Daemon) Del(j Job) error {
	i := c.schedules.Pos(j.ID())
	if i == -1 {
		return ErrJobIDNotExists
	}

	if !c.running {
		c.schedules = append(c.schedules[:i], c.schedules[i+1:]...)
		return nil
	}

	c.del <- &Schedule{
		Job: j,
	}

	return nil
}

func (c *Daemon) Update(j Job) error {
	i := c.schedules.Pos(j.ID())
	if i == -1 {
		return ErrJobIDNotExists
	}

	expression, err := cronexpr.Parse(j.Cronexpr())
	if err != nil {
		return err
	}

	s := &Schedule{
		Expression: expression,
		Job:        j,
		Next:       expression.Next(c.now),
	}

	if !c.running {
		c.schedules = append(c.schedules[:i], c.schedules[i+1:]...)
		c.schedules = append(c.schedules, s)
		return nil
	}

	c.update <- s

	return nil
}

func (c *Daemon) Start() {
	if c.running {
		return
	}

	go c.run()
}

func (c *Daemon) Stop() {
	if !c.running {
		return
	}

	c.stop <- struct{}{}
}

func (c *Daemon) SnapShot() Schedules {
	if c.running {
		c.snapshot <- nil
		ss := <-c.snapshot
		return ss
	}

	return c.snap()
}

func (c *Daemon) snap() Schedules {
	ss := Schedules{}
	for _, s := range c.schedules {
		ss = append(ss, &Schedule{
			Expression: s.Expression,
			Next:       s.Next,
			Prev:       s.Prev,
			Job:        s.Job,
		})
	}
	return ss
}

func (c *Daemon) execute() {
	c.now = c.now.In(c.location)

	for _, s := range c.schedules {
		if s.Next.After(c.now) || s.Next.IsZero() {
			break
		}

		go c.runWithRecovery(s.Job)

		s.Prev = s.Next
		s.Next = s.Expression.Next(c.now)
	}
}

func (c *Daemon) runWithRecovery(j Job) {
	defer func() {
		if r := recover(); r != nil {
			const size = 64 << 10
			buf := make([]byte, size)
			buf = buf[:runtime.Stack(buf, false)]
			log.Printf("cron: panic running job: %v\n%s", r, buf)
		}
	}()
	j.Run()
}

func (c *Daemon) reset(timer *time.Timer, now *time.Time) {
	sort.Sort(Schedules(c.schedules))

	effective := now.AddDate(10, 0, 0)

	if len(c.schedules) != 0 && !c.schedules[0].Next.IsZero() {
		effective = c.schedules[0].Next
	}

	timer.Reset(effective.Sub(*now))
}

func (c *Daemon) run() {
	c.running = true
	c.now = time.Now().In(c.location)
	timer := time.NewTimer(time.Minute)

	for {
		c.reset(timer, &c.now)

		select {
		case c.now = <-timer.C:
			c.execute()

		case s := <-c.add:
			timer.Stop()

			c.schedules = append(c.schedules, s)

		case s := <-c.update:
			timer.Stop()

			i := c.schedules.Pos(s.ID())
			if i == -1 {
				continue
			}

			c.schedules = append(c.schedules[:i], c.schedules[i+1:]...)
			c.schedules = append(c.schedules, s)

		case s := <-c.del:
			timer.Stop()

			i := c.schedules.Pos(s.ID())
			if i == -1 {
				continue
			}

			c.schedules = append(c.schedules[:i], c.schedules[i+1:]...)

		case <-c.snapshot:
			c.snapshot <- c.snap()
			continue

		case <-c.stop:
			c.running = false
			timer.Stop()
			return
		}

		c.now = time.Now().In(c.location)
	}
}
