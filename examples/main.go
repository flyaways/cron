package main

import (
	"fmt"
	"time"

	"github.com/flyaways/cron"
)

type task struct {
	cron.Job
	id   string
	expr string
}

func (t *task) Run() {
	fmt.Println(time.Now().Format(time.RFC3339Nano))
}

func (t *task) ID() string {
	return t.id
}

func (t *task) Cronexpr() string {
	return t.expr
}

func main() {
	c := cron.New()
	c.Start()

	c.Add(&task{
		id:   "AF0783D4-D16A-4CAE-8866-4E0648AE3651",
		expr: "0/2 * * * * ?",
	})
	time.Sleep(time.Second * 10)

	c.Update(&task{
		id:   "AF0783D4-D16A-4CAE-8866-4E0648AE3651",
		expr: "0/5 * * * * ?",
	})
	time.Sleep(time.Second * 10)

	c.Del(&task{
		id: "AF0783D4-D16A-4CAE-8866-4E0648AE3651",
	})
	time.Sleep(time.Second * 10)

	c.Stop()
}
