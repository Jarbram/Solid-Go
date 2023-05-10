package main

import "fmt"

// Definimos una interfaz Animal que tiene un método Describe que devuelve un string con una descripción del animal
type Animal interface {
	Describe() string
}

// Definimos una estructura AnimalBase que incluye un campo Name que representa el nombre del animal
type AnimalBase struct {
	Name string
}

// Implementamos el método Describe de la interfaz Animal en la estructura AnimalBase
func (a AnimalBase) Describe() string {
	return fmt.Sprintf("I am a %s.", a.Name)
}

// Definimos una estructura Dog que incluye un campo Breed que representa la raza del perro
type Dog struct {
	AnimalBase // Embedimos la estructura AnimalBase dentro de la estructura Dog
	Breed      string
}

// Implementamos el método Describe de la interfaz Animal en la estructura Dog
func (d Dog) Describe() string {
	return fmt.Sprintf("%s I am a %s.", d.AnimalBase.Describe(), d.Breed)
}

// Definimos una estructura Cat que incluye un campo Color que representa el color del gato
type Cat struct {
	AnimalBase // Embedimos la estructura AnimalBase dentro de la estructura Cat
	Color      string
}

// Implementamos el método Describe de la interfaz Animal en la estructura Cat
func (c Cat) Describe() string {
	return fmt.Sprintf("%s I am a %s cat.", c.AnimalBase.Describe(), c.Color)
}

// Definimos una función PrintAnimal que acepta cualquier objeto que implemente la interfaz Animal
// y lo imprime en la consola utilizando el método Describe
func PrintAnimal(a Animal) {
	fmt.Println(a.Describe())
}

func main() {
	// Creamos un perro y un gato utilizando las estructuras Dog y Cat respectivamente
	dog := Dog{AnimalBase{"Rex"}, "Viringo"}
	cat := Cat{AnimalBase{"Lety"}, "Atigrada"}

	// Llamamos a la función PrintAnimal con el perro y el gato creados anteriormente
	PrintAnimal(dog)
	PrintAnimal(cat)

	// Podemos crear otros animales que implementen la interfaz Animal y llamar a PrintAnimal con ellos
	// sin tener que modificar la función PrintAnimal
}
