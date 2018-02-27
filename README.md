Golang Cron
=============================

![clock](./clock.jpeg "clock")

<!-- TOC -->

- [Cron Expression](#cron-expression)
- [Install](#install)
- [Usage](#usage)
- [Reference](#reference)
- [License](#license)

<!-- /TOC -->

## Cron Expression

> The reference documentation is found at https://en.wikipedia.org/wiki/Cron#CRON_expression

    Field name     Mandatory?   Allowed values    Allowed special characters
    ----------     ----------   --------------    --------------------------
    Seconds        No           0-59              * / , -
    Minutes        Yes          0-59              * / , -
    Hours          Yes          0-23              * / , -
    Day of month   Yes          1-31              * / , - L W
    Month          Yes          1-12 or JAN-DEC   * / , -
    Day of week    Yes          0-6 or SUN-SAT    * / , - L #
    Year           No           1970–2099         * / , -

## Install

```sh
    go get -u github.com/flyaways/cron
```

## Usage

```go
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
```

## Reference

* [github.com/gorhill/cronexpr](https://github.com/gorhill/cronexpr)
* [github.com/robfig/cron](https://github.com/robfig/cron)

## License

* [Apache License 2.0](http://www.apache.org/licenses/LICENSE-2.0)