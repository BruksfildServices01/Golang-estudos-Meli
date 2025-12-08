package main

import "fmt"

// exEncontrarNumero verifica se um número está dentro de uma lista fixa
func exEncontrarNumero() {
	numeros := []int{3, 7, 10, 15, 20}

	var buscado int
	fmt.Print("Digite um número para buscar na lista [3, 7, 10, 15, 20]: ")
	fmt.Scanln(&buscado)

	encontrou := false
	for _, n := range numeros {
		if n == buscado {
			encontrou = true
			break
		}
	}

	if encontrou {
		fmt.Println("O número", buscado, "FOI encontrado na lista.")
	} else {
		fmt.Println("O número", buscado, "NÃO foi encontrado na lista.")
	}
}
