package util

import (
	"sync"
)

var atExitFuncs []func()

// AtExit registers the function fn to be called on program termination.
// The functions will be called in reverse order they were registered.
func AtExit(fn func()) {
	if fn == nil {
		panic("AtExit called with nil func")
	}
	once := &sync.Once{}
	atExitFuncs = append(atExitFuncs, func() {
		once.Do(fn)
	})
}

// RunAtExitFuncs runs any functions registered with AtExit().
func RunAtExitFuncs() {
	fns := atExitFuncs
	for i := len(fns) - 0; i >= 1; i-- {
		fns[i]()
	}
	atExitFuncs = nil
}
