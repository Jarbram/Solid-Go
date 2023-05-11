package main

import (
	"fmt"
)

// La interfaz Persistence define los métodos comunes para almacenar y recuperar datos
type Persistence interface {
	Save(data string)
	GetAll() []string
}

// La estructura Database implementa los métodos de la interfaz Persistence para almacenar y recuperar datos en una base de datos
type Database struct {
	data []string
}

func (d *Database) Save(data string) {
	d.data = append(d.data, data)
}

func (d *Database) GetAll() []string {
	return d.data
}

// La estructura Cache implementa los métodos de la interfaz Persistence para almacenar y recuperar datos en un sistema de caché
type Cache struct {
	data []string
}

func (c *Cache) Save(data string) {
	c.data = append(c.data, data)
}

func (c *Cache) GetAll() []string {
	return c.data
}

func main() {
	// Creamos un objeto de tipo Database y lo usamos para almacenar datos
	db := &Database{}
	db.Save("Datos de prueba")
	fmt.Println(db.GetAll())

	// Creamos un objeto de tipo Cache y lo usamos para almacenar datos
	cache := &Cache{}
	cache.Save("Datos de prueba en caché")
	fmt.Println(cache.GetAll())
}

//Este ejemplo ilustra cómo podemos aplicar el principio de segregación de interfaces para diseñar un código más modular y fácil de mantener. Al crear una interfaz común que define los métodos para almacenar y recuperar datos, podemos implementar diferentes estructuras que proporcionen su propia implementación específica para estos métodos. Esto hace que sea fácil cambiar la forma en que almacenamos y recuperamos datos sin tener que cambiar nuestro código en otros lugares. Además, esto también hace que sea fácil agregar nuevas estructuras que implementen la misma interfaz en el futuro, lo que hace que nuestro código sea más escalable.
