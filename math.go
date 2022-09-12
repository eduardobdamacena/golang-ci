package main

import "fmt"

func main() {
	fmt.Print("Resultado: ")
	fmt.Println(soma(10, 10))
}

func soma(a int, b int) int {
	return a + b
}
