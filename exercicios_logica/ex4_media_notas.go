package main

import "fmt"

// exMediaNotas lê 4 notas e calcula a média
func exMediaNotas() {
	const quantidade = 4
	notas := make([]float64, 0, quantidade)

	for i := 1; i <= quantidade; i++ {
		var nota float64
		fmt.Printf("Digite a nota %d: ", i)
		fmt.Scanln(&nota)
		notas = append(notas, nota)
	}

	var soma float64
	for _, n := range notas {
		soma += n
	}

	media := soma / float64(len(notas))

	fmt.Println("Notas:", notas)
	fmt.Println("Média:", media)

	if media >= 7 {
		fmt.Println("Situação: Aprovado")
	} else {
		fmt.Println("Situação: Reprovado")
	}
}
