package main

import (
	"backend/codegen"
	"backend/cst"
	"backend/errors"
	"backend/interpreter"
	"backend/parser"
	"fmt"
	"os"

	"github.com/antlr4-go/antlr/v4"

	"encoding/json"
	"net/http"
)

type RunRequest struct {
	Code string `json:"code"`
}

type RunResponse struct {
	Output string `json:"output"`
	Error  bool   `json:"error"`
}

// --------------------------------------------------------------------------------------------
//
// Funcion para Endpoint Ejecucion de Interprete de VLang Cherry y Luego el Generador de Codigo ARM64/AArch64
//
// --------------------------------------------------------------------------------------------

func interpreteYGenerarARMHandler(w http.ResponseWriter, r *http.Request) {
	var req RunRequest
	if err := json.NewDecoder(r.Body).Decode(&req); err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}

	// Paso 1: Ejecutar el intérprete como análisis sintáctico y semántico
	output, huboError := runProgram(req.Code)

	if huboError {
		response := RunResponse{
			Output: output,
			Error:  true,
		}
		w.Header().Set("Content-Type", "application/json")
		json.NewEncoder(w).Encode(response)
		return
	}

	// Paso 2: Ejecutar el generador de código ARM si no hubo errores
	armOutput, armErr := generateARMCode(req.Code)
	if armErr {
		response := RunResponse{
			Output: armOutput,
			Error:  true,
		}
		w.Header().Set("Content-Type", "application/json")
		json.NewEncoder(w).Encode(response)
		return
	}

	// Todo fue exitoso, devolvemos el código ARM como salida
	response := RunResponse{
		Output: armOutput,
		Error:  false,
	}
	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(response)
}

// --------------------------------------------------------------------------------------------
//
// Funcion para Endpoint Ejecucion de Interprete de VLang Cherry
//
// --------------------------------------------------------------------------------------------

func interpreteHandler(w http.ResponseWriter, r *http.Request) {
	var req RunRequest
	if err := json.NewDecoder(r.Body).Decode(&req); err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}

	output, huboError := runProgram(req.Code)

	if huboError {
		response := RunResponse{
			Output: output,
			Error:  true,
		}
		w.Header().Set("Content-Type", "application/json")
		json.NewEncoder(w).Encode(response)
		return
	}

	response := RunResponse{
		Output: output,
		Error:  false,
	}
	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(response)
}

// --------------------------------------------------------------------------------------------
//
// Funcion para Ejecucion de Interprete de VLang Cherry
//
// --------------------------------------------------------------------------------------------

func runProgram(code string) (salida string, err bool) {

	// defer se ejecuta despues de que la funcion principal termine, ya sea de forma normal o por un panic

	defer func() {
		// con recover recupero el panic si es que hubo uno y continuo la ejecucion
		// si no coloco el recover, el programa se detendra al encontrar un panic
		if r := recover(); r != nil {
			// caputuro el mensaje del panic y lo imprimo, ERROR mensaje Panic

			// de esta forma retorno en el deafer ya que con la palabra return no puedo
			// nombro lo que quiero retornar tanto dentro del deafe como fuera de el
			// con el mismo nombre

			salida = fmt.Sprintf("ERROR: %v", r)
			err = true
		}
	}()

	// si el panic se declara donde esta el defer, se ejecutara el defer pero no seguira ejecutandose lo de abajo
	// si el panic se declara en otro lugar, el defer, se ejecutara y luego tambien se ejecutara lo de abajo

	input := antlr.NewInputStream(code)

	// creo el error table para capturar los errores
	errorTable := interpreter.NewErrorTable()

	lexer := parser.NewGramaticaLexer(input)

	// creo el error listener para capturar los errores lexicos
	lexer.RemoveErrorListeners() // eliminar los error listeners por defecto
	lexer.AddErrorListener(errors.NewLexicalErrorListener(errorTable))

	tokens := antlr.NewCommonTokenStream(lexer, 0)
	p := parser.NewGramaticaParser(tokens)
	p.RemoveErrorListeners()
	p.AddErrorListener(errors.NewSyntaxErrorListener(errorTable))
	p.SetErrorHandler(errors.NewCustomErrorStrategy())

	p.BuildParseTrees = true
	tree := p.Program()

	eval := &interpreter.Visitor{
		Env:        interpreter.NewEnv(nil, "Global"), // entorno global
		Console:    "",                                // consola para obtener los prints y mensajes
		ErrorTable: errorTable,                        // tabla de errores
	}

	//ejecutar visitor
	eval.Visit(tree)

	// si hay errores lexicos o sintacticos, los muestro
	if len(errorTable.Errors) > 0 {
		salida = errorTable.String()
		err = true

		// Guardar reporte HTML
		if saveErr := errorTable.SaveAsHTML("Reports/tabla_errores.html"); saveErr != nil {
			fmt.Println("Error al guardar el reporte HTML:", saveErr)
		}

		return
	}

	eval.ExportarTablaSimbolosHTML("Reports/tabla_simbolos.html")
	salida = eval.Console
	err = false
	return
}

// --------------------------------------------------------------------------------------------
//
// Generador de CST en SVG
//
// --------------------------------------------------------------------------------------------

func generateReportCST(code string) (salida string, errorArb bool) {
	svg := cst.CstReport(code)

	fmt.Println("Contenido del SVG:")
	fmt.Println("SVG length:", len(svg))
	fmt.Println("Inicio SVG:", svg[:100])

	fmt.Println("SVG length:", len(svg))
	if len(svg) > 100 {
		fmt.Println("Inicio SVG:", svg[:100])
	} else {
		fmt.Println("SVG completo:", svg)
	}

	err := os.WriteFile("arbol.svg", []byte(svg), 0644)
	if err != nil {
		fmt.Println("Error guardando archivo SVG:", err)
		salida = fmt.Sprintf("Error guardando archivo SVG: %v", err)
		errorArb = true
	} else {
		fmt.Println("Árbol sintáctico guardado en arbol.svg")
		salida = "Árbol sintáctico guardado en arbol.svg"
		errorArb = false
	}
	return
}

func viewHandler(w http.ResponseWriter, r *http.Request) {
	var req RunRequest
	if err := json.NewDecoder(r.Body).Decode(&req); err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}

	output, errorArb := generateReportCST(req.Code)
	response := RunResponse{
		Output: output,
		Error:  errorArb}
	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(response)
}

// --------------------------------------------------------------------------------------------
//
// Generador de AMR64/AArch64
//
// --------------------------------------------------------------------------------------------

func generateARMCode(code string) (salida string, huboError bool) {
	defer func() {
		if r := recover(); r != nil {
			salida = fmt.Sprintf("ERROR: %v", r)
			huboError = true
		}
	}()

	input := antlr.NewInputStream(code)
	lexer := parser.NewGramaticaLexer(input)
	tokens := antlr.NewCommonTokenStream(lexer, 0)
	p := parser.NewGramaticaParser(tokens)
	p.BuildParseTrees = true
	tree := p.Program()

	gen := codegen.NewCodeGenVisitor()

	gen.Visit(tree)

	salida = gen.Data.String() + gen.BSS.String() + gen.Code.String() + gen.Fun.String()

	err := os.WriteFile("ARM64.s", []byte(salida), 0644)
	if err != nil {
		fmt.Println("Error al guardar:", err)
		salida = fmt.Sprintf("Error al guardar archivo: %v", err)
		huboError = true
		return
	}

	fmt.Println("Archivo guardado exitosamente")
	huboError = false
	return

}

func arm64Handler(w http.ResponseWriter, r *http.Request) {
	var req RunRequest
	if err := json.NewDecoder(r.Body).Decode(&req); err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}

	output, huboError := generateARMCode(req.Code)
	response := RunResponse{
		Output: output,
		Error:  huboError,
	}

	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(response)
}

// --------------------------------------------------------------------------------------------
//
// Main
//
// --------------------------------------------------------------------------------------------

func main() {

	http.HandleFunc("/interprete", interpreteHandler)
	http.HandleFunc("/cst", viewHandler)
	http.HandleFunc("/generar-arm64", arm64Handler)
	http.HandleFunc("/interprete_y_generar-arm64", interpreteYGenerarARMHandler)

	fmt.Println("Servidor escuchando en http://localhost:8080/interprete")
	fmt.Println("Servidor escuchando en http://localhost:8080/cst")
	fmt.Println("Servidor escuchando en http://localhost:8080/generar-arm64")
	fmt.Println("Servidor escuchando en http://localhost:8080/interprete_y_generar-arm64")
	http.ListenAndServe(":8080", nil)

}
