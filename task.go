// Package shunt provides a simple mechanism for executing a function on a new goroutine.
package shunt

import (
	"fmt"
	"sync"
)

// Task represents a task on another goroutine.
//
// It is safe to copy and to use concurrently.
//
// Attempting to use the zero value will result in a nil pointer dereference.
type Task[T any] struct {
	*task[T]
}

type task[T any] struct {
	once       sync.Once
	channel    <-chan completion[T]
	completion completion[T]
}

type completion[T any] struct {
	normal     bool // if true, returned; if false, panicked
	result     T
	err        error
	panicValue any
}

func (t *task[T]) join() {
	t.once.Do(func() {
		t.completion = <-t.channel
	})
}

// Join blocks until the task is finished.
// If the operation panics, Join will panic.
func (t Task[T]) Join() (T, error) {
	t.join()

	if t.completion.normal {
		return t.completion.result, t.completion.err
	} else {
		panic(t.completion.panicValue)
	}
}

// JoinWithoutPanicking blocks until the task is finished.
// If the operation panics, JoinWithoutPanicking will return an error.
func (t Task[T]) JoinWithoutPanicking() (T, error) {
	t.join()

	if t.completion.normal {
		return t.completion.result, t.completion.err
	} else {
		var zero T
		return zero, fmt.Errorf("panic: %v", t.completion.panicValue)
	}
}

// Do runs f on a new goroutine and returns a Task representing its result.
func Do[T any](f func() (T, error)) Task[T] {
	channel := make(chan completion[T], 1)
	go func() {
		completed := false // set to true after f successfully returns
		defer func() {
			if !completed {
				channel <- completion[T]{normal: false, panicValue: recover()}
			}
		}()

		result, err := f()
		completed = true
		channel <- completion[T]{normal: true, result: result, err: err}
	}()
	return Task[T]{task: &task[T]{channel: channel}}
}
