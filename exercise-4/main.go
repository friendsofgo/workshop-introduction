package main

import (
	"bufio"
	"encoding/csv"
	"flag"
	"fmt"
	"io"
	"log"
	"os"
	"strconv"
)

const problemsFile = "problems.csv"
const questionTpl = "what %v, sir?"

func main() {
	totalAnswers := 0
	correctAnswers := 0

	fileNamePtr := flag.String("file", problemsFile, "problems file name")
	flag.Parse()

	csvFile, err := os.Open(*fileNamePtr)
	if err != nil {
		log.Panicf("problems file not found: %v", *fileNamePtr)
	}
	csvReader := csv.NewReader(bufio.NewReader(csvFile))
	for {
		line, err := csvReader.Read()
		if err == io.EOF {
			break
		} else if err != nil {
			panic(err)
		} else {
			totalAnswers++

			fmt.Printf(questionTpl, line[0])

			var userAnswer int
			_, err := fmt.Scanf("%d", &userAnswer)
			if err != nil {
				// handle error
			}

			answer, err := strconv.Atoi(line[1]);
			if err != nil {
				// handle error
			}

			if userAnswer == answer {
				correctAnswers++
			}
		}
	}

	fmt.Printf("%v correct answers of %v\n", correctAnswers, totalAnswers)
}