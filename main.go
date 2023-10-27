package main

import (
	_ "embed"
	"fmt"
	"strconv"
	"strings"
)

var (
	//go:embed numbers.txt
	mydata []byte
)

func main() {

	fmt.Println(string(mydata))

	fmt.Println("----------------------")

	lines := strings.Split(string(mydata), "\r\n")

	var sum int
	for _, line := range lines {
		if line != "" {
			val, _ := strconv.Atoi(line)
			sum += val
		}

	}
	fmt.Println("Sum of all numbers: ", sum)

}
