package main

import "fmt"

func exIf() {
	idade := 18

	if idade >= 18 {
		fmt.Println("É maior de idade")
	} else {
		fmt.Println("É menor de idade")
	}

	// if com variável local
	if nota := 7.5; nota >= 7 {
		fmt.Println("Aprovado, nota:", nota)
	} else {
		fmt.Println("Reprovado, nota:", nota)
	}

	temp := 30
	if temp < 15 {
		fmt.Println("Está frio")
	} else if temp < 25 {
		fmt.Println("Clima agradável")
	} else {
		fmt.Println("Está calor")
	}
}
