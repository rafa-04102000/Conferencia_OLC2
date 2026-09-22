package interpreter

import (
	"backend/parser"
	"fmt"
	"os"
	"reflect"
	"strconv"
	"strings"

	"github.com/antlr4-go/antlr/v4"
)

// --------------------------------------------------------------------------------------------
//
//   - Variables de datos primitivos int, f64, string, bool
//   - Slices de datos primitivos int, f64, string, bool
//   - Funciones
//
// --------------------------------------------------------------------------------------------
type Variable struct {
	Nombre string      // opcional, normalmente es la clave del mapa
	Tipo   string      // "int", "float", "bool", "string"
	Valor  interface{} // puede ser cualquier tipo
	Mut    bool        // ¿es mutable?
}

// este sirve para estar retornando valores de cualquier tipo
type Value struct {
	value interface{}
	info  string
}

type Slice struct {
	Nombre    string        // opcional, normalmente es la clave del mapa
	Tipo      string        // "int", "float", "bool", "string"
	Elementos []interface{} // puede ser cualquier tipo
	Mut       bool          // ¿es mutable?
}

type StructDcl struct {
	Nombre string            // opcional, normalmente es la clave del mapa
	Campos map[string]string // campos del struct, clave: nombre del campo, valor: tipo del campo
}
type StructInstance struct {
	Nombre     string                 // opcional, normalmente es la clave del mapa
	StructType string                 // nombre del struct al que pertenece
	Campos     map[string]interface{} // campos del struct, clave: nombre del campo, valor: valor del campo
}

type Function struct {
	Parent      *Env // Entorno donde se declaró la función
	Parametros  []Parametro
	Cuerpo      *parser.BlockContext
	TipoRetorno string // tipo de retorno de la funcion, puede ser "void" si no retorna nada
	TipoSlice   string // tipo de slice si es que es una funcion que retorna un de slice
}
type Parametro struct {
	Nombre    string
	Tipo      string
	TipoSlice string // tipo de slice si es que es un parametro de tipo slice
}

// --------------------------------------------------------------------------------------------

// --------------------------------------------------------------------------------------------
//
//  Valores de Entorno
//
// --------------------------------------------------------------------------------------------

type Env struct {
	name      string
	parent    *Env
	vars      map[string]*Variable
	slicesL   map[string]*Slice
	structDcl map[string]*StructDcl
	structIns map[string]*StructInstance
	funcs     map[string]*Function
}

func NewEnv(parent *Env, name string) *Env {
	return &Env{
		name:      name,
		parent:    parent,
		vars:      make(map[string]*Variable),
		slicesL:   make(map[string]*Slice),
		structDcl: make(map[string]*StructDcl),
		structIns: make(map[string]*StructInstance),
		funcs:     make(map[string]*Function),
	}
}

// Obtener y declarar variables
func (e *Env) GetVar(id string) (*Variable, bool) {
	if val, ok := e.vars[id]; ok {
		return val, true
	}
	if e.parent != nil {
		return e.parent.GetVar(id)
	}
	return nil, false
}

func (e *Env) SetVar(id string, variable *Variable) {
	e.vars[id] = variable
}

// Obtener y declarar slices
func (e *Env) GetSlice(id string) (*Slice, bool) {
	if val, ok := e.slicesL[id]; ok {
		return val, true
	}
	if e.parent != nil {
		return e.parent.GetSlice(id)
	}
	return nil, false
}
func (e *Env) SetSlice(id string, slice *Slice) {
	e.slicesL[id] = slice
}

// Obtener y declarar Structs
func (e *Env) GetStructDefinition(id string) (*StructDcl, bool) {
	if val, ok := e.structDcl[id]; ok {
		return val, true
	}
	if e.parent != nil {
		return e.parent.GetStructDefinition(id)
	}
	return nil, false
}

func (e *Env) SetStructDefinition(id string, def *StructDcl) {
	e.structDcl[id] = def
}

// Instancia de Structs Get y Set
func (e *Env) GetStructInstance(id string) (*StructInstance, bool) {
	if val, ok := e.structIns[id]; ok {
		return val, true
	}
	if e.parent != nil {
		return e.parent.GetStructInstance(id)
	}
	return nil, false
}

func (e *Env) SetStructInstance(id string, inst *StructInstance) {
	e.structIns[id] = inst
}

// Obtener y declarar funciones

func (e *Env) SetFunc(nombre string, fn *Function) {
	e.funcs[nombre] = fn
}

func (e *Env) GetFunc(nombre string) (*Function, bool) {
	if val, ok := e.funcs[nombre]; ok {
		return val, true
	}
	if e.parent != nil {
		return e.parent.GetFunc(nombre)
	}
	return nil, false
}

// --------------------------------------------------------------------------------------------
// --------------------------------------------------------------------------------------------
//
//	Para Sentencias de Transferencia
//
// --------------------------------------------------------------------------------------------
type BreakSignal struct{}

var Break = BreakSignal{}

type ContinueSignal struct{}

var Continue = ContinueSignal{}

type ReturnSignal struct {
	Value interface{} // valor de retorno
	info  string
}

var Return = ReturnSignal{Value: nil, info: "void"} // valor por defecto, si no se retorna nada

// --------------------------------------------------------------------------------------------
// --------------------------------------------------------------------------------------------
//
//	Struct con implementacion de Visitor
//
// --------------------------------------------------------------------------------------------
type Visitor struct {
	parser.GramaticaVisitor
	// parser.GramaticaVisitor Define los métodos que puede implementar un Visitor
	// Embebo el struct Visitor con el parser.GrmaticaVisitor para que pueda implementar los métodos del Visitor
	// Hace que el Visitor satisfaga o reutilice esa interfaz/contrato segun como este generado por ANTLR
	Env *Env // entorno actual (en vez de memory)
	// en el entorno estaran las variables, funciones, slices, structs, etc.
	Console        string
	ControlContext []string // para controlar el contexto de ejecucion, para ver si ando
	// dentro de un switch, for, funcion, armo una pila para saber si sigo entrando o saliendo de contextos
	TablaSimb  []TablaSimbolos // tabla de simbolos, para guardar las variables, slices, structs y funciones
	ErrorTable *ErrorTable     // tabla de errores, para guardar los errores lexicos y sintacticos
}

type TablaSimbolos struct {
	Id          string // nombre de la variable, slice, struct, funcion
	TipoSimbolo string // tipo de simbolo, puede ser "variable", "slice", "struct", "funcion"
	TipoDato    string // tipo de dato de la variable, slice, struct, funcion
	Ambito      string // ambito donde se declara el simbolo, puede ser "Global", "Main", "Local", etc.
	Linea       int    // linea donde se declara el simbolo
	Columna     int    // columna donde se declara el simbolo
}

// Para ver los simbolos que se han declarado
func (v *Visitor) PrintTablaSimbolos() {
	fmt.Println("Tabla de Símbolos:")
	for _, s := range v.TablaSimb {
		fmt.Printf("ID: %s | Tipo: %s | Dato: %s | Línea: %d | Columna: %d | Ámbito: %s\n", s.Id, s.TipoSimbolo, s.TipoDato, s.Linea, s.Columna, s.Ambito)
	}
}

// Crear el reporte de la tabla de simbolos en HTML
func (v *Visitor) ExportarTablaSimbolosHTML(nombreArchivo string) {
	file, err := os.Create(nombreArchivo)
	if err != nil {
		//agregar error a la tabla de errores
		v.ErrorTable.NewRuntimeError(0, 0, "No se pudo crear el archivo HTML: "+err.Error())
		panic("No se pudo crear el archivo HTML: " + err.Error())
	}
	defer file.Close()

	// Inicio del HTML
	fmt.Fprintln(file, `<!DOCTYPE html>
<html lang="es">
<head>
    <meta charset="UTF-8">
    <meta name="viewport" content="width=device-width, initial-scale=1.0">
    <title>Tabla de Símbolos</title>
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
            background: linear-gradient(135deg, #3a7bd5 0%, #00d2ff 100%);
            color: white;
            padding: 2rem 0;
            margin-bottom: 2rem;
            border-radius: 0 0 20px 20px;
            box-shadow: 0 4px 20px 0 rgba(0,0,0,0.12), 0 7px 10px -5px rgba(0,0,0,0.2);
        }
        .logo-container {
            margin-bottom: 1rem;
        }
        .logo {
            width: 80px;
            height: 80px;
            border-radius: 50%;
            object-fit: cover;
            border: 3px solid white;
            box-shadow: 0 4px 8px rgba(0,0,0,0.1);
        }
        .card-panel {
            border-radius: 10px;
            box-shadow: 0 4px 20px 0 rgba(0,0,0,0.12), 0 7px 10px -5px rgba(0,0,0,0.2);
            margin-bottom: 2rem;
        }
        .team-member {
            padding: 10px;
            margin-bottom: 5px;
            border-left: 4px solid #3a7bd5;
            background-color: rgba(58, 123, 213, 0.05);
        }
        .table-container {
            overflow-x: auto;
        }
        th {
            font-weight: 500;
            background-color: #3a7bd5 !important;
            color: white !important;
        }
        tr:hover {
            background-color: rgba(0, 210, 255, 0.05) !important;
        }
        .badge {
            padding: 3px 8px;
            border-radius: 10px;
            font-size: 0.8rem;
            font-weight: 500;
            text-transform: uppercase;
        }
        .badge.variable {
            background-color: #4caf50;
            color: white;
        }
        .badge.function {
            background-color: #ff9800;
            color: white;
        }
        .badge.struct {
            background-color: #9c27b0;
            color: white;
        }
        .badge.slice {
            background-color:rgb(248, 50, 24);
            color: white;
        }
        .ambito-badge {
            padding: 3px 8px;
            border-radius: 10px;
            font-size: 0.8rem;
            font-weight: 500;
            color: white;
        }
        .ambito-global {
            background-color: #2196f3; /* Azul */
        }
        .ambito-main {
            background-color: #673ab7; /* Morado */
        }
        .ambito-function {
            background-color: #ff9800; /* Naranja */
        }
        .ambito-if {
            background-color: #4caf50; /* Verde */
        }
        .ambito-switch {
            background-color: #9c27b0; /* Morado oscuro */
        }
        .ambito-for {
            background-color: #f44336; /* Rojo */
        }
        .ambito-forinslice {
            background-color: #ff5722; /* Naranja oscuro */
        }
        .ambito-individualblock {
            background-color: #607d8b; /* Gris azulado */
        }
    </style>
</head>
<body>
    <div class="header">
        <div class="container">
            <div class="row valign-wrapper">
                <div class="col s12 m2 center-align logo-container">
                    <img src="../assets/tux.jpeg" alt="Logo" class="logo">
                </div>
                <div class="col s12 m10">
                    <h2 class="white-text">Tabla de Símbolos</h2>
                    <p class="flow-text white-text">Simbolos Declarados en la Ejecucion</p>
                </div>
            </div>
        </div>
    </div>

    <div class="container">

	<!-- Tabla de símbolos -->
        <div class="card-panel">
            <div class="table-container">
                <table class="highlight centered responsive-table">
                    <thead>
                        <tr>
                            <th><i class="material-icons left">code</i> Identificador</th>
                            <th><i class="material-icons left">category</i> Tipo</th>
                            <th><i class="material-icons left">data_usage</i> Tipo de Dato</th>
                            <th><i class="material-icons left">layers</i> Ámbito</th>
                            <th><i class="material-icons left">pin_drop</i> Línea</th>
                            <th><i class="material-icons left">view_column</i> Columna</th>
                        </tr>
                    </thead>
                    <tbody>`)

	// Agregar cada símbolo a la tabla con estilos condicionales
	for _, simbolo := range v.TablaSimb {
		// Determinar clase CSS basada en el tipo de símbolo
		badgeClass := "badge "
		switch strings.ToLower(simbolo.TipoSimbolo) {
		case "variable":
			badgeClass += "variable"
		case "slice":
			badgeClass += "slice"
		case "funcion":
			badgeClass += "function"
		case "struct":
			badgeClass += "struct"
		default:
			badgeClass += "blue"
		}

		// Determinar clase CSS para el ámbito
		ambitoClass := "ambito-badge "
		switch strings.ToLower(simbolo.Ambito) {
		case "global":
			ambitoClass += "ambito-global"
		case "main":
			ambitoClass += "ambito-main"
		case "function":
			ambitoClass += "ambito-function"
		case "if":
			ambitoClass += "ambito-if"
		case "switch":
			ambitoClass += "ambito-switch"
		case "for":
			ambitoClass += "ambito-for"
		case "forinslice":
			ambitoClass += "ambito-forinslice"
		case "individualblock":
			ambitoClass += "ambito-individualblock"
		default:
			ambitoClass += "blue-grey"
		}

		fmt.Fprintf(file,
			"\n<tr><td><strong>%s</strong></td><td><span class='%s'>%s</span></td><td>%s</td><td><span class='%s'>%s</span></td><td>%d</td><td>%d</td></tr>",
			simbolo.Id, badgeClass, simbolo.TipoSimbolo, simbolo.TipoDato, ambitoClass, simbolo.Ambito, simbolo.Linea, simbolo.Columna)
	}

	// Cierre del HTML
	fmt.Fprintln(file, `
                    </tbody>
                </table>
            </div>
        </div>
    </div>

    <!-- Footer -->
    <footer class="page-footer blue-grey lighten-5">
        <div class="container">
            <div class="row">
                <div class="col s12 center-align">
                    <p class="blue-grey-text">Conferencia Compiladores 2</p>
                </div>
            </div>
        </div>
    </footer>

    <!-- Materialize JS -->
    <script src="https://cdnjs.cloudflare.com/ajax/libs/materialize/1.0.0/js/materialize.min.js"></script>
    <script>
        document.addEventListener('DOMContentLoaded', function() {
            // Inicializar componentes de Materialize si es necesario
        });
    </script>
</body>
</html>`)

}

// --------------------------------------------------------------------------------------------

// --------------------------------------------------------------------------------------------
//
//  Funciones para devolver tipos de datos
//
// --------------------------------------------------------------------------------------------

func getDefaultValue(tipo string) interface{} {

	switch tipo {
	case "int":
		return int64(0)
	case "f64":
		return float64(0.0)
	case "string":
		return ""
	case "bool":
		return false
	case "&": //& obtengo la direecion de memoria
		return "referencia"
	case "*": //* obtengo valor que esta en esa direccion
		return "desreferencia"
	default:

		panic("Tipo desconocido: " + tipo)
	}
}

func (v *Visitor) validateType(typeValue string, value Value) {

	switch typeValue {
	case "int":
		if _, ok := value.value.(int64); !ok {
			//agregar error a la tabla de errores
			v.ErrorTable.NewRuntimeError(0, 0, "no se puede asignar un valor de tipo "+reflect.TypeOf(value.value).String()+" a una variable 'int'")

			//panic(fmt.Sprintf("no se puede asignar un valor de tipo %T a una variable 'int'", value.value))
		}
	case "f64":
		if _, ok := value.value.(float64); !ok {
			//agregar error a la tabla de errores
			v.ErrorTable.NewRuntimeError(0, 0, "no se puede asignar un valor de tipo "+reflect.TypeOf(value.value).String()+" a una variable 'float'")
			panic(fmt.Sprintf("no se puede asignar un valor de tipo %T a una variable 'float'", value.value))
		}
	case "string":
		if _, ok := value.value.(string); !ok {
			//agregar error a la tabla de errores
			v.ErrorTable.NewRuntimeError(0, 0, "no se puede asignar un valor de tipo "+reflect.TypeOf(value.value).String()+" a una variable 'string'")
			panic(fmt.Sprintf("no se puede asignar un valor de tipo %T a una variable 'string'", value.value))
		}
	case "bool":
		if _, ok := value.value.(bool); !ok {
			//agregar error a la tabla de errores
			v.ErrorTable.NewRuntimeError(0, 0, "no se puede asignar un valor de tipo "+reflect.TypeOf(value.value).String()+" a una variable 'bool'")
			panic(fmt.Sprintf("no se puede asignar un valor de tipo %T a una variable 'bool'", value.value))
		}
	default:
		//agregar error a la tabla de errores
		v.ErrorTable.NewRuntimeError(0, 0, "Tipo desconocido para la asignacion: "+typeValue)
		panic("tipo desconocido para la asignacion: " + typeValue)
	}
}

// --------------------------------------------------------------------------------------------

// --------------------------------------------------------------------------------------------
//
// Inicio del Arbol de Sintaxis Abstracto (AST)
//
// --------------------------------------------------------------------------------------------

func (v *Visitor) Visit(tree antlr.ParseTree) Value {
	switch val := tree.(type) {
	// Iniciales
	case *parser.ProgramContext:
		return Value{value: v.VisitProgram(val), info: "Program"}
	case *parser.BlockContext:
		return Value{value: v.VisitBlock(val), info: "Block"}
	case *parser.StatementContext:
		return Value{value: v.VisitStatement(val), info: "Statement"}

	// Funcion Main
	case *parser.MainfunctionContext:
		return Value{value: v.VisitMainfunction(val), info: "MainFunction"}

	// |sentencias|
	case *parser.DeclarationExplicitVarContext:
		sent := v.VisitDeclarationExplicitVar(val)
		return Value{value: sent.value, info: "Declaracion de Variable"}
	case *parser.DeclarationImplicitVarContext:
		sent := v.VisitDeclarationImplicitVar(val)
		return Value{value: sent.value, info: "Declaracion de Variable"}
	case *parser.DeclarationExplicitSliceContext:
		sent := v.VisitDeclarationExplicitSlice(val)
		//ACA ES EL PROBLEMA
		return Value{value: sent.value, info: "Declaracion de Slice"}
	case *parser.DeclarationImplicitSliceContext:
		sent := v.VisitDeclarationImplicitSlice(val)
		return Value{value: sent.value, info: "Declaracion de Slice"}
	case *parser.StructDeclarationContext:
		sent := v.VisitStructDeclaration(val)
		return Value{value: sent.value, info: sent.info}
	case *parser.StructInstanceDeclarationContext:
		sent := v.VisitStructInstanceDeclaration(val)
		return Value{value: sent.value, info: sent.info}
	case *parser.FunctionDeclarationContext:
		sent := v.VisitFunctionDeclaration(val)
		return Value{value: sent.value, info: "Declaracion de Funcion"}

	// Asignaciones
	// visit Attribute Assignment
	case *parser.AttributeAssignmentExprContext:
		sent := v.VisitAttributeAssignmentExpr(val)
		return Value{value: sent.value, info: sent.info}
	case *parser.AssignmentExprContext:
		// En este caso, no se espera un retorno, solo se ejecuta la sentencia
		sent := v.VisitAssignmentExpr(val)
		return Value{value: sent.value, info: sent.info}
	case *parser.AssignmentPlusExprContext:
		// En este caso, no se espera un retorno, solo se ejecuta la sentencia
		sent := v.VisitAssignmentPlusExpr(val)
		return Value{value: sent.value, info: sent.info}
	case *parser.AssignmentMinusExprContext:
		// En este caso, no se espera un retorno, solo se ejecuta la sentencia
		sent := v.VisitAssignmentMinusExpr(val)
		return Value{value: sent.value, info: sent.info}
	//incremento y decremento AssignmentIncrementExpr
	case *parser.AssignmentIncrementExprContext:
		// En este caso, no se espera un retorno, solo se ejecuta la sentencia
		sent := v.VisitAssignmentIncrement(val)
		return Value{value: sent.value, info: sent.info}
	case *parser.AssignmentDecrementExprContext:
		// En este caso, no se espera un retorno, solo se ejecuta la sentencia
		sent := v.VisitAssignmentDecrement(val)
		return Value{value: sent.value, info: sent.info}
	case *parser.SliceAssignmentExprContext:
		// En este caso, no se espera un retorno, solo se ejecuta la sentencia
		sent := v.VisitSliceAssignmentExpr(val)
		return Value{value: sent.value, info: sent.info}

	case *parser.PrintContext:
		// En este caso, no se espera un retorno, solo se imprime
		v.VisitPrintExpr(val)
		return Value{value: nil, info: "Print"}

	case *parser.IfContext:
		// En este caso, no se espera un retorno, solo se ejecuta la sentencia
		v.VisitIf(val)
		return Value{value: nil, info: "If"}

	case *parser.SwitchContext:
		// En este caso, no se espera un retorno, solo se ejecuta la sentencia
		v.VisitSwitch(val)
		return Value{value: nil, info: "Switch"}

	// nombre de sentencia for ->     : FOR expression '{' block '}' #ForSimple
	case *parser.ForSimpleContext:
		// En este caso, no se espera un retorno, solo se ejecuta la sentencia
		v.VisitForSimple(val)
		return Value{value: nil, info: "ForSimple"}

	// nombre de sentencia for-> FOR declaration ';' expression ';' assignment? '{' block '}' #ForDeclaration
	case *parser.ForDeclarationContext:
		// En este caso, no se espera un retorno, solo se ejecuta la sentencia
		v.VisitForDeclaration(val)
		return Value{value: nil, info: "ForDeclaration"}

	// nombre de sentencia -> FOR expression ',' expression 'in' expression '{' block '}' #ForInSlice
	case *parser.ForInSliceContext:
		// En este caso, no se espera un retorno, solo se ejecuta la sentencia
		v.VisitForInSlice(val)
		return Value{value: nil, info: "ForInSlice"}

	case *parser.IndividualBlockContext:
		// En este caso, no se espera un retorno, solo se ejecuta el bloque
		v.VisitIndividualBlock(val)
		return Value{value: nil, info: "IndividualBlock"}

	case *parser.BreakStatementContext:
		// En este caso, no se espera un retorno, solo se ejecuta la sentencia break
		v.VisitBreakStatement(val)
		return Value{value: nil, info: "Break"}
	case *parser.ContinueStatementContext:
		// En este caso, no se espera un retorno, solo se ejecuta la sentencia continue
		v.VisitContinueStatement(val)
		return Value{value: nil, info: "Continue"}
	case *parser.ReturnStatementContext:
		// En este caso, no se espera un retorno, solo se ejecuta la sentencia return
		v.VisitReturnStatement(val)
		return Value{value: nil, info: "Return"}

	// |expresiones|
	case *parser.NotExprContext:
		expr := v.VisitNotExpr(val)
		return Value{value: expr.value, info: expr.info}
	case *parser.NegExprContext:
		expr := v.VisitNegExpr(val)
		return Value{value: expr.value, info: expr.info}
	case *parser.RefExprContext:
		expr := v.VisitRefExpr(val)
		return Value{value: expr.value, info: expr.info}
	case *parser.DerefExprContext:
		expr := v.VisitDeRefExpr(val)
		return Value{value: expr.value, info: expr.info}
	case *parser.OpExprContext:
		expr := v.VisitOpExpr(val)
		return Value{value: expr.value, info: expr.info}
	case *parser.ParExprContext:
		expr := v.VisitParExpr(val)
		return Value{value: expr.value, info: expr.info}
	// Expresiones de Slices
	case *parser.SlicePrimitiveValueContext:
		expr := v.VisitSlicePrimitiveValue(val)
		return Value{value: expr.value, info: expr.info}
	case *parser.IndexOfFunctionContext:
		expr := v.VisitIndexOfFunction(val)
		return Value{value: expr.value, info: expr.info}
	case *parser.JoinFunctionContext:
		expr := v.VisitJoinFunction(val)
		return Value{value: expr.value, info: expr.info}
	case *parser.LengthFunctionContext:
		expr := v.VisitLengthFunction(val)
		return Value{value: expr.value, info: expr.info}
	case *parser.AppendExprContext:
		expr := v.VisitAppendExpr(val)
		return Value{value: expr.value, info: expr.info}
	// -------------

	// Visit Access to Struct Instance
	case *parser.SliceAccessContext:
		expr := v.VisitSliceAccess(val)
		return Value{value: expr.value, info: expr.info}
	case *parser.StructAccessContext:
		expr := v.VisitStructAccess(val)
		return Value{value: expr.value, info: expr.info}
	case *parser.FunctionCallExprContext:
		expr := v.VisitFunctionCallExpr(val)
		return Value{value: expr.value, info: expr.info}
	case *parser.AtoiExprContext:
		expr := v.VisitAtoiExpr(val)
		return Value{value: expr.value, info: expr.info}
	case *parser.ParseFloatExprContext:
		expr := v.VisitParseFloatExpr(val)
		return Value{value: expr.value, info: expr.info}
	case *parser.TypeOfExprContext:
		expr := v.VisitTypeOfExpr(val)
		return Value{value: expr.value, info: expr.info}
	case *parser.IdExprContext:
		expr := v.VisitIdExpr(val)
		return Value{value: expr.value, info: expr.info}
	case *parser.IntExprContext:
		expr := v.VisitIntExpr(val)
		return Value{value: expr.value, info: expr.info}
	case *parser.FloatExprContext:
		expr := v.VisitFloatExpr(val)
		return Value{value: expr.value, info: expr.info}
	case *parser.BoolExprContext:
		expr := v.VisitBoolExpr(val)
		return Value{value: expr.value, info: expr.info}
	case *parser.StrExprContext:
		expr := v.VisitStrExpr(val)
		return Value{value: expr.value, info: expr.info}

	// LLamadas a funciones
	case *parser.FunctionCallContext:
		expr := v.VisitFunctionCall(val)
		return Value{value: expr.value, info: expr.info}

	default:
		// Captura el tipo de dato y el texto del nodo	)
		// Captura el tipo de dato y el texto del nodo
		tipo := reflect.TypeOf(val).String()
		texto := tree.GetText()

		if token, ok := tree.GetPayload().(antlr.Token); ok {
			// esto es si es un token, entonces puedo obtener la linea y columna
			v.ErrorTable.AddError(
				token.GetLine(),
				token.GetColumn(),
				fmt.Sprintf("Sentencia desconocida: %s (%s)", texto, tipo),
				SemanticError,
			)
		} else {
			// esto es por si no he implementado el contecto, como *parse.BlockContext, *AlgunContext, etc. entonces no tiene linea ni columna, entonces pongo 0,0
			v.ErrorTable.AddError(
				0,
				0,
				fmt.Sprintf("Sentencia desconocida sin posición: %s (%s)", texto, tipo),
				SemanticError,
			)
		}

		fmt.Println("Contexto" + reflect.TypeOf(val).String())
		panic("Sentencia Deconocida :) -> " + val.GetText())
	}

	return Value{value: nil, info: "Visit"}
}

// --------------------------------------------------------------------------------------------

// --------------------------------------------------------------------------------------------
//
// Inicio de Vitias de los nodos del AST
// Inicio con los nodos altos
//
// --------------------------------------------------------------------------------------------

func (v *Visitor) VisitProgram(ctx *parser.ProgramContext) Value {
	// Crea un nuevo entorno que hereda del actual
	// previousEnv := v.env
	// v.env = NewEnv(previousEnv, "entorno")
	// defer func() { v.env = previousEnv }()
	// Al salir del bloque, regreso al entorno anterior

	v.Visit(ctx.Block())
	//	return v.Visit(ctx.Block())

	return Value{value: true, info: "Program"}
}

func (v *Visitor) VisitBlock(ctx *parser.BlockContext) Value {
	// recorro todas las sentencias dentro del bloque

	cont := 0
	var mainCtx parser.IMainfunctionContext = nil

	// Registrar todo y buscar main (no ejecutarla aún) hasta valirdar
	for _, statement := range ctx.AllStatement() {
		if v.Env.name == "Global" && statement.Mainfunction() != nil {
			cont++
			mainCtx = statement.Mainfunction() // Guardamos el nodo para luego
			continue
			// con esto evito que se ejecute main en el bloque global
			// ya que quiero que main se ejecute al final
		}
		// Registrar todo (funciones, vars, etc.)
		v.Visit(statement)
	}

	if v.Env.name == "Global" {
		if cont == 0 {
			//agrego error a la tabla
			v.ErrorTable.AddError(
				0, 0,
				"No se encontró ninguna función 'main()'",
				SemanticError,
			)
			//retorno un error
			panic("No se encontró ninguna función 'main()'")
		} else if cont > 1 {
			//agrego error a la tabla
			v.ErrorTable.AddError(
				0, 0,
				"Solo se permite una función 'main()'",
				SemanticError,
			)
			//retorno un error
			panic("Solo se permite una función 'main()'")
		}

		// Ejecutar main manualmente
		v.Visit(mainCtx)
	}
	return Value{value: true, info: "Block"}
}

func (v *Visitor) VisitStatement(ctx *parser.StatementContext) Value {

	switch {
	case ctx.Mainfunction() != nil:
		return v.Visit(ctx.Mainfunction())
	case ctx.Declaration() != nil:
		return v.Visit(ctx.Declaration())
	case ctx.Assignment() != nil:
		return v.Visit(ctx.Assignment())
	case ctx.Print_() != nil:
		return v.Visit(ctx.Print_())
	case ctx.If_() != nil:
		return v.Visit(ctx.If_())
	case ctx.Switch_() != nil:
		return v.Visit(ctx.Switch_())
	case ctx.For_() != nil:
		return v.Visit(ctx.For_())
	case ctx.IndividualBlock() != nil:
		return v.Visit(ctx.IndividualBlock())
	case ctx.TransferStatemen() != nil:
		return v.Visit(ctx.TransferStatemen())
	case ctx.Call() != nil:
		return v.Visit(ctx.Call())
	default:
		// Captura información del token
		token, ok := ctx.GetPayload().(antlr.Token)
		line, col := 0, 0
		if ok {
			line = token.GetLine()
			col = token.GetColumn()
		}
		// Agrega el error a la tabla
		v.ErrorTable.AddError(
			line,
			col,
			fmt.Sprintf("Sentencia desconocida: %s", ctx.GetText()),
			SemanticError,
		)

		//panic("Sentencia Desconocida: " + ctx.GetText())
		return Value{value: nil, info: "Sentencia Desconocida"}
	}

}

// --------------------------------------------------------------------------------------------
// Sentencias
// --------------------------------------------------------------------------------------------

// Funcion Main
func (v *Visitor) VisitMainfunction(ctx *parser.MainfunctionContext) Value {

	previousEnv := v.Env
	v.Env = NewEnv(previousEnv, "Main")
	defer func() { v.Env = previousEnv }()

	v.Visit(ctx.Block())
	return Value{value: true, info: "MainFunction"}
}

// Declaracion de variables
func (v *Visitor) VisitDeclarationExplicitVar(ctx *parser.DeclarationExplicitVarContext) Value {
	id := ctx.ID().GetText()
	tipoVariable := ctx.PrimitiveType().GetText()
	value := v.Visit(ctx.Expression())
	mutable := ctx.MUT() != nil

	// verifico si la variable ya existe
	if _, ok := v.Env.GetVar(id); ok {
		//agregar error a la tabla de errores de tipo semantic error func (et *ErrorTable) NewSemanticError(token antlr.Token, msg string)
		v.ErrorTable.AddError(
			ctx.GetStart().GetLine(),
			ctx.GetStart().GetColumn(),
			"La variable ya existe: "+id,
			SemanticError,
		)
		panic("La variable ya existe: " + id)
	}

	// verifico si hay un slice con el mismo nombre
	_, ok2 := v.Env.GetSlice(id)
	if ok2 {

		//agregar error a la tabla de errores
		v.ErrorTable.AddError(
			ctx.GetStart().GetLine(),
			ctx.GetStart().GetColumn(),
			"La variable no puede tener el mismo nombre que un slice: "+id,
			SemanticError,
		)
		panic("La variable no puede tener el mismo nombre que un slice: " + id)
	}

	// verifico si no hay una funcion con el mismo nombre
	if _, ok3 := v.Env.GetFunc(id); ok3 {
		//agregar error a la tabla de errores
		v.ErrorTable.AddError(
			ctx.GetStart().GetLine(),
			ctx.GetStart().GetColumn(),
			"La variable no puede tener el mismo nombre que una funcion: "+id,
			SemanticError,
		)
		panic("La variable no puede tener el mismo nombre que una funcion: " + id)
	}

	// si me retorna un info y es slice, entonces no es una variable
	if value.info == "slice" {
		//agregar error a la tabla de errores
		v.ErrorTable.AddError(
			ctx.GetStart().GetLine(),
			ctx.GetStart().GetColumn(),
			fmt.Sprintf("No se puede declarar una variable tipo %s con un valor de tipo slice: %s", tipoVariable, id),
			SemanticError,
		)
		panic(fmt.Sprintf("No se puede declarar una variable tipo %s con un valor de tipo slice: %s", tipoVariable, id))
	}

	// si retorna un struct, entonces no es una variable

	if value.info == "struct" {
		//agregar error a la tabla de errores
		v.ErrorTable.AddError(
			ctx.GetStart().GetLine(),
			ctx.GetStart().GetColumn(),
			fmt.Sprintf("No se puede declarar una variable tipo %s con un valor de tipo struct: %s", tipoVariable, id),
			SemanticError,
		)

		panic(fmt.Sprintf("No se puede declarar una variable tipo %s con un valor de tipo struct: %s", tipoVariable, id))
	}

	v.validateType(tipoVariable, value)

	variable := &Variable{
		Nombre: id,
		Tipo:   tipoVariable,
		Valor:  value.value,
		Mut:    mutable,
	}

	//fmt.Println("Variable declarada en entorno:", v.env.name, "Nombre:", variable.Nombre, "Tipo:", variable.Tipo, "Valor:", variable.Valor)

	v.Env.SetVar(id, variable)

	// Registro la variable en la tabla de simbolos
	v.TablaSimb = append(v.TablaSimb, TablaSimbolos{
		Id:          id,
		TipoSimbolo: "variable",
		TipoDato:    tipoVariable,
		Ambito:      v.Env.name,
		Linea:       ctx.GetStart().GetLine(),
		Columna:     ctx.GetStart().GetColumn(),
	})

	return Value{value: true, info: "Declaracion de Variable"}

}

func (v *Visitor) VisitDeclarationImplicitVar(ctx *parser.DeclarationImplicitVarContext) Value {
	id := ctx.ID().GetText()
	tipoVariable := ctx.PrimitiveType().GetText()
	mutable := ctx.MUT() != nil

	// verifico si la variable ya existe
	if _, ok := v.Env.GetVar(id); ok {
		// agregar error a la tabla de errores
		v.ErrorTable.AddError(
			ctx.GetStart().GetLine(),
			ctx.GetStart().GetColumn(),
			"La variable ya existe: "+id,
			SemanticError,
		)
		panic("La variable ya existe: " + id)
	}

	// verifico si hay un slice con el mismo nombre
	_, ok2 := v.Env.GetSlice(id)
	if ok2 {
		// agregar error a la tabla de errores
		v.ErrorTable.AddError(
			ctx.GetStart().GetLine(),
			ctx.GetStart().GetColumn(),
			"La variable no puede tener el mismo nombre que un slice: "+id,
			SemanticError,
		)
		panic("La variable no puede tener el mismo nombre que un slice: " + id)
	}

	// verifico si no hay una funcion con el mismo nombre
	if _, ok3 := v.Env.GetFunc(id); ok3 {
		// agregar error a la tabla de errores
		v.ErrorTable.AddError(
			ctx.GetStart().GetLine(),
			ctx.GetStart().GetColumn(),
			"La variable no puede tener el mismo nombre que una funcion: "+id,
			SemanticError,
		)
		panic("La variable no puede tener el mismo nombre que una funcion: " + id)
	}

	// Determino el tipo de la variable según el valor
	defaultValue := getDefaultValue(tipoVariable)
	variable := &Variable{
		Nombre: id,
		Tipo:   tipoVariable,
		Valor:  defaultValue,
		Mut:    mutable,
	}

	v.Env.SetVar(id, variable)

	// Registro la variable en la tabla de simbolos
	v.TablaSimb = append(v.TablaSimb, TablaSimbolos{
		Id:          id,
		TipoSimbolo: "variable",
		TipoDato:    tipoVariable,
		Ambito:      v.Env.name,
		Linea:       ctx.GetStart().GetLine(),
		Columna:     ctx.GetStart().GetColumn(),
	})
	return Value{value: true, info: "Declaracion de Variable"}
}

func (v *Visitor) VisitDeclarationExplicitSlice(ctx *parser.DeclarationExplicitSliceContext) Value {
	id := ctx.ID().GetText()

	//verificar que sea tipo slice unicamente
	if ctx.PrimitiveType() == nil {
		// agregar error a la tabla de errores
		v.ErrorTable.AddError(
			ctx.GetStart().GetLine(),
			ctx.GetStart().GetColumn(),
			"El tipo de dato del slice no es valido, debe ser un tipo primitivo: "+id,
			SemanticError,
		)
		panic("El tipo de dato del slice no es valido, debe ser un tipo primitivo: " + id)
	}

	tipoSlice := ctx.PrimitiveType().GetText()
	mutable := ctx.MUT() != nil

	// verifico si el slice tiene el mismo nombre que una variable
	_, ok := v.Env.GetVar(id)
	if ok {
		// agregar error a la tabla de errores
		v.ErrorTable.AddError(
			ctx.GetStart().GetLine(),
			ctx.GetStart().GetColumn(),
			"El slice no puede tener el mismo nombre que una variable: "+id,
			SemanticError,
		)
		panic("El slice no puede tener el mismo nombre que una variable: " + id)
	}

	// verifico si no hay una funcion con el mismo nombre
	if _, ok3 := v.Env.GetFunc(id); ok3 {
		// agregar error a la tabla de errores
		v.ErrorTable.AddError(
			ctx.GetStart().GetLine(),
			ctx.GetStart().GetColumn(),
			"El slice no puede tener el mismo nombre que una funcion: "+id,
			SemanticError,
		)
		panic("El slice no puede tener el mismo nombre que una funcion: " + id)
	}

	// verifico si el slice ya existe
	_, ok2 := v.Env.GetSlice(id)
	if ok2 {
		// agregar error a la tabla de errores
		v.ErrorTable.AddError(
			ctx.GetStart().GetLine(),
			ctx.GetStart().GetColumn(),
			"El slice ya existe: "+id,
			SemanticError,
		)
		panic("El slice ya existe: " + id)
	}

	// verifico el tipo del slice
	if tipoSlice != "int" && tipoSlice != "f64" && tipoSlice != "string" && tipoSlice != "bool" {
		// agregar error a la tabla de errores
		v.ErrorTable.AddError(
			ctx.GetStart().GetLine(),
			ctx.GetStart().GetColumn(),
			"Tipo de slice no soportado: "+tipoSlice,
			SemanticError,
		)
		panic("Tipo de slice no soportado: " + tipoSlice)
	}

	// creo el slice con los elementos
	newSlice := make([]interface{}, 0)

	// con esto yo le digo que quiero y ese me retorna
	// ejemplo si mando string me retorna el tipo string
	// me retorna tipos literalmente
	expectedType := map[string]reflect.Kind{
		"int":    reflect.Int64,
		"f64":    reflect.Float64,
		"string": reflect.String,
		"bool":   reflect.Bool,
	}[tipoSlice]

	// verifico si los elementos del slice son del tipo correcto
	for _, element := range ctx.ListSlice().AllExpression() {
		value := v.Visit(element)
		// esto es valido pero con lo de abajo se mira mas pro XD
		/*if value.info != tipoSlice {
			panic(fmt.Sprintf("Tipo de elemento %s no coincide con el tipo del slice %s", value.info, tipoSlice))
		}
		*/

		if reflect.TypeOf(value.value).Kind() != expectedType {
			// agregar error a la tabla de errores
			v.ErrorTable.AddError(
				ctx.GetStart().GetLine(),
				ctx.GetStart().GetColumn(),
				fmt.Sprintf("El tipo de valor %v tipo %s que se quiere introducir al slice %s no es %s", value.value, value.info, id, tipoSlice),
				SemanticError,
			)
			panic(fmt.Sprintf("El tipo de valor %v tipo %s que se quiere introducir al slice %s no es %s", value.value, value.info, id, tipoSlice))
		}

		newSlice = append(newSlice, value.value)
	}

	slice := &Slice{
		Nombre:    id,
		Tipo:      tipoSlice,
		Elementos: newSlice,
		Mut:       mutable,
	}

	v.Env.SetSlice(id, slice)

	// Registro el slice en la tabla de simbolos
	v.TablaSimb = append(v.TablaSimb, TablaSimbolos{
		Id:          id,
		TipoSimbolo: "slice",
		TipoDato:    tipoSlice,
		Ambito:      v.Env.name,
		Linea:       ctx.GetStart().GetLine(),
		Columna:     ctx.GetStart().GetColumn(),
	})

	//fmt.Println("Slice declarado en entorno:", v.env.name, "Nombre:", slice.Nombre, "Tipo:", slice.Tipo, "Elementos:", slice.Elementos)
	return Value{value: true, info: "Declaracion de Slice"}
}

func (v *Visitor) VisitDeclarationImplicitSlice(ctx *parser.DeclarationImplicitSliceContext) Value {
	id := ctx.ID().GetText()
	tipoSlice := ctx.PrimitiveType().GetText()
	mutable := ctx.MUT() != nil

	// verifico si el slice tiene el mismo nombre que una variable
	_, ok := v.Env.GetVar(id)
	if ok {
		// agregar error a la tabla de errores
		v.ErrorTable.AddError(
			ctx.GetStart().GetLine(),
			ctx.GetStart().GetColumn(),
			"El slice no puede tener el mismo nombre que una variable: "+id,
			SemanticError,
		)
		panic("El slice no puede tener el mismo nombre que una variable: " + id)
	}

	// verifico si no hay una funcion con el mismo nombre
	if _, ok3 := v.Env.GetFunc(id); ok3 {
		// agregar error a la tabla de errores
		v.ErrorTable.AddError(
			ctx.GetStart().GetLine(),
			ctx.GetStart().GetColumn(),
			"El slice no puede tener el mismo nombre que una funcion: "+id,
			SemanticError,
		)
		//
		panic("El slice no puede tener el mismo nombre que una funcion: " + id)
	}

	// verifico si el slice ya existe
	_, ok2 := v.Env.GetSlice(id)
	if ok2 {
		// agregar error a la tabla de errores
		v.ErrorTable.AddError(
			ctx.GetStart().GetLine(),
			ctx.GetStart().GetColumn(),
			"El slice ya existe: "+id,
			SemanticError,
		)
		panic("El slice ya existe: " + id)
	}

	// verifico el tipo del slice
	if tipoSlice != "int" && tipoSlice != "f64" && tipoSlice != "string" && tipoSlice != "bool" {
		// agregar error a la tabla de errores
		v.ErrorTable.AddError(
			ctx.GetStart().GetLine(),
			ctx.GetStart().GetColumn(),
			"Tipo de slice no soportado: "+tipoSlice,
			SemanticError,
		)
		panic("Tipo de slice no soportado: " + tipoSlice)
	}

	// creo el slice con los elementos, en este caso, un slice vacio
	newSlice := make([]interface{}, 0)

	slice := &Slice{
		Nombre:    id,
		Tipo:      tipoSlice,
		Elementos: newSlice,
		Mut:       mutable,
	}

	v.Env.SetSlice(id, slice)

	// Registro el slice en la tabla de simbolos
	v.TablaSimb = append(v.TablaSimb, TablaSimbolos{
		Id:          id,
		TipoSimbolo: "slice",
		TipoDato:    tipoSlice,
		Ambito:      v.Env.name,
		Linea:       ctx.GetStart().GetLine(),
		Columna:     ctx.GetStart().GetColumn(),
	})

	return Value{value: true, info: "Declaracion de Slice"}
}

// Visit Struct dcl
func (v *Visitor) VisitStructDeclaration(ctx *parser.StructDeclarationContext) Value {
	id := ctx.ID().GetText()

	// veo si no hay una variable con el mismo nombre
	if _, exists := v.Env.GetVar(id); exists {
		//agregar error a la tabla de errores
		v.ErrorTable.AddError(
			ctx.GetStart().GetLine(),
			ctx.GetStart().GetColumn(),
			"El struct no puede tener el mismo nombre que una variable: "+id,
			SemanticError,
		)
		panic("El struct no puede tener el mismo nombre que una variable: " + id)
	}

	// veo si no hay un slice con el mismo nombre
	if _, exists := v.Env.GetSlice(id); exists {
		//agregar error a la tabla de errores
		v.ErrorTable.AddError(
			ctx.GetStart().GetLine(),
			ctx.GetStart().GetColumn(),
			"El struct no puede tener el mismo nombre que un slice: "+id,
			SemanticError,
		)
		//panic("El struct no puede tener el mismo nombre que un slice: " + id)
	}
	// veo si no hay una funcion con el mismo nombre
	if _, exists := v.Env.GetFunc(id); exists {
		//agregar error a la tabla de errores
		v.ErrorTable.AddError(
			ctx.GetStart().GetLine(),
			ctx.GetStart().GetColumn(),
			"El struct no puede tener el mismo nombre que una funcion: "+id,
			SemanticError,
		)
		panic("El struct no puede tener el mismo nombre que una funcion: " + id)
	}
	// veo si no hay un struct con el mismo nombre

	if _, exists := v.Env.GetStructDefinition(id); exists {
		//agregar error a la tabla de errores
		v.ErrorTable.AddError(
			ctx.GetStart().GetLine(),
			ctx.GetStart().GetColumn(),
			"El struct ya fue declarado: "+id,
			SemanticError,
		)
		panic("El struct ya fue declarado: " + id)
	}

	campos := make(map[string]string)

	for _, field := range ctx.AllStruct_field() {
		tipo := field.PrimitiveType().GetText()
		nombre := field.ID().GetText()

		if _, ok := campos[nombre]; ok {
			//agregar error a la tabla de errores
			v.ErrorTable.AddError(
				ctx.GetStart().GetLine(),
				ctx.GetStart().GetColumn(),
				"Campo duplicado: "+nombre,
				SemanticError,
			)
			panic("Campo duplicado: " + nombre)
		}
		campos[nombre] = tipo
	}

	structDef := &StructDcl{
		Nombre: id,
		Campos: campos,
	}

	v.Env.SetStructDefinition(id, structDef)

	// Registro el struct en la tabla de simbolos
	v.TablaSimb = append(v.TablaSimb, TablaSimbolos{
		Id:          id,
		TipoSimbolo: "struct",
		TipoDato:    "struct",
		Ambito:      v.Env.name,
		Linea:       ctx.GetStart().GetLine(),
		Columna:     ctx.GetStart().GetColumn(),
	})

	return Value{value: true, info: "Struct definido: " + id}

}

// Visit Struct Assigment
func (v *Visitor) VisitStructInstanceDeclaration(ctx *parser.StructInstanceDeclarationContext) Value {
	instanceID := ctx.ID(0).GetText()
	structName := ctx.ID(1).GetText()

	def, ok := v.Env.GetStructDefinition(structName)
	if !ok {
		//agregar error a la tabla de errores
		v.ErrorTable.AddError(
			ctx.GetStart().GetLine(),
			ctx.GetStart().GetColumn(),
			"Struct no definido: "+structName,
			SemanticError,
		)
		panic("Struct no definido: " + structName)
	}

	if _, exists := v.Env.GetVar(instanceID); exists {
		//agregar error a la tabla de errores
		v.ErrorTable.AddError(
			ctx.GetStart().GetLine(),
			ctx.GetStart().GetColumn(),
			"Ya existe una variable con el nombre: "+instanceID,
			SemanticError,
		)
		panic("Ya existe una variable con el nombre: " + instanceID)
	}

	if _, exists := v.Env.GetStructInstance(instanceID); exists {
		//agregar error a la tabla de errores
		v.ErrorTable.AddError(
			ctx.GetStart().GetLine(),
			ctx.GetStart().GetColumn(),
			"Ya existe una instancia de struct con el nombre: "+instanceID,
			SemanticError,
		)
		panic("Ya existe una instancia de struct con el nombre: " + instanceID)
	}

	camposInst := make(map[string]interface{})

	for _, field := range ctx.AllLstStruct() {
		campo := field.ID().GetText()
		valor := v.Visit(field.Expression())

		expectedType, ok := def.Campos[campo]
		if !ok {
			//agregar error a la tabla de errores
			v.ErrorTable.AddError(
				ctx.GetStart().GetLine(),
				ctx.GetStart().GetColumn(),
				"Campo '"+campo+"' no está en el struct '"+structName+"'",
				SemanticError,
			)
			panic("Campo '" + campo + "' no está en el struct '" + structName + "'")
		}

		actualKind := reflect.TypeOf(valor.value).Kind()
		expectedKind := map[string]reflect.Kind{
			"int":    reflect.Int64,
			"f64":    reflect.Float64,
			"string": reflect.String,
			"bool":   reflect.Bool,
		}[expectedType]

		if actualKind != expectedKind {
			//agregar error a la tabla de errores
			v.ErrorTable.AddError(
				ctx.GetStart().GetLine(),
				ctx.GetStart().GetColumn(),
				"Campo '"+campo+"' espera tipo '"+expectedType+"', se dio tipo '"+actualKind.String()+"'",
				SemanticError,
			)
			panic(fmt.Sprintf("Campo '%s' espera '%s', se dio '%s'", campo, expectedType, actualKind))
		}

		camposInst[campo] = valor.value
	}

	for campo := range def.Campos {
		if _, ok := camposInst[campo]; !ok {
			//agregar error a la tabla de errores
			v.ErrorTable.AddError(
				ctx.GetStart().GetLine(),
				ctx.GetStart().GetColumn(),
				"Falta asignar el campo: "+campo+" en la instancia de struct '"+structName+"'",
				SemanticError,
			)
			panic("Falta asignar el campo: " + campo)
		}
	}

	instance := &StructInstance{
		Nombre:     instanceID,
		StructType: structName,
		Campos:     camposInst,
	}

	v.Env.SetStructInstance(instanceID, instance)

	return Value{value: true, info: "Instancia de " + structName}
}

// Access to Struct Instance
func (v *Visitor) VisitStructAccess(ctx *parser.StructAccessContext) Value {
	// Obtener la parte izquierda (debe ser una instancia de struct)
	instanceVal := v.Visit(ctx.Expression())
	fieldName := ctx.ID().GetText()

	// Asegurarse de que el valor es una instancia de struct
	instancia, ok := instanceVal.value.(*StructInstance)
	if !ok {
		// También puedes permitir el acceso por nombre de instancia guardada en el entorno
		if nombreInst, esString := instanceVal.value.(string); esString {
			if structInst, ok := v.Env.GetStructInstance(nombreInst); ok {
				instancia = structInst
			} else {
				//agregar error a la tabla de errores
				v.ErrorTable.AddError(
					ctx.GetStart().GetLine(),
					ctx.GetStart().GetColumn(),
					"'"+nombreInst+"' no es una instancia de struct válida",
					SemanticError,
				)
				panic(fmt.Sprintf("'%s' no es una instancia de struct válida", nombreInst))
			}
		} else {
			//agregar error a la tabla de errores
			v.ErrorTable.AddError(
				ctx.GetStart().GetLine(),
				ctx.GetStart().GetColumn(),
				"La parte izquierda no es una instancia de struct",
				SemanticError,
			)
			panic("La parte izquierda no es una instancia de struct")
		}
	}

	val, ok := instancia.Campos[fieldName]
	if !ok {
		//agregar error a la tabla de errores
		v.ErrorTable.AddError(
			ctx.GetStart().GetLine(),
			ctx.GetStart().GetColumn(),
			"El campo '"+fieldName+"' no existe en la instancia '"+instancia.Nombre+"'",
			SemanticError,
		)
		panic(fmt.Sprintf("El campo '%s' no existe en la instancia '%s'", fieldName, instancia.Nombre))
	}

	// Determinar el tipo del valor para la info
	var tipo string
	switch val.(type) {
	case int64:
		tipo = "int"
	case float64:
		tipo = "f64"
	case string:
		tipo = "string"
	case bool:
		tipo = "bool"
	default:
		tipo = "desconocido"
	}

	return Value{value: val, info: tipo}
}

func (v *Visitor) VisitFunctionDeclaration(ctx *parser.FunctionDeclarationContext) Value {

	// fmt.Println("Visitando Declaracion de Funcion")

	if v.Env.name != "Global" {

		//agregar el error a la tabla
		v.ErrorTable.AddError(
			ctx.GetStart().GetLine(),
			ctx.GetStart().GetColumn(),
			"Las funciones solo se pueden declarar en el ámbito global",
			SemanticError,
		)
		panic("Las funciones solo se pueden declarar en el ámbito global")
	}

	// Obtener el nombre de la función
	funcName := ctx.ID().GetText()

	// reviso si no hay variable con el mismo nombre
	if _, ok := v.Env.GetVar(funcName); ok {
		//agregar error a la tabla de errores
		v.ErrorTable.AddError(
			ctx.GetStart().GetLine(),
			ctx.GetStart().GetColumn(),
			"Ya existe una variable con el mismo nombre que la función: "+funcName,
			SemanticError,
		)
		panic("Ya existe una variable con el mismo nombre que la función: " + funcName)
	}

	// reviso si no hay slice con el mismo nombre
	_, ok2 := v.Env.GetSlice(funcName)
	if ok2 {
		//agregar error a la tabla de errores
		v.ErrorTable.AddError(
			ctx.GetStart().GetLine(),
			ctx.GetStart().GetColumn(),
			"Ya existe un slice con el mismo nombre que la función: "+funcName,
			SemanticError,
		)
		panic("Ya existe un slice con el mismo nombre que la función: " + funcName)
	}
	// reviso si no hay struct con el mismo nombre
	_, ok3 := v.Env.GetStructDefinition(funcName)
	if ok3 {
		//agregar error a la tabla de errores
		v.ErrorTable.AddError(
			ctx.GetStart().GetLine(),
			ctx.GetStart().GetColumn(),
			"Ya existe un struct con el mismo nombre que la función: "+funcName,
			SemanticError,
		)
		panic("Ya existe un struct con el mismo nombre que la función: " + funcName)
	}

	// reviso si no hay funciones con el mismo nombre
	if _, ok4 := v.Env.GetFunc(funcName); ok4 {
		//agregar error a la tabla de errores
		v.ErrorTable.AddError(
			ctx.GetStart().GetLine(),
			ctx.GetStart().GetColumn(),
			"Ya existe una función con el mismo nombre: "+funcName,
			SemanticError,
		)
		panic("Ya existe una función con el mismo nombre: " + funcName)
	}

	tipoRetorno := "void" // por defecto
	if ctx.PrimitiveType() != nil {
		tipoRetorno = ctx.PrimitiveType().GetText()
	} else if ctx.TiposSlice() != nil {
		tipoRetorno = "[]" + ctx.TiposSlice().PrimitiveType().GetText()
	}

	tipoSlice := "" // por defecto, no es un slice
	// veo si es un retorno de tipo slice
	if tipoRetorno == "[]int" {
		tipoRetorno = "slice"
		tipoSlice = "int"
	} else if tipoRetorno == "[]f64" {
		tipoRetorno = "slice"
		tipoSlice = "f64"
	} else if tipoRetorno == "[]string" {
		tipoRetorno = "slice"
		tipoSlice = "string"
	} else if tipoRetorno == "[]bool" {
		tipoRetorno = "slice"
		tipoSlice = "bool"
	}

	// solo le declaro el entorno donde fue declarado ya que no se ejecuta
	// le creo el nuevo enorno cuando se llama a la funcion

	// Registrar los parámetros
	var parametros []Parametro

	if ctx.Parameters() != nil {
		// Si hay parámetros, los recorremos
		for _, param := range ctx.Parameters().AllParameter() {
			switch p := param.(type) {
			case *parser.NormalTypeParameterContext:
				paramName := p.ID().GetText()
				paramType := p.PrimitiveType().GetText()
				parametros = append(parametros, Parametro{
					Nombre:    paramName,
					Tipo:      paramType,
					TipoSlice: "",
				})

			case *parser.SliceTypeParameterContext:
				paramName := p.ID().GetText()
				paramType := "slice"
				paramSlice := p.PrimitiveType().GetText()
				parametros = append(parametros, Parametro{
					Nombre:    paramName,
					Tipo:      paramType,
					TipoSlice: paramSlice,
				})

			default:
				//agregar error a la tabla de errores
				v.ErrorTable.AddError(
					ctx.GetStart().GetLine(),
					ctx.GetStart().GetColumn(),
					"Tipo de parámetro desconocido: "+param.GetText(),
					SemanticError,
				)
				panic("Tipo de parámetro desconocido")
			}
		}
	}

	if ctx.Block() == nil {
		//agregar error a la tabla de errores
		v.ErrorTable.AddError(
			ctx.GetStart().GetLine(),
			ctx.GetStart().GetColumn(),
			"La función debe tener un bloque de código",
			SemanticError,
		)
		panic("La función debe tener un bloque de código")
	}

	// Indico en que entorno se declara la funcion
	declarationEnv := v.Env

	// Crear la función

	function := &Function{
		Parent:      declarationEnv,
		Cuerpo:      ctx.Block().(*parser.BlockContext),
		TipoRetorno: tipoRetorno,
		TipoSlice:   tipoSlice,
		Parametros:  parametros,
	}

	v.Env.SetFunc(funcName, function)

	// Registro la función en la tabla de simbolos
	v.TablaSimb = append(v.TablaSimb, TablaSimbolos{
		Id:          funcName,
		TipoSimbolo: "funcion",
		TipoDato:    tipoRetorno,
		Ambito:      v.Env.name,
		Linea:       ctx.GetStart().GetLine(),
		Columna:     ctx.GetStart().GetColumn(),
	})

	return Value{value: true, info: "Declaracion de Funcion"}
}

// --------------------------------------------------------------------------------------------
// Asignaciones

// Visit AttributeAssignment
func (v *Visitor) VisitAttributeAssignmentExpr(ctx *parser.AttributeAssignmentExprContext) Value {

	// En el ambito global solo se permiten asignaciones
	if v.Env.name == "Global" {
		//agregar el error a la tabla
		v.ErrorTable.AddError(
			ctx.GetStart().GetLine(),
			ctx.GetStart().GetColumn(),
			"No se puede hacer una asignación en el ámbito global, solo se permiten declaraciones",
			SemanticError,
		)
		panic("No se puede hacer una asignación en el ámbito global, solo se permiten declaraciones")
	}

	leftExpr := ctx.Expression(0)
	right := v.Visit(ctx.Expression(1))

	// Detectar si el lado izquierdo es un acceso a campo (miStruct.campo)
	if accessCtx, ok := leftExpr.(*parser.StructAccessContext); ok {
		structID := accessCtx.Expression().GetText()
		fieldID := accessCtx.ID().GetText()

		// Obtener instancia
		inst, ok := v.Env.GetStructInstance(structID)
		if !ok {
			//agregar error a la tabla de errores
			v.ErrorTable.AddError(
				ctx.GetStart().GetLine(),
				ctx.GetStart().GetColumn(),
				"No existe la instancia de struct: "+structID,
				SemanticError,
			)

			panic("No existe la instancia de struct: " + structID)
		}

		// Obtener definición
		def, ok := v.Env.GetStructDefinition(inst.StructType)
		if !ok {
			//agregar error a la tabla de errores
			v.ErrorTable.AddError(
				ctx.GetStart().GetLine(),
				ctx.GetStart().GetColumn(),
				"No existe la definición del struct: "+inst.StructType,
				SemanticError,
			)
			panic("No existe la definición del struct: " + inst.StructType)
		}

		// Verificar que el campo exista
		expectedType, ok := def.Campos[fieldID]
		if !ok {
			//agregar error a la tabla de errores
			v.ErrorTable.AddError(
				ctx.GetStart().GetLine(),
				ctx.GetStart().GetColumn(),
				"El campo '%s' no existe en el struct '%s'",
				SemanticError,
			)
			panic(fmt.Sprintf("El campo '%s' no existe en el struct '%s'", fieldID, structID))
		}

		// Verificar tipo correcto
		actualKind := reflect.TypeOf(right.value).Kind()
		expectedKind := map[string]reflect.Kind{
			"int":    reflect.Int64,
			"f64":    reflect.Float64,
			"string": reflect.String,
			"bool":   reflect.Bool,
		}[expectedType] // aquí cast porque lo guardaste como interface{}

		if actualKind != expectedKind {
			//agregar error a la tabla de errores
			v.ErrorTable.AddError(
				ctx.GetStart().GetLine(),
				ctx.GetStart().GetColumn(),
				"Tipo incorrecto para el campo '%s': esperado %s pero se obtuvo %s",
				SemanticError,
			)

			panic(fmt.Sprintf("Tipo incorrecto para el campo '%s': esperado %s pero se obtuvo %s",
				fieldID, expectedType, actualKind))
		}

		// Asignar el valor
		inst.Campos[fieldID] = right.value
		return Value{value: true, info: "Campo actualizado exitosamente"}
	}

	// Si no es un acceso a struct, es una asignación de variable normal
	id, ok := leftExpr.(*parser.IdExprContext) // asumiendo que tienes esto
	if !ok {
		//agregar error a la tabla de errores
		v.ErrorTable.AddError(
			ctx.GetStart().GetLine(),
			ctx.GetStart().GetColumn(),
			"Asignación inválida: lado izquierdo no reconocido",
			SemanticError,
		)
		panic("Asignación inválida: lado izquierdo no reconocido")
	}

	varName := id.GetText()
	variable, exists := v.Env.GetVar(varName)
	if !exists {
		//agregar error a la tabla de errores
		v.ErrorTable.AddError(
			ctx.GetStart().GetLine(),
			ctx.GetStart().GetColumn(),
			"Variable no declarada: "+varName,
			SemanticError,
		)
		panic("Variable no declarada: " + varName)
	}

	// Validar tipo
	actualKind := reflect.TypeOf(right.value).Kind()
	expectedKind := map[string]reflect.Kind{
		"int":    reflect.Int64,
		"f64":    reflect.Float64,
		"string": reflect.String,
		"bool":   reflect.Bool,
	}[variable.Tipo]

	if actualKind != expectedKind {
		//agregar error a la tabla de errores
		v.ErrorTable.AddError(
			ctx.GetStart().GetLine(),
			ctx.GetStart().GetColumn(),
			"Tipo incorrecto: esperado %s pero se obtuvo %s",
			SemanticError,
		)

		panic(fmt.Sprintf("Tipo incorrecto: esperado %s pero se obtuvo %s", variable.Tipo, actualKind))
	}

	variable.Valor = right.value
	return Value{value: true, info: "Variable actualizada exitosamente"}
}

func (v *Visitor) VisitAssignmentExpr(ctx *parser.AssignmentExprContext) Value {

	// En el ambito global solo se permiten asignaciones
	if v.Env.name == "Global" {
		//agregar el error a la tabla
		v.ErrorTable.AddError(
			ctx.GetStart().GetLine(),
			ctx.GetStart().GetColumn(),
			"No se puede hacer una asignación en el ámbito global, solo se permiten declaraciones",
			SemanticError,
		)
		panic("No se puede hacer una asignación en el ámbito global, solo se permiten declaraciones")
	}

	id := ctx.ID().GetText()

	// veo si es una variable o un slice

	// si es variable
	variable, ok := v.Env.GetVar(id)
	if ok {

		// Verifico si la variable es mutable
		// if !variable.Mut {
		// 	panic("La variable no es mutable: " + id)
		// }

		// Visito la expresión para obtener el nuevo valor
		newValue := v.Visit(ctx.Expression())

		if newValue.info != "int" && newValue.info != "f64" && newValue.info != "string" && newValue.info != "bool" {
			//agregar error a la tabla de errores
			v.ErrorTable.AddError(
				ctx.GetStart().GetLine(),
				ctx.GetStart().GetColumn(),
				"Tipo de dato no soportado para la asignación a una variable: "+newValue.info,
				SemanticError,
			)
			panic(fmt.Sprintf("Tipo de dato no soportado para la asignación a una variable: %s", newValue.info))
		}

		// Valido el tipo del nuevo valor
		v.validateType(variable.Tipo, newValue)

		//verifico si el tipo de la variable coincide con el tipo del nuevo valor
		if variable.Tipo != newValue.info {
			//agregar error a la tabla de errores
			v.ErrorTable.AddError(
				ctx.GetStart().GetLine(),
				ctx.GetStart().GetColumn(),
				"Tipo de dato de la variable '"+variable.Tipo+"' no coincide con el tipo del nuevo valor '"+newValue.info+"'",
				SemanticError,
			)

			//retorna nulo
			return Value{value: nil, info: "Error en la asignación de variable"}

			//panic(fmt.Sprintf("Tipo de dato de la variable %s no coincide con el tipo del nuevo valor %s", variable.Tipo, newValue.info))
		}

		// Actualizo el valor de la variable en el entorno
		// por el puntero
		variable.Valor = newValue.value

		return Value{value: true, info: "Assingancion/Reasignacion de Valor en Variable"}
	}

	// si es slice
	slice, ok2 := v.Env.GetSlice(id)
	if ok2 {
		// Verifico si el slice es mutable
		// if !slice.Mut {
		// 	panic("El slice no es mutable: " + id)
		// }

		// Visito la expresión para obtener el nuevo valor
		newValue := v.Visit(ctx.Expression())
		if newValue.info != "slice" {
			//agregar error a la tabla de errores
			v.ErrorTable.AddError(
				ctx.GetStart().GetLine(),
				ctx.GetStart().GetColumn(),
				"Tipo de dato no soportado para la asignación a un slice: "+newValue.info,
				SemanticError,
			)

			panic(fmt.Sprintf("Tipo de dato no soportado para la asignación a un slice: %s", newValue.info))
		}

		// veo si el slice es del tipo correcto

		if slice.Tipo != newValue.value.(*Slice).Tipo {
			//agregar error a la tabla de errores
			v.ErrorTable.AddError(
				ctx.GetStart().GetLine(),
				ctx.GetStart().GetColumn(),
				"Tipo de dato del slice '"+slice.Tipo+"' no coincide con el tipo del nuevo valor '"+newValue.value.(*Slice).Tipo+"'",
				SemanticError,
			)

			panic(fmt.Sprintf("Tipo de dato del slice %s no coincide con el tipo del nuevo valor %s", slice.Tipo, newValue.value.(*Slice).Tipo))
		}

		// Actualizo el valor del slice en el entorno
		slice.Elementos = newValue.value.(*Slice).Elementos
		return Value{value: true, info: "Assingancion/Reasignacion de Valor en Slice"}

	}

	// aca poner el struct
	// si no es variable ni slice, entonces es un struct
	structInstance, ok3 := v.Env.GetStructInstance(id)
	if ok3 {
		// Verifico si el struct es mutable
		// if !structInstance.Mut {
		// 	panic("El struct no es mutable: " + id)
		// }
		// Visito la expresión para obtener el nuevo valor
		newValue := v.Visit(ctx.Expression())
		if newValue.info != "struct" {
			//agregar error a la tabla de errores
			v.ErrorTable.AddError(
				ctx.GetStart().GetLine(),
				ctx.GetStart().GetColumn(),
				"Tipo de dato no soportado para la asignación a un struct: "+newValue.info,
				SemanticError,
			)
			panic(fmt.Sprintf("Tipo de dato no soportado para la asignación a un struct: %s", newValue.info))
		}
		// veo si el struct es del tipo correcto
		if structInstance.StructType != newValue.value.(*StructInstance).StructType {
			//agregar error a la tabla de errores
			v.ErrorTable.AddError(
				ctx.GetStart().GetLine(),
				ctx.GetStart().GetColumn(),
				"Tipo de dato del struct '"+structInstance.StructType+"' no coincide con el tipo del nuevo valor '"+newValue.value.(*StructInstance).StructType+"'",
				SemanticError,
			)
			panic(fmt.Sprintf("Tipo de dato del struct %s no coincide con el tipo del nuevo valor %s", structInstance.StructType, newValue.value.(*StructInstance).StructType))
		}
		// Actualizo el valor del struct en el entorno
		structInstance.Campos = newValue.value.(*StructInstance).Campos
		return Value{value: true, info: "Assingancion/Reasignacion de Valor en Struct"}
	}

	// Si no es ni variable, ni slice, ni struct, entonces es un error
	//agregar error a la tabla de errores
	v.ErrorTable.AddError(
		ctx.GetStart().GetLine(),
		ctx.GetStart().GetColumn(),
		"No se encontro el elemento: "+id,
		SemanticError,
	)
	fmt.Println("No se encontro el elemento: " + id)
	//panic("No se encontro el elemento: " + id)
	return Value{value: false, info: "Error en la asignación"}

}

func (v *Visitor) VisitAssignmentPlusExpr(ctx *parser.AssignmentPlusExprContext) Value {

	// En el ambito global solo se permiten declaraciones
	if v.Env.name == "Global" {
		//agregar el error a la tabla
		v.ErrorTable.AddError(
			ctx.GetStart().GetLine(),
			ctx.GetStart().GetColumn(),
			"No se puede hacer una asignación += en el ámbito global, solo se permiten declaraciones",
			SemanticError,
		)
		panic("No se puede hacer una asignación += en el ámbito global, solo se permiten declaraciones")
	}

	id := ctx.ID().GetText()

	// veo si es un slice y si lo fuera le digo que no se puede hacer una asignacion += a un slice
	_, ok := v.Env.GetSlice(id)
	if ok {
		//agregar error a la tabla de errores
		v.ErrorTable.AddError(
			ctx.GetStart().GetLine(),
			ctx.GetStart().GetColumn(),
			"No se puede hacer una asignación += a un slice: "+id,
			SemanticError,
		)

		panic("No se puede hacer una asignación += a un slice: " + id)
	}

	// si es variable
	variable, ok := v.Env.GetVar(id)

	if !ok {
		//agregar error a la tabla de errores
		v.ErrorTable.AddError(
			ctx.GetStart().GetLine(),
			ctx.GetStart().GetColumn(),
			"No se encontro la variable: "+id,
			SemanticError,
		)
		panic("No se encontro la variable: " + id)
	}

	// Verifico si la variable es mutable
	// if !variable.Mut {
	// 	panic("La variable no es mutable: " + id)
	// }

	// Visito la expresión para obtener el valor a sumar
	valueToAdd := v.Visit(ctx.Expression())

	if valueToAdd.info != "int" && valueToAdd.info != "f64" && valueToAdd.info != "string" && valueToAdd.info != "bool" {
		//agregar error a la tabla de errores
		v.ErrorTable.AddError(
			ctx.GetStart().GetLine(),
			ctx.GetStart().GetColumn(),
			"Tipo de dato no soportado para la asignación += a una variable: "+valueToAdd.info,
			SemanticError,
		)

		panic(fmt.Sprintf("Tipo de dato no soportado para la asignación += a una variable: %s", valueToAdd.info))
	}

	// Valido el tipo del nuevo valor
	v.validateType(variable.Tipo, valueToAdd)

	// Actualizo el valor de la variable en el entorno
	switch variable.Tipo {
	case "int":
		variable.Valor = variable.Valor.(int64) + valueToAdd.value.(int64)
	case "f64":
		variable.Valor = variable.Valor.(float64) + valueToAdd.value.(float64)
	case "string":
		variable.Valor = variable.Valor.(string) + valueToAdd.value.(string)
	default:
		//agregar error a la tabla de errores
		v.ErrorTable.AddError(
			ctx.GetStart().GetLine(),
			ctx.GetStart().GetColumn(),
			"Tipo de dato no soportado para la asignación += a una variable: "+variable.Tipo,
			SemanticError,
		)
		panic("Tipo de variable no soportado para la asignación: " + variable.Tipo)
	}

	return Value{value: true, info: "Assingancion/Reasignacion de Valor en Variable"}
}

func (v *Visitor) VisitAssignmentMinusExpr(ctx *parser.AssignmentMinusExprContext) Value {

	// En el ambito global solo se permiten declaraciones
	if v.Env.name == "Global" {
		//agregar el error a la tabla
		v.ErrorTable.AddError(
			ctx.GetStart().GetLine(),
			ctx.GetStart().GetColumn(),
			"No se puede hacer una asignación -= en el ámbito global, solo se permiten declaraciones",
			SemanticError,
		)

		panic("No se puede hacer una asignación -= en el ámbito global, solo se permiten declaraciones")
	}

	id := ctx.ID().GetText()

	// veo si es un slice y si lo fuera le digo que no se puede hacer una asignacion += a un slice
	_, ok := v.Env.GetSlice(id)
	if ok {
		//agregar error a la tabla de errores
		v.ErrorTable.AddError(
			ctx.GetStart().GetLine(),
			ctx.GetStart().GetColumn(),
			"No se puede hacer una asignación -= a un slice: "+id,
			SemanticError,
		)
		panic("No se puede hacer una asignación += a un slice: " + id)
	}

	variable, ok := v.Env.GetVar(id)

	if !ok {
		//agregar error a la tabla de errores
		v.ErrorTable.AddError(
			ctx.GetStart().GetLine(),
			ctx.GetStart().GetColumn(),
			"No se encontro la variable: "+id,
			SemanticError,
		)
		panic("No se encontro la variable: " + id)
	}

	// Verifico si la variable es mutable
	// if !variable.Mut {
	// 	panic("La variable no es mutable: " + id)
	// }

	// Visito la expresión para obtener el valor a restar
	valueToSubtract := v.Visit(ctx.Expression())

	if valueToSubtract.info != "int" && valueToSubtract.info != "f64" && valueToSubtract.info != "string" && valueToSubtract.info != "bool" {
		//agregar error a la tabla de errores
		v.ErrorTable.AddError(
			ctx.GetStart().GetLine(),
			ctx.GetStart().GetColumn(),
			"Tipo de dato no soportado para la asignación -= a una variable: "+valueToSubtract.info,
			SemanticError,
		)

		panic(fmt.Sprintf("Tipo de dato no soportado para la asignación += a una variable: %s", valueToSubtract.info))
	}

	// Valido el tipo del nuevo valor
	v.validateType(variable.Tipo, valueToSubtract)

	// Actualizo el valor de la variable en el entorno
	switch variable.Tipo {
	case "int":
		variable.Valor = variable.Valor.(int64) - valueToSubtract.value.(int64)
	case "f64":
		variable.Valor = variable.Valor.(float64) - valueToSubtract.value.(float64)
	default:
		//agregar error a la tabla de errores
		v.ErrorTable.AddError(
			ctx.GetStart().GetLine(),
			ctx.GetStart().GetColumn(),
			"Tipo de dato no soportado para la asignación -= a una variable: "+variable.Tipo,
			SemanticError,
		)

		panic("Tipo de variable no soportado para la asignación: " + variable.Tipo)
	}

	return Value{value: true, info: "Assingancion/Reasignacion de Valor en Variable"}
}

// AssignmentIncrementExpr
func (v *Visitor) VisitAssignmentIncrement(ctx *parser.AssignmentIncrementExprContext) Value {

	// En el ambito global solo se permiten declaraciones
	if v.Env.name == "Global" {
		//agregar el error a la tabla
		v.ErrorTable.AddError(
			ctx.GetStart().GetLine(),
			ctx.GetStart().GetColumn(),
			"No se puede hacer una asignación ++ en el ámbito global, solo se permiten declaraciones",
			SemanticError,
		)

		panic("No se puede hacer una asignación ++ en el ámbito global, solo se permiten declaraciones")
	}

	id := ctx.ID().GetText()

	// veo si es un slice y si lo fuera le digo que no se puede hacer una asignacion += a un slice
	_, ok := v.Env.GetSlice(id)
	if ok {
		//agregar error a la tabla de errores
		v.ErrorTable.AddError(
			ctx.GetStart().GetLine(),
			ctx.GetStart().GetColumn(),
			"No se puede hacer una asignación ++ a un slice: "+id,
			SemanticError,
		)
		panic("No se puede hacer una asignación ++ a un slice: " + id)
	}

	variable, ok := v.Env.GetVar(id)

	if !ok {
		//agregar error a la tabla de errores
		v.ErrorTable.AddError(
			ctx.GetStart().GetLine(),
			ctx.GetStart().GetColumn(),
			"No se encontro la variable: "+id,
			SemanticError,
		)
		panic("No se encontro la variable: " + id)
	}

	// Verifico si la variable es mutable
	// if !variable.Mut {
	// 	panic("La variable no es mutable: " + id)
	// }

	if variable.Tipo != "int" && variable.Tipo != "f64" {
		//agregar error a la tabla de errores
		v.ErrorTable.AddError(
			ctx.GetStart().GetLine(),
			ctx.GetStart().GetColumn(),
			"Tipo de dato no soportado para la asignación ++ a una variable: "+variable.Tipo,
			SemanticError,
		)
		panic(fmt.Sprintf("Tipo de dato no soportado para la asignación ++ a una variable: %s", variable.Tipo))
	}

	// Actualizo el valor de la variable en el entorno
	switch variable.Tipo {
	case "int":
		variable.Valor = variable.Valor.(int64) + 1
	case "f64":
		variable.Valor = variable.Valor.(float64) + 1.0
	default:
		//agregar error a la tabla de errores
		v.ErrorTable.AddError(
			ctx.GetStart().GetLine(),
			ctx.GetStart().GetColumn(),
			"Tipo de dato no soportado para la asignación ++ a una variable: "+variable.Tipo,
			SemanticError,
		)
		panic("Tipo de variable no soportado para la asignación: " + variable.Tipo)
	}

	return Value{value: true, info: "Assingancion/Reasignacion de Valor en Variable"}
}

func (v *Visitor) VisitAssignmentDecrement(ctx *parser.AssignmentDecrementExprContext) Value {
	// En el ambito global solo se permiten declaraciones
	if v.Env.name == "Global" {
		//agregar el error a la tabla
		v.ErrorTable.AddError(
			ctx.GetStart().GetLine(),
			ctx.GetStart().GetColumn(),
			"No se puede hacer una asignación -- en el ámbito global, solo se permiten declaraciones",
			SemanticError,
		)
		panic("No se puede hacer una asignación -- en el ámbito global, solo se permiten declaraciones")
	}

	id := ctx.ID().GetText()

	// veo si es un slice y si lo fuera le digo que no se puede hacer una asignacion += a un slice
	_, ok := v.Env.GetSlice(id)
	if ok {
		//agregar error a la tabla de errores
		v.ErrorTable.AddError(
			ctx.GetStart().GetLine(),
			ctx.GetStart().GetColumn(),
			"No se puede hacer una asignación -- a un slice: "+id,
			SemanticError,
		)
		panic("No se puede hacer una asignación -- a un slice: " + id)
	}

	variable, ok := v.Env.GetVar(id)

	if !ok {
		//agregar error a la tabla de errores
		v.ErrorTable.AddError(
			ctx.GetStart().GetLine(),
			ctx.GetStart().GetColumn(),
			"No se encontro la variable: "+id,
			SemanticError,
		)

		panic("No se encontro la variable: " + id)
	}

	// Verifico si la variable es mutable
	// if !variable.Mut {
	// 	panic("La variable no es mutable: " + id)
	// }

	if variable.Tipo != "int" && variable.Tipo != "f64" {
		//agregar error a la tabla de errores
		v.ErrorTable.AddError(
			ctx.GetStart().GetLine(),
			ctx.GetStart().GetColumn(),
			"Tipo de dato no soportado para la asignación -- a una variable: "+variable.Tipo,
			SemanticError,
		)
		panic(fmt.Sprintf("Tipo de dato no soportado para la asignación -- a una variable: %s", variable.Tipo))
	}

	// Actualizo el valor de la variable en el entorno
	switch variable.Tipo {
	case "int":
		variable.Valor = variable.Valor.(int64) - 1
	case "f64":
		variable.Valor = variable.Valor.(float64) - 1.0
	default:
		//agregar error a la tabla de errores
		v.ErrorTable.AddError(
			ctx.GetStart().GetLine(),
			ctx.GetStart().GetColumn(),
			"Tipo de dato no soportado para la asignación -- a una variable: "+variable.Tipo,
			SemanticError,
		)
		panic("Tipo de variable no soportado para la asignación: " + variable.Tipo)
	}

	return Value{value: true, info: "Assingancion/Reasignacion de Valor en Variable"}
}

// Acceso a Slice y actualizar su valor
func (v *Visitor) VisitSliceAssignmentExpr(ctx *parser.SliceAssignmentExprContext) Value {
	// En el ambito global solo se permiten asignaciones
	if v.Env.name == "Global" {
		//agregar el error a la tabla
		v.ErrorTable.AddError(
			ctx.GetStart().GetLine(),
			ctx.GetStart().GetColumn(),
			"No se puede hacer una asignación en el ámbito global, solo se permiten declaraciones",
			SemanticError,
		)
		panic("No se puede hacer una asignación en el ámbito global, solo se permiten declaraciones")
	}

	id := ctx.ID().GetText()

	// Busco el slice
	slice, ok := v.Env.GetSlice(id)
	if !ok {
		//agregar error a la tabla de errores
		v.ErrorTable.AddError(
			ctx.GetStart().GetLine(),
			ctx.GetStart().GetColumn(),
			"No se encontro el slice: "+id,
			SemanticError,
		)

		panic("No se encontro el slice: " + id)
	}

	// Visito la expresión para obtener el índice
	index := v.Visit(ctx.Expression(0))
	if index.info != "int" {
		//agregar error a la tabla de errores
		v.ErrorTable.AddError(
			ctx.GetStart().GetLine(),
			ctx.GetStart().GetColumn(),
			"El índice debe ser de tipo int, se obtuvo: "+index.info,
			SemanticError,
		)

		panic(fmt.Sprintf("El índice debe ser de tipo int, se obtuvo: %s", index.info))
	}

	// Visito la expresión para obtener el nuevo valor
	newValue := v.Visit(ctx.Expression(1))
	if newValue.info != slice.Tipo {
		//agregar error a la tabla de errores
		v.ErrorTable.AddError(
			ctx.GetStart().GetLine(),
			ctx.GetStart().GetColumn(),
			"El nuevo valor debe ser de tipo "+slice.Tipo+", se obtuvo: "+newValue.info+" en la asignacion del slice "+id,
			SemanticError,
		)
		panic(fmt.Sprintf("El nuevo valor debe ser de tipo %s, se obtuvo: %s, en la asignacion del slice %s", slice.Tipo, newValue.info, id))
	}

	// Actualizo el elemento del slice
	if int(index.value.(int64)) < 0 || int(index.value.(int64)) >= len(slice.Elementos) {
		//agregar error a la tabla de errores
		v.ErrorTable.AddError(
			ctx.GetStart().GetLine(),
			ctx.GetStart().GetColumn(),
			"Índice fuera de rango para el slice: "+id,
			SemanticError,
		)
		panic("Índice fuera de rango")
	}
	slice.Elementos[int(index.value.(int64))] = newValue.value

	return Value{value: true, info: "Asignación/Reasignacion en Slice"}
}

// Print
func (v *Visitor) VisitPrintExpr(ctx *parser.PrintContext) Value {

	// En el ambito global solo se permiten declaraciones
	if v.Env.name == "Global" {
		//agregar el error a la tabla
		v.ErrorTable.AddError(
			ctx.GetStart().GetLine(),
			ctx.GetStart().GetColumn(),
			"No se puede hacer un print en el ámbito global, solo se permiten declaraciones",
			SemanticError,
		)

		panic("No se puede hacer un print en el ámbito global, solo se permiten declaraciones")
	}

	if ctx.PrintList() == nil {
		fmt.Println("")
		return Value{value: nil, info: "Print"}
	}

	expressions := ctx.PrintList().AllExpression() // incluye la primera y las del grupo (, expression)*

	var parts []string
	for _, expr := range expressions {
		val := v.Visit(expr)
		//fmt.Println("HERE IS IN PRINT")
		//fmt.Println("Valor:", val.value, "Tipo:", val.info)
		if val.info == "int" || val.info == "f64" || val.info == "string" || val.info == "bool" {
			parts = append(parts, fmt.Sprintf("%v", val.value))
		} else if val.info == "slice" {
			slice := val.value.(*Slice)
			// Convierto los elementos del slice a string
			var sliceParts []string
			for _, elem := range slice.Elementos {
				sliceParts = append(sliceParts, fmt.Sprintf("%v", elem))
			}
			parts = append(parts, fmt.Sprintf("[%s]", strings.Join(sliceParts, ", ")))
		} else if val.info == "struct" {
			structIns := val.value.(*StructInstance)
			// Convierto los campos del struct a string
			var structParts []string
			for campo, elem := range structIns.Campos {
				structParts = append(structParts, fmt.Sprintf("%s: %v", campo, elem))
			}
			parts = append(parts, fmt.Sprintf("%s{%s}", structIns.Nombre, strings.Join(structParts, ", ")))
		} else {
			fmt.Println("Tipo no soportado en Print:", val.info)
			fmt.Println("Valor:", val.value)
			//agregar error a la tabla de errores
			v.ErrorTable.AddError(
				ctx.GetStart().GetLine(),
				ctx.GetStart().GetColumn(),
				"Tipo de dato no soportado para imprimir: "+val.info,
				SemanticError,
			)
			panic(fmt.Sprintf("Tipo de dato no soportado para imprimir: %s", val.info))
		}
	}

	// fmt.Println(strings.Join(parts, " "))
	v.Console += strings.Join(parts, " ") + "\n"
	return Value{value: nil, info: "Print"}

}

// If
func (v *Visitor) VisitIf(ctx *parser.IfContext) Value {

	// En el ambito global solo se permiten declaraciones
	if v.Env.name == "Global" {
		//agregar el error a la tabla
		v.ErrorTable.AddError(
			ctx.GetStart().GetLine(),
			ctx.GetStart().GetColumn(),
			"No se puede hacer un if en el ámbito global, solo se permiten declaraciones",
			SemanticError,
		)
		panic("No se puede hacer un if en el ámbito global, solo se permiten declaraciones")
	}

	// Entonrno para el if
	previousEnv := v.Env
	v.Env = NewEnv(previousEnv, "If")
	defer func() { v.Env = previousEnv }()
	// Al salir del if, regreso al entorno anterior
	// --------------

	cond := v.Visit(ctx.Expression())

	if cond.info != "bool" {
		//agregar error a la tabla de errores
		v.ErrorTable.AddError(
			ctx.GetStart().GetLine(),
			ctx.GetStart().GetColumn(),
			"La condición del if debe ser de tipo booleano, se obtuvo: "+cond.info,
			SemanticError,
		)
		panic("La condición del if debe ser booleana")
	}

	if cond.value.(bool) {
		// ejecuto el primer if
		v.Visit(ctx.Block())
		return Value{value: nil, info: "if"}
	}

	// Procesar los else if (si existen)
	for _, elseif := range ctx.AllElseIfBlock() {
		cond := v.Visit(elseif.Expression())
		if cond.info != "bool" {
			//agregar error a la tabla de errores
			v.ErrorTable.AddError(
				ctx.GetStart().GetLine(),
				ctx.GetStart().GetColumn(),
				"La condición del else if debe ser de tipo booleano, se obtuvo: "+cond.info,
				SemanticError,
			)
			panic("La condición del else if debe ser de tipo booleano")
		}
		if cond.value.(bool) {
			v.Visit(elseif.Block())
			return Value{value: nil, info: "if"}
		}
	}

	// Procesar else final (si existe)
	if ctx.ElseBlock() != nil {
		v.Visit(ctx.ElseBlock().Block())
		return Value{value: nil, info: "if"}
	}

	return Value{value: nil, info: "if"}
}

// Switch
func (v *Visitor) VisitSwitch(ctx *parser.SwitchContext) Value {

	// En el ambito global solo se permiten declaraciones
	if v.Env.name == "Global" {
		//agregar el error a la tabla
		v.ErrorTable.AddError(
			ctx.GetStart().GetLine(),
			ctx.GetStart().GetColumn(),
			"No se puede hacer un switch en el ámbito global, solo se permiten declaraciones",
			SemanticError,
		)

		panic("No se puede hacer un switch en el ámbito global, solo se permiten declaraciones")
	}

	// Entonrno para el switch
	previousEnv := v.Env
	v.Env = NewEnv(previousEnv, "Switch")
	defer func() {
		v.Env = previousEnv
		v.ControlContext = v.ControlContext[:len(v.ControlContext)-1]
	}()
	// Al salir del if, regreso al entorno anterior
	// --------------

	// Indico que estoy dentro de un switch
	v.ControlContext = append(v.ControlContext, "switch")

	// Obtengo el valor de la expresión del switch
	exprValue := v.Visit(ctx.Expression())
	// tiene que ser un tipo primitivo, int, f64, string o bool
	// si no, no lo puedo evaluar

	tipoPrimitivo := exprValue.info

	if tipoPrimitivo != "int" && tipoPrimitivo != "f64" && tipoPrimitivo != "string" && tipoPrimitivo != "bool" {
		//agregar error a la tabla de errores
		v.ErrorTable.AddError(
			ctx.GetStart().GetLine(),
			ctx.GetStart().GetColumn(),
			" Tipo de expresión del switch no soportado: "+tipoPrimitivo,
			SemanticError,
		)
		panic(fmt.Sprintf("Tipo de expresión del switch no soportado: %s", tipoPrimitivo))
	}

	// Recorro los case blocks
	for _, caseBlock := range ctx.AllCaseBlock() {
		caseValue := v.Visit(caseBlock.Expression())
		if caseValue.info != tipoPrimitivo {
			//agregar error a la tabla de errores
			v.ErrorTable.AddError(
				caseBlock.GetStart().GetLine(),
				caseBlock.GetStart().GetColumn(),
				"Tipo de case no coincide con el tipo del switch: "+caseValue.info+" != "+tipoPrimitivo,
				SemanticError,
			)
			panic(fmt.Sprintf("Tipo de case %s no coincide con el tipo del switch %s", caseValue.info, tipoPrimitivo))
		}
		if caseValue.value == exprValue.value {

			// Capturo el break
			// en este caso paturo el panic dentro del bloque que esta por ejecutarse
			// si es de tipo BreakSignal, lo ignoro, osea no hacemos nada
			// si es otro tipo de panic, lo relanzo, mostrando el error
			defer func() {
				if r := recover(); r != nil {
					if _, ok := r.(BreakSignal); !ok {
						panic(r) // si no es un break, relanzo
					}
				}
			}()

			v.Visit(caseBlock.Block())

			return Value{value: nil, info: "Switch"}
		}
	}

	// Si hay un default block, lo ejecuto
	if ctx.DefaultBlock() != nil {
		v.Visit(ctx.DefaultBlock().Block())
	}

	return Value{value: nil, info: "Switch"}
}

// For
func (v *Visitor) VisitForSimple(ctx *parser.ForSimpleContext) Value {

	//: FOR expression '{' block '}' #ForSimple , tomar el valor de la expresión como condición del for
	if v.Env.name == "Global" {
		//agregar el error a la tabla
		v.ErrorTable.AddError(
			ctx.GetStart().GetLine(),
			ctx.GetStart().GetColumn(),
			"No se puede hacer un for en el ámbito global, solo se permiten declaraciones",
			SemanticError,
		)
		panic("No se puede hacer un for en el ámbito global, solo se permiten declaraciones")
	}
	// Entonrno para el for
	previousEnv := v.Env
	v.Env = NewEnv(previousEnv, "For")
	defer func() {
		v.Env = previousEnv
		v.ControlContext = v.ControlContext[:len(v.ControlContext)-1]
	}()
	// Al salir del for, regreso al entorno anterior
	v.ControlContext = append(v.ControlContext, "for") // Indico que estoy dentro de un for
	// --------------
	// Capturo el break
	defer func() {
		if r := recover(); r != nil {
			if _, ok := r.(BreakSignal); !ok {
				panic(r) // si no es un break, relanzo
			}
		}
	}()
	/*
		// Capturo el continue
		defer func() {
			if r := recover(); r != nil {
				if _, ok := r.(ContinueSignal); !ok {
					panic(r) // si no es un continue, relanzo
				}
			}
		}()
	*/
	// Al salir del for, regreso al entorno anterior
	// --------------
	// Evaluo la expresión del for
	condValue := v.Visit(ctx.Expression())
	if condValue.info != "bool" {
		//agregar error a la tabla de errores
		v.ErrorTable.AddError(
			ctx.GetStart().GetLine(),
			ctx.GetStart().GetColumn(),
			"Tipo de expresión del for no soportado: "+condValue.info+", se esperaba un booleano",
			SemanticError,
		)
		panic(fmt.Sprintf("Tipo de expresión del for no soportado: %s, se esperaba un booleano", condValue.info))
	}
	// Mientras la condición sea verdadera, ejecuto el bloque
	for condValue.value.(bool) {
		// Capturo el continue
		func() {
			defer func() {
				if r := recover(); r != nil {
					switch r.(type) {
					case BreakSignal:
						// Rompemos completamente el ciclo
						panic(r)
					case ContinueSignal:
						// Interrumpimos solo esta iteración
						return
					default:
						panic(r)
					}
				}
			}()

			v.Visit(ctx.Block())
		}()

		// Evaluo la expresión del for nuevamente
		condValue = v.Visit(ctx.Expression())
		if condValue.info != "bool" {
			//agregar error a la tabla de errores
			v.ErrorTable.AddError(
				ctx.GetStart().GetLine(),
				ctx.GetStart().GetColumn(),
				"Tipo de expresión del for no soportado: "+condValue.info+", se esperaba un booleano",
				SemanticError,
			)
			panic(fmt.Sprintf("Tipo de expresión del for no soportado: %s, se esperaba un booleano", condValue.info))
		}
	}
	return Value{value: nil, info: "ForSimple"}
}

//FOR declaration ';' expression ';' assignment? '{' block '}' #ForDeclaration

func (v *Visitor) VisitForDeclaration(ctx *parser.ForDeclarationContext) Value {
	// En el ambito global no se permite el for
	if v.Env.name == "Global" {
		//agregar el error a la tabla
		v.ErrorTable.AddError(
			ctx.GetStart().GetLine(),
			ctx.GetStart().GetColumn(),
			"No se puede hacer un for en el ámbito global, solo se permiten declaraciones",
			SemanticError,
		)
		panic("No se puede hacer un for en el ámbito global, solo se permiten declaraciones")
	}

	// Entonrno para el for
	previousEnv := v.Env
	v.Env = NewEnv(previousEnv, "For")
	defer func() {
		v.Env = previousEnv
		v.ControlContext = v.ControlContext[:len(v.ControlContext)-1]
	}()
	// Al salir del for, regreso al entorno anterior
	v.ControlContext = append(v.ControlContext, "for") // Indico que estoy dentro de un for
	// --------------
	// Capturo el break
	//--------- esto se agrego en el commit
	defer func() {
		if r := recover(); r != nil {
			if _, ok := r.(BreakSignal); !ok {
				panic(r) // si no es un break, relanzo
			}
		}
	}()
	//----------------

	// Declaración de la variable del for
	if ctx.Declaration() != nil {
		v.Visit(ctx.Declaration())
	}

	// Evaluo la expresión del for
	condValue := v.Visit(ctx.Expression())
	if condValue.info != "bool" {
		//agregar error a la tabla de errores
		v.ErrorTable.AddError(
			ctx.GetStart().GetLine(),
			ctx.GetStart().GetColumn(),
			"Tipo de expresión del for no soportado: "+condValue.info+", se esperaba un booleano",
			SemanticError,
		)
		panic(fmt.Sprintf("Tipo de expresión del for no soportado: %s, se esperaba un booleano", condValue.info))
	}

	for condValue.value.(bool) {

		func() {
			defer func() {
				if r := recover(); r != nil {
					switch r.(type) {
					case BreakSignal:
						panic(r) // salir completamente
					case ContinueSignal:
						// no hacemos nada, simplemente salimos de esta iteración
					default:
						panic(r) // relanzamos si no es break ni continue
					}
				}
			}()

			v.Visit(ctx.Block())
		}()

		// Ejecutar incremento si hay
		if ctx.Assignment() != nil {
			v.Visit(ctx.Assignment())
		}

		// Evaluo la expresión del for nuevamente
		condValue = v.Visit(ctx.Expression())
		if condValue.info != "bool" {
			//agregar error a la tabla de errores
			v.ErrorTable.AddError(
				ctx.GetStart().GetLine(),
				ctx.GetStart().GetColumn(),
				"Tipo de expresión del for no soportado: "+condValue.info+", se esperaba un booleano",
				SemanticError,
			)
			panic(fmt.Sprintf("Tipo de expresión del for no soportado: %s, se esperaba un booleano", condValue.info))
		}

	}
	return Value{value: nil, info: "ForDeclaration"}
}

func (v *Visitor) VisitForInSlice(ctx *parser.ForInSliceContext) Value {
	// En el ambito global no se permite el for
	if v.Env.name == "Global" {
		//agregar el error a la tabla
		v.ErrorTable.AddError(
			ctx.GetStart().GetLine(),
			ctx.GetStart().GetColumn(),
			"No se puede hacer un for en el ámbito global, solo se permiten declaraciones",
			SemanticError,
		)
		panic("No se puede hacer un for en el ámbito global, solo se permiten declaraciones")
	}

	// Entonrno para el for
	previousEnv := v.Env
	v.Env = NewEnv(previousEnv, "For")
	defer func() {
		v.Env = previousEnv
		v.ControlContext = v.ControlContext[:len(v.ControlContext)-1]
	}()
	// Al salir del for, regreso al entorno anterior
	v.ControlContext = append(v.ControlContext, "for") // Indico que estoy dentro de un for
	// --------------

	// Capturo el break
	defer func() {
		if r := recover(); r != nil {
			if _, ok := r.(BreakSignal); !ok {
				panic(r) // si no es un break, relanzo
			}
		}
	}()
	/*
		// Capturo el continue
		defer func() {
			if r := recover(); r != nil {
				if _, ok := r.(ContinueSignal); !ok {
					panic(r) // si no es un continue, relanzo
				}
			}
		}()
	*/
	// ...código existente...
	sliceValue := v.Visit(ctx.Expression())
	// ...código existente...

	if sliceValue.info != "slice" {
		//agregar error a la tabla de errores
		v.ErrorTable.AddError(
			ctx.GetStart().GetLine(),
			ctx.GetStart().GetColumn(),
			"Tipo de expresión del for no soportado: "+sliceValue.info+", se esperaba un slice",
			SemanticError,
		)

		panic(fmt.Sprintf("Tipo de expresión del for no soportado: %s, se esperaba un slice", sliceValue.info))
	}

	slice := sliceValue.value.(*Slice)

	variableNames := []string{
		ctx.GetToken(parser.GramaticaParserID, 0).GetText(),
		ctx.GetToken(parser.GramaticaParserID, 1).GetText(),
	}

	for i, elem := range slice.Elementos {

		func() {
			defer func() {
				if r := recover(); r != nil {
					switch r.(type) {
					case BreakSignal:
						panic(r) // salir completamente del for
					case ContinueSignal:
						// ignoramos para continuar siguiente iteración
					default:
						panic(r)
					}
				}
			}()

			newEnv := NewEnv(v.Env, "ForInSlice")

			for j, varName := range variableNames {
				var valor interface{}
				var tipo string

				if j == 0 {
					valor = int64(i)
					tipo = "int"
				} else {
					valor = elem
					tipo = slice.Tipo
				}

				newVar := &Variable{
					Nombre: varName,
					Tipo:   tipo,
					Valor:  valor,
					Mut:    true,
				}
				newEnv.SetVar(varName, newVar)
			}

			// Cambio entorno temporalmente
			prev := v.Env
			v.Env = newEnv
			defer func() { v.Env = prev }()

			v.Visit(ctx.Block())
		}()
	}

	return Value{value: nil, info: "ForInSlice"}

}

// Bloque Individual
func (v *Visitor) VisitIndividualBlock(ctx *parser.IndividualBlockContext) Value {
	if v.Env.name == "Global" {
		//agregar el error a la tabla
		v.ErrorTable.AddError(
			ctx.GetStart().GetLine(),
			ctx.GetStart().GetColumn(),
			"No se puede hacer un bloque individual en el ámbito global, solo se permiten declaraciones",
			SemanticError,
		)
	}

	previousEnv := v.Env
	v.Env = NewEnv(previousEnv, "IndividualBlock")
	defer func() { v.Env = previousEnv }()

	v.Visit(ctx.Block())
	return Value{value: nil, info: "IndividualBlock"}
}

// Especial para saber si estoy dentro de un loop o switch
func (v *Visitor) isInLoopOrSwitch() bool {
	for i := len(v.ControlContext) - 1; i >= 0; i-- {
		if v.ControlContext[i] == "for" || v.ControlContext[i] == "switch" {
			return true
		}
	}
	return false
}

// Break
func (v *Visitor) VisitBreakStatement(ctx *parser.BreakStatementContext) Value {
	// En el ambito global no se permite el break
	if v.Env.name == "Global" {
		//agregar el error a la tabla
		v.ErrorTable.AddError(
			ctx.GetStart().GetLine(),
			ctx.GetStart().GetColumn(),
			"No se puede hacer un break en el ámbito global, solo se permiten declaraciones",
			SemanticError,
		)

		panic("No se puede hacer un break en el ámbito global, solo se permiten declaraciones")
	}

	if !v.isInLoopOrSwitch() {
		//agregar el error a la tabla
		v.ErrorTable.AddError(
			ctx.GetStart().GetLine(),
			ctx.GetStart().GetColumn(),
			"'break' solo puede usarse dentro de un 'for' o 'switch'",
			SemanticError,
		)

		panic("'break' solo puede usarse dentro de un 'for' o 'switch'")
	}
	panic(BreakSignal{})

	panic(Break) // lanzamos señal para romper
	// Estás lanzando un panic cuyo valor es de tipo BreakSignal. Luego
	// , cuando usás recover() dentro de un defer, podés verificar ese tipo
	// específicamente

	/*
		// Verifico si estoy dentro de un bucle
		if v.Env.name != "For" && v.Env.name != "While" {
			panic("El break solo se puede usar dentro de un bucle")
		}

		// Si estoy dentro de un bucle, simplemente retorno
		return Value{value: nil, info: "Break"}*/
}

// Especial para saber si estoy dentro de un loop
func (v *Visitor) isInLoop() bool {
	for i := len(v.ControlContext) - 1; i >= 0; i-- {
		if v.ControlContext[i] == "for" {
			return true
		}
	}
	return false
}

// Continue
func (v *Visitor) VisitContinueStatement(ctx *parser.ContinueStatementContext) Value {
	// En el ambito global no se permite el continue
	if v.Env.name == "Global" {
		//agregar el error a la tabla
		v.ErrorTable.AddError(
			ctx.GetStart().GetLine(),
			ctx.GetStart().GetColumn(),
			"No se puede hacer un continue en el ámbito global, solo se permiten declaraciones",
			SemanticError,
		)

		panic("No se puede hacer un continue en el ámbito global, solo se permiten declaraciones")
	}
	if !v.isInLoop() {
		//agregar el error a la tabla de errores
		v.ErrorTable.AddError(
			ctx.GetStart().GetLine(),
			ctx.GetStart().GetColumn(),
			"'continue' solo puede usarse dentro de un 'for'",
			SemanticError,
		)
		panic("'continue' solo puede usarse dentro de un 'for'")
	}
	panic(Continue) // lanzás una señal que luego capturás en tu ejecución del for
}

func (v *Visitor) isInFunction() bool {
	for i := len(v.ControlContext) - 1; i >= 0; i-- {
		if v.ControlContext[i] == "function" {
			return true
		}
	}
	return false
}

func (v *Visitor) VisitReturnStatement(ctx *parser.ReturnStatementContext) Value {
	if v.Env.name == "Global" {
		//agregar el error a la tabla
		v.ErrorTable.AddError(
			ctx.GetStart().GetLine(),
			ctx.GetStart().GetColumn(),
			"No se puede hacer un return en el ámbito global, solo se permiten declaraciones",
			SemanticError,
		)
		panic("No se puede hacer un continue en el ámbito global, solo se permiten declaraciones")
	}

	if !v.isInFunction() {
		//agregar el error a la tabla de errores
		v.ErrorTable.AddError(
			ctx.GetStart().GetLine(),
			ctx.GetStart().GetColumn(),
			"'return' solo puede usarse dentro de una función",
			SemanticError,
		)

		panic("'return' solo puede usarse dentro de una función")
	}
	var returnValue interface{} = nil
	var info string = "void"

	if ctx.Expression() != nil {
		value := v.Visit(ctx.Expression())
		returnValue = value.value
		info = value.info
	}

	panic(ReturnSignal{Value: returnValue, info: info})
}

// --------------------------------------------------------------------------------------------

// --------------------------------------------------------------------------------------------
// Expresiones
// --------------------------------------------------------------------------------------------

func (v *Visitor) VisitNotExpr(ctx *parser.NotExprContext) Value {
	val := v.Visit(ctx.Expression())

	if val.info == "bool" {
		return Value{value: !val.value.(bool), info: "bool"} // Negación de un entero
	}

	//agregar error a la tabla de errores
	v.ErrorTable.AddError(
		ctx.GetStart().GetLine(),
		ctx.GetStart().GetColumn(),
		"Operador unario (!) no soportado para el tipo: "+val.info+" con valor "+fmt.Sprintf("%v", val.value),
		SemanticError,
	)
	panic(fmt.Sprintf("Operador unario (!) no soportado para el tipo: %s con valor %v", val.info, val.value))

}

func (v *Visitor) VisitNegExpr(ctx *parser.NegExprContext) Value {
	// Negación de un valor
	val := v.Visit(ctx.Expression())

	if val.info == "int" {
		return Value{value: -val.value.(int64), info: "int"} // Negación de un entero
	}
	if val.info == "f64" {
		return Value{value: -val.value.(float64), info: "f64"} // Negación de un float
	}

	//agregar error a la tabla de errores
	v.ErrorTable.AddError(
		ctx.GetStart().GetLine(),
		ctx.GetStart().GetColumn(),
		"Operador unario (-) no soportado para el tipo: "+val.info+" con valor "+fmt.Sprintf("%v", val.value),
		SemanticError,
	)
	panic(fmt.Sprintf("Operador unario (-) no soportado para el tipo: %s con valor %v", val.info, val.value))

}

func (v *Visitor) VisitRefExpr(ctx *parser.RefExprContext) Value {
	// Retorna una referencia a la variable
	id := ctx.ID().GetText()
	variable, ok := v.Env.GetVar(id)
	if !ok {
		//agregar error a la tabla de errores
		v.ErrorTable.AddError(
			ctx.GetStart().GetLine(),
			ctx.GetStart().GetColumn(),
			"No se encontro la variable a la cual se le quiere aplicar &: "+id,
			SemanticError,
		)
		panic("No se encontro la variable a la cual se le quiere aplicar &: " + id)
	}

	//esto era antes cuando solo eran Variable sin el *Variable
	//return Value{value: &variable, info: "referencia"}
	// Retorna una referencia a la variable
	// aca devuelvo un objeto de tipo *Variable que contiene la direccion de memoria de la variable

	// si le vuelvo a poner el &Vairable estaria haciendo esto **Variable
	// como deje las variables como punteros, no necesito hacer nada especial
	return Value{value: variable, info: "referencia"} // Retorna una referencia a la variable

}

func (v *Visitor) VisitDeRefExpr(ctx *parser.DerefExprContext) Value {
	// Retorna el valor de la variable a la cual se le aplica *
	id := ctx.ID().GetText()

	// en teoria, busca una variable de tipo referencia
	variable, ok := v.Env.GetVar(id)

	if !ok {
		//agregar error a la tabla de errores
		v.ErrorTable.AddError(
			ctx.GetStart().GetLine(),
			ctx.GetStart().GetColumn(),
			"No se encontro la variable a la cual se le quiere aplicar *: "+id,
			SemanticError,
		)

		panic("No se encontro la variable a la cual se le quiere aplicar *: " + id)
	}

	if variable.Tipo != "referencia" {
		//agregar error a la tabla de errores
		v.ErrorTable.AddError(
			ctx.GetStart().GetLine(),
			ctx.GetStart().GetColumn(),
			"Se esperaba una referencia, pero se obtuvo otro tipo: "+variable.Tipo,
			SemanticError,
		)

		panic("Se esperaba una referencia, pero se obtuvo otro tipo: " + variable.Tipo)
	}

	refValue, ok2 := variable.Valor.(*Variable)
	// Desreferencia el valor de la variable

	if !ok2 {
		//agregar error a la tabla de errores
		v.ErrorTable.AddError(
			ctx.GetStart().GetLine(),
			ctx.GetStart().GetColumn(),
			"No se puede desreferenciar un valor que no es una referencia: *"+id,
			SemanticError,
		)

		panic("No se puede desreferenciar un valor que no es una referencia: *" + id)
	}

	return Value{value: refValue.Valor, info: refValue.Tipo} // Retorna el valor de la variable
}

func (v *Visitor) VisitOpExpr(ctx *parser.OpExprContext) Value {
	l := v.Visit(ctx.GetLeft())
	r := v.Visit(ctx.GetRight())
	op := ctx.GetOp().GetText()

	// Puedo operar si son valores primitivos, las aviralbes me retornan su valor
	// valido los tipos de l y r

	// int + int = int
	if l.info == "int" && r.info == "int" {
		valorl := l.value.(int64)
		valorr := r.value.(int64)
		switch op {
		case "+":
			return Value{value: valorl + valorr, info: "int"}
		case "-":
			return Value{value: valorl - valorr, info: "int"}
		case "*":
			return Value{value: valorl * valorr, info: "int"}
		case "/":
			if valorr == 0 {
				//agregar error a la tabla de errores
				v.ErrorTable.AddError(
					ctx.GetStart().GetLine(),
					ctx.GetStart().GetColumn(),
					"División por cero "+fmt.Sprintf("%d / %d", valorl, valorr),
					SemanticError,
				)
				panic("División por cero " + fmt.Sprintf("%d / %d", valorl, valorr))
			}

			//return Value{value: valorl / valorr, info: "int"}

			// con esto veo si la division es exacta o no, si tiene decimales o no, y depende de eso retorno un int o un float64
			if valorl%valorr == 0 {
				return Value{value: valorl / valorr, info: "int"}
			}
			return Value{value: float64(valorl) / float64(valorr), info: "f64"}
		case "%":
			if valorr == 0 {
				//agregar error a la tabla de errores
				v.ErrorTable.AddError(
					ctx.GetStart().GetLine(),
					ctx.GetStart().GetColumn(),
					"Módulo por cero "+fmt.Sprintf("%d %% %d", valorl, valorr),
					SemanticError,
				)
				panic("Módulo por cero " + fmt.Sprintf("%d %% %d", valorl, valorr))
			}
			return Value{value: valorl % valorr, info: "int"}
		case "<":
			return Value{value: valorl < valorr, info: "bool"}
		case "<=":
			return Value{value: valorl <= valorr, info: "bool"}
		case ">":
			return Value{value: valorl > valorr, info: "bool"}
		case ">=":
			return Value{value: valorl >= valorr, info: "bool"}
		case "==":
			return Value{value: valorl == valorr, info: "bool"}
		case "!=":
			return Value{value: valorl != valorr, info: "bool"}
		}
	}

	// int + f64 = f64  (o f64 + int)
	if (l.info == "int" && r.info == "f64") || (l.info == "f64" && r.info == "int") {
		var left, right float64
		if l.info == "int" {
			left = float64(l.value.(int64))
			right = r.value.(float64)
		} else {
			left = l.value.(float64)
			right = float64(r.value.(int64))
		}

		switch op {
		case "+":
			return Value{value: left + right, info: "f64"}
		case "-":
			return Value{value: left - right, info: "f64"}
		case "*":
			return Value{value: left * right, info: "f64"}
		case "/":
			if right == 0 {
				//agregar error a la tabla de errores
				v.ErrorTable.AddError(
					ctx.GetStart().GetLine(),
					ctx.GetStart().GetColumn(),
					"División por cero "+fmt.Sprintf("%f / %f", left, right),
					SemanticError,
				)
				panic("División por cero " + fmt.Sprintf("%f / %f", left, right))
			}
			return Value{value: left / right, info: "f64"}
		case "<":
			return Value{value: left < right, info: "bool"}
		case "<=":
			return Value{value: left <= right, info: "bool"}
		case ">":
			return Value{value: left > right, info: "bool"}
		case ">=":
			return Value{value: left >= right, info: "bool"}
		case "==":
			return Value{value: left == right, info: "bool"}
		case "!=":
			return Value{value: left != right, info: "bool"}
		}
	}

	// f64 + f64 = f64
	if l.info == "f64" && r.info == "f64" {
		valorl := l.value.(float64)
		valorr := r.value.(float64)
		switch op {
		case "+":
			return Value{value: valorl + valorr, info: "f64"}
		case "-":
			return Value{value: valorl - valorr, info: "f64"}
		case "*":
			return Value{value: valorl * valorr, info: "f64"}
		case "/":
			if valorr == 0 {
				//agregar error a la tabla de errores
				v.ErrorTable.AddError(
					ctx.GetStart().GetLine(),
					ctx.GetStart().GetColumn(),
					"División por cero "+fmt.Sprintf("%f / %f", valorl, valorr),
					SemanticError,
				)
				panic("División por cero " + fmt.Sprintf("%f / %f", valorl, valorr))
			}
			return Value{value: valorl / valorr, info: "f64"}
		case "<":
			return Value{value: valorl < valorr, info: "bool"}
		case "<=":
			return Value{value: valorl <= valorr, info: "bool"}
		case ">":
			return Value{value: valorl > valorr, info: "bool"}
		case ">=":
			return Value{value: valorl >= valorr, info: "bool"}
		case "==":
			return Value{value: valorl == valorr, info: "bool"}
		case "!=":
			return Value{value: valorl != valorr, info: "bool"}
		}
	}

	// string + string = string
	if l.info == "string" && r.info == "string" {
		valorl := l.value.(string)
		valorr := r.value.(string)
		switch op {
		case "+":
			return Value{value: valorl + valorr, info: "string"}
		case "<":
			return Value{value: valorl < valorr, info: "bool"}
		case "<=":
			return Value{value: valorl <= valorr, info: "bool"}
		case ">":
			return Value{value: valorl > valorr, info: "bool"}
		case ">=":
			return Value{value: valorl >= valorr, info: "bool"}
		case "==":
			return Value{value: valorl == valorr, info: "bool"}
		case "!=":
			return Value{value: valorl != valorr, info: "bool"}
		}
	}

	// bool con bool = bool
	if l.info == "bool" && r.info == "bool" {
		valorl := l.value.(bool)
		valorr := r.value.(bool)
		switch op {
		case "==":
			return Value{value: valorl == valorr, info: "bool"}
		case "!=":
			return Value{value: valorl != valorr, info: "bool"}
		case "&&":
			return Value{value: valorl && valorr, info: "bool"}
		case "||":
			return Value{value: valorl || valorr, info: "bool"}
		}
	}

	//agregar error a la tabla de errores
	v.ErrorTable.AddError(
		ctx.GetStart().GetLine(),
		ctx.GetStart().GetColumn(),
		"Tipos no soportados para la operación: "+op+" entre "+l.info+" y "+r.info,
		SemanticError,
	)
	panic("Tipos no soportados para la operación: " + op + " entre " + l.info + " y " + r.info)

}

func (v *Visitor) VisitParExpr(ctx *parser.ParExprContext) Value {
	// Visita la expresión dentro de los paréntesis
	val := v.Visit(ctx.Expression())
	return Value{value: val.value, info: val.info}
}

// Exprsiones de Slices

func (v *Visitor) VisitSlicePrimitiveValue(ctx *parser.SlicePrimitiveValueContext) Value {
	// use un alis de SlicePrimtiveValue
	// el No terminal tiene de nombre SlicePrimitiveValueFunction
	// entonces ahi tengo que entrar
	return v.Visit(ctx.SlicePrimitiveValueFunction())
}

func (v *Visitor) VisitIndexOfFunction(ctx *parser.IndexOfFunctionContext) Value {
	id := ctx.ID().GetText()

	// Busco el slice
	slice, ok := v.Env.GetSlice(id)
	if !ok {
		//agregar error a la tabla de errores
		v.ErrorTable.AddError(
			ctx.GetStart().GetLine(),
			ctx.GetStart().GetColumn(),
			"No se encontro el slice: "+id,
			SemanticError,
		)

		panic("No se encontro el slice: " + id)
	}

	// Visito la expresión para obtener el valor a buscar
	valueToFind := v.Visit(ctx.Expression())

	if valueToFind.info != slice.Tipo {
		//agregar error a la tabla de errores
		v.ErrorTable.AddError(
			ctx.GetStart().GetLine(),
			ctx.GetStart().GetColumn(),
			"Tipo de dato "+valueToFind.info+" no coincide con el tipo del slice "+slice.Tipo,
			SemanticError,
		)
		panic(fmt.Sprintf("Tipo de dato %s no coincide con el tipo del slice %s", valueToFind.info, slice.Tipo))
	}

	// Busco el índice del valor en el slice
	for i, elem := range slice.Elementos {
		if elem == valueToFind.value {
			return Value{value: int64(i), info: "int"} // Retorno el índice como int
		}
	}

	return Value{value: int64(-1), info: "int"} // Retorno -1 si no se encuentra
}

func (v *Visitor) VisitJoinFunction(ctx *parser.JoinFunctionContext) Value {
	id := ctx.ID().GetText()

	// Busco el slice
	slice, ok := v.Env.GetSlice(id)
	if !ok {
		//agregar error a la tabla de errores
		v.ErrorTable.AddError(
			ctx.GetStart().GetLine(),
			ctx.GetStart().GetColumn(),
			"No se encontro el slice: "+id,
			SemanticError,
		)
		panic("No se encontro el slice: " + id)
	}

	// SOlo puedo hacer si el slice es de tipo string
	if slice.Tipo != "string" {
		//agregar error a la tabla de errores
		v.ErrorTable.AddError(
			ctx.GetStart().GetLine(),
			ctx.GetStart().GetColumn(),
			"El tipo de slice "+slice.Tipo+" no es soportado para la función join, se esperaba un slice de tipo string",
			SemanticError,
		)
		panic(fmt.Sprintf("El tipo de slice %s no es soportado para la función join, se esperaba un slice de tipo string", slice.Tipo))
	}

	// veo si el separador es un string
	separador := v.Visit(ctx.Expression())
	if separador.info != "string" {
		//agregar error a la tabla de errores
		v.ErrorTable.AddError(
			ctx.GetStart().GetLine(),
			ctx.GetStart().GetColumn(),
			"El separador debe ser de tipo string, se obtuvo: "+separador.info,
			SemanticError,
		)
		panic(fmt.Sprintf("El separador debe ser de tipo string, se obtuvo: %s", separador.info))
	}

	var parts []string
	for _, elem := range slice.Elementos {
		parts = append(parts, fmt.Sprintf("%v", elem))

	}

	// Retorno el resultado como un string
	return Value{value: strings.Join(parts, separador.value.(string)), info: "string"}
}

func (v *Visitor) VisitLengthFunction(ctx *parser.LengthFunctionContext) Value {
	id := ctx.ID().GetText()

	// Busco el slice
	slice, ok := v.Env.GetSlice(id)
	if !ok {
		//agregar error a la tabla de errores
		v.ErrorTable.AddError(
			ctx.GetStart().GetLine(),
			ctx.GetStart().GetColumn(),
			"No se encontro el slice: "+id,
			SemanticError,
		)
		panic("No se encontro el slice: " + id)
	}

	// Retorno la longitud del slice
	return Value{value: int64(len(slice.Elementos)), info: "int"}
}

func (v *Visitor) VisitAppendExpr(ctx *parser.AppendExprContext) Value {
	id := ctx.ID().GetText()

	// Busco el slice
	slice, ok := v.Env.GetSlice(id)
	if !ok {
		//agregar error a la tabla de errores
		v.ErrorTable.AddError(
			ctx.GetStart().GetLine(),
			ctx.GetStart().GetColumn(),
			"No se encontro el slice: "+id,
			SemanticError,
		)

		panic("No se encontro el slice: " + id)
	}

	// Verifico si el slice es mutable
	// if !slice.Mut {
	// 	panic("El slice no es mutable: " + id)
	// }

	// Visito la expresión para obtener el valor a agregar
	valueToAppend := v.Visit(ctx.Expression())

	if valueToAppend.info != slice.Tipo {
		//agregar error a la tabla de errores
		v.ErrorTable.AddError(
			ctx.GetStart().GetLine(),
			ctx.GetStart().GetColumn(),
			"Tipo de dato "+valueToAppend.info+" no coincide con el tipo del slice "+slice.Tipo,
			SemanticError,
		)
		panic(fmt.Sprintf("Tipo de dato %s no coincide con el tipo del slice %s", valueToAppend.info, slice.Tipo))
	}

	// Agrego el valor al slice
	// slice.Elementos = append(slice.Elementos, valueToAppend.value)

	// se crea uno nuevo para no modificar el original
	// ya sea si se lo vuelvo a asiganar al original o a otro slice
	nuevoSlice := &Slice{
		Nombre:    slice.Nombre,
		Tipo:      slice.Tipo,
		Elementos: append(slice.Elementos, valueToAppend.value), // copia + append
		Mut:       slice.Mut,
	}

	return Value{value: nuevoSlice, info: "slice"}
}

// ---------------------

func (v *Visitor) VisitSliceAccess(ctx *parser.SliceAccessContext) Value {
	id := ctx.ID().GetText()

	// Busco el slice
	slice, ok := v.Env.GetSlice(id)
	if !ok {
		//agregar error a la tabla de errores
		v.ErrorTable.AddError(
			ctx.GetStart().GetLine(),
			ctx.GetStart().GetColumn(),
			"No se encontro el slice: "+id,
			SemanticError,
		)
		panic("No se encontro el slice: " + id)
	}

	// Visito la expresión para obtener el índice
	index := v.Visit(ctx.Expression())
	if index.info != "int" {
		//agregar error a la tabla de errores
		v.ErrorTable.AddError(
			ctx.GetStart().GetLine(),
			ctx.GetStart().GetColumn(),
			"El índice debe ser de tipo int, se obtuvo: "+index.info,
			SemanticError,
		)
		panic(fmt.Sprintf("El índice debe ser de tipo int, se obtuvo: %s", index.info))
	}

	// Accedo al elemento del slice
	if int(index.value.(int64)) < 0 || int(index.value.(int64)) >= len(slice.Elementos) {
		//agregar error a la tabla de errores
		v.ErrorTable.AddError(
			ctx.GetStart().GetLine(),
			ctx.GetStart().GetColumn(),
			"Índice fuera de rango para el slice "+id+": "+fmt.Sprintf("%d", index.value.(int64)),
			SemanticError,
		)

		panic("Índice fuera de rango")
	}

	return Value{value: slice.Elementos[int(index.value.(int64))], info: slice.Tipo}
}

func (v *Visitor) VisitFunctionCallExpr(ctx *parser.FunctionCallExprContext) (result Value) {

	funcName := ctx.ID().GetText()

	// Busco la función
	function, ok := v.Env.GetFunc(funcName)
	if !ok {
		// agrego el error a la tabla
		v.ErrorTable.AddError(
			ctx.GetStart().GetLine(),
			ctx.GetStart().GetColumn(),
			"No se encontro la función: "+funcName,
			SemanticError,
		)
		panic("No se encontro la función: " + funcName)
	}

	// veo si hay argumentos y si la cantidad es igual
	// a la cantidad de parámetros de la función

	// Verifico los argumentos

	var args []Value
	if ctx.Arguments() != nil {
		for _, arg := range ctx.Arguments().AllExpression() {
			args = append(args, v.Visit(arg))
		}
	}

	if len(args) != len(function.Parametros) {
		//agregar error a la tabla de errores
		v.ErrorTable.AddError(
			ctx.GetStart().GetLine(),
			ctx.GetStart().GetColumn(),
			"Cantidad de argumentos no coincide con la cantidad de parámetros de la función "+funcName+": "+fmt.Sprintf("%d != %d", len(args), len(function.Parametros)),
			SemanticError,
		)
		panic(fmt.Sprintf("Cantidad de argumentos (%d) no coincide con la cantidad de parámetros de la función %s (%d)", len(args), funcName, len(function.Parametros)))
	}

	result = Value{value: nil, info: "void"} // valor por defecto o inicial de retorno

	// veo si tego que retornar algo
	if function.TipoRetorno == "void" {
		// aca no retorno nada, solo ejecuto la función
		// ejecuto la funcion donde fue declarada y le creo
		// las variables, slices o structs que le pasaron como argumentos

		// creo un nuevo entorno para la función

		callEnv := v.Env // el entonrno donde se esta llamando la función

		creationEnv := function.Parent // el entorno donde fue declarada la función

		// tambien creamos una nueva pila de contextos dentro
		// para tenerla limpia y no meter la pila del entorno anterior
		// osea la funcion tenga una pila de contextos limpia
		contextosPrev := v.ControlContext
		v.ControlContext = make([]string, 0) // reinicio la pila de contextos

		v.Env = NewEnv(creationEnv, "Function")

		defer func() {
			v.Env = callEnv
			v.ControlContext = v.ControlContext[:len(v.ControlContext)-1]
			v.ControlContext = contextosPrev // vuelvo a la pila de contextos anterior
			if r := recover(); r != nil {
				if r, ok := r.(ReturnSignal); ok {
					//panic("Una función no puede retornar un valor si su tipo de retorno es void")
					// if r.info != "void" {
					// 	panic(fmt.Sprintf("La función '%s' es de tipo void, pero retornó un valor de tipo '%s'", funcName, r.info))
					// }
					result.info = "void" // si es un return, lo cambio a void
					result.value = nil   // y el valor a nil
				} else {
					panic(r) // relanzo si no es return
				}
			}
		}()

		// Indico que estoy dentro de un funcion
		v.ControlContext = append(v.ControlContext, "function")

		// Asignar parámetros
		for i, param := range function.Parametros {
			paramValue := args[i]
			if paramValue.info != param.Tipo {
				//agregar error a la tabla de errores
				v.ErrorTable.AddError(
					ctx.GetStart().GetLine(),
					ctx.GetStart().GetColumn(),
					"Tipo del argumento '"+param.Nombre+"' ("+paramValue.info+") no coincide con el parámetro de tipo "+param.Tipo,
					SemanticError,
				)
				panic(fmt.Sprintf("Tipo del argumento '%s' (%s) no coincide con el parámetro de tipo %s", param.Nombre, paramValue.info, param.Tipo))
			}

			if paramValue.info == "int" || paramValue.info == "f64" || paramValue.info == "string" || paramValue.info == "bool" {
				v.Env.SetVar(param.Nombre, &Variable{Nombre: param.Nombre, Tipo: param.Tipo, Valor: paramValue.value, Mut: true})
			}

			if paramValue.info == "slice" {

				// veo si los tipos coinciden
				if paramValue.value.(*Slice).Tipo != param.TipoSlice {
					//agregar error a la tabla de errores
					v.ErrorTable.AddError(
						ctx.GetStart().GetLine(),
						ctx.GetStart().GetColumn(),
						"Tipo del argumento '"+param.Nombre+"' ("+paramValue.value.(*Slice).Tipo+") no coincide con el parámetro de tipo "+param.TipoSlice,
						SemanticError,
					)
					panic(fmt.Sprintf("Tipo del argumento '%s' (%s) no coincide con el parámetro de tipo %s", param.Nombre, paramValue.value.(*Slice).Tipo, param.TipoSlice))
				}

				// si es un slice, lo guardo como un slice
				sliceValue := paramValue.value.(*Slice)
				// creo una copia ya que aca le paso la direccion de memoria
				// y no quiero que se modifique el original
				newSlice := &Slice{
					Nombre:    sliceValue.Nombre,
					Tipo:      sliceValue.Tipo,
					Elementos: make([]interface{}, len(sliceValue.Elementos)),
					Mut:       sliceValue.Mut,
				}
				// Copiar los elementos del slice original al nuevo slice
				copy(newSlice.Elementos, sliceValue.Elementos)
				v.Env.SetSlice(param.Nombre, newSlice) // Asigno el slice al entorno
			}

		}

		v.Visit(function.Cuerpo)

	}

	// Si la función tiene un tipo de retorno, obtengo el valor de retorno
	if function.TipoRetorno != "void" {
		// creo un nuevo entorno para la función

		callEnv := v.Env // el entonrno donde se esta llamando la función

		creationEnv := function.Parent // el entorno donde fue declarada la función

		contextosPrev := v.ControlContext
		v.ControlContext = make([]string, 0) // reinicio la pila de contextos

		v.Env = NewEnv(creationEnv, "Function")
		defer func() {
			v.Env = callEnv
			v.ControlContext = v.ControlContext[:len(v.ControlContext)-1]
			v.ControlContext = contextosPrev // vuelvo a la pila de contextos anterior
			if r := recover(); r != nil {
				if ret, ok := r.(ReturnSignal); ok {
					if ret.info != function.TipoRetorno {
						//agregar error a la tabla de errores
						v.ErrorTable.AddError(
							ctx.GetStart().GetLine(),
							ctx.GetStart().GetColumn(),
							"Tipo de retorno de la función '"+funcName+"' no coincide con el tipo esperado '"+function.TipoRetorno+"'",
							SemanticError,
						)
						panic(fmt.Sprintf("Tipo de retorno de la función '%s' no coincide con el tipo esperado '%s'", funcName, function.TipoRetorno))
					}

					if ret.info == "slice" {
						if ret.Value.(*Slice).Tipo != function.TipoSlice {
							//agregar error a la tabla de errores
							v.ErrorTable.AddError(
								ctx.GetStart().GetLine(),
								ctx.GetStart().GetColumn(),
								"Tipo de retorno de la función '"+funcName+"' no coincide con el tipo esperado slice '"+function.TipoSlice+"'",
								SemanticError,
							)

							panic(fmt.Sprintf("Tipo de retorno de la función '%s' no coincide con el tipo esperado slice '%s'", funcName, function.TipoSlice))
						}
					}

					result = Value{value: ret.Value, info: ret.info}
				} else {
					panic(r) // relanzo si no es return
				}
			}
		}()
		// Indico que estoy dentro de un funcion
		v.ControlContext = append(v.ControlContext, "function")
		// Asignar parámetros
		for i, param := range function.Parametros {
			paramValue := args[i]
			if paramValue.info != param.Tipo {
				//agregar error a la tabla de errores
				v.ErrorTable.AddError(
					ctx.GetStart().GetLine(),
					ctx.GetStart().GetColumn(),
					"Tipo del argumento '"+param.Nombre+"' ("+paramValue.info+") no coincide con el parámetro de tipo "+param.Tipo,
					SemanticError,
				)

				panic(fmt.Sprintf("Tipo del argumento '%s' (%s) no coincide con el parámetro de tipo %s", param.Nombre, paramValue.info, param.Tipo))
			}

			if paramValue.info == "int" || paramValue.info == "f64" || paramValue.info == "string" || paramValue.info == "bool" {
				v.Env.SetVar(param.Nombre, &Variable{Nombre: param.Nombre, Tipo: param.Tipo, Valor: paramValue.value, Mut: true})
			}

			if paramValue.info == "slice" {

				// veo si los tipos coinciden
				if paramValue.value.(*Slice).Tipo != param.TipoSlice {
					//agregar error a la tabla de errores
					v.ErrorTable.AddError(
						ctx.GetStart().GetLine(),
						ctx.GetStart().GetColumn(),
						"Tipo del argumento '"+param.Nombre+"' ("+paramValue.value.(*Slice).Tipo+") no coincide con el parámetro de tipo "+param.TipoSlice,
						SemanticError,
					)
					panic(fmt.Sprintf("Tipo del argumento '%s' (%s) no coincide con el parámetro de tipo %s", param.Nombre, paramValue.value.(*Slice).Tipo, param.TipoSlice))
				}
				// si es un slice, lo guardo como un slice
				sliceValue := paramValue.value.(*Slice)
				// creo una copia ya que aca le paso la direccion de memoria
				// y no quiero que se modifique el original
				newSlice := &Slice{
					Nombre:    sliceValue.Nombre,
					Tipo:      sliceValue.Tipo,
					Elementos: make([]interface{}, len(sliceValue.Elementos)),
					Mut:       sliceValue.Mut,
				}
				// Copiar los elementos del slice original al nuevo slice
				copy(newSlice.Elementos, sliceValue.Elementos)
				v.Env.SetSlice(param.Nombre, newSlice) // Asigno el slice al entorno
			}

		}

		v.Visit(function.Cuerpo)

		//agrego el error a la tabla de errores
		v.ErrorTable.AddError(
			ctx.GetStart().GetLine(),
			ctx.GetStart().GetColumn(),
			"La función '"+funcName+"' no retornó ningún valor",
			SemanticError,
		)

		panic("La función '" + funcName + "' no retornó ningún valor")

	}

	return result
}

func (v *Visitor) VisitAtoiExpr(ctx *parser.AtoiExprContext) Value {
	val := v.Visit(ctx.Expression())
	if val.info != "string" {
		//agregar error a la tabla de errores
		v.ErrorTable.AddError(
			ctx.GetStart().GetLine(),
			ctx.GetStart().GetColumn(),
			"El valor a convertir a entero debe ser de tipo string, se obtuvo: "+val.info,
			SemanticError,
		)
		panic(fmt.Sprintf("El valor a convertir a entero debe ser de tipo string, se obtuvo: %s", val.info))
	}
	// verifico si la cadena es un numero entero, si es decinaml no se pode convertir
	cadena := val.value.(string)
	if numeroConvertido, err := strconv.ParseInt(cadena, 10, 64); err != nil {
		//agregar error a la tabla de errores
		v.ErrorTable.AddError(
			ctx.GetStart().GetLine(),
			ctx.GetStart().GetColumn(),
			"El valor '"+cadena+"' no es un número entero válido para convertir a int",
			SemanticError,
		)
		panic(fmt.Sprintf("El valor '%s' no es un número entero válido", cadena))
	} else {
		return Value{value: numeroConvertido, info: "int"}
	}
}

func (v *Visitor) VisitParseFloatExpr(ctx *parser.ParseFloatExprContext) Value {
	val := v.Visit(ctx.Expression())
	if val.info != "string" {
		//agregar error a la tabla de errores
		v.ErrorTable.AddError(
			ctx.GetStart().GetLine(),
			ctx.GetStart().GetColumn(),
			"El valor a convertir a float debe ser de tipo string, se obtuvo: "+val.info,
			SemanticError,
		)
		panic(fmt.Sprintf("El valor a convertir a float debe ser de tipo string, se obtuvo: %s", val.info))
	}
	// Verifico si la cadena representa un número válido (entero o decimal)
	cadena := val.value.(string)
	if numeroConvertido, err := strconv.ParseFloat(cadena, 64); err != nil {
		//agregar error a la tabla de errores
		v.ErrorTable.AddError(
			ctx.GetStart().GetLine(),
			ctx.GetStart().GetColumn(),
			"El valor '"+cadena+"' no es un número válido para convertir a float",
			SemanticError,
		)
		panic(fmt.Sprintf("El valor '%s' no es un número válido para convertir a float", cadena))
	} else {
		return Value{value: numeroConvertido, info: "f64"}
	}
}

func (v *Visitor) VisitTypeOfExpr(ctx *parser.TypeOfExprContext) Value {
	val := v.Visit(ctx.Expression())
	// Retorna el tipo del valor
	if val.info == "slice" {
		return Value{value: "[]" + val.value.(*Slice).Tipo, info: "string"} // Retorna el tipo del slice
	}

	return Value{value: val.info, info: "string"} // Retorna el tipo del valor
}

func (v *Visitor) VisitIdExpr(ctx *parser.IdExprContext) Value {
	id := ctx.GetText()

	variable, ok := v.Env.GetVar(id)
	if ok {
		return Value{value: variable.Valor, info: variable.Tipo} // Retorna el valor de la variable
	}
	sliceE, ok2 := v.Env.GetSlice(id)
	if ok2 {
		return Value{value: sliceE, info: "slice"} // Retorna el slice
	}
	structS, ok3 := v.Env.GetStructInstance(id)
	if ok3 {
		return Value{value: structS, info: "struct"} // Retorna la instancia del struct
	}

	//agregar error a la tabla de errores
	v.ErrorTable.AddError(
		ctx.GetStart().GetLine(),
		ctx.GetStart().GetColumn(),
		"No se encontro el elemento: "+id,
		SemanticError,
	)
	panic("No se encontro el elemento: " + id)
}

func (v *Visitor) VisitIntExpr(ctx *parser.IntExprContext) Value {
	value, _ := strconv.ParseInt(ctx.GetText(), 10, 64)
	return Value{value: value, info: "int"}
}

func (v *Visitor) VisitFloatExpr(ctx *parser.FloatExprContext) Value {
	value, _ := strconv.ParseFloat(ctx.GetText(), 64)
	return Value{value: value, info: "f64"}
}

func (v *Visitor) VisitBoolExpr(ctx *parser.BoolExprContext) Value {
	value, _ := strconv.ParseBool(ctx.GetText())
	return Value{value: value, info: "bool"}
}

func (v *Visitor) VisitStrExpr(ctx *parser.StrExprContext) Value {
	// value := strings.Trim(ctx.GetText(), "\"")
	// return Value{value: value, info: "string"}
	raw := ctx.GetText()                  // incluye comillas
	unquoted, err := strconv.Unquote(raw) // el strcon.unquote se encarga de interpretar \n, \t, \\, \"
	if err != nil {
		//agregar error a la tabla de errores
		v.ErrorTable.AddError(
			ctx.GetStart().GetLine(),
			ctx.GetStart().GetColumn(),
			"Error al procesar cadena: "+err.Error(),
			SemanticError,
		)
		panic("Error al procesar cadena: " + err.Error())
	}
	return Value{value: unquoted, info: "string"}
}

// --------------------------------------------------------------------------------------------
//
// Llamadas
//
// --------------------------------------------------------------------------------------------

func (v *Visitor) VisitFunctionCall(ctx *parser.FunctionCallContext) (result Value) {
	if v.Env.name == "Global" {
		//agregar error a la tabla de errores
		v.ErrorTable.AddError(
			ctx.GetStart().GetLine(),
			ctx.GetStart().GetColumn(),
			"No se puede hacer una llamada a una función en el ámbito global, solo se permiten declaraciones",
			SemanticError,
		)

		panic("No se puede hacer una llamada a una función en el ámbito global, solo se permiten declaraciones")
	}

	funcName := ctx.ID().GetText()

	// Busco la función
	function, ok := v.Env.GetFunc(funcName)
	if !ok {
		// agrego el error a la tabla
		v.ErrorTable.AddError(
			ctx.GetStart().GetLine(),
			ctx.GetStart().GetColumn(),
			"No se encontro la función: "+funcName,
			SemanticError,
		)
		return Value{value: nil, info: "void"} // Retorno un valor por defecto

		//panic("No se encontro la función: " + funcName)
	}

	// veo si hay argumentos y si la cantidad es igual
	// a la cantidad de parámetros de la función

	// Verifico los argumentos

	var args []Value
	if ctx.Arguments() != nil {
		for _, arg := range ctx.Arguments().AllExpression() {
			args = append(args, v.Visit(arg))
		}
	}

	if len(args) != len(function.Parametros) {
		//agregar error a la tabla de errores
		v.ErrorTable.AddError(
			ctx.GetStart().GetLine(),
			ctx.GetStart().GetColumn(),
			"Cantidad de argumentos no coincide con la cantidad de parámetros de la función "+funcName+": "+fmt.Sprintf("%d != %d", len(args), len(function.Parametros)),
			SemanticError,
		)
		panic(fmt.Sprintf("Cantidad de argumentos (%d) no coincide con la cantidad de parámetros de la función %s (%d)", len(args), funcName, len(function.Parametros)))
	}

	result = Value{value: nil, info: "void"} // valor por defecto o inicial de retorno

	// veo si tego que retornar algo
	if function.TipoRetorno == "void" {
		// aca no retorno nada, solo ejecuto la función
		// ejecuto la funcion donde fue declarada y le creo
		// las variables, slices o structs que le pasaron como argumentos

		// creo un nuevo entorno para la función

		callEnv := v.Env // el entonrno donde se esta llamando la función

		creationEnv := function.Parent // el entorno donde fue declarada la función

		// tambien creamos una nueva pila de contextos dentro
		// para tenerla limpia y no meter la pila del entorno anterior
		// osea la funcion tenga una pila de contextos limpia
		contextosPrev := v.ControlContext
		v.ControlContext = make([]string, 0) // reinicio la pila de contextos

		v.Env = NewEnv(creationEnv, "Function")

		defer func() {
			v.Env = callEnv
			v.ControlContext = v.ControlContext[:len(v.ControlContext)-1]
			v.ControlContext = contextosPrev // vuelvo a la pila de contextos anterior
			if r := recover(); r != nil {
				if r, ok := r.(ReturnSignal); ok {
					//panic("Una función no puede retornar un valor si su tipo de retorno es void")
					// if r.info != "void" {
					// 	panic(fmt.Sprintf("La función '%s' es de tipo void, pero retornó un valor de tipo '%s'", funcName, r.info))
					// }
					result.info = "void" // si es un return, lo cambio a void
					result.value = nil   // y el valor a nil
				} else {
					panic(r) // relanzo si no es return
				}
			}
		}()

		// Indico que estoy dentro de un funcion
		v.ControlContext = append(v.ControlContext, "function")

		// Asignar parámetros
		for i, param := range function.Parametros {
			paramValue := args[i]
			if paramValue.info != param.Tipo {
				//agregar error a la tabla de errores
				v.ErrorTable.AddError(
					ctx.GetStart().GetLine(),
					ctx.GetStart().GetColumn(),
					"Tipo del argumento '"+param.Nombre+"' ("+paramValue.info+") no coincide con el parámetro de tipo "+param.Tipo,
					SemanticError,
				)

				panic(fmt.Sprintf("Tipo del argumento '%s' (%s) no coincide con el parámetro de tipo %s", param.Nombre, paramValue.info, param.Tipo))
			}

			if paramValue.info == "int" || paramValue.info == "f64" || paramValue.info == "string" || paramValue.info == "bool" {
				v.Env.SetVar(param.Nombre, &Variable{Nombre: param.Nombre, Tipo: param.Tipo, Valor: paramValue.value, Mut: true})
			}

			if paramValue.info == "slice" {

				// veo si los tipos coinciden
				if paramValue.value.(*Slice).Tipo != param.TipoSlice {
					//agregar error a la tabla de errores
					v.ErrorTable.AddError(
						ctx.GetStart().GetLine(),
						ctx.GetStart().GetColumn(),
						"Tipo del argumento '"+param.Nombre+"' ("+paramValue.value.(*Slice).Tipo+") no coincide con el parámetro de tipo "+param.TipoSlice,
						SemanticError,
					)
					panic(fmt.Sprintf("Tipo del argumento '%s' (%s) no coincide con el parámetro de tipo %s", param.Nombre, paramValue.value.(*Slice).Tipo, param.TipoSlice))
				}

				// si es un slice, lo guardo como un slice
				sliceValue := paramValue.value.(*Slice)
				// creo una copia ya que aca le paso la direccion de memoria
				// y no quiero que se modifique el original
				newSlice := &Slice{
					Nombre:    sliceValue.Nombre,
					Tipo:      sliceValue.Tipo,
					Elementos: make([]interface{}, len(sliceValue.Elementos)),
					Mut:       sliceValue.Mut,
				}
				// Copiar los elementos del slice original al nuevo slice
				copy(newSlice.Elementos, sliceValue.Elementos)
				v.Env.SetSlice(param.Nombre, newSlice) // Asigno el slice al entorno
			}

		}

		v.Visit(function.Cuerpo)

	}

	// Si la función tiene un tipo de retorno, obtengo el valor de retorno
	if function.TipoRetorno != "void" {
		// creo un nuevo entorno para la función

		callEnv := v.Env // el entonrno donde se esta llamando la función

		creationEnv := function.Parent // el entorno donde fue declarada la función

		contextosPrev := v.ControlContext
		v.ControlContext = make([]string, 0) // reinicio la pila de contextos

		v.Env = NewEnv(creationEnv, "Function")
		defer func() {
			v.Env = callEnv
			v.ControlContext = v.ControlContext[:len(v.ControlContext)-1]
			v.ControlContext = contextosPrev // vuelvo a la pila de contextos anterior
			if r := recover(); r != nil {
				if ret, ok := r.(ReturnSignal); ok {
					if ret.info != function.TipoRetorno {
						//agregar error a la tabla de errores
						v.ErrorTable.AddError(
							ctx.GetStart().GetLine(),
							ctx.GetStart().GetColumn(),
							"Tipo de retorno de la función '"+funcName+"' no coincide con el tipo esperado '"+function.TipoRetorno+"'",
							SemanticError,
						)
						panic(fmt.Sprintf("Tipo de retorno de la función '%s' no coincide con el tipo esperado '%s'", funcName, function.TipoRetorno))
					}

					if ret.info == "slice" {
						if ret.Value.(*Slice).Tipo != function.TipoSlice {
							//agregar error a la tabla de errores
							v.ErrorTable.AddError(
								ctx.GetStart().GetLine(),
								ctx.GetStart().GetColumn(),
								"Tipo de retorno de la función '"+funcName+"' no coincide con el tipo esperado slice '"+function.TipoSlice+"'",
								SemanticError,
							)
							panic(fmt.Sprintf("Tipo de retorno de la función '%s' no coincide con el tipo esperado slice '%s'", funcName, function.TipoSlice))
						}
					}

					result = Value{value: ret.Value, info: ret.info}
				} else {
					panic(r) // relanzo si no es return
				}
			}
		}()
		// Indico que estoy dentro de un funcion
		v.ControlContext = append(v.ControlContext, "function")
		// Asignar parámetros
		for i, param := range function.Parametros {
			paramValue := args[i]
			if paramValue.info != param.Tipo {
				//agregar error a la tabla de errores
				v.ErrorTable.AddError(
					ctx.GetStart().GetLine(),
					ctx.GetStart().GetColumn(),
					"Tipo del argumento '"+param.Nombre+"' ("+paramValue.info+") no coincide con el parámetro de tipo "+param.Tipo,
					SemanticError,
				)
				panic(fmt.Sprintf("Tipo del argumento '%s' (%s) no coincide con el parámetro de tipo %s", param.Nombre, paramValue.info, param.Tipo))
			}

			if paramValue.info == "int" || paramValue.info == "f64" || paramValue.info == "string" || paramValue.info == "bool" {
				v.Env.SetVar(param.Nombre, &Variable{Nombre: param.Nombre, Tipo: param.Tipo, Valor: paramValue.value, Mut: true})
			}

			if paramValue.info == "slice" {

				// veo si los tipos coinciden
				if paramValue.value.(*Slice).Tipo != param.TipoSlice {
					//agregar error a la tabla de errores
					v.ErrorTable.AddError(
						ctx.GetStart().GetLine(),
						ctx.GetStart().GetColumn(),
						"Tipo del argumento '"+param.Nombre+"' ("+paramValue.value.(*Slice).Tipo+") no coincide con el parámetro de tipo "+param.TipoSlice,
						SemanticError,
					)
					panic(fmt.Sprintf("Tipo del argumento '%s' (%s) no coincide con el parámetro de tipo %s", param.Nombre, paramValue.value.(*Slice).Tipo, param.TipoSlice))
				}
				// si es un slice, lo guardo como un slice
				sliceValue := paramValue.value.(*Slice)
				// creo una copia ya que aca le paso la direccion de memoria
				// y no quiero que se modifique el original
				newSlice := &Slice{
					Nombre:    sliceValue.Nombre,
					Tipo:      sliceValue.Tipo,
					Elementos: make([]interface{}, len(sliceValue.Elementos)),
					Mut:       sliceValue.Mut,
				}
				// Copiar los elementos del slice original al nuevo slice
				copy(newSlice.Elementos, sliceValue.Elementos)
				v.Env.SetSlice(param.Nombre, newSlice) // Asigno el slice al entorno
			}

		}

		v.Visit(function.Cuerpo)

		//agrego el error a la tabla
		v.ErrorTable.AddError(
			ctx.GetStart().GetLine(),
			ctx.GetStart().GetColumn(),
			"La función '"+funcName+"' no retornó ningún valor",
			SemanticError,
		)

		// Si no hubo return explícito
		panic("La función '" + funcName + "' no retornó ningún valor")

	}

	return result

}
