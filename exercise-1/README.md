## Swapping Numbers

This basic exercise basically consists on implementing a function that receives an slice of integers and returns
it swapped.

```go
func Nums(nums []int) []int {
	// TODO: Swap numbers
	return nums
}
```

The code provided already contains the boilerplate needed to run the program (see `main.go`) which basically
reads the numbers from the `numbers.txt`, calls the `swap.Nums` function and prints its output.

### How to run the program

Just use:

```shell script
go run main.go
```

And check that it prints the numbers properly swapped.

Otherwise you need to do some work to fix the function.

### Tests

In order to check if you have successfully implemented the `swap.Nums` function, you can run the tests with:

```shell script
$ go test ./...
```