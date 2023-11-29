package shunt

import (
	"errors"
	"testing"
)

func TestSuccess(t *testing.T) {
	task := Do(func() (int, error) {
		return 9 + 10, nil
	})
	result, err := task.Join()
	if result != 19 {
		t.Errorf("result was %v, want 19", result)
	}
	if err != nil {
		t.Errorf("err was %v, want nil", err)
	}
}

func TestError(t *testing.T) {
	task := Do(func() (int, error) {
		return 0, errors.New("hello")
	})
	_, err := task.Join()
	if err == nil || err.Error() != "hello" {
		t.Errorf("err was %v, want hello", err)
	}
}

func TestPanic(t *testing.T) {
	task := Do(func() (int, error) {
		panic("goodbye")
	})
	defer func() {
		v := recover()
		if v != "goodbye" {
			t.Errorf("panic value was %v, want goodbye", v)
		}
	}()
	task.Join()
	t.Error("task.Join() did not panic")
}

func TestNilPanic(t *testing.T) {
	task := Do(func() (int, error) {
		panic(nil)
	})
	defer func() {
		v := recover()
		if v != nil {
			t.Errorf("panic value was %v, want nil", v)
		}
	}()
	task.Join()
	t.Error("task.Join() did not panic")
}

func TestJoinWithoutPanicking(t *testing.T) {
	task := Do(func() (int, error) {
		panic("oops")
	})
	_, err := task.JoinWithoutPanicking()
	if err == nil || err.Error() != "panic: oops" {
		t.Errorf("err was %v, want panic: oops", err)
	}
}

func TestConcurrentJoin(t *testing.T) {
	const expected = 42
	waitChan := make(chan unit)
	task := Do(func() (int, error) {
		<-waitChan
		return expected, nil
	})
	otherTasks := make([]Task[unit], 0, 5)
	for i := 0; i < 5; i++ {
		otherTasks = append(otherTasks, Do(func() (unit, error) {
			result, err := task.Join()
			if result != expected {
				t.Errorf("result was %v, want %v", result, expected)
			}
			if err != nil {
				t.Errorf("err was %v, want nil", err)
			}
			return unit{}, nil
		}))
	}
	close(waitChan)
	for _, t := range otherTasks {
		t.Join()
	}
}
