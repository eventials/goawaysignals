package main

import (
	"fmt"

	"github.com/eventials/goawaysignals"
)

func main() {
	ch1 := make(chan bool)
	ch2 := make(chan bool)

	goawaysignals.Async(func() { <-ch1 })
	goawaysignals.Async(func() { <-ch2 })

	goawaysignals.AfterSignal(func() {
		ch1 <- true
		ch2 <- true
		fmt.Println("Channels notified")
	})

	goawaysignals.Wait()
	fmt.Println("Exiting...")
}
