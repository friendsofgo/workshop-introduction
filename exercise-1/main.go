package main

import (
	"fmt"
	"github.com/friendsofgo/workshop-introduction/exercise-1/sort"
	"io/ioutil"
	"strconv"
	"strings"
)

const (
	filePath = "./numbers.txt"
)
func main() {
	numsFromFile := readNumsFromFile(filePath)
	sortedNums := sort.Nums(numsFromFile)
	fmt.Println(sortedNums)
}

func readNumsFromFile(filePath string) []int {
	data, err := ioutil.ReadFile(filePath)
	if err != nil {
		panic(err)
	}

	lines := strings.Split(string(data), "\n")

	nums := make([]int, 0)
	for _, l := range lines {
		num, _ := strconv.Atoi(l)
		nums = append(nums, num)
	}
	return nums
}
