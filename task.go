// Package shunt provides a simple mechanism for executing a function on a new goroutine.
package shunt

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
	result T
	err    error
}

// Join blocks until the task is finished.
// If the operation panics, Join will panic.
func (t Task[T]) Join() (T, error) {
	<-t.done
	return t.completion.result, t.completion.err
}

// Do runs f on a new goroutine and returns a Task representing its result.
func Do[T any](f func() (T, error)) Task[T] {
	done := make(chan unit)
	completion := new(completion[T])
	go func() {
		completion.result, completion.err = f()
		close(done)
	}()
	return Task[T]{done: done, completion: completion}
}
