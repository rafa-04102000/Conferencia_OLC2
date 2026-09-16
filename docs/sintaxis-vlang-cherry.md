# Sintaxis de V-Lang Cherry

Esta guía describe la **sintaxis soportada por el intérprete de V-Lang Cherry** a partir de la gramática y de los archivos de prueba utilizados durante el desarrollo.

!!! warning "Diferencia entre intérprete y traductor"

    Este documento describe la sintaxis admitida por el **intérprete**. El traductor a ARM64 no implementa necesariamente todas estas características; por eso, su alcance debe comprobarse por separado.

!!! tip "Cómo usar esta referencia"

    Busca una construcción en la tabla de contenido, prueba el ejemplo mínimo y después intégrala en un programa completo dentro de `fn main()`.

---

## 1. Estructura general de un programa

El punto de entrada del programa es:

```v
fn main() {
    // instrucciones
}
```

Las instrucciones que se desean ejecutar deben colocarse dentro de `fn main()`.

También pueden existir declaraciones de `struct` y funciones auxiliares. En los programas de prueba se escribieron antes de `main`:

```v
struct Persona {
    string nombre
    int edad
}

fn sumar(a int, b int) int {
    return a + b
}

fn main() {
    resultado int = sumar(10, 20)
    println(resultado)
}
```

Por claridad se recomienda utilizar esta organización:

```text
structs
   ↓
funciones auxiliares
   ↓
fn main()
```

---

## 2. Comentarios

V-Lang Cherry permite comentarios de una línea:

```v
// Este es un comentario
```

También permite comentarios de bloque:

```v
/*
    Este es un
    comentario de varias líneas
*/
```

---

## 3. Tipos de datos primitivos

Los tipos primitivos disponibles son:

| Tipo | Descripción | Ejemplo |
|---|---|---|
| `int` | Número entero | `42` |
| `f64` | Número decimal | `3.14159` |
| `string` | Cadena de texto | `"Hola"` |
| `bool` | Valor lógico | `true`, `false` |

Ejemplo:

```v
fn main() {
    edad int = 25
    altura f64 = 1.75
    nombre string = "Ana"
    activo bool = true
}
```

---

## 4. Declaración de variables

## 4.1 Variable con tipo y valor

La forma más común es:

```v
nombre tipo = valor
```

Ejemplos:

```v
fn main() {
    entero int = 42
    decimal f64 = 3.14159
    texto string = "Hola, mundo!"
    booleano bool = true
}
```

También se acepta la palabra reservada `mut`:

```v
fn main() {
    mut entero int = 42
    mut texto string = "Hola"
}
```

### Sobre `mut`

La gramática reconoce `mut`, pero en la versión actual del intérprete **no se utiliza para controlar la mutabilidad**.

Por lo tanto, estas dos declaraciones son válidas:

```v
numero int = 10
```

```v
mut numero int = 10
```

En la práctica, `mut` puede considerarse opcional dentro de este proyecto.

---

## 4.2 Variable con tipo y sin valor

La gramática también permite:

```v
nombre tipo
```

Por ejemplo:

```v
fn main() {
    numero int
    decimal f64
    texto string
    bandera bool
}
```

Los valores por defecto manejados por el intérprete son conceptualmente:

```text
int     → 0
f64     → 0.0
string  → ""
bool    → false
```

Para evitar ambigüedades en ejemplos de clase también puede escribirse el valor explícitamente:

```v
numero int = 0
decimal f64 = 0.0
texto string = ""
bandera bool = false
```

---

## 5. Asignación de variables

Una variable existente puede recibir un nuevo valor:

```v
fn main() {
    numero int = 10
    numero = 20

    texto string = "Hola"
    texto = "Texto modificado"
}
```

El intérprete comprueba los tipos. Por ejemplo, asignar un `string` a una variable `int` produce un error semántico:

```v
numero int = 10

// ERROR
numero = "Hola"
```

---

## 6. Operadores de asignación

Se encuentran soportadas las siguientes formas:

```v
x = 10
x += 5
x -= 5
x++
x--
```

Ejemplo:

```v
fn main() {
    contador int = 0

    contador += 5
    contador -= 2
    contador++
    contador--

    println(contador)
}
```

---

## 7. Operaciones aritméticas

Se pueden utilizar:

```text
+   suma
-   resta
*   multiplicación
/   división
%   módulo
```

Ejemplo:

```v
fn main() {
    suma int = 10 + 5
    resta int = 10 - 5
    producto int = 10 * 5
    division int = 10 / 5
    residuo int = 10 % 3

    println(suma)
}
```

El intérprete también permite operaciones entre `int` y `f64` cuando corresponde:

```v
fn main() {
    resultado f64 = 10 + 5.5
    println(resultado)
}
```

---

## 8. Operadores relacionales

Están disponibles:

```text
==   igual
!=   diferente
>    mayor que
<    menor que
>=   mayor o igual
<=   menor o igual
```

Ejemplo:

```v
fn main() {
    numero int = 10

    if numero >= 5 {
        println("El número es mayor o igual a 5")
    }
}
```

---

## 9. Operadores lógicos

Se soportan:

```text
&&   AND
||   OR
!    NOT
```

Ejemplo:

```v
fn main() {
    activo bool = true
    autorizado bool = false

    if activo && !autorizado {
        println("Condición verdadera")
    }
}
```

También pueden utilizarse paréntesis:

```v
if (10 == 10 && 5 == 5) || false {
    println("Correcto")
}
```

---

## 10. Impresión con `println`

Se utiliza:

```v
println(valor)
```

Ejemplos:

```v
fn main() {
    println(42)
    println(3.14)
    println("Hola")
    println(true)
}
```

Puede recibir varios valores:

```v
println("Edad:", 25)
println("Hola", 123, true)
```

También puede llamarse sin argumentos:

```v
println()
```

---

## 11. Punto y coma

El punto y coma es **opcional**.

Estas dos instrucciones son equivalentes:

```v
numero int = 10
```

```v
numero int = 10;
```

Por consistencia, los ejemplos de este repositorio pueden escribirse sin `;`.

---

## 12. Valores nulos

V-Lang Cherry **no implementa un literal `null` o `nil`**.

No existe una sintaxis como:

```v
// NO soportado
dato = null
```

```v
// NO soportado
dato = nil
```

Una variable debe existir antes de ser utilizada. Si se intenta acceder a un identificador que no ha sido declarado, el intérprete debe reportar un error.

Ejemplo incorrecto:

```v
fn main() {
    // ERROR: la variable no existe
    println(variable_inexistente)
}
```

Por tanto, en este intérprete el manejo de errores por variables inexistentes **no equivale a tener valores nulos**.

---

## 13. Condicional `if`

Sintaxis:

```v
if condicion {
    // instrucciones
}
```

Ejemplo:

```v
fn main() {
    edad int = 20

    if edad >= 18 {
        println("Mayor de edad")
    }
}
```

---

## 14. `if`, `else if` y `else`

```v
fn main() {
    numero int = 10

    if numero > 10 {
        println("Mayor")
    } else if numero == 10 {
        println("Igual")
    } else {
        println("Menor")
    }
}
```

---

## 15. Ciclo `for` como `while`

V-Lang Cherry utiliza `for` para expresar un ciclo condicionado:

```v
fn main() {
    i int = 0

    for i < 5 {
        println(i)
        i = i + 1
    }
}
```

También puede utilizarse:

```v
for true {
    // ciclo
}
```

y salir mediante `break`.

---

## 16. `for` clásico

También está soportada la forma:

```v
for inicializacion; condicion; actualizacion {
    // instrucciones
}
```

Ejemplo:

```v
fn main() {
    for i int = 0; i < 5; i++ {
        println(i)
    }
}
```

Los ciclos pueden anidarse:

```v
fn main() {
    for i int = 1; i <= 3; i++ {
        for j int = 1; j <= 3; j++ {
            println(i, "x", j, "=", i * j)
        }
    }
}
```

---

## 17. `for` sobre slices

La sintaxis utilizada es:

```v
for indice, valor in slice {
    // instrucciones
}
```

Ejemplo:

```v
fn main() {
    numeros = []int {10, 20, 30, 40, 50}

    for indice, valor in numeros {
        println("Índice", indice, "=", valor)
    }
}
```

---

## 18. `break`

Permite detener un ciclo:

```v
fn main() {
    contador int = 0

    for true {
        contador++

        if contador >= 5 {
            break
        }
    }
}
```

También está reconocido dentro de `switch`.

---

## 19. `continue`

Permite saltar a la siguiente iteración:

```v
fn main() {
    contador int = 0

    for contador < 10 {
        contador++

        if contador % 2 == 0 {
            continue
        }

        println(contador)
    }
}
```

---

## 20. `switch`

Sintaxis:

```v
switch expresion {
case valor:
    instrucciones
case valor:
    instrucciones
default:
    instrucciones
}
```

Ejemplo:

```v
fn main() {
    dia int = 2

    switch dia {
    case 1:
        println("Lunes")
    case 2:
        println("Martes")
    case 3:
        println("Miércoles")
    default:
        println("Día desconocido")
    }
}
```

Puede utilizarse `break`:

```v
switch numero {
case 1:
    println("Uno")
case 2:
    println("Dos")
    break
default:
    println("Otro")
}
```

---

## 21. Slices

El intérprete soporta slices de una dimensión.

## 21.1 Declaración con valores

```v
fn main() {
    numeros = []int {1, 2, 3, 4, 5}
    palabras = []string {"Hola", "mundo"}
}
```

Los tipos de slice posibles se construyen a partir de los tipos primitivos:

```text
[]int
[]f64
[]string
[]bool
```

---

## 21.2 Declaración sin valores

La gramática permite una declaración como:

```v
fn main() {
    numeros []int
}
```

---

## 21.3 Acceso por índice

```v
fn main() {
    numeros = []int {10, 20, 30}

    primero int = numeros[0]

    println(primero)
}
```

Los índices comienzan en `0`.

---

## 21.4 Modificación de un elemento

```v
fn main() {
    numeros = []int {1, 2, 3}

    numeros[0] = 10

    println(numeros)
}
```

---

## 22. `indexOf`

Busca un elemento dentro de un slice.

Sintaxis:

```v
indexOf(slice, valor)
```

Ejemplo:

```v
fn main() {
    numeros = []int {10, 20, 30, 40, 50}

    indice int = indexOf(numeros, 30)

    println(indice)
}
```

Si el elemento existe, retorna su índice.

Si no existe, en las pruebas del intérprete retorna:

```text
-1
```

Ejemplo:

```v
indice int = indexOf(numeros, 100)
```

---

## 23. `join`

Permite unir los elementos de un slice de cadenas.

Sintaxis:

```v
join(slice, separador)
```

Ejemplo:

```v
fn main() {
    palabras = []string {"Hola", "mundo", "desde", "Go"}

    frase string = join(palabras, " ")

    println(frase)
}
```

Resultado:

```text
Hola mundo desde Go
```

También:

```v
frase string = join(palabras, ", ")
```

Resultado:

```text
Hola, mundo, desde, Go
```

---

## 24. `append`

Permite agregar un elemento a un slice.

```v
fn main() {
    numeros = []int {1, 2, 3}

    numeros = append(numeros, 4)

    println(numeros)
}
```

Resultado esperado:

```text
[1, 2, 3, 4]
```

En esta versión se probó el agregado de **un elemento individual**.

---

## 25. `len`

La palabra `len` está reconocida por la gramática:

```v
len(numeros)
```

Sin embargo, **no se considera una funcionalidad completa del intérprete** para esta documentación.

Por tanto, para los ejercicios de los estudiantes debe considerarse:

```text
len → no implementado / no requerido
```

Aunque existan pruebas parciales en archivos antiguos, no debe asumirse como una característica soportada del intérprete.

---

## 26. Slices o arrays multidimensionales

Los arrays/slices bidimensionales **no fueron implementados**.

Por ejemplo, esta forma no debe utilizarse:

```v
// NO soportado
matriz = [][]int {
    {1, 2, 3},
    {4, 5, 6},
    {7, 8, 9}
}
```

Tampoco:

```v
// NO soportado
valor int = matriz[1][2]
```

ni:

```v
// NO soportado
matriz[0][0] = 100
```

Esta fue una de las funcionalidades pendientes del intérprete.

---

## 27. Funciones

Las funciones se declaran utilizando `fn`.

## 27.1 Función sin parámetros ni retorno

```v
fn saludar() {
    println("¡Hola, mundo!")
}

fn main() {
    saludar()
}
```

---

## 27.2 Función con retorno

El tipo de retorno se escribe después de los paréntesis:

```v
fn obtener_numero() int {
    return 42
}
```

Uso:

```v
fn main() {
    numero int = obtener_numero()
    println(numero)
}
```

---

## 27.3 Función con parámetros

```v
fn sumar(a int, b int) int {
    return a + b
}
```

Uso:

```v
fn main() {
    resultado int = sumar(10, 20)
    println(resultado)
}
```

---

## 27.4 Parámetros de tipo slice

La gramática también contempla parámetros de slice:

```v
fn procesar(valores []int) {
    // instrucciones
}
```

---

## 28. Funciones recursivas

El intérprete permite llamadas recursivas.

Factorial:

```v
fn factorial(n int) int {
    if n <= 1 {
        return 1
    }

    return n * factorial(n - 1)
}
```

Uso:

```v
fn main() {
    resultado int = factorial(5)
    println(resultado)
}
```

Fibonacci:

```v
fn fibonacci(n int) int {
    if n <= 1 {
        return n
    }

    return fibonacci(n - 1) + fibonacci(n - 2)
}
```

---

## 29. `return`

Puede retornar un valor:

```v
return expresion
```

Ejemplo:

```v
fn sumar(a int, b int) int {
    return a + b
}
```

También puede utilizarse sin valor:

```v
fn mostrar() {
    println("Hola")
    return
}
```

---

## 30. Funciones nativas implementadas

## `Atoi`

Convierte un `string` numérico a `int`.

```v
fn main() {
    numero int = Atoi("123")
    println(numero)
}
```

---

## `parseFloat`

La gramática reconoce:

```v
parseFloat(expresion)
```

En algunos archivos de prueba también se utilizó una función auxiliar escrita en el propio lenguaje llamada `parse_float`.

Cuando se quiera utilizar la función nativa definida por la gramática, la forma es:

```v
valor f64 = parseFloat("123.45")
```

---

## `typeOf`

Permite consultar el tipo de un valor:

```v
fn main() {
    t1 string = typeOf(42)
    t2 string = typeOf(3.14)
    t3 string = typeOf("Hola")
    t4 string = typeOf(true)

    println(t1)
    println(t2)
    println(t3)
    println(t4)
}
```

Los valores utilizados en las pruebas fueron:

```text
int
f64
string
bool
```

---

## 31. Structs

V-Lang Cherry implementa estructuras con propiedades primitivas.

## 31.1 Declaración de un `struct`

```v
struct Persona {
    string nombre
    int edad
    f64 estatura
    bool activo
}
```

También pueden utilizarse `;`:

```v
struct Persona {
    string nombre;
    int edad;
}
```

---

## 31.2 Instanciación

```v
fn main() {
    persona = Persona {
        nombre: "Juan",
        edad: 30,
        estatura: 1.75,
        activo: true
    }
}
```

---

## 31.3 Acceso a atributos

Se utiliza el operador `.`:

```v
println(persona.nombre)
println(persona.edad)
```

También se puede guardar un atributo en una variable:

```v
nombre string = persona.nombre
edad int = persona.edad
```

---

## 31.4 Modificación de atributos

```v
persona.nombre = "María"
persona.edad = 25
persona.estatura = 1.65
persona.activo = true
```

En la versión probada los campos del `struct` son de tipos primitivos.

---

## 32. Ejemplo completo

```v
struct Persona {
    string nombre
    int edad
}

fn sumar(a int, b int) int {
    return a + b
}

fn main() {
    nombre string = "Ana"
    edad int = 20

    numeros = []int {10, 20, 30}

    resultado int = sumar(10, 20)

    persona = Persona {
        nombre: nombre,
        edad: edad
    }

    if resultado == 30 {
        println("Suma correcta")
    }

    for indice, valor in numeros {
        println(indice, valor)
    }

    numeros = append(numeros, 40)

    posicion int = indexOf(numeros, 30)

    println("Nombre:", persona.nombre)
    println("Posición:", posicion)
}
```

---

## 33. Resumen de características del intérprete

| Característica | Estado |
|---|---|
| `int` | ✅ |
| `f64` | ✅ |
| `string` | ✅ |
| `bool` | ✅ |
| Declaración de variables | ✅ |
| `mut` | ⚠️ Aceptado, pero sin control real de mutabilidad |
| Asignación | ✅ |
| `+=`, `-=` | ✅ |
| `++`, `--` | ✅ |
| Operaciones aritméticas | ✅ |
| Operaciones relacionales | ✅ |
| Operaciones lógicas | ✅ |
| `println` | ✅ |
| `if` | ✅ |
| `else if` | ✅ |
| `else` | ✅ |
| `for` tipo while | ✅ |
| `for` clásico | ✅ |
| `for indice, valor in slice` | ✅ |
| `switch/case/default` | ✅ |
| `break` | ✅ |
| `continue` | ✅ |
| Funciones | ✅ |
| Funciones con parámetros | ✅ |
| Funciones recursivas | ✅ |
| `return` | ✅ |
| Slices de una dimensión | ✅ |
| Acceso por índice | ✅ |
| Modificación por índice | ✅ |
| `indexOf` | ✅ |
| `join` | ✅ |
| `append` de un elemento | ✅ |
| `len` | ❌ No considerado implementado |
| Arrays/slices bidimensionales | ❌ No implementados |
| Structs con campos primitivos | ✅ |
| Acceso y modificación de atributos | ✅ |
| `Atoi` | ✅ |
| `parseFloat` | ✅ según gramática/implementación |
| `typeOf` | ✅ |
| `null` / `nil` | ❌ No existe |
| Punto y coma obligatorio | ❌ Es opcional |

---

## 34. Recomendación

Para evitar confusiones, se recomienda escribir los programas con esta estructura:

```v
// 1. Structs

struct Ejemplo {
    string nombre
}


// 2. Funciones auxiliares

fn operacion(a int, b int) int {
    return a + b
}


// 3. Punto de entrada

fn main() {

    // Declaraciones

    numero int = 10
    texto string = "Hola"

    // Lógica del programa

    if numero > 5 {
        println(texto)
    }
}
```

De esta manera se diferencia claramente entre:

```text
Definición del programa
        │
        ├── structs
        ├── funciones
        │
        └── fn main()
               │
               └── instrucciones que se ejecutan
```

---

## Nota final

Esta documentación describe **lo que se utilizó y probó en el intérprete basado en Visitor**.

No debe asumirse que todas estas construcciones están disponibles en el **traductor a ARM64**. La sintaxis y las características efectivamente traducibles a ARM64 deben documentarse en un archivo independiente.
