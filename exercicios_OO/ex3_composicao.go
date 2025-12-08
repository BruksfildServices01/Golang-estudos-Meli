package main

import "fmt"

// Em Java seria:
// class Animal {
//     String nome;
//     void falar() { ... }
// }
// class Cachorro extends Animal { ... }

// Em Go: usamos composição (um struct dentro de outro)
type Animal struct {
	Nome string
}

func (a Animal) Falar() {
	fmt.Println("Meu nome é", a.Nome)
}

type Cachorro struct {
	Animal // embedding (composição)
	Raca   string
}

func (c Cachorro) Latir() {
	fmt.Println(c.Nome, "está latindo! Au au!")
}

func exComposicao() {
	dog := Cachorro{
		Animal: Animal{Nome: "Rex"},
		Raca:   "Vira-lata",
	}

	// Acesso ao método de Animal direto pelo Cachorro,
	// parecido com herança em Java, mas é composição.
	dog.Falar()
	dog.Latir()
}
