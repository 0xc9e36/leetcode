package main

import "fmt"

func main() {
	s := "A"
	m := [128]int{}

	for i := range s {
		fmt.Println(s[i])
		m[s[i]]++
	}
}
