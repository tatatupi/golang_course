package main

import "fmt"

func main() {
	intSlice := arange(0, 11, 11)

	for _, number := range intSlice {
		if number%2 == 0 {
			fmt.Printf("%v is Even\n", number)
		} else {
			fmt.Printf("%v is Odd\n", number)
		}
	}
}
