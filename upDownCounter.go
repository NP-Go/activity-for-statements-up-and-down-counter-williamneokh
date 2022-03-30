package main

import (
	"fmt"
	"log"
	"strconv"
)

//Create an application that asks the user to input 2 value
//The application will then count up from the starting number to the ending number and count back
//down to the starting number from the ending number in intervals of 1 step.

func main() {
	//Insert code here
	//Create
	for {
		var input1 string
		var input2 string

		fmt.Println("Please key in first number")
		_, _ = fmt.Scanln(&input1)

		num1, err := strconv.ParseInt(input1, 10, 0)
		if err != nil {
			log.Fatal("Please only key in numbers")
			return
		}
		fmt.Println("Please key in second number")
		_, _ = fmt.Scanln(&input2)
		num2, err := strconv.ParseInt(input2, 10, 0)
		if err != nil {
			log.Fatal("Please only key in numbers")
		}

		for i := num1; i <= num2; i++ {
			fmt.Println(i)
		}
		for i := num2; i >= num1; i-- {
			fmt.Println(i)
		}
	}
}
