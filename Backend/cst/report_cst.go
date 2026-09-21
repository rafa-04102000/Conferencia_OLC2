package cst

import (
	"bytes"
	"encoding/json"
	"fmt"
	"io"
	"net/http"
	"os"
	"path/filepath"
	"runtime"
)

func ReadFile(filename string) string {
	file, err := os.Open(filename)
	if err != nil {
		panic(fmt.Errorf("error abriendo archivo: %w", err))
	}
	defer file.Close()

	content, err := io.ReadAll(file)
	if err != nil {
		panic(fmt.Errorf("error leyendo archivo: %w", err))
	}
	return string(content)
}

func CstReport(input string) string {
	// Obtener ruta relativa a donde está el archivo .go actual
	_, thisFile, _, _ := runtime.Caller(0)
	basePath := filepath.Join(filepath.Dir(thisFile), "..") // sube un nivel desde /cst

	// Leer archivo Gramatica.g4 (parser + lexer en uno)
	grammarPath := filepath.Join(basePath, "Gramatica.g4")
	grammarContent := ReadFile(grammarPath)

	// Codificar a JSON
	grammarJSON, err := json.Marshal(grammarContent)
	if err != nil {
		fmt.Println("Error convirtiendo gramática a JSON:", err)
		return ""
	}

	// Codificar la entrada del usuario
	inputJSON, err := json.Marshal(input)
	if err != nil {
		fmt.Println("Error codificando entrada:", err)
		return ""
	}

	// Crear payload
	payload := []byte(fmt.Sprintf(`{
		"grammar": %s,
		"input": %s,
		"lexgrammar": %s,
		"start": "program"
	}`, grammarJSON, inputJSON, grammarJSON))

	// Enviar solicitud
	req, err := http.NewRequest("POST", "http://lab.antlr.org/parse/", bytes.NewBuffer(payload))
	if err != nil {
		fmt.Println("Error creando solicitud:", err)
		return ""
	}
	req.Header.Set("Content-Type", "application/json")

	client := &http.Client{}
	resp, err := client.Do(req)
	if err != nil {
		fmt.Println("Error enviando solicitud:", err)
		return ""
	}
	defer resp.Body.Close()

	body, err := io.ReadAll(resp.Body)
	if err != nil {
		fmt.Println("Error leyendo respuesta:", err)
		return ""
	}

	// Decodificar JSON
	var data map[string]interface{}
	if err := json.Unmarshal(body, &data); err != nil {
		fmt.Println("Error decodificando JSON:", err)
		return ""
	}

	// Extraer SVGTree
	result, ok := data["result"].(map[string]interface{})
	if !ok {
		fmt.Println("Respuesta no contiene campo 'result'")
		return ""
	}
	svg, ok := result["svgtree"].(string)
	if !ok {
		fmt.Println("Resultado no contiene campo 'svgtree'")
		return ""
	}
	return svg
}
