package main

import (
	"TSwiftCompiler/services"
	"log"
	"net/http"
)

func main() {
	// Manejadores de rutas
	http.HandleFunc("/tscompiler/run", services.TSRun)

	// Iniciar el servidor en el puerto 8080
	log.Fatal(http.ListenAndServe(":8080", nil))
}
