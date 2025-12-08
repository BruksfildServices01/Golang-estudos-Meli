package main

import "fmt"

// Em Java seria:
// interface Forma {
//     double area();
// }

// Em Go:
type Forma interface {
	Area() float64
}

// Em Java:
// class Retangulo implements Forma { ... }
// class Circulo implements Forma { ... }

type Retangulo struct {
	Largura, Altura float64
}

func (r Retangulo) Area() float64 {
	return r.Largura * r.Altura
}

type Circulo struct {
	Raio float64
}

func (c Circulo) Area() float64 {
	const pi = 3.14159
	return pi * c.Raio * c.Raio
}

func imprimirArea(f Forma) {
	// Polimorfismo: qualquer coisa que implemente Forma serve aqui
	fmt.Println("Área:", f.Area())
}

func exInterfaces() {
	r := Retangulo{Largura: 3, Altura: 4}
	c := Circulo{Raio: 2}

	// Em Java: Forma f = new Retangulo(...);
	// Em Go:   var f Forma = r
	imprimirArea(r)
	imprimirArea(c)
}
