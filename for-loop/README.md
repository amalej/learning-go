# For Loop

How to use for loops

## How to run

1. Run `go run main.go`

## Output

```
Incemental loop from 1-3
1
2
3
Incemental loop 7-9
7
8
9
Example on how to exit/break a loop
beak the loop
Loops that continue to the next iteration
1
3
5
Looping through an array
2**0 = 1
2**1 = 2
2**2 = 4
2**3 = 8
2**4 = 16
2**5 = 32
2**6 = 64
2**7 = 128
```

## Notes

```go
package main

import "fmt"

func main() {

	fmt.Println("Incemental loop from 1-3")
	i := 1 // Declare and initialize a variable
	for i <= 3 { // Loop while "i" is less than or equal to 3
		fmt.Println(i)
		i = i + 1 // Increase value of "i"
	}

	fmt.Println("Incemental loop 7-9")
	for j := 7; j <= 9; j++ { // An initial/condition/after format for the "for" loop
		fmt.Println(j)
	}

	fmt.Println("Example on how to exit/break a loop")
	for {
		fmt.Println("beak the loop")
		break // Completely exits out of the for loop
	}

	fmt.Println("Loops that continue to the next iteration")
	for n := 0; n <= 5; n++ {
		if n%2 == 0 { // If the number is divisible by 2, continue to the next iteration
			continue // Continues to the next iteration of the for loop
		}
		fmt.Println(n)
	}

	fmt.Println("Looping through an array")
	for i, v := range []int{1, 2, 4, 8, 16, 32, 64, 128} { // Iterate over each element of the array using "range"
		fmt.Printf("2**%d = %d\n", i, v)
	}
}
```
