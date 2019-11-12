## Sorting Numbers

This basic exercise basically consists on implementing a function that receives an unsorted slice of integers and returns
it sorted.

```go
func Nums(unsortedNums []int) []int {
	// TODO: Sort numbers
	return unsortedNums
}
```

The code provided already contains the boilerplate needed to run the program (see `main.go`) which basically
reads the numbers from the `numbers.txt`, calls the `sort.Nums` function and prints its output.

### How to run the program

Just use:

```shell script
go run main.go
```

And check that it prints the numbers properly sorted.

Otherwise you need to do some work to fix the function.

### Tests

In order to check if you have successfully implemented the `sort.Nums` function, you can run the tests with:

```shell script
$ go test ./...
```