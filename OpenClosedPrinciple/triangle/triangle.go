package triangle

// Creamos un nuevo tipo de forma llamado Triangle que implementa la interfaz Shape
// y define su propio método Area() para calcular el área del triángulo.
type Triangle struct {
	Base   float64
	Height float64
}

func (t Triangle) Area() float64 {
	return 0.5 * t.Base * t.Height
}
