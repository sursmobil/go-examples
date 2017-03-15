package multithreading

import (
	"sync"
	"time"
)

// WaiterWork is a type representing unit of work given to waiter
type WaiterWork func(WaiterWorkResult)

/*
WaiterWorkResult is a channel were result from WaiterWork should be passed.

If value send to the channel is an error it will be returned to the client
as error with nil value
*/
type WaiterWorkResult chan<- interface{}

/*
Waiter is example of running a jobs in background.

It has following capabilities:
* Start - start job in background
* Wait - wait for started job to finish and return its result or error
* Cancel - cancel any routine awaiting result from this waiter
*/
type Waiter struct {
	fn     WaiterWork
	result chan interface{}
	cancel chan bool
	wg     *sync.WaitGroup
}

// NewWaiter creates pointer to new Waiter which will serve given WaiterWork
func NewWaiter(fn WaiterWork) *Waiter {
	return &Waiter{
		fn:     fn,
		wg:     &sync.WaitGroup{},
		cancel: make(chan bool),
		result: make(chan interface{}), // gotcha! comma required
	}
}

/*
Start will make Waiter execute given WaiterWork.

Work result can be obtained using Waiter.Wait function
*/
func (w *Waiter) Start() {
	go func() {
		defer w.wg.Done()
		w.fn(w.result)
	}() // gotcha! function call
	w.wg.Add(1)
}

// Wait for Waiter to finish WaiterWork
func (w *Waiter) Wait(timeout time.Duration) (interface{}, error) {
	timeoutCh := time.After(timeout)
	select {
	case result := <-w.result:
		switch v := result.(type) {
		case error:
			return nil, v
		default:
			w.wg.Wait()
			return result, nil
		}
	case <-w.cancel:
		return nil, ErrWaiterCancel
	case <-timeoutCh:
		return nil, ErrWaiterTimeout
	}
}

// Cancel someone awaiting for Waiter result
func (w *Waiter) Cancel() {
	w.cancel <- true
}
