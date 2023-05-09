package main

import (
	"OpenClosedPrinciple/circle"
	"OpenClosedPrinciple/rectangle"
	"OpenClosedPrinciple/triangle"
	"fmt"
)

// Creamos una interfaz llamada Shape que define el método Area()
// que devuelve el área de la forma geométrica.
type Shape interface {
	Area() float64
}

func main() {
	// Creamos instancias de cada tipo de forma
	rectangle := rectangle.Rectangle{Width: 10, Height: 5}
	circle := circle.Circle{Radius: 7}
	triangle := triangle.Triangle{Base: 8, Height: 3}

	// Creamos un slice de Shape y agregamos cada forma a él.
	shapes := []Shape{rectangle, circle, triangle}

	// Recorremos el slice de formas y llamamos al método Area() de cada una,
	// sin tener que modificar el código fuente original.
	for _, shape := range shapes {
		fmt.Println("El área de la forma es:", shape.Area())
	}
}
