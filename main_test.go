package main_test

import (
	"runtime"
	"sync"
	"testing"

	"github.com/stretchr/testify/require"
)

func TestGoroutineProcs1(t *testing.T) {
	runtime.GOMAXPROCS(1)
	n := 5
	correct := []int{4, 0, 1, 2, 3}

	j := 0
	wg := sync.WaitGroup{}
	for i := range n {
		wg.Add(1)
		go func(i int) {
			defer wg.Done()
			require.Equal(t, correct[j], i)
			j++
		}(i)
	}

	wg.Wait()
}
