package circle

import "math"

// Creamos un tipo de forma llamado Circle que implementa la interfaz Shape
// y define su propio método Area() para calcular el área del círculo.
type Circle struct {
	Radius float64
}

func (c Circle) Area() float64 {
	return math.Pi * c.Radius * c.Radius
}
