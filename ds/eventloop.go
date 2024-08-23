package ds

import "sync/atomic"

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

func (e *Eventloop[T]) Run()
func (e *Eventloop[T]) Send(event T) error
func (e *Eventloop[T]) Close()
func (e *Eventloop[T]) ForceClose()
