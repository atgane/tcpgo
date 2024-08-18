package ds

import (
	"testing"

	"github.com/stretchr/testify/require"
	"go.uber.org/goleak"
)

func TestEventloopGoroutineLeakClose(t *testing.T) {
	defer goleak.VerifyNone(t)

	e := NewEventloop(2, 1024, func(int) {})
	go e.Run()
	e.Close()
}

func TestEventloopGoroutineLeakForceClose(t *testing.T) {
	defer goleak.VerifyNone(t)

	e := NewEventloop(2, 1024, func(int) {})
	go e.Run()
	e.ForceClose()
}

func TestEventloopClose(t *testing.T) {
	e := NewEventloop(2, 1024, func(int) {})
	go e.Run()
	e.Close()
	err := e.Send(1)
	require.Error(t, err)
}
func TestEventloopForceClose(t *testing.T) {
	e := NewEventloop(2, 1024, func(int) {})
	go e.Run()
	e.ForceClose()
	err := e.Send(1)
	require.Error(t, err)
}
