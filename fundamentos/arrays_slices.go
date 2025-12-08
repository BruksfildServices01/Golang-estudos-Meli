package main

import "fmt"

func exArraysESlices() {
	fmt.Println("=== Array (tamanho fixo) ===")
	var notas [3]float64
	notas[0] = 7.5
	notas[1] = 8.0
	notas[2] = 9.2

	fmt.Println("notas:", notas)
	fmt.Println("tamanho do array:", len(notas))

	fmt.Println("\n=== Slice (tamanho variável) ===")
	var frutas []string
	frutas = append(frutas, "maçã")
	frutas = append(frutas, "banana")
	frutas = append(frutas, "laranja")

	fmt.Println("frutas:", frutas)
	fmt.Println("tamanho do slice:", len(frutas))

	fmt.Println("\n=== Slice criado com make ===")
	numeros := make([]int, 0)
	numeros = append(numeros, 1, 2, 3)

	fmt.Println("numeros:", numeros)

	fmt.Println("\n=== Acessando elementos ===")
	fmt.Println("primeira fruta:", frutas[0])
	fmt.Println("segunda fruta:", frutas[1])
}
