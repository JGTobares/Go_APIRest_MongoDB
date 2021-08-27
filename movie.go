// Modelo (struct) para peliculas

package main // Se define dentro el pkg main, ya que todo lo que se defina alli estara disponible para cualquier fichero

type Movie struct {
	Name     string `json:"name"` // Se utilizan backsticks para sustituir las propiedades a valores correctos de JSON
	Year     int    `json:"year"`
	Director string `json:"director"`
}

type Movies []Movie // Se define la lista o array de peliculas
