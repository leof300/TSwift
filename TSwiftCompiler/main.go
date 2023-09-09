package main

import (
	"TSwiftCompiler/services"
	"fmt"
	"log"
	"net/http"
)

func main() {
	// Manejadores de rutas
	http.HandleFunc("/tscompiler/ast", services.TSRun)
	http.HandleFunc("/tscompiler/tree", services.TSOpenTree)
	http.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {
		fmt.Fprintf(w, "¡Hola, mundo!") // Responde con un mensaje simple
	})

	// Iniciar el servidor en el puerto 8080
	log.Fatal(http.ListenAndServe(":8080", nil))

	//interpreter.ProcessFileText("grammar\\ArchivosPrueba\\Basicos\\Intermedias.swift")

}
