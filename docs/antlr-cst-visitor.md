# ANTLR, CST y patrón Visitor

Esta guía explica cómo se utiliza el patrón **Visitor** con ANTLR y Go en V-Lang Cherry. En el proyecto existen **dos Visitors distintos** que recorren el mismo tipo de árbol generado por ANTLR, pero realizan tareas diferentes:

- `interpreter.go` → interpreta directamente el programa.
- `codegen.go` → recorre el mismo CST, pero genera código ARM64.

!!! abstract "Qué aprenderás"

    Comprenderás cómo ANTLR construye el **árbol de sintaxis concreta (CST)**, cómo un Visitor recorre sus nodos y por qué el intérprete y el generador ARM64 pueden utilizar el mismo árbol para obtener resultados distintos.

!!! info "Idea central"

    **ANTLR construye el árbol; cada Visitor determina qué acción debe realizarse al visitar sus nodos.**

---

## 1. Flujo general

ANTLR se encarga de construir la estructura sintáctica del programa:

```text
Código fuente
     ↓
Lexer
     ↓
Tokens
     ↓
Parser
     ↓
CST
```

Después, se selecciona qué Visitor se utilizará:

```text
                     ┌─────────────────────┐
                     │         CST         │
                     └──────────┬──────────┘
                                │
                     mismo árbol generado
                                │
                 ┌──────────────┴──────────────┐
                 │                             │
                 ▼                             ▼
       Visitor intérprete             Visitor generador
        interpreter.go                   codegen.go
                 │                             │
                 ▼                             ▼
      Ejecuta el programa              Genera ARM64.s
```

El mismo CST puede recorrerse más de una vez, seleccionando en cada caso la implementación del Visitor que corresponde al objetivo del programa.

---

## 2. Generación del parser y Visitor

A partir de `Gramatica.g4` se ejecuta:

```bash
antlr4 -Dlanguage=Go -visitor -listener -o parser Gramatica.g4
```

ANTLR genera archivos como:

```text
parser/
├── gramatica_lexer.go
├── gramatica_parser.go
├── gramatica_visitor.go
├── gramatica_base_visitor.go
├── gramatica_listener.go
├── Gramatica.tokens
└── Gramatica.interp
```

Los más importantes para este flujo son:

| Archivo | Función |
|---|---|
| `gramatica_lexer.go` | Convierte el texto en tokens |
| `gramatica_parser.go` | Aplica la gramática y construye el CST |
| `gramatica_visitor.go` | Define los métodos que puede implementar un Visitor |

---

## 3. Construcción del CST

Tanto el intérprete como el generador ARM64 comienzan prácticamente igual:

```go
input := antlr.NewInputStream(code)

lexer := parser.NewGramaticaLexer(input)
tokens := antlr.NewCommonTokenStream(lexer, 0)

p := parser.NewGramaticaParser(tokens)
p.BuildParseTrees = true

tree := p.Program()
```

El flujo es:

```text
code
 ↓
antlr.NewInputStream
 ↓
GramaticaLexer
 ↓
TokenStream
 ↓
GramaticaParser
 ↓
p.Program()
 ↓
CST
```

`Program()` corresponde a la regla inicial de la gramática:

```antlr
program: block EOF;
```

Por eso la raíz del árbol es un `ProgramContext`.

---

## 4. ¿Cómo se recorre el CST?

Supongamos este programa:

```v
fn main() {
    x int = 10
    println(x)
}
```

De forma simplificada, el CST se puede imaginar así:

```text
Program
└── Block
    └── Statement
        └── MainFunction
            └── Block
                ├── Statement
                │   └── Declaration
                │       ├── ID: x
                │       ├── Type: int
                │       └── Expression
                │           └── INT: 10
                │
                └── Statement
                    └── Print
                        └── Expression
                            └── ID: x
```

El Visitor comienza arriba y va entrando en los nodos que necesita:

```text
Program
  ↓
Block
  ↓
Statement
  ↓
Main
  ↓
Block
  ↓
Declaration
  ↓
Expression
  ↓
Literal
```

Después regresa resultados hacia arriba.

---

## 5. El mismo recorrido, dos implementaciones

La parte central del patrón es esta:

```text
            CST
             │
             ▼
          Visit(...)
             │
      detectar tipo de nodo
             │
     ┌───────┼────────┐
     ▼       ▼        ▼
Declaration Print    OpExpr
     │       │        │
     ▼       ▼        ▼
 lógica    lógica    lógica
 propia    propia    propia
```

Los dos Visitors trabajan sobre contextos generados por ANTLR como:

```go
*parser.ProgramContext
*parser.BlockContext
*parser.StatementContext
*parser.PrintContext
*parser.OpExprContext
*parser.IntExprContext
```

Lo que cambia es **qué hacen al visitar cada nodo**.

---

## 6. Visitor del intérprete

En `runProgram` se construye el Visitor del intérprete:

```go
eval := &interpreter.Visitor{
    Env:        interpreter.NewEnv(nil, "Global"),
    Console:    "",
    ErrorTable: errorTable,
}

eval.Visit(tree)
```

Este Visitor interpreta directamente el programa.

Por ejemplo:

```v
x int = 10
println(x)
```

el Visitor puede:

```text
1. Visitar la declaración.
2. Evaluar la expresión 10.
3. Guardar x en el entorno.
4. Visitar println.
5. Buscar x en el entorno.
6. Obtener su valor.
7. Agregar "10" a la consola.
```

No genera ensamblador.

Su resultado principal es la **ejecución semántica del programa**.

---

## 7. Entorno del intérprete

El intérprete sí mantiene un entorno completo:

```go
type Env struct {
    name      string
    parent    *Env
    vars      map[string]*Variable
    slicesL   map[string]*Slice
    structDcl map[string]*StructDcl
    structIns map[string]*StructInstance
    funcs     map[string]*Function
}
```

Este entorno representa el estado del programa mientras se interpreta.

Guarda:

```text
variables
slices
structs
instancias de structs
funciones
ámbitos
```

Por ejemplo:

```v
x int = 10
```

puede representarse conceptualmente como:

```text
Env: Main

x
├── Tipo: int
└── Valor: 10
```

Cuando posteriormente aparece:

```v
println(x)
```

el Visitor consulta el entorno y obtiene el valor `10`.

---

## 8. Visitor del generador ARM64

Para traducir se utiliza otro Visitor:

```go
gen := codegen.NewCodeGenVisitor()

gen.Visit(tree)
```

El recorrido del CST sigue siendo el mismo.

La diferencia es que ahora cada nodo se traduce a texto ARM64.

Ejemplo:

```v
x int = 10
```

en lugar de ejecutarse inmediatamente puede producir:

```asm
.data
x: .word 10
```

Y una operación puede generar instrucciones como:

```asm
ldr w0, [x1]
add w2, w0, w3
str w2, [x4]
```

---

## 9. El generador ARM64 no usa el entorno como un intérprete

Aquí conviene hacer una precisión.

El generador **sí posee una estructura `Env` interna**, pero no funciona como el entorno completo de ejecución del intérprete.

En `codegen.go` se tiene una estructura mucho más pequeña:

```go
type Env struct {
    name   string
    parent *Env
    vars   map[string]*Variable
}
```

Se utiliza principalmente para recordar información necesaria durante la traducción:

```text
qué variables existen
qué tipo poseen
en qué ámbito fueron declaradas
información útil para generar el ensamblador
```

!!! note "Entorno del generador"

    El generador ARM64 no necesita reproducir todo el estado de ejecución del programa como el intérprete. Mantiene únicamente la información auxiliar de símbolos necesaria para traducir el CST a ensamblador.

Es decir:

```text
INTERPRETER
Env = estado real de ejecución
      variables
      valores
      slices
      structs
      funciones
      ámbitos
```

mientras que:

```text
CODEGEN
Env = apoyo para traducción
      variables
      tipos
      ámbitos
      información necesaria para generar ARM64
```

El resultado final del generador no es el entorno:

```text
resultado → ARM64.s
```

---

## 10. Estado propio del CodeGenVisitor

El generador necesita otro tipo de información:

```go
type CodeGenVisitor struct {
    parser.GramaticaVisitor

    Env        *Env
    TemporalId int
    SalidaId   int

    Data strings.Builder
    BSS  strings.Builder
    Code strings.Builder
    Fun  strings.Builder

    LabelCnt int
}
```

Mientras recorre el árbol va construyendo:

```text
Data → .data
BSS  → .bss
Code → .text
Fun  → funciones auxiliares
```

Finalmente:

```go
salida =
    gen.Data.String() +
    gen.BSS.String() +
    gen.Code.String() +
    gen.Fun.String()
```

y se escribe:

```go
os.WriteFile("ARM64.s", []byte(salida), 0644)
```

---

## 11. Ejemplo de la diferencia

Código V-Lang Cherry:

```v
x int = 10
y int = 20
z int = x + y

println(z)
```

## Visitor intérprete

Conceptualmente realiza:

```text
x = 10
y = 20

buscar x → 10
buscar y → 20

10 + 20 → 30

z = 30

println(z) → "30"
```

Resultado:

```text
30
```

---

## Visitor generador ARM64

Conceptualmente realiza:

```text
declaración x
      ↓
escribir .word

declaración y
      ↓
escribir .word

operación x + y
      ↓
generar ldr/add/str

println
      ↓
generar llamada a rutina ARM64
```

Resultado:

```text
ARM64.s
```

El cálculo real lo realizará posteriormente el procesador ARM64 o QEMU.

---

## 12. `Visit` como despachador

En ambos Visitors existe una lógica similar:

```go
func (v *CodeGenVisitor) Visit(tree antlr.ParseTree) interface{} {

    switch val := tree.(type) {

    case *parser.ProgramContext:
        return v.VisitProgram(val)

    case *parser.BlockContext:
        return v.VisitBlock(val)

    case *parser.PrintContext:
        return v.VisitPrint(val)

    case *parser.OpExprContext:
        return v.VisitOpExpr(val)
    }

    return nil
}
```

Puede entenderse como:

```text
¿Qué nodo estoy viendo?
        │
        ├── Program → VisitProgram
        ├── Block → VisitBlock
        ├── Print → VisitPrint
        ├── If → VisitIf
        ├── For → VisitFor
        └── OpExpr → VisitOpExpr
```

---

## 13. Recorrido de expresiones

Para:

```v
10 + 5
```

el árbol simplificado sería:

```text
OpExpr (+)
├── IntExpr (10)
└── IntExpr (5)
```

El Visitor hace:

```go
left := v.Visit(ctx.GetLeft())
right := v.Visit(ctx.GetRight())

op := ctx.GetOp().GetText()
```

Primero visita:

```text
10
```

después:

```text
5
```

y finalmente procesa:

```text
+
```

---

## 14. Retorno de información entre nodos

En el intérprete se utiliza una estructura como:

```go
type Value struct {
    value interface{}
    info  string
}
```

Por ejemplo:

```text
VisitIntExpr(10)
        ↓
Value{
    value: 10,
    info: "int"
}
```

En el generador se utiliza:

```go
type Attr struct {
    addrNUM string
    addrID  string
    info    string
    value   interface{}
}
```

Esto permite transportar información entre nodos durante la traducción.

Ejemplo:

```text
Literal 10
   ↓
addrNUM = "10"
info    = "int"
```

o una variable:

```text
x
 ↓
addrID = "x"
info   = "int"
```

---

## 15. Direcciones y punteros al generar ARM64

Cuando se genera ARM64 no basta con conocer el valor.

También se necesita conocer **dónde estará almacenado**.

Ejemplo:

```asm
ldr x0, =texto
```

Aquí `x0` obtiene la dirección del buffer `texto`.

Después:

```asm
ldr x1, =tmp_1
```

`x1` apunta al origen.

Y se puede llamar:

```asm
bl copiar_sin_limpiar
```

Conceptualmente:

```text
x0 → destino
x1 → origen
```

---

## 16. Uso de un puntero base

Para concatenar strings se utiliza una idea como:

```asm
ldr x0, =salida_1
mov x20, x0
```

Después `x0` puede avanzar:

```text
salida_1
   │
   ▼
[H][o][l][a][ ][...]
 ↑             ↑
x20            x0
```

`x20` conserva la dirección inicial.

`x0` puede avanzar mientras se escriben bytes.

Cuando se necesita recuperar el inicio:

```asm
mov x0, x20
```

o para imprimir:

```asm
mov x1, x20
bl println
```

---

## 17. Resumen de ambos Visitors

| | Intérprete | Generador ARM64 |
|---|---|---|
| Entrada | CST | CST |
| Recorrido | Visitor | Visitor |
| Lenguaje de implementación | Go | Go |
| Ejecuta semántica | Sí | No directamente |
| Mantiene valores en ejecución | Sí | Solo información auxiliar |
| Entorno | Completo | Reducido |
| Resultado | Salida del programa | `ARM64.s` |
| Operaciones | Go calcula | ARM64 calculará |
| `println` | Escribe en `Console` | Genera instrucciones |
| Variables | Valores reales | Símbolos/direcciones/tipos |

---

## 18. Diagrama para la presentación

La imagen anterior sigue siendo válida porque representa el flujo general:

```text
Código
  ↓
Lexer
  ↓
Parser
  ↓
CST
  ↓
Visitor
  ↓
Implementación elegida
```

Con el CST existen **dos caminos**, usando un Visitor para interpretar o generar ARM64:

```text
                    CST
                     │
             ┌───────┴───────┐
             ▼               ▼
      Interpreter        Code Generator
             │               │
             ▼               ▼
         Console          ARM64.s
```

---

## 19. Explicación corta

Para explicarlo durante la conferencia puede utilizarse este guion:

!!! quote "Mismo árbol, diferentes procesos"

    Se puede utilizar el mismo parser y el mismo CST para dos procesos distintos. En el primer caso, se crea un Visitor que interpreta el programa directamente en Go. Este Visitor mantiene un entorno completo con variables, slices, structs y funciones.

!!! quote "Traducción a ARM64"

    Para la traducción se recorre nuevamente el mismo tipo de árbol, pero con otro Visitor. En lugar de ejecutar cada operación, este Visitor escribe instrucciones ARM64 en acumuladores para `.data`, `.bss`, `.text` y las funciones auxiliares.

!!! quote "Diferencia entre los entornos"

    El generador también conserva información sobre las variables, pero no necesita mantener un entorno de ejecución tan completo como el intérprete. Solo guarda los datos necesarios para traducir correctamente cada nodo.

---

## 20. Idea final

```text
                   ANTLR
                     │
                     ▼
                    CST
                     │
            ┌────────┴────────┐
            │                 │
            ▼                 ▼
       Visitor 1          Visitor 2
      Interpreter          CodeGen
            │                 │
            ▼                 ▼
       Ejecutar             Traducir
            │                 │
            ▼                 ▼
        Console             ARM64.s
```

!!! success "Conclusión"

    **El patrón Visitor no cambia el árbol: cambia la acción que se realiza al visitar cada nodo.**
