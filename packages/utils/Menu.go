package utils

import "fmt"

func PrintMenu() {
	fmt.Print("===================================\n")
	fmt.Println("Choose Mode")
	fmt.Println("1. Pure Bruteforce (MAX 8x8 Grid)")
	fmt.Println("2. Optimized Bruteforce")
	fmt.Println("3. Back to Input Selection")
	fmt.Print("===================================\n")
}
