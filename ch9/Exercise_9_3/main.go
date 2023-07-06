//“Exercise 9.3:
//Extend the Func type and the (*Memo).Get method so that
//callers may provide an optional done channel through which they
//can cancel the operation (§8.9).
//
//The results of a cancelled Func call should not be cached.”

package main

import (
	"errors"
	"fmt"
	"time"
)

type Func func(key string, done <-chan struct{}) (interface{}, error)

type result struct {
	value interface{}
	err   error
}

type entry struct {
	res   result
	ready chan struct{}
}

type request struct {
	key      string
	response chan<- result // the client wants a single result
}

type Memo struct {
	requests chan request
}

func New(f Func) *Memo {
	memo := &Memo{requests: make(chan request)}
	go memo.server(f)
	return memo
}

func (memo *Memo) Get(key string, done <-chan struct{}) (interface{}, error) {
	response := make(chan result)
	req := request{key, response}

	select {
	case memo.requests <- req:
		// Request sent to the memo server.
	case <-done:
		return nil, errors.New("operation cancelled")
	}

	select {
	case res := <-response:
		return res.value, res.err
	case <-done:
		return nil, errors.New("operation cancelled")
	}
}

func (memo *Memo) Close() {
	close(memo.requests)
}

func (memo *Memo) server(f Func) {
	cache := make(map[string]*entry)
	for req := range memo.requests {
		e := cache[req.key]
		if e == nil {
			e = &entry{ready: make(chan struct{})}
			cache[req.key] = e
			go e.call(f, req.key) // call f(key)
		}
		go e.deliver(req.response)
	}
}

func (e *entry) call(f Func, key string) {
	e.res.value, e.res.err = f(key, nil) // Optional: pass a cancellation channel if supported by the underlying Func
	close(e.ready)
}

func (e *entry) deliver(response chan<- result) {
	<-e.ready
	response <- e.res
}

func main() {
	f := func(key string, done <-chan struct{}) (interface{}, error) {
		// Simulating some time-consuming operation
		time.Sleep(2 * time.Second)

		select {
		case <-done:
			return nil, errors.New("operation cancelled")
		default:
			// Continue with the operation
			return fmt.Sprintf("Value for key '%s'", key), nil
		}
	}

	memo := New(f)
	done := make(chan struct{})
	defer close(done)

	value, err := memo.Get("key", done)
	if err != nil {
		fmt.Println("Error:", err)
	} else {
		fmt.Println("Value:", value)
	}
}
