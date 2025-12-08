package main

import "fmt"

func exVariaveis() {
	// Declarando variáveis com tipo explícito
	var idade int = 25
	var nome string = "Lucas"
	var ativo bool = true

	// Declarando sem valor inicial -> valor zero
	var numero int      // 0
	var texto string    // ""
	var verdade bool    // false
	var salario float64 // 0.0

	// Declaração curta (só dentro de função)
	cidade := "São Paulo"
	ano := 2025

	fmt.Println("=== Variáveis com valor inicial ===")
	fmt.Println("idade:", idade)
	fmt.Println("nome:", nome)
	fmt.Println("ativo:", ativo)

	fmt.Println("\n=== Variáveis com valor zero ===")
	fmt.Println("numero:", numero)
	fmt.Println("texto:", texto)
	fmt.Println("verdade:", verdade)
	fmt.Println("salario:", salario)

	fmt.Println("\n=== Declaração curta ===")
	fmt.Println("cidade:", cidade)
	fmt.Println("ano:", ano)
}
