package main

import "fmt"

// exTabuada mostra a tabuada de um número de 1 a 10
func exTabuada() {
	var n int
	fmt.Print("Digite um número para ver a tabuada: ")
	fmt.Scanln(&n)

	fmt.Println("Tabuada de", n)
	for i := 1; i <= 10; i++ {
		fmt.Printf("%d x %d = %d\n", n, i, n*i)
	}
}
