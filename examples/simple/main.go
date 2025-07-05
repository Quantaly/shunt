package main

import (
	"fmt"
	"log"
	"time"

	"github.com/Quantaly/shunt/v2"
)

const taskCount = 30

func main() {
	tasks := make([]shunt.Task[int], taskCount)
	for i := 0; i < taskCount; i++ {
		i := i // give the closure its own i
		tasks[i] = shunt.Do(func() (int, error) {
			time.Sleep(time.Second) // some long computation or I/O work...
			return i, nil
		})
	}

	sum := 0
	for _, task := range tasks {
		result, err := task.Join()
		if err != nil {
			log.Fatalln(err)
		}

		sum += result
	}

	fmt.Println(sum)
}
