package main

import (
	"bufio"
	"calc"
	"fmt"
	"math"
	"os"
	"strings"
)

func main() {

	reader := bufio.NewReader(os.Stdin)
	fmt.Println("Calculator")
	for {
		fmt.Print("-> ")
		input, err := reader.ReadString('\n')
		if err != nil {
			panic(err)
		}

		input = stripNewLines(input)

		result, err := calc.Calculate(input)
		if err != nil {
			fmt.Printf("ERR:%s\n", err.Error())
		} else {
			isWholeNumber := result == math.Trunc(result)
			if isWholeNumber {
				fmt.Printf("%d\n", int64(result))
			} else {
				fmt.Printf("%f\n", result)
			}
		}
	}
}

func stripNewLines(s string) string {
	s = strings.Replace(s, "\r\n", "", -1)
	s = strings.Replace(s, "\n", "", -1)

	return s
}
