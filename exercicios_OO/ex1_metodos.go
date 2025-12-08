package main

import "fmt"

// Em Java seria algo como:
// class Pessoa {
//     String nome;
//     int idade;
//     void apresentar() { ... }
// }

// Em Go usamos struct + método com receiver
type Pessoa struct {
	Nome  string
	Idade int
}

// Método de valor: recebe uma cópia de Pessoa
func (p Pessoa) Apresentar() {
	fmt.Printf("Olá, meu nome é %s e tenho %d anos.\n", p.Nome, p.Idade)
}

// Método de ponteiro: pode alterar o valor original
func (p *Pessoa) FazerAniversario() {
	p.Idade++
}

func exMetodos() {
	p := Pessoa{Nome: "Lucas", Idade: 25}

	p.Apresentar() // chama método como em Java: objeto.metodo()

	p.FazerAniversario()
	p.Apresentar()
}
