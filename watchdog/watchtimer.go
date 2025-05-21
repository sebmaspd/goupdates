package main

import (
	"fmt"
	"math/rand"
	"sync"
	"time"
)

type Watchdog struct {
	timeout     time.Duration
	resetChan   chan bool
	stopChan    chan struct{}
	resetAction func()
}

func NewWatchdog(timeout time.Duration, resetAction func()) *Watchdog {
	return &Watchdog{
		timeout:     timeout,
		resetChan:   make(chan bool),
		stopChan:    make(chan struct{}),
		resetAction: resetAction,
	}
}

func (w *Watchdog) Start() {
	go func() {
		timer := time.NewTimer(w.timeout)

		for {
			select {
			case <-w.resetChan:
				if !timer.Stop() {
					<-timer.C
				}
				timer.Reset(w.timeout)
				fmt.Println("Watchdog kicked")

			case <-timer.C:
				fmt.Println("Watchdog timeout! System unresponsive.")
				w.resetAction()
				timer.Reset(w.timeout)

			case <-w.stopChan:
				fmt.Println("Watchdog stopped.")
				return
			}
		}
	}()
}

func (w *Watchdog) Kick() {
	w.resetChan <- true
}

func (w *Watchdog) Stop() {
	close(w.stopChan)
}

func main() {
	var wg sync.WaitGroup

	resetAction := func() {
		fmt.Println(">>> System reset triggered by watchdog.")
	}

	wd := NewWatchdog(3*time.Second, resetAction)
	wd.Start()

	wg.Add(1)
	go func() {
		defer wg.Done()
		for i := 1; i <= 10; i++ {
			// Simulate random system behavior
			delay := time.Duration(rand.Intn(5)) * time.Second
			fmt.Printf("Iteration %d: sleeping for %v\n", i, delay)
			time.Sleep(delay)

			if delay < 3*time.Second {
				wd.Kick()
			} else {
				fmt.Println("Skipped watchdog kick — simulating hang.")
			}
		}
		wd.Stop()
	}()

	wg.Wait()
}
