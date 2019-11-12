## Workshop Introduction

This repository contains some exercises for the introductory level of the 
[Friends of Go](https://friendsofgo.tech/formacion-empresas/) workshops.

Feel free to open a new PR with your solutions.

### Instructions

The exercises' level is incremental, so to avoid confusion, just do the exercises in a sorted way:

```
exercise-1
exercise-2
exercise-3
exercise-4
```

Each exercise has a its own `README.md` file with the specific instructions and some notes.

Feel free to open a new PR to contribute with us (typos, missing info, etc).

### Requirements

This repository is tested under Go 1.12+ and using [Go Modules](https://blog.golang.org/using-go-modules).

The dependencies needed are:

- [`benchcmp`](golang.org/x/tools/cmd/benchcmp`)

    ```shell script
    $ go get -u golang.org/x/tools/cmd/benchcmp
    ```

- [`graphviz`](https://graphviz.gitlab.io/download/)

    See the installation details for each platform on the link.

### Exercises

** 1. Sorting Numbers**

The first exercise basically consists on implementing a function to sort numbers:

```go
func Nums(unsortedNums []int) []int {
	// TODO: Sort numbers
	return unsortedNums
}
``` 

** 2. Stupid Cache**

The second exercise let you do your own implementation of an stupid cache to satisfy the given cache interface:

```go
type StupidCache interface {
	Get(key string) []byte
	Set(key string, data []byte)
}
```

** 3. Benchmarking & Profiling**

The third exercise is focused on benchmarking and profiling the `Marshal` and `Unmarshal` functions of 
the `encoding/json` package.

** 4. Concurrent Quiz**

This fourth (and the last) exercise consists on adding concurrency on a quiz simple application in order
to extend it and make it more funny.

It has has some "bonus tracks" to make it even funnier.