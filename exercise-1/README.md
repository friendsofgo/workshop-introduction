## Swapping Numbers

This basic exercise basically consists on implementing a function that receives an slice of integers and returns
it with each pair of values swapped, last value should remain as is when the length of the input is odd.

```go
func Nums(nums []int) []int {
	// TODO: Swap numbers
	return nums
}
```
So, the idea is considering the the following input:

```
[11, 76, 21, 51, 70, 7, 99, 69, 17, 63]
```

Then it should return the following output:

```
[76, 11, 51, 21, 7, 70, 69, 99, 63, 17]
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