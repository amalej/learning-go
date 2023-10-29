# For Loop

How to use for loops

## How to run

1. Run `go run main.go`

## Output

```
7 is odd
8 is divisible by 4
either 8 or 7 are even
both 4 and 8 are even
9 has 1 digit
```

## Notes

```go
package main

import "fmt"

func main() {

	if num := 7; num%2 == 0 {      // Checks if the number is divisible by 2
		fmt.Println("7 is even")   // Do this if condition is true
	} else {
		fmt.Println("7 is odd")    // This if condition above is false
	}

	if num := 8; num%4 == 0 {                  // Checks if the number is divisible by 2
		fmt.Println("8 is divisible by 4")     // Do this if condition is true
	}

	if 7%2 == 0 || 8%2 == 0 {                   // Checks if condition-A OR condition-B is true
		fmt.Println("either 8 or 7 are even")   // Do if EITHER conditions are true
	}

	if 4%2 == 0 && 8%2 == 0 {                   // Checks if condition-A AND condition-B are true
		fmt.Println("both 4 and 8 are even")    // Do if BOTH conditions are true
	}

	if num := 9; num < 0 {
		fmt.Println(num, "is negative")
	} else if num < 10 {
		fmt.Println(num, "has 1 digit")
	} else {
		fmt.Println(num, "has multiple digits")
	}
}
```
