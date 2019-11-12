## Concurrent Quiz

This exercise consists on implementing concurrency on a quiz simple application.

The given program loads a battery of questions and answers from a CSV file (see `problems.csv`).

Then it prints each question through the standard output and reads the user's answer through the standard input.

At the end, it prints the number of the correct answers (over total).

```shell script
$ go run main.go

what 5+5, sir?10
what 1+1, sir?
...
...
1 correct answers of 12
```

### How to run the program

Just use:

```shell script
go run main.go
```

And start playing.

### Concurrency

As you can see, there's an infinite amount of time to answer the questions, so it has no difficulty. 

So, we want you to add a mechanism to define a maximum time to answer all the questions (i.e. 30 seconds).

In order to do it, you can use [`channels`](https://tour.golang.org/concurrency/2)
and [`goroutines`](https://tour.golang.org/concurrency/1), but not [`tickers`](https://gobyexample.com/tickers).

**Remember the following requirements:**

- When the maximum time is reached, then the program's execution must finish **immediately**.
- When the program's execution finish (either by timeout or all the questions answered), 
**it should still printing the number of correct answers (over total).** 

*NOTE: Remind the `select`'s usage to wait for multiple channels.*

### Bonus track (1)

Was it too easy?

Try to modify your current implementation by shuffling the questions.

It'd add an extra difficulty, because currently you're able to remember all the correct answers, but if we sort it...

### Bonus track (2)

Did you already finished the exercise (even the bonus track) and it seemed easy to you?

Then, go to the [flag](https://golang.org/pkg/flag/) package documentation and investigate how to 
extend your current implementation of the concurrent quiz in order to let the user define the timeout
through a command-line argument.

Would you be able achieve to achieve it? 