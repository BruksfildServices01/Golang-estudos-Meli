package main

import "fmt"

func exLoops() {
	fmt.Println("=== for clássico (i := 0; i < 5; i++) ===")
	for i := 0; i < 5; i++ {
		fmt.Println("i =", i)
	}

	fmt.Println("\n=== for como while ===")
	j := 0
	for j < 3 {
		fmt.Println("j =", j)
		j++
	}

	fmt.Println("\n=== for infinito com break ===")
	k := 0
	for {
		fmt.Println("k =", k)
		k++
		if k == 3 {
			break
		}
	}

	fmt.Println("\n=== for range em slice ===")
	numeros := []int{10, 20, 30}
	for indice, valor := range numeros {
		fmt.Println("indice:", indice, "valor:", valor)
	}
}
