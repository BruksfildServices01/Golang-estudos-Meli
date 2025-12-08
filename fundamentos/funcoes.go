package main

import "fmt"

// soma recebe dois ints e retorna a soma deles
func soma(a int, b int) int {
	return a + b
}

// par verifica se um número é par
func par(n int) bool {
	return n%2 == 0
}

// nomeCompleto demonstra retorno múltiplo
func nomeCompleto(nome string, sobrenome string) (string, int) {
	completo := nome + " " + sobrenome
	tamanho := len(completo)
	return completo, tamanho
}

func exFuncoes() {
	fmt.Println("=== Função soma ===")
	r := soma(3, 5)
	fmt.Println("3 + 5 =", r)

	fmt.Println("\n=== Função par ===")
	fmt.Println("2 é par?", par(2))
	fmt.Println("3 é par?", par(3))

	fmt.Println("\n=== Função com múltiplos retornos ===")
	n, tam := nomeCompleto("Lucas", "Bezerra")
	fmt.Println("nome completo:", n)
	fmt.Println("tamanho:", tam)
}
