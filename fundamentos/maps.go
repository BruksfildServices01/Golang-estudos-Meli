package main

import "fmt"

func exMap() {
	fmt.Println("=== Map (tipo dicionário) ===")

	idades := make(map[string]int)

	idades["João"] = 25
	idades["Maria"] = 30
	idades["Ana"] = 22

	fmt.Println("idades:", idades)
	fmt.Println("Idade do João:", idades["João"])

	fmt.Println("\n=== Verificando se a chave existe ===")
	idade, existe := idades["Pedro"]
	if existe {
		fmt.Println("Idade do Pedro:", idade)
	} else {
		fmt.Println("Pedro não está no map")
	}

	fmt.Println("\n=== Percorrendo o map ===")
	for nome, idade := range idades {
		fmt.Println("nome:", nome, "idade:", idade)
	}

	fmt.Println("\n=== Deletando uma chave ===")
	delete(idades, "Ana")
	fmt.Println("idades após deletar Ana:", idades)
}
