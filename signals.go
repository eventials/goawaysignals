package goawaysignals

import (
	"os"
	"os/signal"
	"sync"
	"syscall"
)

var (
	wg           sync.WaitGroup
	after        func()
	closeChannel = make(chan bool, 1)
)

// Async runs the given function in a goroutine.
//
// Before invoking the given function, adds an unit to the WaitGroup.
// After the execution of the given function, mark this unit of work as done.
func Async(funct func()) {
	go func() {
		wg.Add(1)
		funct()
		wg.Done()
	}()
}

// AfterSignal register the given function to run after the OS signal.
//
// Often used to initiate your services shutdown.
func AfterSignal(funct func()) {
	after = funct
}

// Close will call your AfterSignal function right away (without waiting the OS signal)
// and waits the end of the WaitGroup.
func Close() {
	closeChannel <- true
}

// Wait for the OS signal then execute your AfterSignal function. After that,
// waits for the end of the WaitGroup.
func Wait() {
	sigChannel := make(chan os.Signal, 1)
	signal.Notify(sigChannel, syscall.SIGHUP, syscall.SIGINT, syscall.SIGTERM, syscall.SIGQUIT)

	select {
	case <-sigChannel:
	case <-closeChannel:
	}

	if after != nil {
		after()
	}

	wg.Wait()
}
