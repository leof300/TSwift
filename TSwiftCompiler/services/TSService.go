package services

import (
	"TSwiftCompiler/interpreter"
	"encoding/json"
	"net/http"
)

type TSRequest struct {
	InputText string `json:"InputText"`
}

// Manejador para obtener todos los elementos
func TSRun(w http.ResponseWriter, r *http.Request) {
	var tsinput TSRequest
	err := json.NewDecoder(r.Body).Decode(&tsinput)

	if err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}

	response := interpreter.ProcessInputText(tsinput.InputText)

	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(response)
}
