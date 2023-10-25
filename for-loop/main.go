package main

import "fmt"

func main() {

	fmt.Println("Incemental loop from 1-3")
	i := 1
	for i <= 3 {
		fmt.Println(i)
		i = i + 1
	}

	fmt.Println("Incemental loop 7-9")
	for j := 7; j <= 9; j++ {
		fmt.Println(j)
	}

	fmt.Println("Example on how to exit/break a loop")
	for {
		fmt.Println("beak the loop")
		break
	}

	fmt.Println("Loops that continue to the next iteration")
	for n := 0; n <= 5; n++ {
		if n%2 == 0 { // If the number is divisible by 2, continue to the next iteration
			continue
		}
		fmt.Println(n)
	}

	fmt.Println("Looping through an array")
	for i, v := range []int{1, 2, 4, 8, 16, 32, 64, 128} {
		fmt.Printf("2**%d = %d\n", i, v)
	}
}
