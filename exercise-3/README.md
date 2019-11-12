## Benchmarking & Profiling

The third exercise is focused on benchmarking and profiling the `Marshal` and `Unmarshal` functions of 
the `encoding/json` package.

The code provided already contains the boilerplate needed to run the program (see `main.go`) which basically
checks that there are no errors while marhsalling and unmarshalling some dummy data.

### How to run the program

Just use:

```shell script
go run main.go
```

And check the output to see if everything is working as expected.

Otherwise you need to do some work to fix it.

### `encoding/json` performance issues

As you probably already know, the `encoding/json` package has some performance issues. 

So we recommend you to replace the provided `JsonMarshal` and `JsonUnmarhsal` functions to
use [`https://github.com/json-iterator/go`](https://github.com/json-iterator/go) instead of the standard library.

So, to do it properly, follow the steps below:

1. Do a `profile` (an save it on `data/`) to demonstrate that the `encoding/json` library is a bottleneck.
2. Write a benchmark for the `JsonMarshal` and `JsonUnmarhsal` functions.
3. Do a benchmark (an save it on `data/`) of the current implementation.
4. Replace the `encoding/json` library for `json-iterator` (imports, calls, etc).
5. Do a benchmark (an save it on `data/`) of the new implementation.
6. Do a `profile` (an save it on `data/`) to demonstrate that the `encoding/json` is not a bottleneck anymore.

*NOTE: Use the `data()` method (or the data inside it) as you consider.* 