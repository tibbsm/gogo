package main

import "fmt"

func main() {
	a := make([]int, 11)
	for i := range a {
		if i%2 == 0 {
			fmt.Printf("%d is even.\n", i)
		} else {
			fmt.Printf("%d is odd.\n", i)
		}
	}
}

// Pretty straight forward execution. I googled the best way (?) to do a for
// in Go and used that. Then I was left with a modulo operator problem or
// finding out if it was even or odd. Had some trouble with fmt.Println(), so I
// Googled some more and found that printf() was better for this case. Println()
// doesn't do formatting and Printf() does.
