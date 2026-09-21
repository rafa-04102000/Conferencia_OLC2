package interpreter

import (
	"fmt"
	"os"
	"strings"

	"github.com/antlr4-go/antlr/v4"
)

const (
	LexicalError  = "Error léxico"
	SyntaxError   = "Error sintáctico"
	SemanticError = "Error semántico"
	RuntimeError  = "Error de ejecución"
)

type Error struct {
	Line   int
	Column int
	Msg    string
	Type   string
}

type ErrorTable struct {
	Errors []Error
}

func (et *ErrorTable) AddError(line int, column int, msg string, errorType string) {
	et.Errors = append(et.Errors, Error{line, column, msg, errorType})
}

func (et *ErrorTable) NewLexicalError(line int, column int, msg string) {
	et.AddError(line, column, msg, LexicalError)
}

func (et *ErrorTable) NewSyntaxError(line int, column int, msg string) {
	et.AddError(line, column, msg, SyntaxError)
}

func (et *ErrorTable) NewSemanticError(token antlr.Token, msg string) {
	et.AddError(token.GetLine(), token.GetColumn(), msg, SemanticError)
}

func (et *ErrorTable) NewRuntimeError(line int, column int, msg string) {
	et.AddError(line, column, msg, RuntimeError)
}

func NewErrorTable() *ErrorTable {
	return &ErrorTable{}
}

func (et *ErrorTable) String() string {
	var sb strings.Builder
	for _, e := range et.Errors {
		var tipo string
		switch e.Type {
		case LexicalError:
			tipo = "Lexical Error"
		case SyntaxError:
			tipo = "Syntax Error"
		case SemanticError:
			tipo = "Semantic Error"
		default:
			tipo = "Unknown Error"
		}
		sb.WriteString(fmt.Sprintf("[%s] (%d:%d) %s\n", tipo, e.Line, e.Column, e.Msg))
	}
	return sb.String()
}

func (et *ErrorTable) SaveAsHTML(filename string) error {
	html := `<!DOCTYPE html>
	<html lang="es">
	<head>
		<meta charset="UTF-8">
		<title>Reporte de Errores</title>
		<!-- Material Icons -->
		<link href="https://fonts.googleapis.com/icon?family=Material+Icons" rel="stylesheet">
		<!-- Materialize CSS -->
		<link rel="stylesheet" href="https://cdnjs.cloudflare.com/ajax/libs/materialize/1.0.0/css/materialize.min.css">
		<!-- Google Fonts - Roboto -->
		<link href="https://fonts.googleapis.com/css2?family=Roboto:wght@300;400;500;700&display=swap" rel="stylesheet">
		<style>
			body {
				font-family: 'Roboto', sans-serif;
				background-color: #f5f5f5;
			}
			.header {
				background: linear-gradient(135deg, #ff6a00 0%, #ee0979 100%);
				color: white;
				padding: 2rem 0;
				margin-bottom: 2rem;
				border-radius: 0 0 20px 20px;
				box-shadow: 0 4px 20px rgba(0,0,0,0.12);
			}
			.card-panel {
				border-radius: 10px;
				box-shadow: 0 4px 20px rgba(0,0,0,0.12);
			}
			th {
				background-color: #ee0979 !important;
				color: white !important;
				font-weight: 500;
			}
			.error-lexical {
				color: white;
				background-color: #ff5722;
				padding: 4px 10px;
				border-radius: 8px;
				font-size: 0.8rem;
				text-transform: uppercase;
			}
			.error-syntax {
				color: white;
				background-color: #fbc02d;
				padding: 4px 10px;
				border-radius: 8px;
				font-size: 0.8rem;
				text-transform: uppercase;
			}
			.error-semantic {
				color: white;
				background-color: #2196f3;
				padding: 4px 10px;
				border-radius: 8px;
				font-size: 0.8rem;
				text-transform: uppercase;
			}
			.error-runtime {
				color: white;
				background-color: #9c27b0;
				padding: 4px 10px;
				border-radius: 8px;
				font-size: 0.8rem;
				text-transform: uppercase;
			}
		</style>
	</head>
	<body>

		<div class="header center-align">
			<h3 class="white-text">Reporte de Errores</h3>
			<p class="flow-text white-text">Errores encontrados durante la ejecución</p>
		</div>

		<div class="container">
			<div class="card-panel">
				<table class="highlight centered responsive-table">
					<thead>
						<tr>
							<th>#</th>
							<th>Descripción</th>
							<th>Línea</th>
							<th>Columna</th>
							<th>Tipo</th>
						</tr>
					</thead>
					<tbody>`

	for i, err := range et.Errors {
		var class string
		switch err.Type {
		case LexicalError:
			class = "error-lexical"
		case SyntaxError:
			class = "error-syntax"
		case SemanticError:
			class = "error-semantic"
		case RuntimeError:
			class = "error-runtime"
		default:
			class = "grey lighten-2"
		}

		html += fmt.Sprintf(`
						<tr>
							<td>%d</td>
							<td>%s</td>
							<td>%d</td>
							<td>%d</td>
							<td><span class="%s">%s</span></td>
						</tr>`, i+1, err.Msg, err.Line, err.Column, class, err.Type)
	}

	html += `
					</tbody>
				</table>
			</div>
		</div>

		<footer class="page-footer pink lighten-5">
			<div class="container">
				<div class="row center-align">
					<p class="pink-text text-darken-2">Conferencia Compiladores 2</p>
				</div>
			</div>
		</footer>

		<!-- Materialize JS -->
		<script src="https://cdnjs.cloudflare.com/ajax/libs/materialize/1.0.0/js/materialize.min.js"></script>
	</body>
	</html>`

	return os.WriteFile(filename, []byte(html), 0644)
}
