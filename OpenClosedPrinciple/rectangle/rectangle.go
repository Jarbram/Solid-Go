package rectangle

// Creamos un tipo de forma llamado Rectangle que implementa la interfaz Shape
// y define su propio método Area() para calcular el área del rectángulo.
type Rectangle struct {
	Width  float64
	Height float64
}

func (r Rectangle) Area() float64 {
	return r.Width * r.Height
}
