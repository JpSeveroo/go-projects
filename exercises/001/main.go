package main

import (
	"fmt"
)

func main() {
	greet("William")
	fmt.Println(getMessage())
	n1 := getOneNumber()
	n2, n3 := getTwoNumbers()
	
	total := sumThree(n1, n2, n3)
	fmt.Println("O resultado da combinação é:", total)
}