package main

import (
	"fmt"
	"log"
	"time"

	"github.com/Quantaly/shunt"
)

const taskCount = 30

func main() {
	tasks := make([]shunt.Task[int], taskCount)
	for i := 0; i < taskCount; i++ {
		i := i // give the closure its own i
		tasks[i] = shunt.Do(func() (int, error) {
			time.Sleep(time.Second) // some long computation or I/O work...
			// panic("this panic will appear in the initial goroutine")
			return i, nil
		})
	}

	sum := 0
	for i := 0; i < taskCount; i++ {
		result, err := tasks[i].Join()
		if err != nil {
			log.Fatalln(err)
		}

		sum += result
	}

	fmt.Println(sum)
}
