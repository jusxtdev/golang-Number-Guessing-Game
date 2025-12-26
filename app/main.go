package main

import (
	"fmt"
	"math/rand"
)

func generate_rand_num() int {
	min := 0
	max := 100
	num := rand.Intn(max-min+1) + min
	return num
}

func input_num_between_0_100() int {
	for {
		var input_num int
		fmt.Println("-------------\nEnter Number :")
		_, err := fmt.Scanln(&input_num)
		if err != nil {
			continue
		}
		if input_num >= 0 && input_num <= 100 {
			return input_num
		}
		fmt.Println("Please enter a number between 0 and 100.")
	}
}

func main() {
	// generate random number (between 0 - 100)
	random_number := generate_rand_num()

	// game loop
	for {

		// input number
		in_num := input_num_between_0_100()

		// match number
		// if input == rn, exit
		if in_num == random_number {
			fmt.Println("YOU GUESSED RIGHT !!")
			break
		} else if in_num < random_number {
			// elif input < rn, you are behind
			fmt.Println("You are behind")
		} else {
			// elif input > rn, you are ahead
			fmt.Println("You are ahead")
		}
	}
}
