package main

import "fmt"

// exMaiorDeTres lê três números e mostra qual é o maior
func exMaiorDeTres() {
	var a, b, c int

	fmt.Print("Digite o primeiro número: ")
	fmt.Scanln(&a)

	fmt.Print("Digite o segundo número: ")
	fmt.Scanln(&b)

	fmt.Print("Digite o terceiro número: ")
	fmt.Scanln(&c)

	maior := a

	if b > maior {
		maior = b
	}
	if c > maior {
		maior = c
	}

	fmt.Println("O maior número é:", maior)
}
