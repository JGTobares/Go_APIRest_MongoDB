// Se definen todas las rutas del proyecto

package main

import (
	"net/http" //libreria para implementar el servidor

	"github.com/gorilla/mux" //importamos la libreria gorilla mux para el routing de webs en GO
	//permite realizar respuestas con json
)

type Route struct {
	Name       string
	Method     string
	Pattern    string
	HandleFunc http.HandlerFunc // Se emplea hadleFunc dentro de struct para simplificar la definicion de cada ruta
}

type Routes []Route // Se genera un array para manejar una coleccion de rutas

func NewRouter() *mux.Router {
	router := mux.NewRouter().StrictSlash(true)

	for _, route := range routes {
		router.
			Name(route.Name).
			Methods(route.Method).
			Path(route.Pattern).
			Handler(route.HandleFunc)
	}

	return router
}

// Cada nueva ruta se va agregando a este array
var routes = Routes{
	Route{
		"Index",
		"GET",
		"/",
		Index,
	},
	Route{
		"MovieList",
		"GET",
		"/movies",
		MovieList,
	},
	Route{
		"MovieShow",
		"GET",
		"/movie/{id}",
		MovieShow,
	},
	Route{
		"MovieAdd",
		"POST",
		"/movie",
		MovieAdd,
	},
	Route{
		"MovieUpdate",
		"PUT",
		"/movie/{id}",
		MovieUpdate,
	},
	Route{
		"MovieRemove",
		"DELETE",
		"/movie/{id}",
		MovieRemove,
	},
}
