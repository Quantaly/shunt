// Package shunt provides a simple mechanism for executing a function on a new goroutine.
package shunt

import (
	"fmt"
)

type unit struct{}

// Task represents a task on another goroutine.
//
// It is safe to copy and to use concurrently.
//
// The zero value is not safe to use.
// In the current implementation, using the zero value will cause the program to receive on a nil channel.
type Task[T any] struct {
	done       <-chan unit
	completion *completion[T]
}

type completion[T any] struct {
	normal     bool // if true, returned; if false, panicked
	result     T
	err        error
	panicValue any
}

// Join blocks until the task is finished.
// If the operation panics, Join will panic.
func (t Task[T]) Join() (T, error) {
	<-t.done

	if t.completion.normal {
		return t.completion.result, t.completion.err
	} else {
		panic(t.completion.panicValue)
	}
}

// JoinWithoutPanicking blocks until the task is finished.
// If the operation panics, JoinWithoutPanicking will return an error.
func (t Task[T]) JoinWithoutPanicking() (T, error) {
	<-t.done

	if t.completion.normal {
		return t.completion.result, t.completion.err
	} else {
		var zero T
		return zero, fmt.Errorf("panic: %v", t.completion.panicValue)
	}
}

// Do runs f on a new goroutine and returns a Task representing its result.
func Do[T any](f func() (T, error)) Task[T] {
	done := make(chan unit)
	completion := new(completion[T])
	go func() {
		defer func() {
			if !completion.normal {
				completion.panicValue = recover()
				close(done)
			}
		}()

		completion.normal = false
		completion.result, completion.err = f()
		completion.normal = true
		close(done)
	}()
	return Task[T]{done: done, completion: completion}
}
