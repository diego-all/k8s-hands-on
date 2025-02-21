package main

import (
	"encoding/json"
	"fmt"
	"log"
	"net/http"
	"os"
	"regexp"
	"time"
)

type HandsOn struct {
	Time     time.Time `json:"time"`
	Hostname string    `json:"hostname"`
}

func ServeHTTP(w http.ResponseWriter, r *http.Request) {
	// Aquí validamos todos los paths para detectar los payloads maliciosos
	if !validatePath(r.URL.Path) {
		// Si la validación falla, retornamos un Bad Request
		http.Error(w, "Bad Request", http.StatusBadRequest)
		return
	}

	// Si el path es '/', procesamos la respuesta correctamente
	if r.URL.Path == "/" {
		resp := HandsOn{
			Time:     time.Now(),
			Hostname: os.Getenv("HOSTNAME"),
		}

		jsonResp, err := json.Marshal(&resp)
		if err != nil {
			http.Error(w, "Error", http.StatusInternalServerError)
			return
		}

		w.Header().Set("Content-Type", "application/json")
		w.WriteHeader(http.StatusOK)
		w.Write(jsonResp)
	} else {
		// Para todas las demás rutas, retornamos un 404 Not Found
		http.NotFound(w, r)
	}
}

func validatePath(path string) bool {
	// Validar Path Traversal (buscar secuencias ../)
	pathTraversalPattern := `(\.\./|\.\.\\)`
	if matched, _ := regexp.MatchString(pathTraversalPattern, path); matched {
		// Imprimir en consola cuando se detecta un Path Traversal
		fmt.Println("Path traversal detected: ", path)
		return false
	}

	// Validar SQL Injection (Patrones comunes de SQLi)
	sqlInjectionPattern := `['";--]`
	if matched, _ := regexp.MatchString(sqlInjectionPattern, path); matched {
		// Imprimir en consola cuando se detecta un intento de SQL Injection
		fmt.Println("SQL Injection detected: ", path)
		return false
	}

	// Si no se detectó nada malicioso, devolver true
	return true
}

func main() {
	http.HandleFunc("/", ServeHTTP)
	log.Println("Servidor escuchando en http://localhost:9001")
	log.Fatal(http.ListenAndServe(":9001", nil))
}
