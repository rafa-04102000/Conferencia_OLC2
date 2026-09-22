package codegen

import (
	"backend/parser"
	"fmt"
	"os"
	"strconv"
	"strings"

	"github.com/antlr4-go/antlr/v4"
)

// Struct TIpo Visitor, seguirmeos con GramataicaVisitor
type CodeGenVisitor struct {
	parser.GramaticaVisitor
	Env        *Env            // Entorno actual (para variables, funciones, etc.)
	TemporalId int             // Contador para generar temporales únicos
	SalidaId   int             // Contador para generar salidas únicas// prints unicos
	Data       strings.Builder // Acumulador de datos (variables, constantes, etc.)
	BSS        strings.Builder // Acumulador de datos no inicializados
	Code       strings.Builder // Acumulador de las instrucciones ARM generadas
	Fun        strings.Builder // Acumulador de funciones
	LabelCnt   int             // Contador para generar etiquetas únicas

	loopStack []LoopContext // Pila para manejar bucles anidados
}

type LoopContext struct {
	// para manejar break y continue
	startLabel string // inicio del loop
	endLabel   string // fin del loop
}

type Env struct {
	name   string
	parent *Env
	vars   map[string]*Variable
}

func NewEnv(parent *Env, name string) *Env {
	return &Env{
		name:   name,
		parent: parent,
		vars:   make(map[string]*Variable),
	}
}

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
	default:
		panic("Tipo desconocido: " + tipo)
	}
}

type Variable struct {
	Nombre string
	Tipo   string      // "int", "float", "bool", "string"
	Valor  interface{} // puede ser cualquier tipo
	Addr   string      // dirección en memoria
	IsMut  bool        // si es mutable o no
}

func (v *CodeGenVisitor) validateType(typeValue string, value Attr) {

	switch typeValue {
	case "int":
		if _, ok := value.value.(int64); !ok {
			panic(fmt.Sprintf("no se puede asignar un valor de tipo %T a una variable 'int'", value.value))
		}
	case "f64":
		if _, ok := value.value.(float64); !ok {
			panic(fmt.Sprintf("no se puede asignar un valor de tipo %T a una variable 'float'", value.value))
		}
	case "string":
		if _, ok := value.value.(string); !ok {
			panic(fmt.Sprintf("no se puede asignar un valor de tipo %T a una variable 'string'", value.value))
		}
	case "bool":
		if _, ok := value.value.(bool); !ok {
			panic(fmt.Sprintf("no se puede asignar un valor de tipo %T a una variable 'bool'", value.value))
		}
	default:
		panic("tipo desconocido para la asignacion: " + typeValue)
	}
}

func NewCodeGenVisitor() *CodeGenVisitor {
	return &CodeGenVisitor{
		Env:        NewEnv(nil, "Global"), // Entorno global
		TemporalId: 0,
		SalidaId:   0,
		Data:       strings.Builder{},
		BSS:        strings.Builder{},
		Code:       strings.Builder{},
		LabelCnt:   0,
	}
}

type Attr struct {
	addrNUM string // si es un valor inmediato (como 5)
	addrID  string // si es una variable, como "x"
	// si es un valor inmediato de que tipo es
	info  string      // "int", "float", "bool", "str"
	value interface{} // valor real, puede ser int, float, bool o string
}

func (v *CodeGenVisitor) writeExit() {
	v.Data.WriteString("\n")
	v.BSS.WriteString("\n")

	v.Code.WriteString("\n\t// Salida del programa\n")
	v.Code.WriteString("\tmov x0, #0\t// código de salida 0\n")
	// x0 es el primer argumento de la syscall exit, que es el código de
	v.Code.WriteString("\tmov x8, #93\t// syscall exit\n")
	v.Code.WriteString("\tsvc #0\n")
}

func (v *CodeGenVisitor) NextLabel() string {
	v.LabelCnt++
	return fmt.Sprintf("label_%d", v.LabelCnt)
}

// nextTemp genera un nombre de etiqueta temporal único
func (v *CodeGenVisitor) nextTemp() string {
	v.TemporalId++
	return fmt.Sprintf("tmp_%d", v.TemporalId)
}

func (v *CodeGenVisitor) nextSal() string {
	v.SalidaId++
	return fmt.Sprintf("salida_%d", v.SalidaId)
}

// ----------------------------------------------------------------
func (v *CodeGenVisitor) Visit(tree antlr.ParseTree) interface{} {

	switch val := tree.(type) {

	// |Bloques Iniciales|
	case *parser.ProgramContext:
		v.VisitProgram(val)
		return nil
	case *parser.BlockContext:
		v.VisitBlock(val)
		return nil
	case *parser.StatementContext:
		v.VisitStatement(val)
		return nil

	// |Main|
	case *parser.MainfunctionContext:
		v.VisitMainfunction(val)
		return nil

	// |Declaraciones|
	case *parser.DeclarationExplicitVarContext:
		v.VisitDeclarationExplicitVar(val)
		return nil
	case *parser.DeclarationImplicitVarContext:
		v.VisitDeclarationImplicitVar(val)
		return nil

		// | Print |
	case *parser.PrintContext:
		v.VisitPrint(val)
		return nil

		// |Expresiones|
	case *parser.NotExprContext:
		attr := v.VisitNotExpr(val).(Attr)
		return attr
	case *parser.OpExprContext:
		attr := v.VisitOpExpr(val).(Attr)
		return attr
	case *parser.IdExprContext:
		attr := v.VisitIdExpr(val).(Attr)
		return attr
	case *parser.ParExprContext:
		attr := v.VisitParExpr(val).(Attr)
		return attr
	case *parser.IntExprContext:
		attr := v.VisitIntExpr(val).(Attr)
		return attr
	case *parser.FloatExprContext:
		attr := v.VisitFloatExpr(val).(Attr)
		return attr
	case *parser.BoolExprContext:
		attr := v.VisitBoolExpr(val).(Attr)
		return attr
	case *parser.StrExprContext:
		attr := v.VisitStrExpr(val).(Attr)
		return attr

	// | If |
	case *parser.IfContext:
		v.VisitIf(val)
		return nil

	// | Asignaciones |
	case *parser.AssignmentExprContext:
		v.VisitAssignmentExpr(val)

	// | Incremento |
	case *parser.AssignmentIncrementExprContext:
		v.VisitAssignmentIncrementExpr(val)

	// | Decremento |
	case *parser.AssignmentDecrementExprContext:
		v.VisitAssignmentDecrementExpr(val)

	// | Break |
	case *parser.BreakStatementContext:
		v.VisitBreak(val)

	// | Continue |
	case *parser.ContinueStatementContext:
		v.VisitContinue(val)

	// | For |
	case *parser.ForSimpleContext:
		v.VisitForSimple(val)

	default:
		panic(fmt.Sprintf("Tipo de nodo inesperado: %T", val))

	}

	return nil
}

//---------------------------------------------------------------
// |Program|
//---------------------------------------------------------------

func (v *CodeGenVisitor) VisitProgram(ctx *parser.ProgramContext) interface{} {
	v.Data.WriteString("// ------------------------------------------------------------------------------\n") // Sección de variables globales
	v.Data.WriteString("// 	| VARIABLES GLOBALES DATA Y BSS |\n\n")
	v.Data.WriteString(".data\n")                       // Sección de datos inicializados
	v.Data.WriteString("newline: .asciz \"\\n\"\n")     // Nueva línea
	v.Data.WriteString("flt_100: .double 100.0\n")      // Doble de 100.0
	v.Data.WriteString("str_true: .asciz \"true\"\n")   // Cadena para true
	v.Data.WriteString("str_false: .asciz \"false\"\n") // Cadena para false

	v.BSS.WriteString(".bss\n") // Sección de datos no inicializados

	v.Code.WriteString("// ------------------------------------------------------------------------------\n") // Sección de código
	v.Code.WriteString("// 	| CODIGO |\n\n")                                                                  // Sección de Funciones

	v.Code.WriteString(".text\n.global _start\n\n")
	v.Code.WriteString("_start:\n") // Punto de entrada del programa

	v.Fun.WriteString("\n// ------------------------------------------------------------------------------\n") // Sección de Funciones
	v.Fun.WriteString("// 	| FUNCIONES |\n\n")                                                                 // Sección de Funciones

	// Pongo la funcion para copiar strings
	// limpia y verifica nulos
	v.Fun.WriteString("\n//---- Función para copiar strings ----\n")
	v.Fun.WriteString("copiar_string:\n")
	v.Fun.WriteString("\tmov x3, x0          // Guardar puntero original del destino\n")
	v.Fun.WriteString("\tmov x9, x4          // Copiar tamaño para limpieza\n\n")

	// Limpieza del buffer destino con ceros
	v.Fun.WriteString("clean_loop:\n")
	v.Fun.WriteString("\tcbz x9, copy_start\n")
	v.Fun.WriteString("\tmov w10, #0\n")
	v.Fun.WriteString("\tstrb w10, [x0], #1\n")
	v.Fun.WriteString("\tsubs x9, x9, #1\n")
	v.Fun.WriteString("\tb clean_loop\n\n")

	// Restaurar puntero y contador para copia
	v.Fun.WriteString("copy_start:\n")
	v.Fun.WriteString("\tmov x0, x3          // Restaurar puntero destino\n")
	v.Fun.WriteString("\tmov x5, x4          // Tamaño máximo para copia\n\n")

	// Copiar el string carácter por carácter
	v.Fun.WriteString("copy_loop:\n")
	v.Fun.WriteString("\tldrb w6, [x1], #1   // Cargar byte del origen\n")
	v.Fun.WriteString("\tcbz w6, done_copy   // Terminar si es null\n")
	v.Fun.WriteString("\tstrb w6, [x0], #1   // Escribir byte en destino\n")
	v.Fun.WriteString("\tsubs x5, x5, #1     // Decrementar contador\n")
	v.Fun.WriteString("\tb.gt copy_loop\n\n")

	// Si se llenó el buffer sin encontrar null, forzar terminación
	v.Fun.WriteString("force_null:\n")
	v.Fun.WriteString("\tmov w6, #0\n")
	v.Fun.WriteString("\tstrb w6, [x0]\n")

	v.Fun.WriteString("done_copy:\n")
	v.Fun.WriteString("\tret\n\n")

	v.Fun.WriteString("\n//---- Fin Función para copiar strings ----\n")

	// Funcion para pasar int a string
	v.Fun.WriteString("\n//---- Función para convertir int a string ----\n")

	v.Fun.WriteString("int_to_str:\n")
	v.Fun.WriteString("\tmov x11, x1       // puntero de escritura\n")
	v.Fun.WriteString("\tmov x3, #10      // base decimal\n")
	v.Fun.WriteString("\tmov x10, #0       // contador de dígitos\n\n")
	v.Fun.WriteString("reverse_digits:\n")
	v.Fun.WriteString("\tudiv w5, w0, w3  // w5 = w0 / 10\n")
	v.Fun.WriteString("\tmsub w6, w5, w3, w0  // w6 = w0 - w5*10 → w6 = w0 % 10\n")
	v.Fun.WriteString("\tadd w6, w6, #48     // convertir a ASCII\n")
	v.Fun.WriteString("\tstrb w6, [x11], #1\n")
	v.Fun.WriteString("\tmov w0, w5\n")
	v.Fun.WriteString("\tadd x10, x10, #1\n")
	v.Fun.WriteString("\tcbnz w0, reverse_digits\n\n")
	v.Fun.WriteString("\tmov w6, #0\n")
	v.Fun.WriteString("\tstrb w6, [x11]\n\n")
	v.Fun.WriteString("\t// Invertir string (porque lo construimos al revés)\n")
	v.Fun.WriteString("\tsub x11, x11, #1\n")
	v.Fun.WriteString("\tmov x5, x1\n")
	v.Fun.WriteString("reverse_loop:\n")
	v.Fun.WriteString("\tcmp x5, x11\n")
	v.Fun.WriteString("\tbge reverse_done\n")
	v.Fun.WriteString("\tldrb w6, [x5]\n")
	v.Fun.WriteString("\tldrb w7, [x11]\n")
	v.Fun.WriteString("\tstrb w7, [x5]\n")
	v.Fun.WriteString("\tstrb w6, [x11]\n")
	v.Fun.WriteString("\tadd x5, x5, #1\n")
	v.Fun.WriteString("\tsub x11, x11, #1\n")
	v.Fun.WriteString("\tb reverse_loop\n")
	v.Fun.WriteString("reverse_done:\n")
	v.Fun.WriteString("\tret\n\n")
	v.Fun.WriteString("\n//---- Fin Función para convertir int a string ----\n")

	// funcion para convertir float a string
	v.Fun.WriteString("\n//---- Función para convertir float a string ----\n")

	v.Fun.WriteString("float_to_str:\n")
	v.Fun.WriteString("\t// Preservar registros que vamos a modificar\n")
	v.Fun.WriteString("\tstp x29, x30, [sp, #-16]!  // Guardar frame pointer y return address\n")
	v.Fun.WriteString("\tmov x29, sp\n")
	v.Fun.WriteString("\n")
	v.Fun.WriteString("\tmov x4, x1            // Guardar puntero destino original\n")
	v.Fun.WriteString("\n")
	v.Fun.WriteString("\t// Parte entera: truncar float\n")
	v.Fun.WriteString("\tfcvtzs w0, d0         // Convertir a entero (parte entera)\n")
	v.Fun.WriteString("\tbl int_to_str         // Convertir parte entera a string\n")
	v.Fun.WriteString("\n")
	v.Fun.WriteString("\t// Avanzar al final del string para agregar punto decimal\n")
	v.Fun.WriteString("\tmov x5, x4            // x5 = puntero que avanza\n")
	v.Fun.WriteString("find_end:\n")
	v.Fun.WriteString("\tldrb w6, [x5]\n")
	v.Fun.WriteString("\tcbz w6, add_dot\n")
	v.Fun.WriteString("\tadd x5, x5, #1\n")
	v.Fun.WriteString("\tb find_end\n")
	v.Fun.WriteString("\n")
	v.Fun.WriteString("add_dot:\n")
	v.Fun.WriteString("\tmov w6, #'.'\n")
	v.Fun.WriteString("\tstrb w6, [x5], #1     // Agregar punto decimal\n")
	v.Fun.WriteString("\n")
	v.Fun.WriteString("\t// Calcular parte decimal: (valor - parte_entera) * 100\n")
	v.Fun.WriteString("\tscvtf d1, w0          // Convertir parte entera a float\n")
	v.Fun.WriteString("\tfsub d2, d0, d1       // Obtener parte decimal\n")
	v.Fun.WriteString("\tldr x0, =flt_100      // Cargar 100.0\n")
	v.Fun.WriteString("\tldr d3, [x0]\n")
	v.Fun.WriteString("\tfmul d2, d2, d3       // Multiplicar por 100 para obtener 2 decimales\n")
	v.Fun.WriteString("\n")
	v.Fun.WriteString("\t// Convertir decimales a entero\n")
	v.Fun.WriteString("\tfcvtzu w7, d2         // Convertir a entero sin signo\n")
	v.Fun.WriteString("\n")
	v.Fun.WriteString("\t// Manejar redondeo (ej. 33.4999999 → 33.5 → 50)\n")
	v.Fun.WriteString("\t// Agregamos 0.5 antes de convertir para redondear correctamente\n")
	v.Fun.WriteString("\tfmov d4, #0.5\n")
	v.Fun.WriteString("\tfadd d2, d2, d4\n")
	v.Fun.WriteString("\tfcvtzu w7, d2\n")
	v.Fun.WriteString("\n")
	v.Fun.WriteString("\t// Convertir parte decimal a string\n")
	v.Fun.WriteString("\tmov x1, x5            // Puntero donde escribir los decimales\n")
	v.Fun.WriteString("\tmov w0, w7\n")
	v.Fun.WriteString("\tbl int_to_str\n")
	v.Fun.WriteString("\n")
	v.Fun.WriteString("\t// Asegurar exactamente 2 dígitos\n")
	v.Fun.WriteString("\tmov x6, x5            // Puntero al inicio de los decimales\n")
	v.Fun.WriteString("\tmov x7, x1            // Puntero al final del string\n")
	v.Fun.WriteString("\tsub x7, x7, x6        // Longitud de los decimales\n")
	v.Fun.WriteString("\n")
	v.Fun.WriteString("\tcmp x7, #1\n")
	v.Fun.WriteString("\tbne check_length\n")
	v.Fun.WriteString("\n")
	v.Fun.WriteString("\t// Solo 1 dígito, insertar '0' delante\n")
	v.Fun.WriteString("\tldrb w8, [x5]\n")
	v.Fun.WriteString("\tmov w9, #'0'\n")
	v.Fun.WriteString("\tstrb w9, [x5]\n")
	v.Fun.WriteString("\tstrb w8, [x5, #1]\n")
	v.Fun.WriteString("\tmov w8, #0\n")
	v.Fun.WriteString("\tstrb w8, [x5, #2]\n")
	v.Fun.WriteString("\tb done_float\n")
	v.Fun.WriteString("\n")
	v.Fun.WriteString("check_length:\n")
	v.Fun.WriteString("\tcmp x7, #0\n")
	v.Fun.WriteString("\tbne done_float\n")
	v.Fun.WriteString("\n")
	v.Fun.WriteString("\t// Ningún dígito, insertar \"00\"\n")
	v.Fun.WriteString("\tmov w8, #'0'\n")
	v.Fun.WriteString("\tstrb w8, [x5], #1\n")
	v.Fun.WriteString("\tstrb w8, [x5], #1\n")
	v.Fun.WriteString("\tmov w8, #0\n")
	v.Fun.WriteString("\tstrb w8, [x5]\n")
	v.Fun.WriteString("\n")
	v.Fun.WriteString("done_float:\n")
	v.Fun.WriteString("\tldp x29, x30, [sp], #16  // Restaurar frame pointer y return address\n")
	v.Fun.WriteString("\tret\n\n")

	v.Fun.WriteString("\n//---- Fin Función para convertir float a string ----\n")

	// Funcion para convertir bool a string
	v.Fun.WriteString("\n//---- Función para convertir bool a string ----\n")

	v.Fun.WriteString("bool_to_str:\n")
	v.Fun.WriteString("\t// w0 = valor booleano (1=true, 0=false)\n")
	v.Fun.WriteString("\t// x1 = dirección destino\n")
	v.Fun.WriteString("\n")
	v.Fun.WriteString("\tcmp w0, #1\n")
	v.Fun.WriteString("\tbeq store_true\n")
	v.Fun.WriteString("\n")
	v.Fun.WriteString("\t// Guardar 'false'\n")
	v.Fun.WriteString("\tldr x2, =str_false\n")
	v.Fun.WriteString("\tb copy_string\n")
	v.Fun.WriteString("\n")
	v.Fun.WriteString("store_true:\n")
	v.Fun.WriteString("\t// Guardar 'true'\n")
	v.Fun.WriteString("\tldr x2, =str_true\n")
	v.Fun.WriteString("\n")
	v.Fun.WriteString("copy_string:\n")
	v.Fun.WriteString("\tldrb w3, [x2], #1\n")
	v.Fun.WriteString("\tstrb w3, [x1], #1\n")
	v.Fun.WriteString("\tcbnz w3, copy_string\n")
	v.Fun.WriteString("\n")
	v.Fun.WriteString("\tret\n")

	v.Fun.WriteString("\n//---- Fin Función para convertir bool a string ----\n")

	// Funcion para avanzar al final de la cadena
	v.Fun.WriteString("\n//---- Función para avanzar al final de una cadena ----\n")

	v.Fun.WriteString("advance_to_end:\n")
	v.Fun.WriteString("\tldrb w3, [x0]\n")
	v.Fun.WriteString("\tcbz w3, done_advance\n")
	v.Fun.WriteString("\tadd x0, x0, #1\n")
	v.Fun.WriteString("\tb advance_to_end\n")
	v.Fun.WriteString("done_advance:\n")
	v.Fun.WriteString("\tret\n\n")

	v.Fun.WriteString("\n//---- Fin Función para avanzar al final de una cadena ----\n")

	// Funcion para el println
	v.Fun.WriteString("\n//---- Función para imprimir en consola ----\n")

	v.Fun.WriteString("println:\n")
	v.Fun.WriteString("\tmov x2, #0          // Contador de longitud\n")
	v.Fun.WriteString("\tmov x3, x1          // Guardar puntero original\n")
	v.Fun.WriteString("\ncount_loop:\n")
	v.Fun.WriteString("\tldrb w4, [x3, x2]    // Cargar byte actual\n")
	v.Fun.WriteString("\tcbz w4, do_print      // Si es nulo, salta a imprimir\n")
	v.Fun.WriteString("\tadd x2, x2, #1        // Incrementar contador\n")
	v.Fun.WriteString("\tb count_loop          // Repetir hasta encontrar nulo\n")
	v.Fun.WriteString("\ndo_print:\n")
	v.Fun.WriteString("\tmov x0, #1		// stdout\n")
	v.Fun.WriteString("\tmov x8, #64         // syscall write\n")
	v.Fun.WriteString("\tsvc #0\n")
	v.Fun.WriteString("\tret\n")

	v.Fun.WriteString("\n//---- Fin Función para imprimir en consola ----\n")

	// Funcion copiar sin limpiar
	v.Fun.WriteString("\n//---- Función para copiar strings sin limpiar ----\n")

	v.Fun.WriteString("copiar_sin_limpiar:\n")
	v.Fun.WriteString("\t// x0 = destino\n")
	v.Fun.WriteString("\t// x1 = origen\n")
	v.Fun.WriteString("copiar_sin_loop:\n")
	v.Fun.WriteString("\tldrb w6, [x1], #1\n")
	v.Fun.WriteString("\tcbz w6, fin_copiar_sin\n")
	v.Fun.WriteString("\tstrb w6, [x0], #1\n")
	v.Fun.WriteString("\tb copiar_sin_loop\n")
	v.Fun.WriteString("fin_copiar_sin:\n")
	v.Fun.WriteString("\tret\n")

	v.Fun.WriteString("\n//---- Fin Función para copiar strings sin limpiar ----\n")

	// Funcion para comparar strings ==

	v.Fun.WriteString("\n//---- Función para comparar strings (==) ----\n")

	v.Fun.WriteString("strcmp_eq:\n")
	v.Fun.WriteString("\tcmp_loop_eq:\n")
	v.Fun.WriteString("\tldrb w2, [x0], #1\n")
	v.Fun.WriteString("\tldrb w3, [x1], #1\n")
	v.Fun.WriteString("\tcmp w2, w3\n")
	v.Fun.WriteString("\tb.ne not_equal\n")
	v.Fun.WriteString("\tcbz w2, strings_equal\n")
	v.Fun.WriteString("\tb cmp_loop_eq\n")
	v.Fun.WriteString("not_equal:\n")
	v.Fun.WriteString("\tmov w0, #0\n")
	v.Fun.WriteString("\tret\n")
	v.Fun.WriteString("strings_equal:\n")
	v.Fun.WriteString("\tmov w0, #1\n")
	v.Fun.WriteString("\tret\n\n")

	v.Fun.WriteString("\n//---- Fin Función para comparar strings (==) ----\n")

	// Funcion para comparar strings !=
	v.Fun.WriteString("\n//---- Función para comparar strings (!=) ----\n")

	// x0 = dirección de string 1
	// x1 = dirección de string 2
	// Retorna w0 = 1 si son distintos, 0 si son iguales

	v.Fun.WriteString("strcmp_ne:\n")
	v.Fun.WriteString("\tcmp_loop_ne:\n")
	v.Fun.WriteString("\tldrb w2, [x0], #1\n")  // Leer byte de string 1
	v.Fun.WriteString("\tldrb w3, [x1], #1\n")  // Leer byte de string 2
	v.Fun.WriteString("\tcmp w2, w3\n")         // Comparar bytes
	v.Fun.WriteString("\tb.ne str_not_equal\n") // Si son diferentes, retornar 1
	v.Fun.WriteString("\tcbz w2, str_equal\n")  // Si llegamos a fin, son iguales
	v.Fun.WriteString("\tb cmp_loop_ne\n")      // Continuar comparación

	v.Fun.WriteString("str_not_equal:\n")
	v.Fun.WriteString("\tmov w0, #1\n") // Son distintos
	v.Fun.WriteString("\tret\n")

	v.Fun.WriteString("str_equal:\n")
	v.Fun.WriteString("\tmov w0, #0\n") // Son iguales
	v.Fun.WriteString("\tret\n\n")

	v.Fun.WriteString("\n//---- Fin Función para comparar strings (!=) ----\n")

	// Funcion para comparar strings <

	v.Fun.WriteString("\n//---- Función para comparar strings (<) ----\n")

	v.Fun.WriteString("strcmp_lt:\n")
	v.Fun.WriteString("\tcmp_loop_lt:\n")
	v.Fun.WriteString("\tldrb w2, [x0], #1\n")
	v.Fun.WriteString("\tldrb w3, [x1], #1\n")
	v.Fun.WriteString("\tcmp w2, w3\n")
	v.Fun.WriteString("\tb.lo str1_less\n")
	v.Fun.WriteString("\tb.hi str1_greater\n")
	v.Fun.WriteString("\tcbz w2, strings_equal_lt\n")
	v.Fun.WriteString("\tb cmp_loop_lt\n")
	v.Fun.WriteString("str1_less:\n")
	v.Fun.WriteString("\tmov w0, #1\n")
	v.Fun.WriteString("\tret\n")
	v.Fun.WriteString("str1_greater:\n")
	v.Fun.WriteString("strings_equal_lt:\n")
	v.Fun.WriteString("\tmov w0, #0\n")
	v.Fun.WriteString("\tret\n\n")

	v.Fun.WriteString("\n//---- Fin Función para comparar strings (<) ----\n")

	// Funcion para comparar strings <=

	v.Fun.WriteString("\n//---- Función para comparar strings (<=) ----\n")

	v.Fun.WriteString("strcmp_le:\n")
	v.Fun.WriteString("\tmov x2, x0\n")
	v.Fun.WriteString("\tmov x3, x1\n")
	v.Fun.WriteString("\tbl strcmp_lt\n")
	v.Fun.WriteString("\tcmp w0, #1\n")
	v.Fun.WriteString("\tbeq return_1_le\n")
	v.Fun.WriteString("\tmov x0, x2\n")
	v.Fun.WriteString("\tmov x1, x3\n")
	v.Fun.WriteString("\tbl strcmp_eq\n")
	v.Fun.WriteString("\tret\n")
	v.Fun.WriteString("return_1_le:\n")
	v.Fun.WriteString("\tmov w0, #1\n")
	v.Fun.WriteString("\tret\n\n")

	v.Fun.WriteString("\n//---- Fin Función para comparar strings (<=) ----\n")

	// Funcion para comparar strings >

	v.Fun.WriteString("\n//---- Función para comparar strings (>) ----\n")

	v.Fun.WriteString("strcmp_gt:\n")
	v.Fun.WriteString("\tmov x2, x0\n")
	v.Fun.WriteString("\tmov x3, x1\n")
	v.Fun.WriteString("\tmov x0, x1\n")
	v.Fun.WriteString("\tmov x1, x2\n")
	v.Fun.WriteString("\tbl strcmp_lt\n")
	v.Fun.WriteString("\tret\n\n")

	v.Fun.WriteString("\n//---- Fin Función para comparar strings (>) ----\n")

	// Funcion para comparar strings >=

	v.Fun.WriteString("\n//---- Función para comparar strings (>=) ----\n")

	v.Fun.WriteString("strcmp_ge:\n")
	v.Fun.WriteString("\tmov x2, x0\n")
	v.Fun.WriteString("\tmov x3, x1\n")
	v.Fun.WriteString("\tmov x0, x1\n")
	v.Fun.WriteString("\tmov x1, x2\n")
	v.Fun.WriteString("\tbl strcmp_lt\n")
	v.Fun.WriteString("\tcmp w0, #1\n")
	v.Fun.WriteString("\tbeq return_0_ge\n")
	v.Fun.WriteString("\tmov x0, x2\n")
	v.Fun.WriteString("\tmov x1, x3\n")
	v.Fun.WriteString("\tbl strcmp_eq\n")
	v.Fun.WriteString("\tret\n")
	v.Fun.WriteString("return_0_ge:\n")
	v.Fun.WriteString("\tmov w0, #0\n")
	v.Fun.WriteString("\tret\n\n")

	v.Fun.WriteString("\n//---- Fin Función para comparar strings (>=) ----\n")

	v.Visit(ctx.Block())

	return nil
}

//---------------------------------------------------------------
// |Block|
//---------------------------------------------------------------

func (v *CodeGenVisitor) VisitBlock(ctx *parser.BlockContext) interface{} {
	// recorro todas las sentencias dentro del bloque

	cont := 0
	var mainCtx parser.IMainfunctionContext = nil

	// Registrar todo y buscar main (no ejecutarla aún) hasta validar
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
			panic("No se encontró ninguna función 'main()'")
		} else if cont > 1 {
			panic("Solo se permite una función 'main()'")
		}

		// Ejecuto main si existe la funcion main
		v.Visit(mainCtx)
	}
	return nil
}

// ---------------------------------------------------------------
// |Statement|
// ---------------------------------------------------------------
func (v *CodeGenVisitor) VisitStatement(ctx *parser.StatementContext) interface{} {
	// fmt.Println("Visitando sentencia:", ctx.GetText())
	switch {
	case ctx.Mainfunction() != nil:
		return v.Visit(ctx.Mainfunction())
	case ctx.Declaration() != nil:
		return v.Visit(ctx.Declaration())
	case ctx.Print_() != nil:
		return v.Visit(ctx.Print_())
	case ctx.If_() != nil:
		return v.Visit(ctx.If_())
	case ctx.Assignment() != nil:
		return v.Visit(ctx.Assignment())
	case ctx.For_() != nil:
		return v.Visit(ctx.For_())
	default:
		panic(fmt.Sprintf("Tipo de sentencia inesperada: %T", ctx))
	}
	return nil
}

//---------------------------------------------------------------
// |Main|
//---------------------------------------------------------------

func (v *CodeGenVisitor) VisitMainfunction(ctx *parser.MainfunctionContext) interface{} {

	previousEnv := v.Env
	v.Env = NewEnv(previousEnv, "Main")
	defer func() { v.Env = previousEnv }()

	v.Visit(ctx.Block())

	v.writeExit()
	return nil
}

// ---------------------------------------------------------------
// |Declaration Explicit Var|
// ---------------------------------------------------------------
func (v *CodeGenVisitor) VisitDeclarationExplicitVar(ctx *parser.DeclarationExplicitVarContext) interface{} {
	// fmt.Println("Visitando declaración explícita de variable")
	id := ctx.ID().GetText()

	// Validar si ya existe
	if _, ok := v.Env.GetVar(id); ok {
		panic(fmt.Sprintf("Variable ya declarada: %s", id))
	}

	// mensaje de declaracion
	v.Code.WriteString(fmt.Sprintf("\t// ---------- Declaración de variable: %s ----------\n\n", id))

	// Evaluar expresión
	expr := v.Visit(ctx.Expression()).(Attr)
	tipo := ctx.PrimitiveType().GetText()
	v.validateType(tipo, expr)

	if isNUM(expr) {
		// Caso: valor inmediato (literal)
		switch tipo {
		case "int", "bool":
			v.Data.WriteString(fmt.Sprintf("%s: .word %s\n", id, expr.addrNUM))
		case "f64":
			v.Data.WriteString(fmt.Sprintf("%s: .double %s\n", id, expr.addrNUM))
		case "string":
			// declaro la cadena en .bss con skip de 64
			v.BSS.WriteString(fmt.Sprintf("%s: .skip %d\n", id, 64))

			// guardo la cadena en  el id en bss
			// ocupo su direccion de memoria
			v.Code.WriteString(fmt.Sprintf("\tldr x0, =%s\t//destiono\n", id))
			// creo temporal para la cadena
			tmp := v.nextTemp()
			// declaro la cadena en .data
			v.Data.WriteString(fmt.Sprintf("%s: .asciz \"%s\"\n", tmp, expr.addrNUM))
			// copio la cadena a la direccion de memoria
			v.Code.WriteString(fmt.Sprintf("\tldr x1, =%s\t//origen\n", tmp))
			//v.Code.WriteString("\tmov x4, #64\t//Tamaño del buffer destino\n") // Tamaño de la cadena
			// copiar_string si pide cantidad del buffer
			v.Code.WriteString("\tbl copiar_sin_limpiar\n")

			// en copiar string x0 es el destino y x1 es el origen
			// y x4 es el tamaño del buffer destino
		}
	} else {
		// Caso: valor en memoria (resultado de operación o variable)

		// Declaración en .data o .bss según tipo
		switch tipo {
		case "int", "bool":
			v.Data.WriteString(fmt.Sprintf("%s: .word 0\n", id))
		case "f64":
			v.Data.WriteString(fmt.Sprintf("%s: .double 0.0\n", id))
		case "string":
			// fmt.Println(len(expr.value.(string)), "longitud de la cadena")
			//length := len(expr.value.(string)) + 2 // +2 para el carácter nulo
			// sumo dos porque la cadena normal no cuenta el \n, y aparte tambien le agrego el nulo
			// por eso sumo dos
			v.BSS.WriteString(fmt.Sprintf("%s: .skip %d\n", id, 64))
		}

		// Cargar dirección destino
		v.Code.WriteString(fmt.Sprintf("\tldr x0, =%s\n", id))

		// Cargar valor del origen (expr.addrID)
		v.Code.WriteString(fmt.Sprintf("\tldr x1, =%s\n", expr.addrID))
		switch tipo {
		case "int", "bool":
			v.Code.WriteString("\tldr w2, [x1]\n")
			v.Code.WriteString("\tstr w2, [x0]\n\n")
		case "f64":
			v.Code.WriteString("\tldr d1, [x1]\n")
			v.Code.WriteString("\tstr d1, [x0]\n\n")
		case "string":
			// Copiar string byte a byte
			// en copiar sin limpiar x0 es el destion y x1 es el origen

			v.Code.WriteString("\tbl copiar_sin_limpiar\n\n")
		}
	}

	// fin de la declaracion
	v.Code.WriteString(fmt.Sprintf("\n\t// ---------- Variable %s de tipo %s declarada correctamente ----------\n\n", id, tipo))

	// Guardar variable en entorno
	variable := &Variable{
		Nombre: id,
		Tipo:   tipo,
		Valor:  expr.value,
	}
	v.Env.SetVar(id, variable)

	return nil
}

func isNUM(attr Attr) bool {
	return attr.addrNUM != "" && attr.addrID == ""
}

// ---------------------------------------------------------------
// |DeclarationImplicitVar|
// ---------------------------------------------------------------
func (v *CodeGenVisitor) VisitDeclarationImplicitVar(ctx *parser.DeclarationImplicitVarContext) interface{} {
	id := ctx.ID().GetText()
	tipo := ctx.PrimitiveType().GetText()

	// Validar si ya existe la variable
	if _, ok := v.Env.GetVar(id); ok {
		panic(fmt.Sprintf("Variable ya declarada: %s", id))
	}

	// mensaje de declaracion
	//v.Code.WriteString(fmt.Sprintf("\t// ---------- Declaración de variable: %s ----------\n\n", id))

	// Escribe en .data con valor por defecto según el tipo

	switch tipo {
	case "int", "bool":
		v.Data.WriteString(fmt.Sprintf("%s: .word 0\n", id))
	case "f64":
		v.Data.WriteString(fmt.Sprintf("%s: .double 0.0\n", id)) //
	case "string":
		// declaro la cadena en .bss con skip de 64
		v.BSS.WriteString(fmt.Sprintf("%s: .skip %d\n", id, 64))
		// v.Data.WriteString(fmt.Sprintf("%s: .asciz \"\"\n", id)) // cadena vacía
	default:
		fmt.Fprintf(os.Stderr, "Tipo no soportado en declaración implícita: %s\n", tipo)
	}

	// fin de la declaracion
	//v.Code.WriteString(fmt.Sprintf("\n\t// ---------- Variable %s de tipo %s declarada correctamente ----------\n\n", id, tipo))

	defaultValue := getDefaultValue(tipo)
	variable := &Variable{
		Nombre: id,
		Tipo:   tipo,
		Valor:  defaultValue,
	}
	v.Env.SetVar(id, variable)

	return nil
}

// ---------------------------------------------------------------
// |Print|
// ---------------------------------------------------------------
func (v *CodeGenVisitor) VisitPrint(ctx *parser.PrintContext) interface{} {

	if ctx.PrintList() == nil {
		return nil
	}

	// declaro la cadena de salida con .skip 128
	esal := v.nextSal()
	v.BSS.WriteString(fmt.Sprintf("%s: .skip %d\n", esal, 128))

	// aca ire metiendo los id de las partes de la cadena de salida

	// imprimo donde empieza el proceso
	v.Code.WriteString(fmt.Sprintf("\t// ---------- Print para %s ----------\n\n", esal))

	var parts []string
	for _, expr := range ctx.PrintList().AllExpression() {
		val := v.Visit(expr).(Attr)
		// ver si es addrNUM o addrID
		// fmt.Println(fmt.Sprintf("Valor de addrNUM: %s, addrID: %s, info: %s", val.addrNUM, val.addrID, val.info))
		if val.addrNUM != "" {

			if val.info == "string" {
				// Si es una cadena, la declaro en .data
				temp := v.nextTemp()
				v.Data.WriteString(fmt.Sprintf("%s: .asciz \"%s\"\n", temp, val.value.(string)))
				// añado el temporal a las partes
				parts = append(parts, temp)
			} else if val.info == "int" {
				// Si es un int, lo declaro como .word
				temp := v.nextTemp()
				v.Data.WriteString(fmt.Sprintf("%s: .word %s\n", temp, val.addrNUM))
				// ahora lo convierto a string
				//creo un temporal para la cadena en .bss
				tempStr := v.nextTemp()
				v.BSS.WriteString(fmt.Sprintf("%s: .skip %d\n", tempStr, 64)) // espacio para la cadena
				// en .code llamo a la funcion int_to_str
				v.Code.WriteString(fmt.Sprintf("\tldr x2, =%s\t// direccion del valor a comvertir\n", temp))
				v.Code.WriteString("\tldr w0, [x2]\n") // cargar el valor del int
				// destino el tempStr
				v.Code.WriteString(fmt.Sprintf("\tldr x1, =%s\t// destino de la cadena\n", tempStr))
				v.Code.WriteString("\tbl int_to_str\n") // llamar a la funcion para convertir

				// añado el temporal a las partes, el temporal de destino
				parts = append(parts, tempStr)
			} else if val.info == "f64" {
				// Si es un float, lo declaro como .double
				temp := v.nextTemp()
				v.Data.WriteString(fmt.Sprintf("%s: .double %s\n", temp, val.addrNUM))
				// ahora lo convierto a string
				// creo un temporal para la cadena en .bss
				tempStr := v.nextTemp()
				v.BSS.WriteString(fmt.Sprintf("%s: .skip %d\n", tempStr, 64)) // espacio para la cadena
				// en .code llamo a la funcion float_to_str
				v.Code.WriteString(fmt.Sprintf("\tldr x0, =%s\t// direccion del valor a comvertir\n", temp))
				v.Code.WriteString("\tldr d0, [x0]\n") // cargar el valor del float
				// destino el tempStr
				v.Code.WriteString(fmt.Sprintf("\tldr x1, =%s\t// destino de la cadena\n", tempStr))
				v.Code.WriteString("\tbl float_to_str\n") // llamar a la funcion para convertir

				// añado el temporal a las partes, el temporal de destino
				parts = append(parts, tempStr)

			} else if val.info == "bool" {
				// Si es un bool, lo declaro como .word
				temp := v.nextTemp()
				v.Data.WriteString(fmt.Sprintf("%s: .word %s\n", temp, val.addrNUM))

				// evaluo el valor del bool si es 1 es true y si es 0 es false
				// creo la cadena en .BSS para guardar true o false
				tempStr := v.nextTemp()
				v.BSS.WriteString(fmt.Sprintf("%s: .skip %d\n", tempStr, 64)) // espacio para la cadena

				//empiezo a evaluar
				v.Code.WriteString(fmt.Sprintf("\tldr x7, =%s\n", temp)) // cargar el valor del bool
				v.Code.WriteString("\tldr w0, [x7]\n")                   // cargar
				v.Code.WriteString(fmt.Sprintf("\tldr x1, =%s\t// destino de la cadena\n", tempStr))
				// llamo a la funcion bool_to_str
				v.Code.WriteString("\tbl bool_to_str\n") // llamar a la funcion para convertir

				// añado el temporal a las partes
				parts = append(parts, tempStr)
			} else {
				panic("Tipo de dato no soportado en print: " + val.info)
			}
		} else if val.addrID != "" {
			if val.info == "string" {
				// ya se que viene un id
				// en este caso, solo guardo el id

				parts = append(parts, val.addrID) // añado el id a las partes

			} else if val.info == "int" {
				// hagarro el valor del id y lo comvierto a string

				v.Code.WriteString(fmt.Sprintf("\tldr x2, =%s\t// direccion del valor a comvertir\n", val.addrID))
				v.Code.WriteString("\tldr w0, [x2]\n") // cargar el valor del int
				// creo un temporal para la cadena en .bss
				tempStr := v.nextTemp()
				v.BSS.WriteString(fmt.Sprintf("%s: .skip %d\n", tempStr, 64)) // espacio para la cadena
				// en .code llamo a la funcion int_to_str
				v.Code.WriteString(fmt.Sprintf("\tldr x1, =%s\t// destino de la cadena\n", tempStr))
				v.Code.WriteString("\tbl int_to_str\n") // llamar a la funcion para convertir
				// añado el temporal a las partes, el temporal de destino
				parts = append(parts, tempStr)

			} else if val.info == "f64" {
				// ya se que viene un id
				v.Code.WriteString(fmt.Sprintf("\tldr x0, =%s\t// direccion del valor a comvertir\n", val.addrID))
				v.Code.WriteString("\tldr d0, [x0]\n") // cargar el valor del float
				// creo un temporal para la cadena en .bss
				tempStr := v.nextTemp()
				v.BSS.WriteString(fmt.Sprintf("%s: .skip %d\n", tempStr, 64)) // espacio para la cadena
				// en .code llamo a la funcion float_to_str
				v.Code.WriteString(fmt.Sprintf("\tldr x1, =%s\t// destino de la cadena\n", tempStr))
				v.Code.WriteString("\tbl float_to_str\n") // llamar a la funcion para convertir
				// añado el temporal a las partes, el temporal de destino
				parts = append(parts, tempStr)
			} else if val.info == "bool" {
				// ya se que viene un id
				v.Code.WriteString(fmt.Sprintf("\tldr x7, =%s\t// direccion del valor a comvertir\n", val.addrID))
				v.Code.WriteString("\tldr w0, [x7]\n") // cargar el valor del bool
				// creo un temporal para la cadena en .bss
				tempStr := v.nextTemp()
				v.BSS.WriteString(fmt.Sprintf("%s: .skip %d\n", tempStr, 64)) // espacio para la cadena
				// en .code llamo a la funcion bool_to_str
				v.Code.WriteString(fmt.Sprintf("\tldr x1, =%s\t// destino de la cadena\n", tempStr))
				v.Code.WriteString("\tbl bool_to_str\n") // llamar a la funcion para convertir
				// añado el temporal a las partes, el temporal de destino
				parts = append(parts, tempStr)
			} else {
				panic("Tipo de dato no soportado en print: " + val.info)
			}
		}
	}

	// armo la cadena de salida, necesito los temporales de las partes
	// estes temporales ya los declare y estan en parts
	// Mensaje de que empieza a concatenar todo
	v.Code.WriteString("\n\t// --- Inicio Concatenacion cadena ---\n\n")
	for i, part := range parts {
		// fmt.Println("Parte de cadena:", part)
		if i == 0 {
			// al inicio solo copio el valor a la cadena luego me voy corriendo
			v.Code.WriteString(fmt.Sprintf("\tldr x0, =%s\t// destino\n", esal)) // cargar la direccion de memoria de salida

			v.Code.WriteString("\tmov x20, x0\t// guardar puntero base\n") // ← guardar dirección base

			v.Code.WriteString(fmt.Sprintf("\tldr x1, =%s\t// origen\n", part))  // cargar la primera parte
			v.Code.WriteString("\tmov x4, #128\t// Tamaño del buffer destino\n") // Tamaño de la cadena
			v.Code.WriteString("\tbl copiar_sin_limpiar\n")                      // llamar a la funcion para copiar
		} else {
			v.Code.WriteString("\tmov x0, x20\t// restaurar puntero base\n")

			// avanzar puntero al final
			v.Code.WriteString("\tbl advance_to_end\n") // avanzar al final de la cadena de salida
			// agrego espacio " "
			v.Code.WriteString("\tmov w6, #' '\n") // agregar espacio
			v.Code.WriteString("\tstrb w6, [x0], #1\n")

			// agrego el otro valor a la cadena de salida
			v.Code.WriteString(fmt.Sprintf("\tldr x1, =%s\t// origen\n", part)) // cargar la direccion de memoria de salida
			v.Code.WriteString("\tbl copiar_sin_limpiar\n")                     // llamar a la funcion para copiar
		}
	}
	v.Code.WriteString("\tmov x0, x20\t// restaurar puntero base\n")

	// de ultimo agrego el salto de linea
	v.Code.WriteString("\tbl advance_to_end\n") // avanzar al final de la cadena de salida
	// agrego salto de linea
	// cargo ya el string newline
	v.Code.WriteString("\tldr x1, =newline\n")      // cargar la direccion de memoria de salida
	v.Code.WriteString("\tbl copiar_sin_limpiar\n") // llamar a la funcion para copiar

	v.Code.WriteString("\tmov x1, x20\t// restaurar puntero base en x1\n") // para println
	// llamo a println para imprimir la cadena de salida
	//v.Code.WriteString(fmt.Sprintf("\tldr x1, =%s\n", esal)) // cargar la direccion de memoria de salida
	v.Code.WriteString("\tbl println\n") // llamar a la funcion println

	v.Code.WriteString("\n\t// --- Fin Concatenacion de cadena ---\n\n")

	v.Code.WriteString(fmt.Sprintf("\t// ---------- Fin Print para %s ----------\n\n", esal))

	return nil
}

// ---------------------------------------------------------------
// |NotExpr|
// ---------------------------------------------------------------

func (v *CodeGenVisitor) VisitNotExpr(ctx *parser.NotExprContext) interface{} {
	val := v.Visit(ctx.Expression()).(Attr)

	if val.info == "bool" {
		if isNUM(val) {
			// quiere decir que es un valor inmediato
			// declaro el valor del inmediato
			temp := v.nextTemp()
			v.Data.WriteString(fmt.Sprintf("%s: .word %s\n", temp, val.addrNUM))
			// obtengo el valor desde memoria y cambio el valor
			// val.addrNUM trae 0 o 1, no son enteros son strings entonces solo cambio a 1 o 0
			v.Code.WriteString(fmt.Sprintf("\tldr x0, =%s\n", temp)) // cargar la direccion de memoria del inmediato
			v.Code.WriteString("\tldr w1, [x0]\n")                   // Cargar valor del inmediato
			v.Code.WriteString("\tcmp w1, #0\n")                     // Comparar con 0

			v.Code.WriteString("\tcset w2, eq\t// Si w1 == 0, entonces w2 = 1; si no, w2 = 0\n") // Setear w2 a 1 si es igual a 0, 0 si no

			// el valor de w2 lo meto en temp
			//v.Code.WriteString(fmt.Sprintf("\tldr x0, =%s\n", temp)) // cargar la direccion de memoria del inmediato
			v.Code.WriteString("\tstr w2, [x0]\n")

			return Attr{
				addrID: temp, // Guardar el resultado en la dirección de memoria del inmediato
				info:   "bool",
				value:  !val.value.(bool), // Negación del valor booleano
			} // // Guardar el resultado en la dirección de memoria del inmediato

		} else if !isNUM(val) {
			// quiere decir que es un id
			// obtengo el valor desde memoria y cambio el valor
			v.Code.WriteString(fmt.Sprintf("\tldr x0, =%s\n", val.addrID)) // cargar la direccion de memoria del id
			v.Code.WriteString("\tldr w1, [x0]\n")                         // Cargar valor del id
			v.Code.WriteString("\tcmp w1, #0\n")                           // Comparar con 0

			v.Code.WriteString("\tcset w2, eq\t// Si w1 == 0, entonces w2 = 1; si no, w2 = 0\n") // Setear w2 a 1 si es igual a 0, 0 si no

			// creo un temporal para no moficar el id origen

			temp := v.nextTemp()
			v.Data.WriteString(fmt.Sprintf("%s: .word 0\n", temp))
			// en este tempo guado el resultado y ese id emvio

			// cargo la direccion de memoria del temporal
			v.Code.WriteString(fmt.Sprintf("\tldr x0, =%s\n", temp)) // cargar la direccion de memoria del temporal
			// ahora guardo el valor de w2 en el temporal
			v.Code.WriteString("\tstr w2, [x0]\n") // Guardar el resultado en la dirección de memoria del temporal

			return Attr{
				addrID: temp,
				info:   "bool",
				value:  !val.value.(bool), // Negación del valor booleano
			}

		}

		//return Value{value: !val.value.(bool), info: "bool"} // Negación de un entero
	}

	panic(fmt.Sprintf("Operador unario (!) no soportado para el tipo: %s con valor %v", val.info, val.value))

}

// ---------------------------------------------------------------
// |OpExpr|
// ---------------------------------------------------------------
func (v *CodeGenVisitor) VisitOpExpr(ctx *parser.OpExprContext) interface{} {
	left := v.Visit(ctx.GetLeft()).(Attr)
	right := v.Visit(ctx.GetRight()).(Attr)
	op := ctx.GetOp().GetText()

	// si izquierdo es un valor inmediato y derecho es un id
	if isNUM(left) && !isNUM(right) {
		// sabemos que lado derecho es un id

		// int + int = int
		if left.info == "int" && right.info == "int" {
			lval := left.value.(int64)
			rval := right.value.(int64)
			var result interface{}
			tipoExp := "int"
			switch op {
			case "+":
				result = lval + rval
			case "-":
				result = lval - rval
			case "*":
				result = lval * rval
			case "/":
				if rval == 0 {
					panic("División por cero " + fmt.Sprintf("%d / %d", lval, rval))
				}
				if lval%rval == 0 {
					result = lval / rval
				} else {
					result = float64(lval) / float64(rval)
					tipoExp = "f64"
				}

			case "%":
				if rval == 0 {
					panic("Módulo por cero " + fmt.Sprintf("%d %% %d", lval, rval))
				}
				result = lval % rval
			case "<":
				result = lval < rval
				tipoExp = "bool"
			case "<=":
				result = lval <= rval
				tipoExp = "bool"
			case ">":
				result = lval > rval
				tipoExp = "bool"
			case ">=":
				result = lval >= rval
				tipoExp = "bool"
			case "==":
				result = lval == rval
				tipoExp = "bool"
			case "!=":
				result = lval != rval
				tipoExp = "bool"
			}

			if tipoExp == "int" {

				// ya se que el lado derecho es un id
				// obtengo el valor de memoria del id
				v.Code.WriteString(fmt.Sprintf("\tldr x0, =%s\n", right.addrID))
				v.Code.WriteString("\tldr w1, [x0]\n") // Cargar valor de la variable

				switch op {
				case "+":
					// se que el lado izquierdo es un valor inmediato
					// seraua x0, #valor inmediato
					v.Code.WriteString(fmt.Sprintf("\tadd w2, w1, #%s\n", left.addrNUM))
				case "-":
					// si es negativo el resultado me mostrara numeros raros

					// tengo que restar izquierdo - derecho

					// como sub no permite inmediato en el primer valor, creo un registro temporal
					// rt := v.nextTemp()
					// v.Data.WriteString(fmt.Sprintf("%s: .word %s\n", rt, left.addrNUM))
					// v.Code.WriteString(fmt.Sprintf("\tldr x0, =%s\n", rt))
					// v.Code.WriteString("\tldr w0, [x0]\n") // Cargar valor del inmediato

					// cargo mejor el valor en un registro, con mov
					v.Code.WriteString(fmt.Sprintf("\tmov w0, #%s\n", left.addrNUM))

					v.Code.WriteString("\tsub w2, w0, w1\n")
				case "*":
					// como mul solo permite multiplicacion entre registros, declaro el registro del inmediato
					v.Code.WriteString(fmt.Sprintf("\tmov w3, #%s\n", left.addrNUM))
					v.Code.WriteString("\tmul w2, w3, w1\n")
				case "/":
					v.Code.WriteString(fmt.Sprintf("\tmov w3, #%s\n", left.addrNUM))
					v.Code.WriteString("\tsdiv w2, w3, w1\n")
					// sdiv hacepta valores negativos y decimales

				}

				tpm := v.nextTemp()
				// Temporal donde se guradara el resultado
				v.Data.WriteString(fmt.Sprintf("%s: .word 0\n", tpm))
				v.Code.WriteString(fmt.Sprintf("\tldr x1, =%s\n", tpm))
				v.Code.WriteString("\tstr w2, [x1]\n\n")

				return Attr{
					addrID: tpm,
					info:   tipoExp,
					value:  result,
				}

			} else if tipoExp == "f64" {
				// en este punto ya se que el resultado de los dos ints da un double

				// el lado derecho es un id
				// obtengo el valor de memoria del id y lo comvierto a double
				v.Code.WriteString(fmt.Sprintf("\tldr x0, =%s\n", right.addrID))
				v.Code.WriteString("\tldr w1, [x0]\n") // Cargar valor de la variable
				// uso scvtf para convertir el valor inmediato a double
				v.Code.WriteString("scvtf d1, w1\n") // Convertir int a double

				// ya se yo que solo se da en la divicion

				switch op {
				case "/":
					// asi que comvierto el lado izquierdo que es un valor inmediato intero a double
					// en arm no puedo declarar un float inmediato
					// asi que lo declaro en un registro temporal
					etq := v.nextTemp()
					v.Data.WriteString(fmt.Sprintf("%s: .double %s.0\n", etq, left.addrNUM))
					v.Code.WriteString(fmt.Sprintf("\tldr x0, =%s\n", etq))
					v.Code.WriteString("\tldr d2, [x0]\n")
					v.Code.WriteString("\tfdiv d2, d2, d1\n")
				}

				tpm := v.nextTemp()
				// Temporal donde se guradara el resultado f64
				v.Data.WriteString(fmt.Sprintf("%s: .double 0.0\n", tpm))
				v.Code.WriteString(fmt.Sprintf("\tldr x1, =%s\n", tpm))
				v.Code.WriteString("\tstr d2, [x1]\n\n")

				return Attr{
					addrID: tpm,
					info:   tipoExp,
					value:  result,
				}

			} else if tipoExp == "bool" {
				// ya se que el lado derecho es un id
				// obtengo el valor de memoria del id
				v.Code.WriteString(fmt.Sprintf("\tldr x1, =%s\n", right.addrID))
				v.Code.WriteString("\tldr w1, [x1]\n") // Cargar valor de la variable

				switch op {
				case "<":
					// necesito poner el lado izquierdo como registro
					v.Code.WriteString(fmt.Sprintf("\tmov w2, #%s\n", left.addrNUM))
					v.Code.WriteString("\tcmp w2, w1\n") // Comparar con el inmediato
					v.Code.WriteString("\tcset w3, lt\n")
				case "<=":
					v.Code.WriteString(fmt.Sprintf("\tmov w2, #%s\n", left.addrNUM))
					v.Code.WriteString("\tcmp w2, w1\n")
					v.Code.WriteString("\tcset w3, le\n")
				case ">":
					v.Code.WriteString(fmt.Sprintf("\tmov w2, #%s\n", left.addrNUM))
					v.Code.WriteString("\tcmp w2, w1\n")
					v.Code.WriteString("\tcset w3, gt\n")
				case ">=":
					v.Code.WriteString(fmt.Sprintf("\tmov w2, #%s\n", left.addrNUM))
					v.Code.WriteString("\tcmp w2, w1\n")
					v.Code.WriteString("\tcset w3, ge\n")
				case "==":
					v.Code.WriteString(fmt.Sprintf("\tmov w2, #%s\n", left.addrNUM))
					v.Code.WriteString("\tcmp w2, w1\n")
					v.Code.WriteString("\tcset w3, eq\n")
				case "!=":
					v.Code.WriteString(fmt.Sprintf("\tmov w2, #%s\n", left.addrNUM))
					v.Code.WriteString("\tcmp w2, w1\n")
					v.Code.WriteString("\tcset w3, ne\n")
				}

				tpm := v.nextTemp()
				// Temporal donde se guradara el resultado bool
				v.Data.WriteString(fmt.Sprintf("%s: .word 0\n", tpm))
				v.Code.WriteString(fmt.Sprintf("\tldr x2, =%s\n", tpm))
				v.Code.WriteString("\tstr w3, [x2]\n\n")

				return Attr{
					addrID: tpm,
					info:   tipoExp,
					value:  result,
				}
			}

		}

		// int + float = float | float + int = float
		if (left.info == "int" && right.info == "f64") || (left.info == "f64" && right.info == "int") {
			// Caso: int y float, siempre se convierte a float
			var lval, rval float64
			if left.info == "int" {
				lval = float64(left.value.(int64))
			} else {
				lval = left.value.(float64)
			}

			if right.info == "int" {
				rval = float64(right.value.(int64))
			} else {
				rval = right.value.(float64)
			}

			var result interface{}
			tipoExp := "f64"
			switch op {
			case "+":
				result = lval + rval
			case "-":
				result = lval - rval
			case "*":
				result = lval * rval
			case "/":
				if rval == 0.0 {
					panic("División por cero " + fmt.Sprintf("%f / %f", lval, rval))
				}
				result = lval / rval
			case "<":
				result = lval < rval
				tipoExp = "bool"
			case "<=":
				result = lval <= rval
				tipoExp = "bool"
			case ">":
				result = lval > rval
				tipoExp = "bool"
			case ">=":
				result = lval >= rval
				tipoExp = "bool"
			case "==":
				result = lval == rval
				tipoExp = "bool"
			case "!=":
				result = lval != rval
				tipoExp = "bool"
			}

			if tipoExp == "f64" {

				// veo cual de los dos lados es el f64
				if left.info == "f64" {
					// el lado izquierdo es un f64

					// declaro mi registro temporal para el inmediato del lado izquierdo
					etq := v.nextTemp()
					v.Data.WriteString(fmt.Sprintf("%s: .double %s\n", etq, left.addrNUM))
					v.Code.WriteString(fmt.Sprintf("\tldr x0, =%s\n", etq))
					v.Code.WriteString("\tldr d1, [x0]\n")

					// lado derecho es un id int
					v.Code.WriteString(fmt.Sprintf("\tldr x1, =%s\n", right.addrID))
					v.Code.WriteString("\tldr w1, [x1]\n")
					v.Code.WriteString("scvtf d2, w1\n") // Convertir int a double

					switch op {
					case "+":
						v.Code.WriteString(("\tfadd d3, d1, d2\n"))
					case "-":
						v.Code.WriteString(("\tfsub d3, d1, d2\n"))
					case "*":
						v.Code.WriteString(("\tfmul d3, d1, d2\n"))
					case "/":
						v.Code.WriteString(("\tfdiv d3, d1, d2\n"))
					}

					tpm := v.nextTemp()
					// Temporal donde se guradara el resultado f64
					v.Data.WriteString(fmt.Sprintf("%s: .double 0.0\n", tpm))
					v.Code.WriteString(fmt.Sprintf("\tldr x2, =%s\n", tpm))
					v.Code.WriteString("\tstr d3, [x2]\n\n")

					return Attr{
						addrID: tpm,
						info:   tipoExp,
						value:  result,
					}

				} else if right.info == "f64" {
					// el lado derecho es un f64

					// declaro mi registro temporal para el inmediato y lo casteo a double
					// etq := v.nextTemp() + "ss"
					// v.Data.WriteString(fmt.Sprintf("%s: .double %s\n", etq, left.addrNUM))

					// se que trae un int del lado izquierdo, lo asigno a un registro w con mov
					v.Code.WriteString("\tmov w1, #" + left.addrNUM + "\n")
					// lo casteo a double
					v.Code.WriteString("scvtf d1, w1\n") // Convertir int a double

					// lado derecho es un id double
					v.Code.WriteString(fmt.Sprintf("\tldr x1, =%s\n", right.addrID))
					v.Code.WriteString("\tldr d2, [x1]\n") // Cargar valor de la variable

					switch op {
					case "+":
						v.Code.WriteString(("\tfadd d3, d1, d2\n"))
					case "-":
						v.Code.WriteString(("\tfsub d3, d1, d2\n"))
					case "*":
						v.Code.WriteString(("\tfmul d3, d1, d2\n"))
					case "/":
						v.Code.WriteString(("\tfdiv d3, d1, d2\n"))
					}

					tpm := v.nextTemp()
					// Temporal donde se guradara el resultado f64
					v.Data.WriteString(fmt.Sprintf("%s: .double 0.0\n", tpm))
					v.Code.WriteString(fmt.Sprintf("\tldr x2, =%s\n", tpm))
					v.Code.WriteString("\tstr d3, [x2]\n\n")

					return Attr{
						addrID: tpm,
						info:   tipoExp,
						value:  result,
					}

				}

			} else if tipoExp == "bool" {
				// veo cual de los dos lados es el f64
				if left.info == "f64" {
					// el lado izquierdo es un f64

					// declaro mi registro temporal para el inmediato
					etq := v.nextTemp()
					v.Data.WriteString(fmt.Sprintf("%s: .double %s\n", etq, left.addrNUM))
					v.Code.WriteString(fmt.Sprintf("\tldr x0, =%s\n", etq))
					v.Code.WriteString("\tldr d1, [x0]\n")

					// lado derecho es un id
					v.Code.WriteString(fmt.Sprintf("\tldr x1, =%s\n", right.addrID))
					v.Code.WriteString("\tldr w1, [x1]\n")
					v.Code.WriteString("scvtf d2, w1\n") // Convertir int a double

					switch op {
					case "<":
						v.Code.WriteString("\tfcmp d1, d2\n")
						v.Code.WriteString("\tcset w3, lt\n")
					case "<=":
						v.Code.WriteString("\tfcmp d1, d2\n")
						v.Code.WriteString("\tcset w3, le\n")
					case ">":
						v.Code.WriteString("\tfcmp d1, d2\n")
						v.Code.WriteString("\tcset w3, gt\n")
					case ">=":
						v.Code.WriteString("\tfcmp d1, d2\n")
						v.Code.WriteString("\tcset w3, ge\n")
					case "==":
						v.Code.WriteString("\tfcmp d1, d2\n")
						v.Code.WriteString("\tcset w3, eq\n")
					case "!=":
						v.Code.WriteString("\tfcmp d1, d2\n")
						v.Code.WriteString("\tcset w3, ne\n")
					}

					tpm := v.nextTemp()
					// Temporal donde se guradara el resultado bool
					v.Data.WriteString(fmt.Sprintf("%s: .word 0\n", tpm))
					v.Code.WriteString(fmt.Sprintf("\tldr x2, =%s\n", tpm))
					v.Code.WriteString("\tstr w3, [x2]\n\n")

					return Attr{
						addrID: tpm,
						info:   tipoExp,
						value:  result,
					}

				} else if right.info == "f64" {
					// el lado derecho es un f64

					// declaro mi registro temporal para el inmediato y lo casteo a double
					etq := v.nextTemp()
					v.Data.WriteString(fmt.Sprintf("%s: .double %s\n", etq, left.addrNUM))
					// se que trae un int del lado izquierdo, lo asigno a un registro w con mov
					v.Code.WriteString("\tmov w1, #" + left.addrNUM + "\n")
					// lo casteo a double
					v.Code.WriteString("scvtf d1, w1\n") // Convertir int a double

					// lado derecho es un id double
					v.Code.WriteString(fmt.Sprintf("\tldr x1, =%s\n", right.addrID))
					v.Code.WriteString("\tldr d2, [x1]\n") // Cargar valor de la variable

					switch op {
					case "<":
						v.Code.WriteString("\tfcmp d1, d2\n")
						v.Code.WriteString("\tcset w3, lt\n")
					case "<=":
						v.Code.WriteString("\tfcmp d1, d2\n")
						v.Code.WriteString("\tcset w3, le\n")
					case ">":
						v.Code.WriteString("\tfcmp d1, d2\n")
						v.Code.WriteString("\tcset w3, gt\n")
					case ">=":
						v.Code.WriteString("\tfcmp d1, d2\n")
						v.Code.WriteString("\tcset w3, ge\n")
					case "==":
						v.Code.WriteString("\tfcmp d1, d2\n")
						v.Code.WriteString("\tcset w3, eq\n")
					case "!=":
						v.Code.WriteString("\tfcmp d1, d2\n")
						v.Code.WriteString("\tcset w3, ne\n")
					}

					tpm := v.nextTemp()
					// Temporal donde se guradara el resultado bool
					v.Data.WriteString(fmt.Sprintf("%s: .word 0\n", tpm))
					v.Code.WriteString(fmt.Sprintf("\tldr x2, =%s\n", tpm))
					v.Code.WriteString("\tstr w3, [x2]\n\n")

					return Attr{
						addrID: tpm,
						info:   tipoExp,
						value:  result,
					}

				}
			}

		}

		// f64 + f64 = f64
		if left.info == "f64" && right.info == "f64" {
			var lval, rval float64
			lval = left.value.(float64)
			rval = right.value.(float64)

			var result interface{}
			tipoExp := "f64"
			switch op {
			case "+":
				result = lval + rval
			case "-":
				result = lval - rval
			case "*":
				result = lval * rval
			case "/":
				if rval == 0.0 {
					panic("División por cero " + fmt.Sprintf("%f / %f", lval, rval))
				}
				result = lval / rval
			case "<":
				result = lval < rval
				tipoExp = "bool"
			case "<=":
				result = lval <= rval
				tipoExp = "bool"
			case ">":
				result = lval > rval
				tipoExp = "bool"
			case ">=":
				result = lval >= rval
				tipoExp = "bool"
			case "==":
				result = lval == rval
				tipoExp = "bool"
			case "!=":
				result = lval != rval
				tipoExp = "bool"
			}

			if tipoExp == "f64" {

				// ya no necesito ver cual lado es el f64, ya que ambos son f64
				// se que el lado derecho es un id

				// el lado izquierdo es un f64

				// declaro mi registro temporal para el inmediato
				etq := v.nextTemp()
				v.Data.WriteString(fmt.Sprintf("%s: .double %s\n", etq, left.addrNUM))
				v.Code.WriteString(fmt.Sprintf("\tldr x0, =%s\n", etq))
				v.Code.WriteString("\tldr d1, [x0]\n")

				// lado derecho es un id double
				v.Code.WriteString(fmt.Sprintf("\tldr x1, =%s\n", right.addrID))
				v.Code.WriteString("\tldr d2, [x1]\n") // Cargar valor de la variable

				switch op {
				case "+":
					v.Code.WriteString(("\tfadd d3, d1, d2\n"))
				case "-":
					v.Code.WriteString(("\tfsub d3, d1, d2\n"))
				case "*":
					v.Code.WriteString(("\tfmul d3, d1, d2\n"))
				case "/":
					v.Code.WriteString(("\tfdiv d3, d1, d2\n"))
				}

				tpm := v.nextTemp()
				// Temporal donde se guradara el resultado f64
				v.Data.WriteString(fmt.Sprintf("%s: .double 0.0\n", tpm))
				v.Code.WriteString(fmt.Sprintf("\tldr x2, =%s\n", tpm))
				v.Code.WriteString("\tstr d3, [x2]\n\n")

				return Attr{
					addrID: tpm,
					info:   tipoExp,
					value:  result,
				}

			} else if tipoExp == "bool" {
				// ya no necesito ver cual lado es el f64, ya que ambos son f64
				// se que el lado derecho es un id

				// el lado izquierdo es un f64

				// declaro mi registro temporal para el inmediato
				etq := v.nextTemp()
				v.Data.WriteString(fmt.Sprintf("%s: .double %s\n", etq, left.addrNUM))
				v.Code.WriteString(fmt.Sprintf("\tldr x0, =%s\n", etq))
				v.Code.WriteString("\tldr d1, [x0]\n")

				// lado derecho es un id double
				v.Code.WriteString(fmt.Sprintf("\tldr x1, =%s\n", right.addrID))
				v.Code.WriteString("\tldr d2, [x1]\n") // Cargar valor de la variable

				switch op {
				case "<":
					v.Code.WriteString("\tfcmp d1, d2\n")
					v.Code.WriteString("\tcset w3, lt\n")
				case "<=":
					v.Code.WriteString("\tfcmp d1, d2\n")
					v.Code.WriteString("\tcset w3, le\n")
				case ">":
					v.Code.WriteString("\tfcmp d1, d2\n")
					v.Code.WriteString("\tcset w3, gt\n")
				case ">=":
					v.Code.WriteString("\tfcmp d1, d2\n")
					v.Code.WriteString("\tcset w3, ge\n")
				case "==":
					v.Code.WriteString("\tfcmp d1, d2\n")
					v.Code.WriteString("\tcset w3, eq\n")
				case "!=":
					v.Code.WriteString("\tfcmp d1, d2\n")
					v.Code.WriteString("\tcset w3, ne\n")
				}

				tpm := v.nextTemp()
				// Temporal donde se guradara el resultado bool
				v.Data.WriteString(fmt.Sprintf("%s: .word 0\n", tpm))
				v.Code.WriteString(fmt.Sprintf("\tldr x2, =%s\n", tpm))
				v.Code.WriteString("\tstr w3, [x2]\n\n")

				return Attr{
					addrID: tpm,
					info:   tipoExp,
					value:  result,
				}

			}
		}

		// string + string = string
		if left.info == "string" && right.info == "string" {
			var lval, rval string
			lval = left.value.(string)
			rval = right.value.(string)

			var result interface{}
			tipoExp := "string"
			switch op {
			case "+":
				result = lval + rval
			case "<":
				result = lval < rval
				tipoExp = "bool"
			case "<=":
				result = lval <= rval
				tipoExp = "bool"
			case ">":
				result = lval > rval
				tipoExp = "bool"
			case ">=":
				result = lval >= rval
				tipoExp = "bool"
			case "==":
				result = lval == rval
				tipoExp = "bool"
			case "!=":
				result = lval != rval
				tipoExp = "bool"
			}

			if tipoExp == "string" {
				// uso copiar sin limpiar para concatenar cadenas
				// muevo el puntero para copiar donde termine la otra cadena

				// cadena final, sera de 64
				cfinal := v.nextTemp()
				v.BSS.WriteString(fmt.Sprintf("%s: .skip %d\n", cfinal, 64))           // espacio para la cadena de salida
				v.Code.WriteString(fmt.Sprintf("\tldr x0, =%s\t// destino\n", cfinal)) // cargar la direccion de memoria de salida
				v.Code.WriteString("\tmov x20, x0\t// guardar puntero base\n")
				// lado izquierdo es un inmediato string
				etq := v.nextTemp()
				v.Data.WriteString(fmt.Sprintf("%s: .asciz \"%s\"\n", etq, lval))
				v.Code.WriteString(fmt.Sprintf("\tldr x1, =%s\t// origen\n", etq))
				v.Code.WriteString("\tmov x4, #64\t// Tamaño del buffer destino\n")
				v.Code.WriteString("\tbl copiar_sin_limpiar\n")
				v.Code.WriteString("\tmov x0, x20\t// restaurar puntero base\n")
				v.Code.WriteString("\tbl advance_to_end\n") // avanzar al final de la cadena de salida

				// se que lado derecho es un id
				v.Code.WriteString(fmt.Sprintf("\tldr x1, =%s\t// origen\n", right.addrID)) // cargar la direccion de memoria del id
				v.Code.WriteString("\tbl copiar_sin_limpiar\n")

				return Attr{
					addrID: cfinal,
					info:   tipoExp,
					value:  result,
				}
			} else if tipoExp == "bool" {

				// resultado de comparacion de strings bool
				resBool := v.nextTemp()
				v.Data.WriteString(fmt.Sprintf("%s: .word 0\n", resBool))

				// obtengo las dos direcciones de memoria de los strings
				// lado izquierdo es un inmediato string
				etq := v.nextTemp()
				v.Data.WriteString(fmt.Sprintf("%s: .asciz \"%s\"\n", etq, lval))
				v.Code.WriteString(fmt.Sprintf("\tldr x0, =%s\n", etq)) // cargar la direccion de memoria del inmediato

				// ya se que el lado derecho es un id
				// obtengo el valor de memoria del id
				v.Code.WriteString(fmt.Sprintf("\tldr x1, =%s\n", right.addrID))

				switch op {
				case "<":
					// mando a llamar a la funcion strcmp_lt
					v.Code.WriteString("\tbl strcmp_lt\n")
					// el resultado se guarda en w0
					v.Code.WriteString(fmt.Sprintf("\tldr x2, =%s\n", resBool)) // cargar la direccion de memoria del resultado
					v.Code.WriteString("\tstr w0, [x2]\n")                      // guardar el resultado en la direccion de memoria

				case "<=":
					// mando a llamar a la funcion strcmp_le
					v.Code.WriteString("\tbl strcmp_le\n")
					// el resultado se guarda en w0
					v.Code.WriteString(fmt.Sprintf("\tldr x2, =%s\n", resBool)) // cargar la direccion de memoria del resultado
					v.Code.WriteString("\tstr w0, [x2]\n")                      // guardar el resultado en la direccion de memoria

				case ">":
					// mando a llamar a la funcion strcmp_gt
					v.Code.WriteString("\tbl strcmp_gt\n")
					// el resultado se guarda en w0
					v.Code.WriteString(fmt.Sprintf("\tldr x2, =%s\n", resBool)) // cargar la direccion de memoria del resultado
					v.Code.WriteString("\tstr w0, [x2]\n")                      // guardar el resultado en la direccion de memoria

				case ">=":
					// mando a llamar a la funcion strcmp_ge
					v.Code.WriteString("\tbl strcmp_ge\n")
					// el resultado se guarda en w0
					v.Code.WriteString(fmt.Sprintf("\tldr x2, =%s\n", resBool)) // cargar la direccion de memoria del resultado
					v.Code.WriteString("\tstr w0, [x2]\n")                      // guardar el resultado en la direccion de memoria

				case "==":
					// mando a llamar a la funcion strcmp_eq
					v.Code.WriteString("\tbl strcmp_eq\n")
					// el resultado se guarda en w0
					v.Code.WriteString(fmt.Sprintf("\tldr x2, =%s\n", resBool)) // cargar la direccion de memoria del resultado
					v.Code.WriteString("\tstr w0, [x2]\n")                      // guardar el resultado en la direccion de memoria

				case "!=":
					// mando a llamar a la funcion strcmp_ne
					v.Code.WriteString("\tbl strcmp_ne\n")
					// el resultado se guarda en w0
					v.Code.WriteString(fmt.Sprintf("\tldr x2, =%s\n", resBool)) // cargar la direccion de memoria del resultado
					v.Code.WriteString("\tstr w0, [x2]\n")                      // guardar el resultado en la direccion de memoria

				}

				// tpm := v.nextTemp()
				// // Temporal donde se guradara el resultado bool
				// v.Data.WriteString(fmt.Sprintf("%s: .word 0\n", tpm))
				// v.Code.WriteString(fmt.Sprintf("\tldr x2, =%s\n", tpm))
				// v.Code.WriteString("\tstr w3, [x2]\n\n")

				return Attr{
					addrID: resBool,
					info:   tipoExp,
					value:  result,
				}
			}

		}

		// bool + bool = bool
		if left.info == "bool" && right.info == "bool" {

			var lval, rval bool
			lval = left.value.(bool)
			rval = right.value.(bool)

			var result interface{}
			tipoExp := "bool"
			switch op {
			case "==":
				result = lval == rval
			case "!=":
				result = lval != rval
			case "&&":
				result = lval && rval
			case "||":
				result = lval || rval
			}

			if tipoExp == "bool" {
				// lado izquierdo es un inmediato bool
				etq := v.nextTemp()
				v.Data.WriteString(fmt.Sprintf("%s: .word %s\n", etq, left.addrNUM))
				v.Code.WriteString(fmt.Sprintf("\tldr x0, =%s\n", etq)) // cargar la direccion de memoria del inmediato
				v.Code.WriteString("\tldr w1, [x0]\n")                  // cargar el valor del inmediato

				// lado derecho es un id
				v.Code.WriteString(fmt.Sprintf("\tldr x0, =%s\n", right.addrID)) // cargar la direccion de memoria del id
				v.Code.WriteString("\tldr w2, [x0]\n")                           // cargar el valor del id
				switch op {
				case "==":
					v.Code.WriteString("\tcmp w1, w2\n")  // comparar los dos valores
					v.Code.WriteString("\tcset w3, eq\n") // establecer w2 a 1 si son iguales, 0 si no
				case "!=":
					v.Code.WriteString("\tcmp w1, w2\n")  // comparar los dos valores
					v.Code.WriteString("\tcset w3, ne\n") // establecer w2 a 1 si son diferentes, 0 si no
				case "&&":
					v.Code.WriteString("\tand w3, w1, w2\n") // establecer w2 a 1 si ambos son verdaderos, 0 si alguno es falso
				case "||":
					v.Code.WriteString("\torr w3, w1, w2\n") // establecer w2 a 1 si alguno es verdadero, 0 si ambos son falsos

				}

				tpm := v.nextTemp()
				// Temporal donde se guradara el resultado bool
				v.Data.WriteString(fmt.Sprintf("%s: .word 0\n", tpm))
				v.Code.WriteString(fmt.Sprintf("\tldr x2, =%s\n", tpm))
				v.Code.WriteString("\tstr w3, [x2]\n\n")

				fmt.Println("Resultado de la operación booleano:", result)
				return Attr{
					addrID: tpm,
					info:   tipoExp,
					value:  result,
				}
			}

		}

	}

	// si izquierdo es un id y dercho es un valor inmediato
	if !isNUM(left) && isNUM(right) {
		// sabemos que lado izquierdo es un id

		// int + int = int
		if left.info == "int" && right.info == "int" {
			lval := left.value.(int64)
			rval := right.value.(int64)
			var result interface{}
			tipoExp := "int"
			switch op {
			case "+":
				result = lval + rval
			case "-":
				result = lval - rval
			case "*":
				result = lval * rval
			case "/":
				if rval == 0 {
					panic("División por cero " + fmt.Sprintf("%d / %d", lval, rval))
				}
				if lval%rval == 0 {
					result = lval / rval
				} else {
					result = float64(lval) / float64(rval)
					tipoExp = "f64"
				}

			case "%":
				if rval == 0 {
					panic("Módulo por cero " + fmt.Sprintf("%d %% %d", lval, rval))
				}
				result = lval % rval
			case "<":
				result = lval < rval
				tipoExp = "bool"
			case "<=":
				result = lval <= rval
				tipoExp = "bool"
			case ">":
				result = lval > rval
				tipoExp = "bool"
			case ">=":
				result = lval >= rval
				tipoExp = "bool"
			case "==":
				result = lval == rval
				tipoExp = "bool"
			case "!=":
				result = lval != rval
				tipoExp = "bool"
			}

			if tipoExp == "int" {
				// se que el lado izquierdo es un id

				// obtengo el valor de memoria del id
				v.Code.WriteString(fmt.Sprintf("\tldr x0, =%s\n", left.addrID))
				v.Code.WriteString("\tldr w1, [x0]\n") // Cargar valor de la variable

				switch op {
				case "+":
					// se que el lado derecho es un valor inmediato
					// sera x0, #valor inmediato
					v.Code.WriteString(fmt.Sprintf("\tadd w2, w1, #%s\n", right.addrNUM))
				case "-":
					// si es negativo el resultado me mostrara numeros raros

					// tengo que restar izquierdo - derecho

					// cargo mejor el valor en un registro, con mov
					v.Code.WriteString(fmt.Sprintf("\tmov w0, #%s\n", right.addrNUM))

					v.Code.WriteString("\tsub w2, w1, w0\n")
				case "*":
					// como mul solo permite multiplicacion entre registros, declaro el registro del inmediato
					v.Code.WriteString(fmt.Sprintf("\tmov w3, #%s\n", right.addrNUM))
					v.Code.WriteString("\tmul w2, w3, w1\n")
				case "/":
					v.Code.WriteString(fmt.Sprintf("\tmov w3, #%s\n", right.addrNUM))
					v.Code.WriteString("\tsdiv w2, w1, w3\n")
					// sdiv hacepta valores negativos y decimales

				}

				tpm := v.nextTemp()
				// Temporal donde se guradara el resultado
				v.Data.WriteString(fmt.Sprintf("%s: .word 0\n", tpm))
				v.Code.WriteString(fmt.Sprintf("\tldr x1, =%s\n", tpm))
				v.Code.WriteString("\tstr w2, [x1]\n\n")

				return Attr{
					addrID: tpm,
					info:   tipoExp,
					value:  result,
				}
			} else if tipoExp == "f64" {
				// en este punto ya se que el resultado de los dos ints da un double

				// el lado izquierdo es un id
				// obtengo el valor de memoria del id y lo comvierto a double
				v.Code.WriteString(fmt.Sprintf("\tldr x0, =%s\n", left.addrID))
				v.Code.WriteString("\tldr w1, [x0]\n") // Cargar valor de la variable
				// uso scvtf para convertir el valor inmediato a double
				v.Code.WriteString("scvtf d1, w1\n") // Convertir int a double

				// ya se yo que solo se da en la divicion

				switch op {
				case "/":
					// asi que comvierto el lado derecho que es un valor inmediato intero a double
					// en arm no puedo declarar un float inmediato
					// asi que lo declaro en un registro temporal
					etq := v.nextTemp()
					v.Data.WriteString(fmt.Sprintf("%s: .double %s.0\n", etq, right.addrNUM))
					v.Code.WriteString(fmt.Sprintf("\tldr x0, =%s\n", etq))
					v.Code.WriteString("\tldr d2, [x0]\n")
					v.Code.WriteString("\tfdiv d2, d1, d2\n")
				}

				tpm := v.nextTemp()
				// Temporal donde se guradara el resultado f64
				v.Data.WriteString(fmt.Sprintf("%s: .double 0.0\n", tpm))
				v.Code.WriteString(fmt.Sprintf("\tldr x1, =%s\n", tpm))
				v.Code.WriteString("\tstr d2, [x1]\n\n")

				return Attr{
					addrID: tpm,
					info:   tipoExp,
					value:  result,
				}

			} else if tipoExp == "bool" {
				// ya se que el lado izquierdo es un id
				// obtengo el valor de memoria del id
				v.Code.WriteString(fmt.Sprintf("\tldr x1, =%s\n", left.addrID))
				v.Code.WriteString("\tldr w1, [x1]\n") // Cargar valor de la variable

				switch op {
				case "<":
					// necesito poner el lado izquierdo como registro
					v.Code.WriteString(fmt.Sprintf("\tmov w2, #%s\n", right.addrNUM))
					v.Code.WriteString("\tcmp w1, w2\n") // Comparar con el inmediato
					v.Code.WriteString("\tcset w3, lt\n")
				case "<=":
					v.Code.WriteString(fmt.Sprintf("\tmov w2, #%s\n", right.addrNUM))
					v.Code.WriteString("\tcmp w1, w2\n")
					v.Code.WriteString("\tcset w3, le\n")
				case ">":
					v.Code.WriteString(fmt.Sprintf("\tmov w2, #%s\n", right.addrNUM))
					v.Code.WriteString("\tcmp w1, w2\n")
					v.Code.WriteString("\tcset w3, gt\n")
				case ">=":
					v.Code.WriteString(fmt.Sprintf("\tmov w2, #%s\n", right.addrNUM))
					v.Code.WriteString("\tcmp w1, w2\n")
					v.Code.WriteString("\tcset w3, ge\n")
				case "==":
					v.Code.WriteString(fmt.Sprintf("\tmov w2, #%s\n", right.addrNUM))
					v.Code.WriteString("\tcmp w1, w2\n")
					v.Code.WriteString("\tcset w3, eq\n")
				case "!=":
					v.Code.WriteString(fmt.Sprintf("\tmov w2, #%s\n", right.addrNUM))
					v.Code.WriteString("\tcmp w1, w2\n")
					v.Code.WriteString("\tcset w3, ne\n")
				}

				tpm := v.nextTemp()
				// Temporal donde se guradara el resultado bool
				v.Data.WriteString(fmt.Sprintf("%s: .word 0\n", tpm))
				v.Code.WriteString(fmt.Sprintf("\tldr x2, =%s\n", tpm))
				v.Code.WriteString("\tstr w3, [x2]\n\n")

				return Attr{
					addrID: tpm,
					info:   tipoExp,
					value:  result,
				}
			}
		}

		// int + float = float | float + int = float
		if (left.info == "int" && right.info == "f64") || (left.info == "f64" && right.info == "int") {
			// Caso: int y float, siempre se convierte a float
			var lval, rval float64
			if left.info == "int" {
				lval = float64(left.value.(int64))
			} else {
				lval = left.value.(float64)
			}

			if right.info == "int" {
				rval = float64(right.value.(int64))
			} else {
				rval = right.value.(float64)
			}

			var result interface{}
			tipoExp := "f64"

			switch op {
			case "+":
				result = lval + rval
			case "-":
				result = lval - rval
			case "*":
				result = lval * rval
			case "/":
				if rval == 0.0 {
					panic("División por cero " + fmt.Sprintf("%f / %f", lval, rval))
				}
				result = lval / rval
			case "<":
				result = lval < rval
				tipoExp = "bool"
			case "<=":
				result = lval <= rval
				tipoExp = "bool"
			case ">":
				result = lval > rval
				tipoExp = "bool"
			case ">=":
				result = lval >= rval
				tipoExp = "bool"
			case "==":
				result = lval == rval
				tipoExp = "bool"
			case "!=":
				result = lval != rval
				tipoExp = "bool"
			}

			if tipoExp == "f64" {
				// veo cual de los dos lados es el f64
				if left.info == "f64" {
					// el lado izquierdo es un f64

					// declaro mi registro temporal para el inmediato
					// que ahora el int inmediato es el lado derecho
					v.Code.WriteString("\tmov w1, #" + right.addrNUM + "\n")
					// lo casteo a double
					v.Code.WriteString("scvtf d1, w1\n") // Convertir int a double

					// lado izquierdo es un id f64
					v.Code.WriteString(fmt.Sprintf("\tldr x1, =%s\n", left.addrID))
					v.Code.WriteString("\tldr d2, [x1]\n")

					switch op {
					case "+":
						v.Code.WriteString(("\tfadd d3, d2, d1\n"))
					case "-":
						v.Code.WriteString(("\tfsub d3, d2, d1\n"))
					case "*":
						v.Code.WriteString(("\tfmul d3, d2, d1\n"))
					case "/":
						v.Code.WriteString(("\tfdiv d3, d2, d1\n"))
					}

					tpm := v.nextTemp()
					// Temporal donde se guradara el resultado f64
					v.Data.WriteString(fmt.Sprintf("%s: .double 0.0\n", tpm))
					v.Code.WriteString(fmt.Sprintf("\tldr x2, =%s\n", tpm))
					v.Code.WriteString("\tstr d3, [x2]\n\n")

					return Attr{
						addrID: tpm,
						info:   tipoExp,
						value:  result,
					}

				} else if right.info == "f64" {
					// el lado derecho es un f64

					// declaro mi registro temporal para el inmediato y lo casteo a double
					// el inmediato es el lado derecho
					// etq := v.nextTemp()
					// v.Data.WriteString(fmt.Sprintf("%s: .double %s\n", etq, right.addrNUM))

					// se que trae un f64 inmediato del lado derecho, creo el temporal
					etq := v.nextTemp()
					v.Data.WriteString(fmt.Sprintf("%s: .double %s\n", etq, right.addrNUM))
					v.Code.WriteString(fmt.Sprintf("\tldr x0, =%s\n", etq))
					v.Code.WriteString("\tldr d1, [x0]\n")

					// lado izquierdo es un id int
					v.Code.WriteString(fmt.Sprintf("\tldr x1, =%s\n", left.addrID))
					v.Code.WriteString("\tldr w1, [x1]\n")
					v.Code.WriteString("scvtf d2, w1\n") // Convertir int a double

					switch op {
					case "+":
						v.Code.WriteString(("\tfadd d3, d2, d1\n"))
					case "-":
						v.Code.WriteString(("\tfsub d3, d2, d1\n"))
					case "*":
						v.Code.WriteString(("\tfmul d3, d2, d1\n"))
					case "/":
						v.Code.WriteString(("\tfdiv d3, d2, d1\n"))
					}

					tpm := v.nextTemp()
					// Temporal donde se guradara el resultado f64
					v.Data.WriteString(fmt.Sprintf("%s: .double 0.0\n", tpm))
					v.Code.WriteString(fmt.Sprintf("\tldr x2, =%s\n", tpm))
					v.Code.WriteString("\tstr d3, [x2]\n\n")

					return Attr{
						addrID: tpm,
						info:   tipoExp,
						value:  result,
					}

				}
			} else if tipoExp == "bool" {
				// veo cual de los dos lados es el f64
				if left.info == "f64" {
					// el lado izquierdo es un f64

					// declaro mi registro temporal para el inmediato
					// que ahora el int inmediato es el lado derecho
					v.Code.WriteString("\tmov w1, #" + right.addrNUM + "\n")
					// lo casteo a double
					v.Code.WriteString("scvtf d1, w1\n") // Convertir int a double

					// lado izquierdo es un id f64
					v.Code.WriteString(fmt.Sprintf("\tldr x1, =%s\n", left.addrID))
					v.Code.WriteString("\tldr d2, [x1]\n")

					switch op {
					case "<":
						v.Code.WriteString("\tfcmp d2, d1\n")
						v.Code.WriteString("\tcset w3, lt\n")
					case "<=":
						v.Code.WriteString("\tfcmp d2, d1\n")
						v.Code.WriteString("\tcset w3, le\n")
					case ">":
						v.Code.WriteString("\tfcmp d2, d1\n")
						v.Code.WriteString("\tcset w3, gt\n")
					case ">=":
						v.Code.WriteString("\tfcmp d2, d1\n")
						v.Code.WriteString("\tcset w3, ge\n")
					case "==":
						v.Code.WriteString("\tfcmp d2, d1\n")
						v.Code.WriteString("\tcset w3, eq\n")
					case "!=":
						v.Code.WriteString("\tfcmp d2, d1\n")
						v.Code.WriteString("\tcset w3, ne\n")
					}

					tpm := v.nextTemp()
					// Temporal donde se guradara el resultado bool
					v.Data.WriteString(fmt.Sprintf("%s: .word 0\n", tpm))
					v.Code.WriteString(fmt.Sprintf("\tldr x2, =%s\n", tpm))
					v.Code.WriteString("\tstr w3, [x2]\n\n")

					return Attr{
						addrID: tpm,
						info:   tipoExp,
						value:  result,
					}

				} else if right.info == "f64" {
					// el lado derecho es un f64

					// declaro mi registro temporal para el inmediato y lo casteo a double
					// el inmediato es el lado derecho
					// etq := v.nextTemp()
					// v.Data.WriteString(fmt.Sprintf("%s: .double %s\n", etq, right.addrNUM))

					// se que trae un f64 inmediato del lado derecho, creo el temporal
					etq := v.nextTemp()
					v.Data.WriteString(fmt.Sprintf("%s: .double %s\n", etq, right.addrNUM))
					v.Code.WriteString(fmt.Sprintf("\tldr x0, =%s\n", etq))
					v.Code.WriteString("\tldr d1, [x0]\n")

					// lado izquierdo es un id int
					v.Code.WriteString(fmt.Sprintf("\tldr x1, =%s\n", left.addrID))
					v.Code.WriteString("\tldr w1, [x1]\n")
					v.Code.WriteString("scvtf d2, w1\n") // Convertir int a double

					switch op {
					case "<":
						v.Code.WriteString("\tfcmp d2, d1\n")
						v.Code.WriteString("\tcset w3, lt\n")
					case "<=":
						v.Code.WriteString("\tfcmp d2, d1\n")
						v.Code.WriteString("\tcset w3, le\n")
					case ">":
						v.Code.WriteString("\tfcmp d2, d1\n")
						v.Code.WriteString("\tcset w3, gt\n")
					case ">=":
						v.Code.WriteString("\tfcmp d2, d1\n")
						v.Code.WriteString("\tcset w3, ge\n")
					case "==":
						v.Code.WriteString("\tfcmp d2, d1\n")
						v.Code.WriteString("\tcset w3, eq\n")
					case "!=":
						v.Code.WriteString("\tfcmp d2, d1\n")
						v.Code.WriteString("\tcset w3, ne\n")
					}

					tpm := v.nextTemp()
					// Temporal donde se guradara el resultado bool
					v.Data.WriteString(fmt.Sprintf("%s: .word 0\n", tpm))
					v.Code.WriteString(fmt.Sprintf("\tldr x2, =%s\n", tpm))
					v.Code.WriteString("\tstr w3, [x2]\n\n")

					return Attr{
						addrID: tpm,
						info:   tipoExp,
						value:  result,
					}

				}
			}
		}

		// f64 + f64 = f64
		if left.info == "f64" && right.info == "f64" {
			var lval, rval float64
			lval = left.value.(float64)
			rval = right.value.(float64)

			var result interface{}
			tipoExp := "f64"
			switch op {
			case "+":
				result = lval + rval
			case "-":
				result = lval - rval
			case "*":
				result = lval * rval
			case "/":
				if rval == 0.0 {
					panic("División por cero " + fmt.Sprintf("%f / %f", lval, rval))
				}
				result = lval / rval
			case "<":
				result = lval < rval
				tipoExp = "bool"
			case "<=":
				result = lval <= rval
				tipoExp = "bool"
			case ">":
				result = lval > rval
				tipoExp = "bool"
			case ">=":
				result = lval >= rval
				tipoExp = "bool"
			case "==":
				result = lval == rval
				tipoExp = "bool"
			case "!=":
				result = lval != rval
				tipoExp = "bool"
			}

			if tipoExp == "f64" {
				// ya no necesito ver cual lado es el f64, ya que ambos son f64
				// se que el lado izquierdo es un id

				// el lado derecho es un f64, inmediato

				// declaro mi registro temporal para el inmediato
				etq := v.nextTemp()
				v.Data.WriteString(fmt.Sprintf("%s: .double %s\n", etq, right.addrNUM))
				v.Code.WriteString(fmt.Sprintf("\tldr x0, =%s\n", etq))
				v.Code.WriteString("\tldr d1, [x0]\n")

				// lado izquierdo es un id double
				v.Code.WriteString(fmt.Sprintf("\tldr x1, =%s\n", left.addrID))
				v.Code.WriteString("\tldr d2, [x1]\n") // Cargar valor de la variable

				switch op {
				case "+":
					v.Code.WriteString(("\tfadd d3, d2, d1\n"))
				case "-":
					v.Code.WriteString(("\tfsub d3, d2, d1\n"))
				case "*":
					v.Code.WriteString(("\tfmul d3, d2, d1\n"))
				case "/":
					v.Code.WriteString(("\tfdiv d3, d2, d1\n"))
				}

				tpm := v.nextTemp()
				// Temporal donde se guradara el resultado f64
				v.Data.WriteString(fmt.Sprintf("%s: .double 0.0\n", tpm))
				v.Code.WriteString(fmt.Sprintf("\tldr x2, =%s\n", tpm))
				v.Code.WriteString("\tstr d3, [x2]\n\n")

				return Attr{
					addrID: tpm,
					info:   tipoExp,
					value:  result,
				}
			} else if tipoExp == "bool" {
				// el lado derecho es un f64, inmediato

				// declaro mi registro temporal para el inmediato
				etq := v.nextTemp()
				v.Data.WriteString(fmt.Sprintf("%s: .double %s\n", etq, right.addrNUM))
				v.Code.WriteString(fmt.Sprintf("\tldr x0, =%s\n", etq))
				v.Code.WriteString("\tldr d1, [x0]\n")

				// lado izquierdo es un id double
				v.Code.WriteString(fmt.Sprintf("\tldr x1, =%s\n", left.addrID))
				v.Code.WriteString("\tldr d2, [x1]\n") // Cargar valor de la variable

				switch op {
				case "<":
					v.Code.WriteString("\tfcmp d2, d1\n")
					v.Code.WriteString("\tcset w3, lt\n")
				case "<=":
					v.Code.WriteString("\tfcmp d2, d1\n")
					v.Code.WriteString("\tcset w3, le\n")
				case ">":
					v.Code.WriteString("\tfcmp d2, d1\n")
					v.Code.WriteString("\tcset w3, gt\n")
				case ">=":
					v.Code.WriteString("\tfcmp d2, d1\n")
					v.Code.WriteString("\tcset w3, ge\n")
				case "==":
					v.Code.WriteString("\tfcmp d2, d1\n")
					v.Code.WriteString("\tcset w3, eq\n")
				case "!=":
					v.Code.WriteString("\tfcmp d2, d1\n")
					v.Code.WriteString("\tcset w3, ne\n")
				}

				tpm := v.nextTemp()
				// Temporal donde se guradara el resultado bool
				v.Data.WriteString(fmt.Sprintf("%s: .word 0\n", tpm))
				v.Code.WriteString(fmt.Sprintf("\tldr x2, =%s\n", tpm))
				v.Code.WriteString("\tstr w3, [x2]\n\n")

				return Attr{
					addrID: tpm,
					info:   tipoExp,
					value:  result,
				}
			}
		}

		// string + string = string
		if left.info == "string" && right.info == "string" {
			var lval, rval string
			lval = left.value.(string)
			rval = right.value.(string)

			var result interface{}
			tipoExp := "string"
			switch op {
			case "+":
				result = lval + rval
			case "<":
				result = lval < rval
				tipoExp = "bool"
			case "<=":
				result = lval <= rval
				tipoExp = "bool"
			case ">":
				result = lval > rval
				tipoExp = "bool"
			case ">=":
				result = lval >= rval
				tipoExp = "bool"
			case "==":
				result = lval == rval
				tipoExp = "bool"
			case "!=":
				result = lval != rval
				tipoExp = "bool"
			}

			if tipoExp == "string" {
				// uso copiar sin limpiar para concatenar cadenas
				// muevo el puntero para copiar donde termine la otra cadena

				// cadena final, sera de 64
				cfinal := v.nextTemp()
				v.BSS.WriteString(fmt.Sprintf("%s: .skip %d\n", cfinal, 64))           // espacio para la cadena de salida
				v.Code.WriteString(fmt.Sprintf("\tldr x0, =%s\t// destino\n", cfinal)) // cargar la direccion de memoria de salida
				v.Code.WriteString("\tmov x20, x0\t// guardar puntero base\n")

				// se que lado izquierdo es un id
				v.Code.WriteString(fmt.Sprintf("\tldr x1, =%s\t// origen\n", left.addrID)) // cargar la direccion de memoria del id
				v.Code.WriteString("\tmov x4, #64\t// Tamaño del buffer destino\n")
				v.Code.WriteString("\tbl copiar_sin_limpiar\n")
				v.Code.WriteString("\tmov x0, x20\t// restaurar puntero base\n")
				v.Code.WriteString("\tbl advance_to_end\n") // avanzar al final de la cadena de salida

				// lado derecho es un inmediato string
				etq := v.nextTemp()
				v.Data.WriteString(fmt.Sprintf("%s: .asciz \"%s\"\n", etq, rval))
				v.Code.WriteString(fmt.Sprintf("\tldr x1, =%s\t// origen\n", etq))
				v.Code.WriteString("\tbl copiar_sin_limpiar\n")

				return Attr{
					addrID: cfinal,
					info:   tipoExp,
					value:  result,
				}
			} else if tipoExp == "bool" {
				// resultado de comparacion de strings bool
				resBool := v.nextTemp()
				v.Data.WriteString(fmt.Sprintf("%s: .word 0\n", resBool))

				// ya se que el lado izquierdo es un id
				// obtengo el valor de memoria del id
				v.Code.WriteString(fmt.Sprintf("\tldr x0, =%s\n", left.addrID))

				// obtengo las dos direcciones de memoria de los strings
				// lado derecho es un inmediato string
				etq := v.nextTemp()
				v.Data.WriteString(fmt.Sprintf("%s: .asciz \"%s\"\n", etq, rval))
				v.Code.WriteString(fmt.Sprintf("\tldr x1, =%s\n", etq)) // cargar la direccion de memoria del inmediato

				switch op {
				case "<":
					// mando a llamar a la funcion strcmp_lt
					v.Code.WriteString("\tbl strcmp_lt\n")
					// el resultado se guarda en w0
					v.Code.WriteString(fmt.Sprintf("\tldr x2, =%s\n", resBool)) // cargar la direccion de memoria del resultado
					v.Code.WriteString("\tstr w0, [x2]\n")                      // guardar el resultado en la direccion de memoria

				case "<=":
					// mando a llamar a la funcion strcmp_le
					v.Code.WriteString("\tbl strcmp_le\n")
					// el resultado se guarda en w0
					v.Code.WriteString(fmt.Sprintf("\tldr x2, =%s\n", resBool)) // cargar la direccion de memoria del resultado
					v.Code.WriteString("\tstr w0, [x2]\n")                      // guardar el resultado en la direccion de memoria

				case ">":
					// mando a llamar a la funcion strcmp_gt
					v.Code.WriteString("\tbl strcmp_gt\n")
					// el resultado se guarda en w0
					v.Code.WriteString(fmt.Sprintf("\tldr x2, =%s\n", resBool)) // cargar la direccion de memoria del resultado
					v.Code.WriteString("\tstr w0, [x2]\n")                      // guardar el resultado en la direccion de memoria

				case ">=":
					// mando a llamar a la funcion strcmp_ge
					v.Code.WriteString("\tbl strcmp_ge\n")
					// el resultado se guarda en w0
					v.Code.WriteString(fmt.Sprintf("\tldr x2, =%s\n", resBool)) // cargar la direccion de memoria del resultado
					v.Code.WriteString("\tstr w0, [x2]\n")                      // guardar el resultado en la direccion de memoria

				case "==":
					// mando a llamar a la funcion strcmp_eq
					v.Code.WriteString("\tbl strcmp_eq\n")
					// el resultado se guarda en w0
					v.Code.WriteString(fmt.Sprintf("\tldr x2, =%s\n", resBool)) // cargar la direccion de memoria del resultado
					v.Code.WriteString("\tstr w0, [x2]\n")                      // guardar el resultado en la direccion de memoria

				case "!=":
					// mando a llamar a la funcion strcmp_ne
					v.Code.WriteString("\tbl strcmp_ne\n")
					// el resultado se guarda en w0
					v.Code.WriteString(fmt.Sprintf("\tldr x2, =%s\n", resBool)) // cargar la direccion de memoria del resultado
					v.Code.WriteString("\tstr w0, [x2]\n")                      // guardar el resultado en la direccion de memoria

				}

				// tpm := v.nextTemp()
				// // Temporal donde se guradara el resultado bool
				// // el resultado de todas las funciones se guarda en w0
				// v.Data.WriteString(fmt.Sprintf("%s: .word 0\n", tpm))
				// v.Code.WriteString(fmt.Sprintf("\tldr x2, =%s\n", tpm))
				// v.Code.WriteString("\tstr w3, [x2]\n\n")

				return Attr{
					addrID: resBool,
					info:   tipoExp,
					value:  result,
				}
			}
		}

		// bool + bool = bool
		if left.info == "bool" && right.info == "bool" {
			var lval, rval bool
			lval = left.value.(bool)
			rval = right.value.(bool)

			var result interface{}
			tipoExp := "bool"
			switch op {
			case "==":
				result = lval == rval
			case "!=":
				result = lval != rval
			case "&&":
				result = lval && rval
			case "||":
				result = lval || rval
			}

			if tipoExp == "bool" {
				// lado derecho es un inmediato bool
				etq := v.nextTemp()
				v.Data.WriteString(fmt.Sprintf("%s: .word %s\n", etq, right.addrNUM))
				v.Code.WriteString(fmt.Sprintf("\tldr x0, =%s\n", etq)) // cargar la direccion de memoria del inmediato
				v.Code.WriteString("\tldr w2, [x0]\n")                  // cargar el valor del inmediato

				// lado izquierdo es un id
				v.Code.WriteString(fmt.Sprintf("\tldr x0, =%s\n", left.addrID)) // cargar la direccion de memoria del id
				v.Code.WriteString("\tldr w1, [x0]\n")                          // cargar el valor del id
				switch op {
				case "==":
					v.Code.WriteString("\tcmp w1, w2\n")  // comparar los dos valores
					v.Code.WriteString("\tcset w3, eq\n") // establecer w2 a 1 si son iguales, 0 si no
				case "!=":
					v.Code.WriteString("\tcmp w1, w2\n")  // comparar los dos valores
					v.Code.WriteString("\tcset w3, ne\n") // establecer w2 a 1 si son diferentes, 0 si no
				case "&&":
					v.Code.WriteString("\tand w3, w1, w2\n") // establecer w2 a 1 si ambos son verdaderos, 0 si alguno es falso
				case "||":
					v.Code.WriteString("\torr w3, w1, w2\n") // establecer w2 a 1 si alguno es verdadero, 0 si ambos son falsos

				}

				tpm := v.nextTemp()
				// Temporal donde se guradara el resultado bool
				v.Data.WriteString(fmt.Sprintf("%s: .word 0\n", tpm))
				v.Code.WriteString(fmt.Sprintf("\tldr x2, =%s\n", tpm))
				v.Code.WriteString("\tstr w3, [x2]\n\n")

				// fmt.Println("Resultado de la operación booleano:", result)
				return Attr{
					addrID: tpm,
					info:   tipoExp,
					value:  result,
				}
			}
		}
	}

	// si ambos lados son inmediatos
	if isNUM(left) && isNUM(right) {

		// int + int = int
		if left.info == "int" && right.info == "int" {
			lval := left.value.(int64)
			rval := right.value.(int64)
			var result interface{}
			tipoExp := "int"
			switch op {
			case "+":
				result = lval + rval
			case "-":
				result = lval - rval
			case "*":
				result = lval * rval
			case "/":
				if rval == 0 {
					panic("División por cero " + fmt.Sprintf("%d / %d", lval, rval))
				}
				if lval%rval == 0 {
					result = lval / rval
				} else {
					result = float64(lval) / float64(rval)
					tipoExp = "f64"
				}

			case "%":
				if rval == 0 {
					panic("Módulo por cero " + fmt.Sprintf("%d %% %d", lval, rval))
				}
				result = lval % rval
			case "<":
				result = lval < rval
				tipoExp = "bool"
			case "<=":
				result = lval <= rval
				tipoExp = "bool"
			case ">":
				result = lval > rval
				tipoExp = "bool"
			case ">=":
				result = lval >= rval
				tipoExp = "bool"
			case "==":
				result = lval == rval
				tipoExp = "bool"
			case "!=":
				result = lval != rval
				tipoExp = "bool"
			}

			if tipoExp == "int" {

				// ya se que el lado derecho es inmediato
				// declaro el inmediato con mov w1, #valor inmediato
				v.Code.WriteString(fmt.Sprintf("\tmov w1, #%s\n", right.addrNUM))

				switch op {
				case "+":
					// se que el lado izquierdo es un valor inmediato
					// seraua x0, #valor inmediato
					v.Code.WriteString(fmt.Sprintf("\tadd w2, w1, #%s\n", left.addrNUM))
				case "-":
					// si es negativo el resultado me mostrara numeros raros

					// tengo que restar izquierdo - derecho

					// cargo mejor el valor en un registro, con mov
					v.Code.WriteString(fmt.Sprintf("\tmov w0, #%s\n", left.addrNUM))

					v.Code.WriteString("\tsub w2, w0, w1\n")
				case "*":
					// como mul solo permite multiplicacion entre registros, declaro el registro del inmediato
					v.Code.WriteString(fmt.Sprintf("\tmov w3, #%s\n", left.addrNUM))
					v.Code.WriteString("\tmul w2, w3, w1\n")
				case "/":
					v.Code.WriteString(fmt.Sprintf("\tmov w3, #%s\n", left.addrNUM))
					v.Code.WriteString("\tsdiv w2, w3, w1\n")
					// sdiv hacepta valores negativos y decimales

				}

				tpm := v.nextTemp()
				// Temporal donde se guradara el resultado
				v.Data.WriteString(fmt.Sprintf("%s: .word 0\n", tpm))
				v.Code.WriteString(fmt.Sprintf("\tldr x1, =%s\n", tpm))
				v.Code.WriteString("\tstr w2, [x1]\n\n")

				return Attr{
					addrID: tpm,
					info:   tipoExp,
					value:  result,
				}

			} else if tipoExp == "f64" {
				// en este punto ya se que el resultado de los dos ints da un double

				// el lado derecho es un inmediato int
				// declaro el inmediato con mov w1, #valor inmediato
				v.Code.WriteString(fmt.Sprintf("\tmov w1, #%s\n", right.addrNUM))
				// y lo convierto a double
				v.Code.WriteString("scvtf d1, w1\n") // Convertir int a double

				// ya se yo que solo se da en la divicion

				switch op {
				case "/":
					// asi que comvierto el lado izquierdo que es un valor inmediato intero a double
					// en arm no puedo declarar un float inmediato
					// asi que lo declaro en un registro temporal
					etq := v.nextTemp()
					v.Data.WriteString(fmt.Sprintf("%s: .double %s.0\n", etq, left.addrNUM))
					v.Code.WriteString(fmt.Sprintf("\tldr x0, =%s\n", etq))
					v.Code.WriteString("\tldr d2, [x0]\n")
					v.Code.WriteString("\tfdiv d2, d2, d1\n")
				}

				tpm := v.nextTemp()
				// Temporal donde se guradara el resultado f64
				v.Data.WriteString(fmt.Sprintf("%s: .double 0.0\n", tpm))
				v.Code.WriteString(fmt.Sprintf("\tldr x1, =%s\n", tpm))
				v.Code.WriteString("\tstr d2, [x1]\n\n")

				return Attr{
					addrID: tpm,
					info:   tipoExp,
					value:  result,
				}

			} else if tipoExp == "bool" {
				// ya se que el lado derecho es un inmediato
				// declaro el inmediato con mov w1, #valor inmediato
				v.Code.WriteString(fmt.Sprintf("\tmov w1, #%s\n", right.addrNUM))

				switch op {
				case "<":
					// necesito poner el lado izquierdo como registro
					v.Code.WriteString(fmt.Sprintf("\tmov w2, #%s\n", left.addrNUM))
					v.Code.WriteString("\tcmp w2, w1\n") // Comparar con el inmediato
					v.Code.WriteString("\tcset w3, lt\n")
				case "<=":
					v.Code.WriteString(fmt.Sprintf("\tmov w2, #%s\n", left.addrNUM))
					v.Code.WriteString("\tcmp w2, w1\n")
					v.Code.WriteString("\tcset w3, le\n")
				case ">":
					v.Code.WriteString(fmt.Sprintf("\tmov w2, #%s\n", left.addrNUM))
					v.Code.WriteString("\tcmp w2, w1\n")
					v.Code.WriteString("\tcset w3, gt\n")
				case ">=":
					v.Code.WriteString(fmt.Sprintf("\tmov w2, #%s\n", left.addrNUM))
					v.Code.WriteString("\tcmp w2, w1\n")
					v.Code.WriteString("\tcset w3, ge\n")
				case "==":
					v.Code.WriteString(fmt.Sprintf("\tmov w2, #%s\n", left.addrNUM))
					v.Code.WriteString("\tcmp w2, w1\n")
					v.Code.WriteString("\tcset w3, eq\n")
				case "!=":
					v.Code.WriteString(fmt.Sprintf("\tmov w2, #%s\n", left.addrNUM))
					v.Code.WriteString("\tcmp w2, w1\n")
					v.Code.WriteString("\tcset w3, ne\n")
				}

				tpm := v.nextTemp()
				// Temporal donde se guradara el resultado bool
				v.Data.WriteString(fmt.Sprintf("%s: .word 0\n", tpm))
				v.Code.WriteString(fmt.Sprintf("\tldr x2, =%s\n", tpm))
				v.Code.WriteString("\tstr w3, [x2]\n\n")

				return Attr{
					addrID: tpm,
					info:   tipoExp,
					value:  result,
				}
			}

		}

		// int + f64 = f64 | f64 + int = f64
		if (left.info == "int" && right.info == "f64") || (left.info == "f64" && right.info == "int") {
			// Caso: int y float, siempre se convierte a float
			var lval, rval float64
			if left.info == "int" {
				lval = float64(left.value.(int64))
			} else {
				lval = left.value.(float64)
			}

			if right.info == "int" {
				rval = float64(right.value.(int64))
			} else {
				rval = right.value.(float64)
			}

			var result interface{}
			tipoExp := "f64"

			switch op {
			case "+":
				result = lval + rval
			case "-":
				result = lval - rval
			case "*":
				result = lval * rval
			case "/":
				if rval == 0.0 {
					panic("División por cero " + fmt.Sprintf("%f / %f", lval, rval))
				}
				result = lval / rval
			case "<":
				result = lval < rval
				tipoExp = "bool"
			case "<=":
				result = lval <= rval
				tipoExp = "bool"
			case ">":
				result = lval > rval
				tipoExp = "bool"
			case ">=":
				result = lval >= rval
				tipoExp = "bool"
			case "==":
				result = lval == rval
				tipoExp = "bool"
			case "!=":
				result = lval != rval
				tipoExp = "bool"
			}

			if tipoExp == "f64" {

				// veo cual de los dos lados es el f64
				if left.info == "f64" {
					// el lado izquierdo es un f64

					// declaro mi registro temporal para el inmediato del lado izquierdo
					etq := v.nextTemp()
					v.Data.WriteString(fmt.Sprintf("%s: .double %s\n", etq, left.addrNUM))
					v.Code.WriteString(fmt.Sprintf("\tldr x0, =%s\n", etq))
					v.Code.WriteString("\tldr d1, [x0]\n")

					// lado derecho es un inmediato int
					// declaro el inmediato con mov w1, #valor inmediato
					v.Code.WriteString(fmt.Sprintf("\tmov w1, #%s\n", right.addrNUM))
					// lo convierto a double
					v.Code.WriteString("scvtf d2, w1\n") // Convertir int a double

					switch op {
					case "+":
						v.Code.WriteString(("\tfadd d3, d1, d2\n"))
					case "-":
						v.Code.WriteString(("\tfsub d3, d1, d2\n"))
					case "*":
						v.Code.WriteString(("\tfmul d3, d1, d2\n"))
					case "/":
						v.Code.WriteString(("\tfdiv d3, d1, d2\n"))
					}

					tpm := v.nextTemp()
					// Temporal donde se guradara el resultado f64
					v.Data.WriteString(fmt.Sprintf("%s: .double 0.0\n", tpm))
					v.Code.WriteString(fmt.Sprintf("\tldr x2, =%s\n", tpm))
					v.Code.WriteString("\tstr d3, [x2]\n\n")

					return Attr{
						addrID: tpm,
						info:   tipoExp,
						value:  result,
					}

				} else if right.info == "f64" {
					// el lado derecho es un f64

					// declaro mi registro temporal para el inmediato y lo casteo a double
					// etq := v.nextTemp() + "ss"
					// v.Data.WriteString(fmt.Sprintf("%s: .double %s\n", etq, left.addrNUM))

					// se que trae un int del lado izquierdo, lo asigno a un registro w con mov
					v.Code.WriteString("\tmov w1, #" + left.addrNUM + "\n")
					// lo casteo a double
					v.Code.WriteString("scvtf d1, w1\n") // Convertir int a double

					// lado derecho es un inmediato double
					// creo el temporal
					etq := v.nextTemp()
					v.Data.WriteString(fmt.Sprintf("%s: .double %s\n", etq, right.addrNUM))
					v.Code.WriteString(fmt.Sprintf("\tldr x1, =%s\n", etq))
					v.Code.WriteString("\tldr d2, [x1]\n")

					switch op {
					case "+":
						v.Code.WriteString(("\tfadd d3, d1, d2\n"))
					case "-":
						v.Code.WriteString(("\tfsub d3, d1, d2\n"))
					case "*":
						v.Code.WriteString(("\tfmul d3, d1, d2\n"))
					case "/":
						v.Code.WriteString(("\tfdiv d3, d1, d2\n"))
					}

					tpm := v.nextTemp()
					// Temporal donde se guradara el resultado f64
					v.Data.WriteString(fmt.Sprintf("%s: .double 0.0\n", tpm))
					v.Code.WriteString(fmt.Sprintf("\tldr x2, =%s\n", tpm))
					v.Code.WriteString("\tstr d3, [x2]\n\n")

					return Attr{
						addrID: tpm,
						info:   tipoExp,
						value:  result,
					}

				}

			} else if tipoExp == "bool" {
				// veo cual de los dos lados es el f64
				if left.info == "f64" {
					// el lado izquierdo es un f64

					// declaro mi registro temporal para el inmediato del lado izquierdo
					etq := v.nextTemp()
					v.Data.WriteString(fmt.Sprintf("%s: .double %s\n", etq, left.addrNUM))
					v.Code.WriteString(fmt.Sprintf("\tldr x0, =%s\n", etq))
					v.Code.WriteString("\tldr d1, [x0]\n")

					// lado derecho es un inmediato int
					// declaro el inmediato con mov w1, #valor inmediato
					v.Code.WriteString(fmt.Sprintf("\tmov w1, #%s\n", right.addrNUM))
					// lo convierto a double
					v.Code.WriteString("scvtf d2, w1\n") // Convertir int a double

					switch op {
					case "<":
						v.Code.WriteString("\tfcmp d1, d2\n")
						v.Code.WriteString("\tcset w3, lt\n")
					case "<=":
						v.Code.WriteString("\tfcmp d1, d2\n")
						v.Code.WriteString("\tcset w3, le\n")
					case ">":
						v.Code.WriteString("\tfcmp d1, d2\n")
						v.Code.WriteString("\tcset w3, gt\n")
					case ">=":
						v.Code.WriteString("\tfcmp d1, d2\n")
						v.Code.WriteString("\tcset w3, ge\n")
					case "==":
						v.Code.WriteString("\tfcmp d1, d2\n")
						v.Code.WriteString("\tcset w3, eq\n")
					case "!=":
						v.Code.WriteString("\tfcmp d1, d2\n")
						v.Code.WriteString("\tcset w3, ne\n")
					}

					tpm := v.nextTemp()
					// Temporal donde se guradara el resultado bool
					v.Data.WriteString(fmt.Sprintf("%s: .word 0\n", tpm))
					v.Code.WriteString(fmt.Sprintf("\tldr x2, =%s\n", tpm))
					v.Code.WriteString("\tstr w3, [x2]\n\n")

					return Attr{
						addrID: tpm,
						info:   tipoExp,
						value:  result,
					}

				} else if right.info == "f64" {
					// el lado derecho es un f64

					// se que trae un int del lado izquierdo, lo asigno a un registro w con mov
					v.Code.WriteString("\tmov w1, #" + left.addrNUM + "\n")
					// lo casteo a double
					v.Code.WriteString("scvtf d1, w1\n") // Convertir int a double

					// lado derecho es un inmediato double
					// creo el temporal
					etq := v.nextTemp()
					v.Data.WriteString(fmt.Sprintf("%s: .double %s\n", etq, right.addrNUM))
					v.Code.WriteString(fmt.Sprintf("\tldr x1, =%s\n", etq))
					v.Code.WriteString("\tldr d2, [x1]\n")

					switch op {
					case "<":
						v.Code.WriteString("\tfcmp d1, d2\n")
						v.Code.WriteString("\tcset w3, lt\n")
					case "<=":
						v.Code.WriteString("\tfcmp d1, d2\n")
						v.Code.WriteString("\tcset w3, le\n")
					case ">":
						v.Code.WriteString("\tfcmp d1, d2\n")
						v.Code.WriteString("\tcset w3, gt\n")
					case ">=":
						v.Code.WriteString("\tfcmp d1, d2\n")
						v.Code.WriteString("\tcset w3, ge\n")
					case "==":
						v.Code.WriteString("\tfcmp d1, d2\n")
						v.Code.WriteString("\tcset w3, eq\n")
					case "!=":
						v.Code.WriteString("\tfcmp d1, d2\n")
						v.Code.WriteString("\tcset w3, ne\n")
					}

					tpm := v.nextTemp()
					// Temporal donde se guradara el resultado bool
					v.Data.WriteString(fmt.Sprintf("%s: .word 0\n", tpm))
					v.Code.WriteString(fmt.Sprintf("\tldr x2, =%s\n", tpm))
					v.Code.WriteString("\tstr w3, [x2]\n\n")

					return Attr{
						addrID: tpm,
						info:   tipoExp,
						value:  result,
					}

				}
			}
		}

		// f64 + f64 = f64
		if left.info == "f64" && right.info == "f64" {
			var lval, rval float64
			lval = left.value.(float64)
			rval = right.value.(float64)

			var result interface{}
			tipoExp := "f64"

			switch op {
			case "+":
				result = lval + rval
			case "-":
				result = lval - rval
			case "*":
				result = lval * rval
			case "/":
				if rval == 0.0 {
					panic("División por cero " + fmt.Sprintf("%f / %f", lval, rval))
				}
				result = lval / rval
			case "<":
				result = lval < rval
				tipoExp = "bool"
			case "<=":
				result = lval <= rval
				tipoExp = "bool"
			case ">":
				result = lval > rval
				tipoExp = "bool"
			case ">=":
				result = lval >= rval
				tipoExp = "bool"
			case "==":
				result = lval == rval
				tipoExp = "bool"
			case "!=":
				result = lval != rval
				tipoExp = "bool"
			}

			if tipoExp == "f64" {

				// ya no necesito ver cual lado es el f64, ya que ambos son f64
				// se que el lado derecho es un inmediato

				// el lado izquierdo es un f64

				// declaro mi registro temporal para el inmediato
				etq := v.nextTemp()
				v.Data.WriteString(fmt.Sprintf("%s: .double %s\n", etq, left.addrNUM))
				v.Code.WriteString(fmt.Sprintf("\tldr x0, =%s\n", etq))
				v.Code.WriteString("\tldr d1, [x0]\n")

				// lado derecho es un inmediato double tambien
				// creo el temporal
				etq2 := v.nextTemp()
				v.Data.WriteString(fmt.Sprintf("%s: .double %s\n", etq2, right.addrNUM))
				v.Code.WriteString(fmt.Sprintf("\tldr x1, =%s\n", etq2))
				v.Code.WriteString("\tldr d2, [x1]\n")

				switch op {
				case "+":
					v.Code.WriteString(("\tfadd d3, d1, d2\n"))
				case "-":
					v.Code.WriteString(("\tfsub d3, d1, d2\n"))
				case "*":
					v.Code.WriteString(("\tfmul d3, d1, d2\n"))
				case "/":
					v.Code.WriteString(("\tfdiv d3, d1, d2\n"))
				}

				tpm := v.nextTemp()
				// Temporal donde se guradara el resultado f64
				v.Data.WriteString(fmt.Sprintf("%s: .double 0.0\n", tpm))
				v.Code.WriteString(fmt.Sprintf("\tldr x2, =%s\n", tpm))
				v.Code.WriteString("\tstr d3, [x2]\n\n")

				return Attr{
					addrID: tpm,
					info:   tipoExp,
					value:  result,
				}

			} else if tipoExp == "bool" {
				// ya no necesito ver cual lado es el f64, ya que ambos son f64
				// se que el lado derecho es un id

				// el lado izquierdo es un f64

				// declaro mi registro temporal para el inmediato
				etq := v.nextTemp()
				v.Data.WriteString(fmt.Sprintf("%s: .double %s\n", etq, left.addrNUM))
				v.Code.WriteString(fmt.Sprintf("\tldr x0, =%s\n", etq))
				v.Code.WriteString("\tldr d1, [x0]\n")

				// lado derecho es un inmediato double tambien
				// creo el temporal
				etq2 := v.nextTemp()
				v.Data.WriteString(fmt.Sprintf("%s: .double %s\n", etq2, right.addrNUM))
				v.Code.WriteString(fmt.Sprintf("\tldr x1, =%s\n", etq2))
				v.Code.WriteString("\tldr d2, [x1]\n")
				switch op {
				case "<":
					v.Code.WriteString("\tfcmp d1, d2\n")
					v.Code.WriteString("\tcset w3, lt\n")
				case "<=":
					v.Code.WriteString("\tfcmp d1, d2\n")
					v.Code.WriteString("\tcset w3, le\n")
				case ">":
					v.Code.WriteString("\tfcmp d1, d2\n")
					v.Code.WriteString("\tcset w3, gt\n")
				case ">=":
					v.Code.WriteString("\tfcmp d1, d2\n")
					v.Code.WriteString("\tcset w3, ge\n")
				case "==":
					v.Code.WriteString("\tfcmp d1, d2\n")
					v.Code.WriteString("\tcset w3, eq\n")
				case "!=":
					v.Code.WriteString("\tfcmp d1, d2\n")
					v.Code.WriteString("\tcset w3, ne\n")
				}

				tpm := v.nextTemp()
				// Temporal donde se guradara el resultado bool
				v.Data.WriteString(fmt.Sprintf("%s: .word 0\n", tpm))
				v.Code.WriteString(fmt.Sprintf("\tldr x2, =%s\n", tpm))
				v.Code.WriteString("\tstr w3, [x2]\n\n")

				return Attr{
					addrID: tpm,
					info:   tipoExp,
					value:  result,
				}

			}
		}

		// string + string = string
		if left.info == "string" && right.info == "string" {
			var lval, rval string
			lval = left.value.(string)
			rval = right.value.(string)

			var result interface{}
			tipoExp := "string"

			switch op {
			case "+":
				result = lval + rval
			case "<":
				result = lval < rval
				tipoExp = "bool"
			case "<=":
				result = lval <= rval
				tipoExp = "bool"
			case ">":
				result = lval > rval
				tipoExp = "bool"
			case ">=":
				result = lval >= rval
				tipoExp = "bool"
			case "==":
				result = lval == rval
				tipoExp = "bool"
			case "!=":
				result = lval != rval
				tipoExp = "bool"
			}

			if tipoExp == "string" {
				// uso copiar sin limpiar para concatenar cadenas
				// muevo el puntero para copiar donde termine la otra cadena

				// cadena final, sera de 64
				cfinal := v.nextTemp()
				v.BSS.WriteString(fmt.Sprintf("%s: .skip %d\n", cfinal, 64))           // espacio para la cadena de salida
				v.Code.WriteString(fmt.Sprintf("\tldr x0, =%s\t// destino\n", cfinal)) // cargar la direccion de memoria de salida
				v.Code.WriteString("\tmov x20, x0\t// guardar puntero base\n")
				// lado izquierdo es un inmediato string
				etq := v.nextTemp()
				v.Data.WriteString(fmt.Sprintf("%s: .asciz \"%s\"\n", etq, lval))
				v.Code.WriteString(fmt.Sprintf("\tldr x1, =%s\t// origen\n", etq))
				v.Code.WriteString("\tmov x4, #64\t// Tamaño del buffer destino\n")
				v.Code.WriteString("\tbl copiar_sin_limpiar\n")
				v.Code.WriteString("\tmov x0, x20\t// restaurar puntero base\n")
				v.Code.WriteString("\tbl advance_to_end\n") // avanzar al final de la cadena de salida

				// se que lado derecho es un inmediato string tambien
				etq2 := v.nextTemp()
				v.Data.WriteString(fmt.Sprintf("%s: .asciz \"%s\"\n", etq2, rval))
				v.Code.WriteString(fmt.Sprintf("\tldr x1, =%s\t// origen\n", etq2)) // cargar la direccion de memoria del id
				v.Code.WriteString("\tbl copiar_sin_limpiar\n")

				return Attr{
					addrID: cfinal,
					info:   tipoExp,
					value:  result,
				}
			} else if tipoExp == "bool" {

				// resultado de comparacion de strings bool
				resBool := v.nextTemp()
				v.Data.WriteString(fmt.Sprintf("%s: .word 0\n", resBool))

				// obtengo las dos direcciones de memoria de los strings
				// lado izquierdo es un inmediato string
				etq := v.nextTemp()
				v.Data.WriteString(fmt.Sprintf("%s: .asciz \"%s\"\n", etq, lval))
				v.Code.WriteString(fmt.Sprintf("\tldr x0, =%s\n", etq)) // cargar la direccion de memoria del inmediato

				// ya se que el lado derecho tambien es un inmediato string
				// obtengo el valor de memoria del inmediato
				etq2 := v.nextTemp()
				v.Data.WriteString(fmt.Sprintf("%s: .asciz \"%s\"\n", etq2, rval))
				v.Code.WriteString(fmt.Sprintf("\tldr x1, =%s\n", etq2))

				switch op {
				case "<":
					// mando a llamar a la funcion strcmp_lt
					v.Code.WriteString("\tbl strcmp_lt\n")
					// el resultado se guarda en w0
					v.Code.WriteString(fmt.Sprintf("\tldr x2, =%s\n", resBool)) // cargar la direccion de memoria del resultado
					v.Code.WriteString("\tstr w0, [x2]\n")                      // guardar el resultado en la direccion de memoria

				case "<=":
					// mando a llamar a la funcion strcmp_le
					v.Code.WriteString("\tbl strcmp_le\n")
					// el resultado se guarda en w0
					v.Code.WriteString(fmt.Sprintf("\tldr x2, =%s\n", resBool)) // cargar la direccion de memoria del resultado
					v.Code.WriteString("\tstr w0, [x2]\n")                      // guardar el resultado en la direccion de memoria

				case ">":
					// mando a llamar a la funcion strcmp_gt
					v.Code.WriteString("\tbl strcmp_gt\n")
					// el resultado se guarda en w0
					v.Code.WriteString(fmt.Sprintf("\tldr x2, =%s\n", resBool)) // cargar la direccion de memoria del resultado
					v.Code.WriteString("\tstr w0, [x2]\n")                      // guardar el resultado en la direccion de memoria

				case ">=":
					// mando a llamar a la funcion strcmp_ge
					v.Code.WriteString("\tbl strcmp_ge\n")
					// el resultado se guarda en w0
					v.Code.WriteString(fmt.Sprintf("\tldr x2, =%s\n", resBool)) // cargar la direccion de memoria del resultado
					v.Code.WriteString("\tstr w0, [x2]\n")                      // guardar el resultado en la direccion de memoria

				case "==":
					// mando a llamar a la funcion strcmp_eq
					v.Code.WriteString("\tbl strcmp_eq\n")
					// el resultado se guarda en w0
					v.Code.WriteString(fmt.Sprintf("\tldr x2, =%s\n", resBool)) // cargar la direccion de memoria del resultado
					v.Code.WriteString("\tstr w0, [x2]\n")                      // guardar el resultado en la direccion de memoria

				case "!=":
					// mando a llamar a la funcion strcmp_ne
					v.Code.WriteString("\tbl strcmp_ne\n")
					// el resultado se guarda en w0
					v.Code.WriteString(fmt.Sprintf("\tldr x2, =%s\n", resBool)) // cargar la direccion de memoria del resultado
					v.Code.WriteString("\tstr w0, [x2]\n")                      // guardar el resultado en la direccion de memoria

				}

				// tpm := v.nextTemp()
				// // Temporal donde se guradara el resultado bool
				// v.Data.WriteString(fmt.Sprintf("%s: .word 0\n", tpm))
				// v.Code.WriteString(fmt.Sprintf("\tldr x2, =%s\n", tpm))
				// v.Code.WriteString("\tstr w3, [x2]\n\n")

				return Attr{
					addrID: resBool,
					info:   tipoExp,
					value:  result,
				}
			}

		}

		// bool + bool = bool
		if left.info == "bool" && right.info == "bool" {
			var lval, rval bool
			lval = left.value.(bool)
			rval = right.value.(bool)

			var result interface{}
			tipoExp := "bool"
			switch op {
			case "==":
				result = lval == rval
			case "!=":
				result = lval != rval
			case "&&":
				result = lval && rval
			case "||":
				result = lval || rval
			}

			if tipoExp == "bool" {
				// lado izquierdo es un inmediato bool
				etq := v.nextTemp()
				v.Data.WriteString(fmt.Sprintf("%s: .word %s\n", etq, left.addrNUM))
				v.Code.WriteString(fmt.Sprintf("\tldr x0, =%s\n", etq)) // cargar la direccion de memoria del inmediato
				v.Code.WriteString("\tldr w1, [x0]\n")                  // cargar el valor del inmediato

				// lado derecho es un tambien es un inmediato bool
				etq2 := v.nextTemp()
				v.Data.WriteString(fmt.Sprintf("%s: .word %s\n", etq2, right.addrNUM))
				v.Code.WriteString(fmt.Sprintf("\tldr x0, =%s\n", etq2)) // cargar la direccion de memoria del inmediato
				v.Code.WriteString("\tldr w2, [x0]\n")                   // cargar el valor del inmediato
				switch op {
				case "==":
					v.Code.WriteString("\tcmp w1, w2\n")  // comparar los dos valores
					v.Code.WriteString("\tcset w3, eq\n") // establecer w2 a 1 si son iguales, 0 si no
				case "!=":
					v.Code.WriteString("\tcmp w1, w2\n")  // comparar los dos valores
					v.Code.WriteString("\tcset w3, ne\n") // establecer w2 a 1 si son diferentes, 0 si no
				case "&&":
					v.Code.WriteString("\tand w3, w1, w2\n") // establecer w2 a 1 si ambos son verdaderos, 0 si alguno es falso
				case "||":
					v.Code.WriteString("\torr w3, w1, w2\n") // establecer w2 a 1 si alguno es verdadero, 0 si ambos son falsos

				}

				tpm := v.nextTemp()
				// Temporal donde se guradara el resultado bool
				v.Data.WriteString(fmt.Sprintf("%s: .word 0\n", tpm))
				v.Code.WriteString(fmt.Sprintf("\tldr x2, =%s\n", tpm))
				v.Code.WriteString("\tstr w3, [x2]\n\n")

				fmt.Println("Resultado de la operación booleano:", result)
				return Attr{
					addrID: tpm,
					info:   tipoExp,
					value:  result,
				}
			}
		}
	}

	// si ambos son id

	if !isNUM(left) && !isNUM(right) {

		// int + int = int | int - int = int | int * int = int | int / int = f64

		if left.info == "int" && right.info == "int" {
			var lval, rval int64
			lval = left.value.(int64)
			rval = right.value.(int64)

			var result interface{}
			tipoExp := "int"

			switch op {
			case "+":
				result = lval + rval
			case "-":
				result = lval - rval
			case "*":
				result = lval * rval
			case "/":
				if rval == 0 {
					panic("División por cero " + fmt.Sprintf("%d / %d", lval, rval))
				}
				if lval%rval == 0 {
					result = lval / rval
				} else {
					result = float64(lval) / float64(rval)
					tipoExp = "f64"
				}

			case "%":
				if rval == 0 {
					panic("Módulo por cero " + fmt.Sprintf("%d %% %d", lval, rval))
				}
				result = lval % rval
			case "<":
				result = lval < rval
				tipoExp = "bool"
			case "<=":
				result = lval <= rval
				tipoExp = "bool"
			case ">":
				result = lval > rval
				tipoExp = "bool"
			case ">=":
				result = lval >= rval
				tipoExp = "bool"
			case "==":
				result = lval == rval
				tipoExp = "bool"
			case "!=":
				result = lval != rval
				tipoExp = "bool"
			}

			if tipoExp == "int" {
				// los dos lados son id int

				// obtengo la direccion del lado derecho int
				v.Code.WriteString(fmt.Sprintf("\tldr x0, =%s\n", right.addrID))
				v.Code.WriteString("\tldr w1, [x0]\n") // Cargar valor de la variable

				switch op {
				case "+":

					// obtengo la direccion del lado izquierdo int
					v.Code.WriteString(fmt.Sprintf("\tldr x0, =%s\n", left.addrID))
					v.Code.WriteString("\tldr w2, [x0]\n")

					v.Code.WriteString("\tadd w2, w2, w1\n") // Sumar los dos registros
				case "-":
					// si es negativo el resultado me mostrara numeros raros

					// tengo que restar izquierdo - derecho

					// cargo el registro del lado izquierdo
					v.Code.WriteString(fmt.Sprintf("\tldr x0, =%s\n", left.addrID))
					v.Code.WriteString("\tldr w0, [x0]\n")

					v.Code.WriteString("\tsub w2, w0, w1\n")
				case "*":
					// como mul solo permite multiplicacion entre registros, declaro el registro del inmediato
					// cargo el registro izquierdo en w3
					v.Code.WriteString(fmt.Sprintf("\tldr x0, =%s\n", left.addrID))
					v.Code.WriteString("\tldr w3, [x0]\n")

					v.Code.WriteString("\tmul w2, w3, w1\n")
				case "/":
					v.Code.WriteString(fmt.Sprintf("\tldr x0, =%s\n", left.addrID))
					v.Code.WriteString("\tldr w3, [x0]\n")

					v.Code.WriteString("\tsdiv w2, w3, w1\n")
				}
				tpm := v.nextTemp()
				// Temporal donde se guradara el resultado int
				v.Data.WriteString(fmt.Sprintf("%s: .word 0\n", tpm))
				v.Code.WriteString(fmt.Sprintf("\tldr x0, =%s\n", tpm))
				v.Code.WriteString("\tstr w2, [x0]\n\n")
				return Attr{
					addrID: tpm,
					info:   tipoExp,
					value:  result,
				}

			} else if tipoExp == "f64" {
				// en este punto ya se que el resultado de los dos ints da un double

				// el lado derecho es un id
				// obtengo el valor de memoria del id y lo comvierto a double
				v.Code.WriteString(fmt.Sprintf("\tldr x0, =%s\n", right.addrID))
				v.Code.WriteString("\tldr w1, [x0]\n") // Cargar valor de la variable
				// uso scvtf para convertir el valor inmediato a double
				v.Code.WriteString("scvtf d1, w1\n") // Convertir int a double

				// ya se yo que solo se da en la divicion

				switch op {
				case "/":
					// asi que comvierto el lado izquierdo que es un valor id intero a double
					// en arm no puedo declarar un float inmediato
					// asi que lo declaro en un registro temporal
					v.Code.WriteString(fmt.Sprintf("\tldr x0, =%s\n", left.addrID))
					v.Code.WriteString("\tldr w2, [x0]\n") // Cargar valor de la variable
					// lo convierto a double
					v.Code.WriteString("scvtf d2, w2\n") // Convertir

					v.Code.WriteString("\tfdiv d2, d2, d1\n")
				}

				tpm := v.nextTemp()
				// Temporal donde se guradara el resultado f64
				v.Data.WriteString(fmt.Sprintf("%s: .double 0.0\n", tpm))
				v.Code.WriteString(fmt.Sprintf("\tldr x0, =%s\n", tpm))
				v.Code.WriteString("\tstr d2, [x0]\n\n")

				return Attr{
					addrID: tpm,
					info:   tipoExp,
					value:  result,
				}

			} else if tipoExp == "bool" {
				// ya se que el lado derecho es un id
				// obtengo el valor de memoria del id
				v.Code.WriteString(fmt.Sprintf("\tldr x1, =%s\n", right.addrID))
				v.Code.WriteString("\tldr w1, [x1]\n") // Cargar valor de la variable

				switch op {
				case "<":
					// necesito cargar la direccion de memoria del lado izquierdo
					v.Code.WriteString(fmt.Sprintf("\tldr x0, =%s\n", left.addrID))
					v.Code.WriteString("\tldr w2, [x0]\n") // Cargar valor de la variable

					v.Code.WriteString("\tcmp w2, w1\n") // Comparar con el inmediato
					v.Code.WriteString("\tcset w3, lt\n")
				case "<=":
					// necesito cargar la direccion de memoria del lado izquierdo
					v.Code.WriteString(fmt.Sprintf("\tldr x0, =%s\n", left.addrID))
					v.Code.WriteString("\tldr w2, [x0]\n") // Cargar valor de la variable

					v.Code.WriteString("\tcmp w2, w1\n")
					v.Code.WriteString("\tcset w3, le\n")
				case ">":
					// necesito cargar la direccion de memoria del lado izquierdo
					v.Code.WriteString(fmt.Sprintf("\tldr x0, =%s\n", left.addrID))
					v.Code.WriteString("\tldr w2, [x0]\n") // Cargar valor de la variable

					v.Code.WriteString("\tcmp w2, w1\n")
					v.Code.WriteString("\tcset w3, gt\n")
				case ">=":
					// necesito cargar la direccion de memoria del lado izquierdo
					v.Code.WriteString(fmt.Sprintf("\tldr x0, =%s\n", left.addrID))
					v.Code.WriteString("\tldr w2, [x0]\n") // Cargar valor de la variable

					v.Code.WriteString("\tcmp w2, w1\n")
					v.Code.WriteString("\tcset w3, ge\n")
				case "==":
					// necesito cargar la direccion de memoria del lado izquierdo
					v.Code.WriteString(fmt.Sprintf("\tldr x0, =%s\n", left.addrID))
					v.Code.WriteString("\tldr w2, [x0]\n") // Cargar valor de la variable

					v.Code.WriteString("\tcmp w2, w1\n")
					v.Code.WriteString("\tcset w3, eq\n")
				case "!=":
					// necesito cargar la direccion de memoria del lado izquierdo
					v.Code.WriteString(fmt.Sprintf("\tldr x0, =%s\n", left.addrID))
					v.Code.WriteString("\tldr w2, [x0]\n") // Cargar valor de la variable

					v.Code.WriteString("\tcmp w2, w1\n")
					v.Code.WriteString("\tcset w3, ne\n")
				}

				tpm := v.nextTemp()
				// Temporal donde se guradara el resultado bool
				v.Data.WriteString(fmt.Sprintf("%s: .word 0\n", tpm))
				v.Code.WriteString(fmt.Sprintf("\tldr x2, =%s\n", tpm))
				v.Code.WriteString("\tstr w3, [x2]\n\n")

				return Attr{
					addrID: tpm,
					info:   tipoExp,
					value:  result,
				}
			}
		}

		// int + f64 = f64 | f64 + int = f64
		if (left.info == "int" && right.info == "f64") || (left.info == "f64" && right.info == "int") {
			var lval, rval float64
			if left.info == "int" {
				lval = float64(left.value.(int64))
			} else {
				lval = left.value.(float64)
			}

			if right.info == "int" {
				rval = float64(right.value.(int64))
			} else {
				rval = right.value.(float64)
			}

			var result interface{}
			tipoExp := "f64"
			switch op {
			case "+":
				result = lval + rval
			case "-":
				result = lval - rval
			case "*":
				result = lval * rval
			case "/":
				if rval == 0.0 {
					panic("División por cero " + fmt.Sprintf("%f / %f", lval, rval))
				}
				result = lval / rval
			case "<":
				result = lval < rval
				tipoExp = "bool"
			case "<=":
				result = lval <= rval
				tipoExp = "bool"
			case ">":
				result = lval > rval
				tipoExp = "bool"
			case ">=":
				result = lval >= rval
				tipoExp = "bool"
			case "==":
				result = lval == rval
				tipoExp = "bool"
			case "!=":
				result = lval != rval
				tipoExp = "bool"
			}

			if tipoExp == "f64" {
				// veo cual de los dos lados es el f64
				if left.info == "f64" {
					// el lado izquierdo es un f64

					// mi lado izquierdo es un id
					v.Code.WriteString(fmt.Sprintf("\tldr x0, =%s\n", left.addrID))
					v.Code.WriteString("\tldr d1, [x0]\n") // Cargar valor de la variable

					// lado derecho es un id int
					v.Code.WriteString(fmt.Sprintf("\tldr x1, =%s\n", right.addrID))
					v.Code.WriteString("\tldr w1, [x1]\n")
					v.Code.WriteString("scvtf d2, w1\n") // Convertir int a double

					switch op {
					case "+":
						v.Code.WriteString(("\tfadd d3, d1, d2\n"))
					case "-":
						v.Code.WriteString(("\tfsub d3, d1, d2\n"))
					case "*":
						v.Code.WriteString(("\tfmul d3, d1, d2\n"))
					case "/":
						v.Code.WriteString(("\tfdiv d3, d1, d2\n"))
					}

					tpm := v.nextTemp()
					// Temporal donde se guradara el resultado f64
					v.Data.WriteString(fmt.Sprintf("%s: .double 0.0\n", tpm))
					v.Code.WriteString(fmt.Sprintf("\tldr x2, =%s\n", tpm))
					v.Code.WriteString("\tstr d3, [x2]\n\n")

					return Attr{
						addrID: tpm,
						info:   tipoExp,
						value:  result,
					}

				} else if right.info == "f64" {
					// el lado derecho es un f64

					// declaro mi registro temporal para el inmediato y lo casteo a double
					// etq := v.nextTemp() + "ss"
					// v.Data.WriteString(fmt.Sprintf("%s: .double %s\n", etq, left.addrNUM))

					// se que trae un int del lado izquierdo y es un id
					v.Code.WriteString(fmt.Sprintf("\tldr x0, =%s\n", left.addrID))
					v.Code.WriteString("\tldr w1, [x0]\n") // Cargar valor de la variable
					// lo convierto a double
					// lo casteo a double
					v.Code.WriteString("scvtf d1, w1\n") // Convertir

					// lado derecho es un id double
					v.Code.WriteString(fmt.Sprintf("\tldr x1, =%s\n", right.addrID))
					v.Code.WriteString("\tldr d2, [x1]\n") // Cargar valor de la variable

					switch op {
					case "+":
						v.Code.WriteString(("\tfadd d3, d1, d2\n"))
					case "-":
						v.Code.WriteString(("\tfsub d3, d1, d2\n"))
					case "*":
						v.Code.WriteString(("\tfmul d3, d1, d2\n"))
					case "/":
						v.Code.WriteString(("\tfdiv d3, d1, d2\n"))
					}

					tpm := v.nextTemp()
					// Temporal donde se guradara el resultado f64
					v.Data.WriteString(fmt.Sprintf("%s: .double 0.0\n", tpm))
					v.Code.WriteString(fmt.Sprintf("\tldr x2, =%s\n", tpm))
					v.Code.WriteString("\tstr d3, [x2]\n\n")

					return Attr{
						addrID: tpm,
						info:   tipoExp,
						value:  result,
					}

				}
			} else if tipoExp == "bool" {
				// veo cual de los dos lados es el f64
				if left.info == "f64" {
					// el lado izquierdo es un f64

					// mi lado izquierdo es un id
					v.Code.WriteString(fmt.Sprintf("\tldr x0, =%s\n", left.addrID))
					v.Code.WriteString("\tldr d1, [x0]\n") // Cargar valor de la variable

					// lado derecho es un id int
					v.Code.WriteString(fmt.Sprintf("\tldr x1, =%s\n", right.addrID))
					v.Code.WriteString("\tldr w1, [x1]\n")
					v.Code.WriteString("scvtf d2, w1\n") // Convertir int a double

					switch op {
					case "<":
						v.Code.WriteString("\tfcmp d1, d2\n")
						v.Code.WriteString("\tcset w3, lt\n")
					case "<=":
						v.Code.WriteString("\tfcmp d1, d2\n")
						v.Code.WriteString("\tcset w3, le\n")
					case ">":
						v.Code.WriteString("\tfcmp d1, d2\n")
						v.Code.WriteString("\tcset w3, gt\n")
					case ">=":
						v.Code.WriteString("\tfcmp d1, d2\n")
						v.Code.WriteString("\tcset w3, ge\n")
					case "==":
						v.Code.WriteString("\tfcmp d1, d2\n")
						v.Code.WriteString("\tcset w3, eq\n")
					case "!=":
						v.Code.WriteString("\tfcmp d1, d2\n")
						v.Code.WriteString("\tcset w3, ne\n")
					}

					tpm := v.nextTemp()
					// Temporal donde se guradara el resultado bool
					v.Data.WriteString(fmt.Sprintf("%s: .word 0\n", tpm))
					v.Code.WriteString(fmt.Sprintf("\tldr x2, =%s\n", tpm))
					v.Code.WriteString("\tstr w3, [x2]\n\n")

					return Attr{
						addrID: tpm,
						info:   tipoExp,
						value:  result,
					}

				} else if right.info == "f64" {
					// el lado derecho es un f64

					// se que trae un int del lado izquierdo y es un id
					v.Code.WriteString(fmt.Sprintf("\tldr x0, =%s\n", left.addrID))
					v.Code.WriteString("\tldr w1, [x0]\n") // Cargar valor de la variable
					// lo convierto a double
					// lo casteo a double
					v.Code.WriteString("scvtf d1, w1\n") // Convertir

					// lado derecho es un id double
					v.Code.WriteString(fmt.Sprintf("\tldr x1, =%s\n", right.addrID))
					v.Code.WriteString("\tldr d2, [x1]\n") // Cargar valor de la variable

					switch op {
					case "<":
						v.Code.WriteString("\tfcmp d1, d2\n")
						v.Code.WriteString("\tcset w3, lt\n")
					case "<=":
						v.Code.WriteString("\tfcmp d1, d2\n")
						v.Code.WriteString("\tcset w3, le\n")
					case ">":
						v.Code.WriteString("\tfcmp d1, d2\n")
						v.Code.WriteString("\tcset w3, gt\n")
					case ">=":
						v.Code.WriteString("\tfcmp d1, d2\n")
						v.Code.WriteString("\tcset w3, ge\n")
					case "==":
						v.Code.WriteString("\tfcmp d1, d2\n")
						v.Code.WriteString("\tcset w3, eq\n")
					case "!=":
						v.Code.WriteString("\tfcmp d1, d2\n")
						v.Code.WriteString("\tcset w3, ne\n")
					}

					tpm := v.nextTemp()
					// Temporal donde se guradara el resultado bool
					v.Data.WriteString(fmt.Sprintf("%s: .word 0\n", tpm))
					v.Code.WriteString(fmt.Sprintf("\tldr x2, =%s\n", tpm))
					v.Code.WriteString("\tstr w3, [x2]\n\n")

					return Attr{
						addrID: tpm,
						info:   tipoExp,
						value:  result,
					}

				}
			}

		}

		// f64 + f64 = f64

		if left.info == "f64" && right.info == "f64" {
			var lval, rval float64
			lval = left.value.(float64)
			rval = right.value.(float64)

			var result interface{}
			tipoExp := "f64"

			switch op {
			case "+":
				result = lval + rval
			case "-":
				result = lval - rval
			case "*":
				result = lval * rval
			case "/":
				if rval == 0.0 {
					panic("División por cero " + fmt.Sprintf("%f / %f", lval, rval))
				}
				result = lval / rval
			case "<":
				result = lval < rval
				tipoExp = "bool"
			case "<=":
				result = lval <= rval
				tipoExp = "bool"
			case ">":
				result = lval > rval
				tipoExp = "bool"
			case ">=":
				result = lval >= rval
				tipoExp = "bool"
			case "==":
				result = lval == rval
				tipoExp = "bool"
			case "!=":
				result = lval != rval
				tipoExp = "bool"
			}

			if tipoExp == "f64" {

				// ya no necesito ver cual lado es el f64, ya que ambos son f64
				// se que el lado derecho es un id

				// el lado izquierdo es un tambien es un id f64

				// declaro mi registro temporal para el id izquierdo
				v.Code.WriteString(fmt.Sprintf("\tldr x0, =%s\n", left.addrID))
				v.Code.WriteString("\tldr d1, [x0]\n") // C

				// lado derecho es un id double
				v.Code.WriteString(fmt.Sprintf("\tldr x1, =%s\n", right.addrID))
				v.Code.WriteString("\tldr d2, [x1]\n") // Cargar valor de la variable

				switch op {
				case "+":
					v.Code.WriteString(("\tfadd d3, d1, d2\n"))
				case "-":
					v.Code.WriteString(("\tfsub d3, d1, d2\n"))
				case "*":
					v.Code.WriteString(("\tfmul d3, d1, d2\n"))
				case "/":
					v.Code.WriteString(("\tfdiv d3, d1, d2\n"))
				}

				tpm := v.nextTemp()
				// Temporal donde se guradara el resultado f64
				v.Data.WriteString(fmt.Sprintf("%s: .double 0.0\n", tpm))
				v.Code.WriteString(fmt.Sprintf("\tldr x2, =%s\n", tpm))
				v.Code.WriteString("\tstr d3, [x2]\n\n")

				return Attr{
					addrID: tpm,
					info:   tipoExp,
					value:  result,
				}

			} else if tipoExp == "bool" {
				// ya no necesito ver cual lado es el f64, ya que ambos son f64
				// se que el lado derecho es un id

				// el lado izquierdo es un f64

				// el lado izquierdo es un tambien es un id f64

				// declaro mi registro temporal para el id izquierdo
				v.Code.WriteString(fmt.Sprintf("\tldr x0, =%s\n", left.addrID))
				v.Code.WriteString("\tldr d1, [x0]\n") // C

				// lado derecho es un id double
				v.Code.WriteString(fmt.Sprintf("\tldr x1, =%s\n", right.addrID))
				v.Code.WriteString("\tldr d2, [x1]\n") // Cargar valor de la variable

				switch op {
				case "<":
					v.Code.WriteString("\tfcmp d1, d2\n")
					v.Code.WriteString("\tcset w3, lt\n")
				case "<=":
					v.Code.WriteString("\tfcmp d1, d2\n")
					v.Code.WriteString("\tcset w3, le\n")
				case ">":
					v.Code.WriteString("\tfcmp d1, d2\n")
					v.Code.WriteString("\tcset w3, gt\n")
				case ">=":
					v.Code.WriteString("\tfcmp d1, d2\n")
					v.Code.WriteString("\tcset w3, ge\n")
				case "==":
					v.Code.WriteString("\tfcmp d1, d2\n")
					v.Code.WriteString("\tcset w3, eq\n")
				case "!=":
					v.Code.WriteString("\tfcmp d1, d2\n")
					v.Code.WriteString("\tcset w3, ne\n")
				}

				tpm := v.nextTemp()
				// Temporal donde se guradara el resultado bool
				v.Data.WriteString(fmt.Sprintf("%s: .word 0\n", tpm))
				v.Code.WriteString(fmt.Sprintf("\tldr x2, =%s\n", tpm))
				v.Code.WriteString("\tstr w3, [x2]\n\n")

				return Attr{
					addrID: tpm,
					info:   tipoExp,
					value:  result,
				}

			}
		}

		// string + string = string
		if left.info == "string" && right.info == "string" {
			var lval, rval string
			lval = left.value.(string)
			rval = right.value.(string)

			var result interface{}
			tipoExp := "string"

			switch op {
			case "+":
				result = lval + rval
			case "<":
				result = lval < rval
				tipoExp = "bool"
			case "<=":
				result = lval <= rval
				tipoExp = "bool"
			case ">":
				result = lval > rval
				tipoExp = "bool"
			case ">=":
				result = lval >= rval
				tipoExp = "bool"
			case "==":
				result = lval == rval
				tipoExp = "bool"
			case "!=":
				result = lval != rval
				tipoExp = "bool"
			}

			if tipoExp == "string" {
				// uso copiar sin limpiar para concatenar cadenas
				// muevo el puntero para copiar donde termine la otra cadena

				// cadena final, sera de 64
				cfinal := v.nextTemp()
				v.BSS.WriteString(fmt.Sprintf("%s: .skip %d\n", cfinal, 64))           // espacio para la cadena de salida
				v.Code.WriteString(fmt.Sprintf("\tldr x0, =%s\t// destino\n", cfinal)) // cargar la direccion de memoria de salida
				v.Code.WriteString("\tmov x20, x0\t// guardar puntero base\n")
				// lado izquierdo es un id
				v.Code.WriteString(fmt.Sprintf("\tldr x1, =%s\t// origen\n", left.addrID)) // cargar la direccion de memoria del id
				v.Code.WriteString("\tmov x4, #64\t// Tamaño del buffer destino\n")
				v.Code.WriteString("\tbl copiar_sin_limpiar\n")
				v.Code.WriteString("\tmov x0, x20\t// restaurar puntero base\n")
				v.Code.WriteString("\tbl advance_to_end\n") // avanzar al final de la cadena de salida

				// se que lado derecho es un id
				v.Code.WriteString(fmt.Sprintf("\tldr x1, =%s\t// origen\n", right.addrID)) // cargar la direccion de memoria del id
				v.Code.WriteString("\tbl copiar_sin_limpiar\n")

				return Attr{
					addrID: cfinal,
					info:   tipoExp,
					value:  result,
				}
			} else if tipoExp == "bool" {

				// resultado de comparacion de strings bool
				resBool := v.nextTemp()
				v.Data.WriteString(fmt.Sprintf("%s: .word 0\n", resBool))

				// obtengo las dos direcciones de memoria de los strings
				// lado izquierdo es un id string
				v.Code.WriteString(fmt.Sprintf("\tldr x0, =%s\n", left.addrID)) // cargar la direccion de memoria del id

				// ya se que el lado derecho es un id
				// obtengo el valor de memoria del id
				v.Code.WriteString(fmt.Sprintf("\tldr x1, =%s\n", right.addrID))

				switch op {
				case "<":
					// mando a llamar a la funcion strcmp_lt
					v.Code.WriteString("\tbl strcmp_lt\n")
					// el resultado se guarda en w0
					v.Code.WriteString(fmt.Sprintf("\tldr x2, =%s\n", resBool)) // cargar la direccion de memoria del resultado
					v.Code.WriteString("\tstr w0, [x2]\n")                      // guardar el resultado en la direccion de memoria

				case "<=":
					// mando a llamar a la funcion strcmp_le
					v.Code.WriteString("\tbl strcmp_le\n")
					// el resultado se guarda en w0
					v.Code.WriteString(fmt.Sprintf("\tldr x2, =%s\n", resBool)) // cargar la direccion de memoria del resultado
					v.Code.WriteString("\tstr w0, [x2]\n")                      // guardar el resultado en la direccion de memoria

				case ">":
					// mando a llamar a la funcion strcmp_gt
					v.Code.WriteString("\tbl strcmp_gt\n")
					// el resultado se guarda en w0
					v.Code.WriteString(fmt.Sprintf("\tldr x2, =%s\n", resBool)) // cargar la direccion de memoria del resultado
					v.Code.WriteString("\tstr w0, [x2]\n")                      // guardar el resultado en la direccion de memoria

				case ">=":
					// mando a llamar a la funcion strcmp_ge
					v.Code.WriteString("\tbl strcmp_ge\n")
					// el resultado se guarda en w0
					v.Code.WriteString(fmt.Sprintf("\tldr x2, =%s\n", resBool)) // cargar la direccion de memoria del resultado
					v.Code.WriteString("\tstr w0, [x2]\n")                      // guardar el resultado en la direccion de memoria

				case "==":
					// mando a llamar a la funcion strcmp_eq
					v.Code.WriteString("\tbl strcmp_eq\n")
					// el resultado se guarda en w0
					v.Code.WriteString(fmt.Sprintf("\tldr x2, =%s\n", resBool)) // cargar la direccion de memoria del resultado
					v.Code.WriteString("\tstr w0, [x2]\n")                      // guardar el resultado en la direccion de memoria

				case "!=":
					// mando a llamar a la funcion strcmp_ne
					v.Code.WriteString("\tbl strcmp_ne\n")
					// el resultado se guarda en w0
					v.Code.WriteString(fmt.Sprintf("\tldr x2, =%s\n", resBool)) // cargar la direccion de memoria del resultado
					v.Code.WriteString("\tstr w0, [x2]\n")                      // guardar el resultado en la direccion de memoria

				}

				// tpm := v.nextTemp()
				// // Temporal donde se guradara el resultado bool
				// v.Data.WriteString(fmt.Sprintf("%s: .word 0\n", tpm))
				// v.Code.WriteString(fmt.Sprintf("\tldr x2, =%s\n", tpm))
				// v.Code.WriteString("\tstr w3, [x2]\n\n")

				return Attr{
					addrID: resBool,
					info:   tipoExp,
					value:  result,
				}
			}
		}

		// bool + bool = bool
		if left.info == "bool" && right.info == "bool" {
			var lval, rval bool
			lval = left.value.(bool)
			rval = right.value.(bool)

			var result interface{}
			tipoExp := "bool"
			switch op {
			case "==":
				result = lval == rval
			case "!=":
				result = lval != rval
			case "&&":
				result = lval && rval
			case "||":
				result = lval || rval
			}

			if tipoExp == "bool" {
				// lado izquierdo es un id bool
				v.Code.WriteString(fmt.Sprintf("\tldr x0, =%s\n", left.addrID)) // cargar la direccion de memoria del id
				v.Code.WriteString("\tldr w1, [x0]\n")                          //

				// lado derecho es un id
				v.Code.WriteString(fmt.Sprintf("\tldr x0, =%s\n", right.addrID)) // cargar la direccion de memoria del id
				v.Code.WriteString("\tldr w2, [x0]\n")                           // cargar el valor del id
				switch op {
				case "==":
					v.Code.WriteString("\tcmp w1, w2\n")  // comparar los dos valores
					v.Code.WriteString("\tcset w3, eq\n") // establecer w2 a 1 si son iguales, 0 si no
				case "!=":
					v.Code.WriteString("\tcmp w1, w2\n")  // comparar los dos valores
					v.Code.WriteString("\tcset w3, ne\n") // establecer w2 a 1 si son diferentes, 0 si no
				case "&&":
					v.Code.WriteString("\tand w3, w1, w2\n") // establecer w2 a 1 si ambos son verdaderos, 0 si alguno es falso
				case "||":
					v.Code.WriteString("\torr w3, w1, w2\n") // establecer w2 a 1 si alguno es verdadero, 0 si ambos son falsos

				}

				tpm := v.nextTemp()
				// Temporal donde se guradara el resultado bool
				v.Data.WriteString(fmt.Sprintf("%s: .word 0\n", tpm))
				v.Code.WriteString(fmt.Sprintf("\tldr x2, =%s\n", tpm))
				v.Code.WriteString("\tstr w3, [x2]\n\n")

				fmt.Println("Resultado de la operación booleano:", result)
				return Attr{
					addrID: tpm,
					info:   tipoExp,
					value:  result,
				}
			}
		}

	}

	panic("Operación no soportada o tipos incompatibles: " + left.info + " " + op + " " + right.info)
}

// ---------------------------------------------------------------
// |ParExpr|
// ---------------------------------------------------------------
func (v *CodeGenVisitor) VisitParExpr(ctx *parser.ParExprContext) interface{} {
	// Visita la expresión dentro de los paréntesis
	val := v.Visit(ctx.Expression()).(Attr)
	return val
}

// ---------------------------------------------------------------
// |IdExpr|
// ---------------------------------------------------------------
func (v *CodeGenVisitor) VisitIdExpr(ctx *parser.IdExprContext) interface{} {
	id := ctx.GetText()

	variable, ok := v.Env.GetVar(id)
	if ok {
		return Attr{addrID: id, info: variable.Tipo, value: variable.Valor}
	}

	panic("No se encontro el elemento: " + id)
}

// ---------------------------------------------------------------
// | IntExpr |
// ---------------------------------------------------------------

func (v *CodeGenVisitor) VisitIntExpr(ctx *parser.IntExprContext) interface{} {
	val := ctx.GetText()
	value, _ := strconv.ParseInt(ctx.GetText(), 10, 64)
	return Attr{addrNUM: val, info: "int", value: value}
}

// ---------------------------------------------------------------
// | FloatExpr |
// ---------------------------------------------------------------

func (v *CodeGenVisitor) VisitFloatExpr(ctx *parser.FloatExprContext) interface{} {
	val := ctx.GetText()
	// Podríamos validar que realmente es un número
	value, err := strconv.ParseFloat(val, 64)
	if err != nil {
		panic("Número flotante inválido: " + val)
	}
	return Attr{addrNUM: val, info: "f64", value: value}
	// por ahora como string, decidirás si lo manejas con VFP (registro flotante)
}

// ---------------------------------------------------------------
// | BoolExpr |
// ---------------------------------------------------------------
func (v *CodeGenVisitor) VisitBoolExpr(ctx *parser.BoolExprContext) interface{} {
	val := ctx.GetText()
	numeric := "0"
	if val == "true" {
		numeric = "1"
	}
	value, _ := strconv.ParseBool(ctx.GetText())
	return Attr{addrNUM: numeric, info: "bool", value: value} // Usamos .word 1 o 0 en memoria
}

func (v *CodeGenVisitor) VisitStrExpr(ctx *parser.StrExprContext) interface{} {
	raw := ctx.GetText() // Incluye comillas
	unquoted, err := strconv.Unquote(raw)
	if err != nil {
		panic("Error al procesar string: " + err.Error())
	}

	// aca lo regreso con un \n al final
	// return Attr{addrNUM: unquoted + "\\n", info: "string", value: unquoted}

	// lo regreso sin \n al final, ya que el print lo agrega
	return Attr{addrNUM: unquoted, info: "string", value: unquoted}
}

// ---------------------------------------------------------------
// | If |
// ---------------------------------------------------------------
func (v *CodeGenVisitor) VisitIf(ctx *parser.IfContext) interface{} {
	//Etiquetas para if
	endLabel := v.NextLabel()

	//Condicion if principal
	cond := v.Visit(ctx.Expression()).(Attr)
	v.Code.WriteString(fmt.Sprintf("\tldr x1, =%s\n", cond.addrID)) // Cargar la dirección de memoria del id
	v.Code.WriteString("\tldr w1, [x1]\n")                          // Cargar
	v.Code.WriteString("\tcmp w1, #0\n")                            // Comparar con 0

	if ctx.AllElseIfBlock() == nil && ctx.ElseBlock() == nil {
		// Si no hay else if ni else, solo salto al final
		v.Code.WriteString(fmt.Sprintf("\tb.eq %s\n", endLabel)) // Saltar si es falso
		v.Visit(ctx.Block())
		v.Code.WriteString(fmt.Sprintf("%s:\n", endLabel)) // Etiqueta de fin del if
		return nil
	}

	elseIfLabels := make([]string, len(ctx.AllElseIfBlock()))
	for i := range elseIfLabels {
		elseIfLabels[i] = v.NextLabel() // Generar etiquetas para cada else if
	}

	elseLabel := ""
	if ctx.ElseBlock() != nil {
		elseLabel = v.NextLabel()
	} else {
		elseLabel = endLabel
	}

	if len(elseIfLabels) > 0 {
		v.Code.WriteString(fmt.Sprintf("\tb.eq %s\n", elseIfLabels[0]))
	} else {
		v.Code.WriteString(fmt.Sprintf("\tb.eq %s\n", elseLabel))
	}

	//Se ejecuta el bloque del if
	v.Visit(ctx.Block())
	v.Code.WriteString(fmt.Sprintf("\tb %s\n", endLabel))

	//si hay else y luego ifs se ejecutaria lo siguiente
	for i, elseIfCtx := range ctx.AllElseIfBlock() {
		// Etiqueta para este else if
		v.Code.WriteString(fmt.Sprintf("%s:\n", elseIfLabels[i]))

		// Generar condición del else if
		elseIfCond := v.Visit(elseIfCtx.Expression()).(Attr)
		v.Code.WriteString(fmt.Sprintf("\tldr x1, =%s\n", elseIfCond.addrID))
		v.Code.WriteString("\tldr w1, [x1]\n")
		v.Code.WriteString("\tcmp w1, #0\n")

		// Determinar a dónde saltar si es falso
		nextLabel := elseLabel
		if i < len(elseIfLabels)-1 {
			nextLabel = elseIfLabels[i+1]
		} else if ctx.ElseBlock() != nil {
			nextLabel = elseLabel
		} else {
			nextLabel = endLabel
		}

		v.Code.WriteString(fmt.Sprintf("\tb.eq %s\n", nextLabel))

		// Bloque del else if
		v.Visit(elseIfCtx.Block())
		v.Code.WriteString(fmt.Sprintf("\tb %s\n", endLabel))
	}

	// Ejecutamos el else
	if ctx.ElseBlock() != nil {
		v.Code.WriteString(fmt.Sprintf("%s:\n", elseLabel))
		v.Visit(ctx.ElseBlock().Block())
	}

	// Etiqueta de fin
	v.Code.WriteString(fmt.Sprintf("%s:\n", endLabel))

	return nil
}

// ---------------------------------------------------------------
// ASIGNACION
// ---------------------------------------------------------------
func (v *CodeGenVisitor) VisitAssignmentExpr(ctx *parser.AssignmentExprContext) interface{} {
	varName := ctx.ID().GetText()
	exprAttr := v.Visit(ctx.Expression()).(Attr)

	// Buscar variable en la tabla de símbolos
	varSym, exists := v.Env.GetVar(varName)
	if !exists {
		panic(fmt.Sprintf("Variable no declarada: %s", varName))
	}

	// Verificar si es mutable
	/*if !varSym.IsMut {
		panic(fmt.Sprintf("No se puede asignar a variable inmutable: %s", varName))
	}*/

	// Verificar compatibilidad de tipos
	if !v.areTypesCompatible(varSym.Tipo, exprAttr.info) {
		panic(fmt.Sprintf("Tipos incompatibles en asignación: %s = %s", varSym.Tipo, exprAttr.info))
	}

	// Generar código ARM64 según el tipo
	v.generateAssignment(varSym, exprAttr)

	return nil
}

func (v *CodeGenVisitor) generateAssignment(dest *Variable, src Attr) {
	v.Code.WriteString(fmt.Sprintf("\t// Asignando a %s (%s)\n", dest.Nombre, dest.Tipo))
	//Cargaremos la direccion de memoria del destino
	v.Code.WriteString(fmt.Sprintf("\tldr x0, =%s\n", dest.Nombre))

	// Determinar si es un valor literal (numérico o string)
	isLiteral := src.addrNUM != "" && src.addrID == ""

	if isLiteral {
		// Caso 1: Es un valor literal
		switch dest.Tipo {
		case "int", "bool":
			v.Code.WriteString(fmt.Sprintf("\tmov w1, #%s\n", src.addrNUM))
			v.Code.WriteString("\tstr w1, [x0]\n")
		case "f64":
			// Para floats necesitamos cargar el valor desde memoria
			tempLabel := v.nextTemp()
			v.Data.WriteString(fmt.Sprintf("%s: .double %s\n", tempLabel, src.addrNUM))
			v.Code.WriteString(fmt.Sprintf("\tldr x1, =%s\n", tempLabel))
			v.Code.WriteString("\tldr d1, [x1]\n")
			v.Code.WriteString("\tstr d1, [x0]\n")
		case "string":
			// Manejo de strings literales
			v.BSS.WriteString(fmt.Sprintf("%s_buffer: .skip 64\n", dest.Nombre))
			v.Code.WriteString(fmt.Sprintf("\tldr x0, =%s\n", dest.Nombre))

			// Almacenar el string literal en .data
			tempStrLabel := v.nextTemp()
			strContent := strings.Trim(src.addrNUM, "\"") // Eliminar comillas si existen
			v.Data.WriteString(fmt.Sprintf("%s: .asciz \"%s\"\n", tempStrLabel, strContent))

			// Copiar el string
			v.Code.WriteString(fmt.Sprintf("\tldr x1, =%s\n", tempStrLabel))
			v.Code.WriteString("\tbl copiar_sin_limpiar\n")
		}
	} else {
		// Caso 2: Es una variable o expresión (tiene dirección en memoria)
		v.Code.WriteString(fmt.Sprintf("\tldr x1, =%s\n", src.addrID))

		switch dest.Tipo {
		case "int", "bool":
			v.Code.WriteString("\tldr w2, [x1]\n")
			v.Code.WriteString("\tstr w2, [x0]\n")
		case "f64":
			v.Code.WriteString("\tldr d1, [x1]\n")
			v.Code.WriteString("\tstr d1, [x0]\n")
		case "string":
			v.Code.WriteString("\tbl copiar_sin_limpiar\n")
		default:
			if strings.HasPrefix(dest.Tipo, "[]") {
				// Copiar slice (puntero + longitud + capacidad)
				v.Code.WriteString("\tldr x2, [x1]\n")
				v.Code.WriteString("\tstr x2, [x0]\n")
				v.Code.WriteString("\tldr x2, [x1, #8]\n")
				v.Code.WriteString("\tstr x2, [x0, #8]\n")
				v.Code.WriteString("\tldr x2, [x1, #16]\n")
				v.Code.WriteString("\tstr x2, [x0, #16]\n")
			} else {
				panic(fmt.Sprintf("Tipo no soportado en asignación: %s", dest.Tipo))
			}
		}
	}

	v.Code.WriteString(fmt.Sprintf("\t// Fin asignación a %s\n\n", dest.Nombre))

}

func (v *CodeGenVisitor) areTypesCompatible(t1, t2 string) bool {
	if t1 == t2 {
		return true
	}
	// Permitir algunas conversiones implícitas
	if t1 == "f64" && t2 == "int" {
		return true
	}
	return false
}

// ---------------------------------------------------------------
// |Incremento y Decremento|
// ---------------------------------------------------------------
// Incremento
func (v *CodeGenVisitor) VisitAssignmentIncrementExpr(ctx *parser.AssignmentIncrementExprContext) interface{} {
	varName := ctx.ID().GetText()
	varSym, exists := v.Env.GetVar(varName)
	if !exists {
		panic(fmt.Sprintf("Variable no declarada: %s", varName))
	}
	/*if !varSym.IsMut {
		panic(fmt.Sprintf("No se puede incrementar variable inmutable: %s", varName))
	}*/

	v.generateIncrementDecrement(varSym, 1)
	return nil
}

// Decremento
func (v *CodeGenVisitor) VisitAssignmentDecrementExpr(ctx *parser.AssignmentDecrementExprContext) interface{} {
	varName := ctx.ID().GetText()
	varSym, exists := v.Env.GetVar(varName)
	if !exists {
		panic(fmt.Sprintf("Variable no declarada: %s", varName))
	}
	/*if !varSym.IsMut {
		panic(fmt.Sprintf("No se puede incrementar variable inmutable: %s", varName))
	}*/

	v.generateIncrementDecrement(varSym, -1)
	return nil
}

// Funcio auxiliar para la generacion del incremento y decremento
// delta > 0 para incremento, delta < 0 para decremento
// para incrementar debe enviarse delta = 1 y para decrementar delta = -1
func (v *CodeGenVisitor) generateIncrementDecrement(varSym *Variable, delta int) {
	v.Code.WriteString(fmt.Sprintf("\tldr x1, =%s\n", varSym.Nombre))

	switch varSym.Tipo {
	case "int":
		v.Code.WriteString("\tldr w2, [x1]\n")
		if delta > 0 {
			v.Code.WriteString(fmt.Sprintf("\tadd w2, w2, #%d\n", delta))
		} else {
			v.Code.WriteString(fmt.Sprintf("\tsub w2, w2, #%d\n", -delta))
		}
		v.Code.WriteString("\tstr w2, [x1]\n")
	case "f64":
		v.Code.WriteString("\tldr d2, [x1]\n")
		temp := v.nextTemp()
		v.Data.WriteString(fmt.Sprintf("%s: .double %d.0\n", temp, delta))
		v.Code.WriteString(fmt.Sprintf("\tldr x3, =%s\n", temp))
		v.Code.WriteString("\tldr d3, [x3]\n")
		if delta > 0 {
			v.Code.WriteString("\tfadd d2, d2, d3\n")
		} else {
			v.Code.WriteString("\tfsub d2, d2, d3\n")
		}
		v.Code.WriteString("\tstr d2, [x1]\n")
	default:
		panic(fmt.Sprintf("No se puede incrementar/decrementar tipo: %s", varSym.Tipo))
	}
}

// ---------------------------------------------------------------
// | FOR |
// ---------------------------------------------------------------
func (v *CodeGenVisitor) VisitForSimple(ctx *parser.ForSimpleContext) interface{} {
	startLabel := v.NextLabel()
	endLabel := v.NextLabel()

	//metemos el contexto del bucle actual
	v.loopStack = append(v.loopStack, LoopContext{
		startLabel: startLabel,
		endLabel:   endLabel,
	})

	// Aqui inicia el bucle
	v.Code.WriteString(fmt.Sprintf("%s:\n", startLabel))

	cond := v.Visit(ctx.Expression()).(Attr)
	v.Code.WriteString(fmt.Sprintf("\tldr x1, =%s\n", cond.addrID))
	v.Code.WriteString("\tldr w1, [x1]\n")
	v.Code.WriteString("\tcmp w1, #0\n")
	v.Code.WriteString(fmt.Sprintf("\tb.eq %s\n", endLabel)) // salta al final si es falso

	//Ejecucion del codigo dentro del bloque del for
	v.Visit(ctx.Block())
	// Regresamos al inicio
	v.Code.WriteString(fmt.Sprintf("\tb %s\n", startLabel))
	// fin del bucle
	v.Code.WriteString(fmt.Sprintf("%s:\n", endLabel))

	v.loopStack = v.loopStack[:len(v.loopStack)-1] // Quitamos el contexto del bucle actual
	return nil
}

// ---------------------------------------------------------------
// | Break |
// ---------------------------------------------------------------
func (v *CodeGenVisitor) VisitBreak(ctx *parser.BreakStatementContext) interface{} {
	if len(v.loopStack) == 0 {
		panic("break fuera de bucle")
	}
	currentLoop := v.loopStack[len(v.loopStack)-1]
	v.Code.WriteString(fmt.Sprintf("\tb %s\n", currentLoop.endLabel))
	return nil
}

// ---------------------------------------------------------------
// | Continue |
// ---------------------------------------------------------------
func (v *CodeGenVisitor) VisitContinue(ctx *parser.ContinueStatementContext) interface{} {
	if len(v.loopStack) == 0 {
		panic("continue fuera de bucle")
	}
	currentLoop := v.loopStack[len(v.loopStack)-1]
	v.Code.WriteString(fmt.Sprintf("\tb %s\n", currentLoop.startLabel))
	return nil
}
