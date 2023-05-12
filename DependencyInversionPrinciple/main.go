package main

import "fmt"

// Definimos una interface para nuestro datastore genérico
type Datastore interface {
	Store(key string, value string)
	Get(key string) string
}

// Implementamos nuestra base de datos
type Database struct {
	data map[string]string
}

// Implementamos la función Store() de la interface Datastore
func (db *Database) Store(key string, value string) {
	db.data[key] = value
}

// Implementamos la función Get() de la interface Datastore
func (db *Database) Get(key string) string {
	return db.data[key]
}

// Implementamos nuestro cache
type Cache struct {
	data map[string]string
}

// Implementamos la función Store() de la interface Datastore
func (cache *Cache) Store(key string, value string) {
	cache.data[key] = value
}

// Implementamos la función Get() de la interface Datastore
func (cache *Cache) Get(key string) string {
	return cache.data[key]
}

func main() {
	// Creamos una instancia de nuestra base de datos
	db := &Database{data: make(map[string]string)}
	// Creamos una instancia de nuestro cache
	cache := &Cache{data: make(map[string]string)}

	// Creamos una lista de Datastores que incluyen nuestra base de datos y nuestro cache
	datastore := []Datastore{db, cache}

	// Iteramos sobre cada datastore en nuestra lista y llamamos Store() con los mismos argumentos
	// En este caso "key" y "value"
	for _, store := range datastore {
		store.Store("key", "value")
		// Luego llamamos Get() para obtener el valor que almacenamos
		fmt.Println(store.Get("key"))
	}
}
