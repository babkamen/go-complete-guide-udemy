package main

import "fmt"

func main() {
	a := make([]int, 11)
	for i := range a {
		if i%2 == 0 {
			fmt.Println("even")
		} else {
			fmt.Println("od")
		}
	}
}
