# Shunt

Shunt is a simple Go library for easily executing a function on a new goroutine and accessing the result.

## Installation

```sh
go get github.com/Quantaly/shunt/v2
```

## Usage

```go
import "github.com/Quantaly/shunt/v2"

task := shunt.Do(func() (int, error) {
	return 42, nil
})
answer, err := task.Join()
```

Also see the [examples](./examples).

## Contributing

Pull requests are welcome. For major changes, please open an issue first to discuss what you would like to change.

Please make sure to update tests as appropriate.

## License

This is free and unencumbered software released into the public domain. See the [Unlicense](https://unlicense.org/) for full details.
