package main

import (
	"bufio"
	"calc"
	"fmt"
	"os"
	"strings"
)

func main() {

	reader := bufio.NewReader(os.Stdin)
	fmt.Println("Calculator")
	for {
		fmt.Print("-> ")
		text, err := reader.ReadString('\n')
		if err != nil {
			panic(err)
		}

		text = strings.Replace(text, "\r\n", "", -1)

		result, err := calc.Calculate(text)
		if err != nil {
			fmt.Println("ERROR:" + err.Error())
		} else {
			fmt.Printf("%f\n", result)
		}
	}

}
