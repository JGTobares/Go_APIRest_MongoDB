package main

import (
	"log"      //imprime log de errores
	"net/http" //libreria para implementar el servidor
)

func main() {
	router := NewRouter()

	server := http.ListenAndServe(":8080", router)
	log.Fatal(server)
}
