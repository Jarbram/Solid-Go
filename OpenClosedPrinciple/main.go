package main

import (
	"fmt"
	"math"
)

// Creamos una interfaz llamada Shape que define el método Area()
// que devuelve el área de la forma geométrica.
type Shape interface {
	Area() float64
}

// Creamos un tipo de forma llamado Rectangle que implementa la interfaz Shape
// y define su propio método Area() para calcular el área del rectángulo.
type Rectangle struct {
	Width  float64
	Height float64
}

func (r Rectangle) Area() float64 {
	return r.Width * r.Height
}

// Creamos un tipo de forma llamado Circle que implementa la interfaz Shape
// y define su propio método Area() para calcular el área del círculo.
type Circle struct {
	Radius float64
}

func (c Circle) Area() float64 {
	return math.Pi * c.Radius * c.Radius
}

// Creamos un nuevo tipo de forma llamado Triangle que implementa la interfaz Shape
// y define su propio método Area() para calcular el área del triángulo.
type Triangle struct {
	Base   float64
	Height float64
}

func (t Triangle) Area() float64 {
	return 0.5 * t.Base * t.Height
}

func main() {
	// Creamos instancias de cada tipo de forma
	rectangle := Rectangle{Width: 10, Height: 5}
	circle := Circle{Radius: 7}
	triangle := Triangle{Base: 8, Height: 3}

	// Creamos un slice de Shape y agregamos cada forma a él.
	shapes := []Shape{rectangle, circle, triangle}

	// Recorremos el slice de formas y llamamos al método Area() de cada una,
	// sin tener que modificar el código fuente original.
	for _, shape := range shapes {
		fmt.Println("El área de la forma es:", shape.Area())
	}
}
