# goawaysignals

Go OS signals handler

## About

`goawaysignals` handles `SIGHUP, SIGINT, SIGTERM, SIGQUIT` and then waits for the registered async functions to return before handing back the control to the callee. This helps you to achieve graceful shutdown on many kinds of services.

## How to use

```go
package main

import (
	"fmt"

	"github.com/eventials/goawaysignals"
)

func main() {
	// fake services that need graceful shutdown.
	producer := NewExampleProducer()
	consumer := NewExampleConsumer()

	goawaysignals.Async(func() { <-producer.NotifyClose() })
	goawaysignals.Async(func() { <-consumer.NotifyClose() })

	goawaysignals.AfterSignal(func() {
		producer.Close()
		consumer.Close()
	})

	goawaysignals.Wait()
	fmt.Println("Exiting...")
}
```

If you need to exit before an OS signal, you can call `goawaysignals.Close()` and still provide graceful shutdown to your services.