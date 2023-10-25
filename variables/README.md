# Variables

How to declare variables

## How to run

1. Run `go run main.go`

## Output

It should output

```
initial
Some Message
0

apple
```

## Notes

```go
package main

import "fmt"

func main() {

	var a = "initial" // Declare a variable named "a"
	fmt.Println(a)

	var b, c string = "Some", "Message"  // Declare multiple variables
	fmt.Println(b, c)

	var d int // Declare a variable with the type of "int". An int will default to zero
	fmt.Println(d)

	var e string // Declare a variable with the type of "int". A string will default to a blank
	fmt.Println(e)

	f := "apple" // ":=" is a syntax used for declaring and initializing a variable
	fmt.Println(f)
}

```
