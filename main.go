package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
	"strings"
)

func main() {
	reader := bufio.NewReader(os.Stdin)
	fmt.Println("Enter an integer number")

	for {
		fmt.Print("-> ")
		rawInput, _ := reader.ReadString('\n')
		input := strings.Trim(rawInput, "\r\n")
		number, err := strconv.Atoi(input)
		if err != nil {
			fmt.Println("That's not an integer. Try again")
		} else {
			resultArray := []string{}
			for i := 1; i <= number; i++ {
				if i%3 == 0 {
					if i%5 == 0 {
						resultArray = append(resultArray, "Fizz Buzz")
						continue
					}

					resultArray = append(resultArray, "Fizz")
					continue
				}

				if i%5 == 0 {
					resultArray = append(resultArray, "Buzz")
				} else {
					resultArray = append(resultArray, fmt.Sprintf("%v", i))
				}
			}

			fmt.Println(strings.Join(resultArray[:], ", "))
		}
	}
}
