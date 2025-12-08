package main

import "fmt"

// exParOuImpar lê um número do usuário e diz se é par ou ímpar
func exParOuImpar() {
	var n int
	fmt.Print("Digite um número inteiro: ")
	fmt.Scanln(&n)

	if n%2 == 0 {
		fmt.Println("O número", n, "é PAR")
	} else {
		fmt.Println("O número", n, "é ÍMPAR")
	}
}
