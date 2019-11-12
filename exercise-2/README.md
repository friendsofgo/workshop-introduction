## Stupid Cache

This exercise consists on implementing your own stupid cache implementation satisfying the already defined interface.

```go
type StupidCache interface {
	Get(key string) []byte
	Set(key string, data []byte)
}
```
*NOTE: Be careful with the way the interfaces work on Go.*

The code provided already contains the boilerplate needed to run the program (see `main.go`) which basically
inserts some values into the given cache and then returns the cached value (and the expected one) for each key.

*NOTE: Be careful with the usage (or not) of pointers.*

### How to run the program

Just use:

```shell script
go run main.go
```

And check the output to see if your cache is working properly.

Otherwise you need to do some work to fix your implementation.

### Tests

**Please, write some tests for your stupid cache implementation.**

There is an already created file for this purpose (see `cache/cache_test.go`).

In order to run the tests use:

```shell script
$ go test ./...
```