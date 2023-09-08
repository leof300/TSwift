package services

import (
	"TSwiftCompiler/interpreter"
	"bufio"
	"encoding/json"
	"fmt"
	"net/http"
	"os"
	"os/exec"
)

type TSRequest struct {
	InputText string `json:"InputText"`
}

const TEMP_FILE_PATH = "services\\tempfiles\\temp.swift"

// Manejador para obtener todos los elementos
func TSRun(w http.ResponseWriter, r *http.Request) {
	var tsinput TSRequest
	err := json.NewDecoder(r.Body).Decode(&tsinput)
	if err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}

	programa := tsinput.InputText

	err = WriteText(TEMP_FILE_PATH, programa)
	if err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}

	response := interpreter.ProcessInputText(programa)

	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(response)
}

func TSOpenTree(w http.ResponseWriter, r *http.Request) {
	//path actual
	absPath, err := os.Getwd()
	if err != nil {
		fmt.Println("Error al obtener el directorio actual:", err)
		return
	}
	fmt.Println(absPath)

	//absPath := "C:/Users/erikf/OneDrive/Documentos/GitHub/TSwift/TSwiftCompiler/services/"
	absPath += "\\services"
	filePath := absPath + "\\antrTree.ps1"

	cmd := exec.Command("powershell.exe", "-File", filePath, "-pathToChangeTo", absPath)

	// Ejecutar el comando
	erre := cmd.Run()
	if erre != nil {

		fmt.Println(erre.Error())
		http.Error(w, erre.Error(), http.StatusBadRequest)
	}

	response := Response{Message: "Arbol generado exitosamente"}
	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(response)

}

func WriteText(path string, content string) error {
	// Crear o abrir el archivo para escritura
	file, err := os.Create(path)
	if err != nil {
		fmt.Println("Error al crear el archivo:", err)
		return err
	}
	defer file.Close()

	// Crear un escritor de archivo
	writer := bufio.NewWriter(file)

	// Escribir el contenido en el archivo
	_, err = writer.WriteString(content)
	if err != nil {
		fmt.Println("Error al escribir en el archivo:", err)
		return err
	}

	// Asegurarse de que los datos se escriban en el archivo
	err = writer.Flush()
	if err != nil {
		fmt.Println("Error al hacer flush en el archivo:", err)
		return err
	}
	return nil
}

type Response struct {
	Message string `json:"message"`
}
