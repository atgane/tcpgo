package ds

import (
	"errors"
	"sync"
	"sync/atomic"
)

type Eventloop[T any] struct {
	queue         chan T
	handler       func(T)
	dispatchCount int
	closeCh       chan struct{}
	closed        atomic.Bool
	state         atomic.Int32
}

func NewEventloop[T any](dispatchCount, queueSize int, handler func(T)) *Eventloop[T] {
	e := &Eventloop[T]{
		queue:         make(chan T, queueSize),
		handler:       handler,
		dispatchCount: dispatchCount,
		closeCh:       make(chan struct{}),
	}
	return e
}

func (e *Eventloop[T]) Run() {
	wg := sync.WaitGroup{}
	for range e.dispatchCount {
		wg.Add(1)
		go func() {
			defer wg.Done()
			e.dispatch()
		}()
	}
	wg.Wait()
}

// ⚠️
func (e *Eventloop[T]) Send(event T) error {
	select {
	case <-e.closeCh:
		return ErrAlreadyClosedLoop
	case e.queue <- event:
		return nil
	}
}

// ⚠️
func (e *Eventloop[T]) Close() {
	if !e.closed.CompareAndSwap(false, true) {
		return
	}

	close(e.closeCh)
}

// ⚠️
func (e *Eventloop[T]) ForceClose() {
	if !e.closed.CompareAndSwap(false, true) {
		return
	}

	close(e.closeCh)
	close(e.queue)
}

func (e *Eventloop[T]) dispatch() {
	for {
		select {
		case event, ok := <-e.queue:
			if !ok {
				return
			}
			e.handler(event)
		case <-e.closeCh:
			return
		}
	}
}

var ErrAlreadyClosedLoop = errors.New("already closed loop")
