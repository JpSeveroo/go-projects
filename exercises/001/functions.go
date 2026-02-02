package main

import "fmt"

// 1. Aceita parâmetro e exibe (sem sobrescrever!)
func greet(name string) {
	fmt.Println("Hello,", name)
}

// 2. Retorna uma mensagem
func getMessage() string {
	return "How are you?"
}

// 3. Soma 3 números passados como argumento
func sumThree(a, b, c int) int {
	return a + b + c
}

// 4. Retorna qualquer número
func getOneNumber() int {
	return 10
}

// 5. Retorna DOIS números
func getTwoNumbers() (int, int) {
	return 5, 15
}
