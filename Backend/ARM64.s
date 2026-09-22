// ------------------------------------------------------------------------------
// 	| VARIABLES GLOBALES DATA Y BSS |

.data
newline: .asciz "\n"
flt_100: .double 100.0
str_true: .asciz "true"
str_false: .asciz "false"
numero: .word 20
tmp_1: .word 0
condicion: .word 0
tmp_2: .asciz "Numero mayor que 10"
tmp_3: .asciz "Fin"

.bss
salida_1: .skip 128
salida_2: .skip 128

// ------------------------------------------------------------------------------
// 	| CODIGO |

.text
.global _start

_start:
	// ---------- Declaración de variable: numero ----------


	// ---------- Variable numero de tipo int declarada correctamente ----------

	// ---------- Declaración de variable: condicion ----------

	ldr x1, =numero
	ldr w1, [x1]
	mov w2, #10
	cmp w1, w2
	cset w3, gt
	ldr x2, =tmp_1
	str w3, [x2]

	ldr x0, =condicion
	ldr x1, =tmp_1
	ldr w2, [x1]
	str w2, [x0]


	// ---------- Variable condicion de tipo bool declarada correctamente ----------

	ldr x1, =condicion
	ldr w1, [x1]
	cmp w1, #0
	b.eq label_1
	// ---------- Print para salida_1 ----------


	// --- Inicio Concatenacion cadena ---

	ldr x0, =salida_1	// destino
	mov x20, x0	// guardar puntero base
	ldr x1, =tmp_2	// origen
	mov x4, #128	// Tamaño del buffer destino
	bl copiar_sin_limpiar
	mov x0, x20	// restaurar puntero base
	bl advance_to_end
	ldr x1, =newline
	bl copiar_sin_limpiar
	mov x1, x20	// restaurar puntero base en x1
	bl println

	// --- Fin Concatenacion de cadena ---

	// ---------- Fin Print para salida_1 ----------

	b label_1
label_1:
	// ---------- Print para salida_2 ----------


	// --- Inicio Concatenacion cadena ---

	ldr x0, =salida_2	// destino
	mov x20, x0	// guardar puntero base
	ldr x1, =tmp_3	// origen
	mov x4, #128	// Tamaño del buffer destino
	bl copiar_sin_limpiar
	mov x0, x20	// restaurar puntero base
	bl advance_to_end
	ldr x1, =newline
	bl copiar_sin_limpiar
	mov x1, x20	// restaurar puntero base en x1
	bl println

	// --- Fin Concatenacion de cadena ---

	// ---------- Fin Print para salida_2 ----------


	// Salida del programa
	mov x0, #0	// código de salida 0
	mov x8, #93	// syscall exit
	svc #0

// ------------------------------------------------------------------------------
// 	| FUNCIONES |


//---- Función para copiar strings ----
copiar_string:
	mov x3, x0          // Guardar puntero original del destino
	mov x9, x4          // Copiar tamaño para limpieza

clean_loop:
	cbz x9, copy_start
	mov w10, #0
	strb w10, [x0], #1
	subs x9, x9, #1
	b clean_loop

copy_start:
	mov x0, x3          // Restaurar puntero destino
	mov x5, x4          // Tamaño máximo para copia

copy_loop:
	ldrb w6, [x1], #1   // Cargar byte del origen
	cbz w6, done_copy   // Terminar si es null
	strb w6, [x0], #1   // Escribir byte en destino
	subs x5, x5, #1     // Decrementar contador
	b.gt copy_loop

force_null:
	mov w6, #0
	strb w6, [x0]
done_copy:
	ret


//---- Fin Función para copiar strings ----

//---- Función para convertir int a string ----
int_to_str:
	mov x11, x1       // puntero de escritura
	mov x3, #10      // base decimal
	mov x10, #0       // contador de dígitos

reverse_digits:
	udiv w5, w0, w3  // w5 = w0 / 10
	msub w6, w5, w3, w0  // w6 = w0 - w5*10 → w6 = w0 % 10
	add w6, w6, #48     // convertir a ASCII
	strb w6, [x11], #1
	mov w0, w5
	add x10, x10, #1
	cbnz w0, reverse_digits

	mov w6, #0
	strb w6, [x11]

	// Invertir string (porque lo construimos al revés)
	sub x11, x11, #1
	mov x5, x1
reverse_loop:
	cmp x5, x11
	bge reverse_done
	ldrb w6, [x5]
	ldrb w7, [x11]
	strb w7, [x5]
	strb w6, [x11]
	add x5, x5, #1
	sub x11, x11, #1
	b reverse_loop
reverse_done:
	ret


//---- Fin Función para convertir int a string ----

//---- Función para convertir float a string ----
float_to_str:
	// Preservar registros que vamos a modificar
	stp x29, x30, [sp, #-16]!  // Guardar frame pointer y return address
	mov x29, sp

	mov x4, x1            // Guardar puntero destino original

	// Parte entera: truncar float
	fcvtzs w0, d0         // Convertir a entero (parte entera)
	bl int_to_str         // Convertir parte entera a string

	// Avanzar al final del string para agregar punto decimal
	mov x5, x4            // x5 = puntero que avanza
find_end:
	ldrb w6, [x5]
	cbz w6, add_dot
	add x5, x5, #1
	b find_end

add_dot:
	mov w6, #'.'
	strb w6, [x5], #1     // Agregar punto decimal

	// Calcular parte decimal: (valor - parte_entera) * 100
	scvtf d1, w0          // Convertir parte entera a float
	fsub d2, d0, d1       // Obtener parte decimal
	ldr x0, =flt_100      // Cargar 100.0
	ldr d3, [x0]
	fmul d2, d2, d3       // Multiplicar por 100 para obtener 2 decimales

	// Convertir decimales a entero
	fcvtzu w7, d2         // Convertir a entero sin signo

	// Manejar redondeo (ej. 33.4999999 → 33.5 → 50)
	// Agregamos 0.5 antes de convertir para redondear correctamente
	fmov d4, #0.5
	fadd d2, d2, d4
	fcvtzu w7, d2

	// Convertir parte decimal a string
	mov x1, x5            // Puntero donde escribir los decimales
	mov w0, w7
	bl int_to_str

	// Asegurar exactamente 2 dígitos
	mov x6, x5            // Puntero al inicio de los decimales
	mov x7, x1            // Puntero al final del string
	sub x7, x7, x6        // Longitud de los decimales

	cmp x7, #1
	bne check_length

	// Solo 1 dígito, insertar '0' delante
	ldrb w8, [x5]
	mov w9, #'0'
	strb w9, [x5]
	strb w8, [x5, #1]
	mov w8, #0
	strb w8, [x5, #2]
	b done_float

check_length:
	cmp x7, #0
	bne done_float

	// Ningún dígito, insertar "00"
	mov w8, #'0'
	strb w8, [x5], #1
	strb w8, [x5], #1
	mov w8, #0
	strb w8, [x5]

done_float:
	ldp x29, x30, [sp], #16  // Restaurar frame pointer y return address
	ret


//---- Fin Función para convertir float a string ----

//---- Función para convertir bool a string ----
bool_to_str:
	// w0 = valor booleano (1=true, 0=false)
	// x1 = dirección destino

	cmp w0, #1
	beq store_true

	// Guardar 'false'
	ldr x2, =str_false
	b copy_string

store_true:
	// Guardar 'true'
	ldr x2, =str_true

copy_string:
	ldrb w3, [x2], #1
	strb w3, [x1], #1
	cbnz w3, copy_string

	ret

//---- Fin Función para convertir bool a string ----

//---- Función para avanzar al final de una cadena ----
advance_to_end:
	ldrb w3, [x0]
	cbz w3, done_advance
	add x0, x0, #1
	b advance_to_end
done_advance:
	ret


//---- Fin Función para avanzar al final de una cadena ----

//---- Función para imprimir en consola ----
println:
	mov x2, #0          // Contador de longitud
	mov x3, x1          // Guardar puntero original

count_loop:
	ldrb w4, [x3, x2]    // Cargar byte actual
	cbz w4, do_print      // Si es nulo, salta a imprimir
	add x2, x2, #1        // Incrementar contador
	b count_loop          // Repetir hasta encontrar nulo

do_print:
	mov x0, #1		// stdout
	mov x8, #64         // syscall write
	svc #0
	ret

//---- Fin Función para imprimir en consola ----

//---- Función para copiar strings sin limpiar ----
copiar_sin_limpiar:
	// x0 = destino
	// x1 = origen
copiar_sin_loop:
	ldrb w6, [x1], #1
	cbz w6, fin_copiar_sin
	strb w6, [x0], #1
	b copiar_sin_loop
fin_copiar_sin:
	ret

//---- Fin Función para copiar strings sin limpiar ----

//---- Función para comparar strings (==) ----
strcmp_eq:
	cmp_loop_eq:
	ldrb w2, [x0], #1
	ldrb w3, [x1], #1
	cmp w2, w3
	b.ne not_equal
	cbz w2, strings_equal
	b cmp_loop_eq
not_equal:
	mov w0, #0
	ret
strings_equal:
	mov w0, #1
	ret


//---- Fin Función para comparar strings (==) ----

//---- Función para comparar strings (!=) ----
strcmp_ne:
	cmp_loop_ne:
	ldrb w2, [x0], #1
	ldrb w3, [x1], #1
	cmp w2, w3
	b.ne str_not_equal
	cbz w2, str_equal
	b cmp_loop_ne
str_not_equal:
	mov w0, #1
	ret
str_equal:
	mov w0, #0
	ret


//---- Fin Función para comparar strings (!=) ----

//---- Función para comparar strings (<) ----
strcmp_lt:
	cmp_loop_lt:
	ldrb w2, [x0], #1
	ldrb w3, [x1], #1
	cmp w2, w3
	b.lo str1_less
	b.hi str1_greater
	cbz w2, strings_equal_lt
	b cmp_loop_lt
str1_less:
	mov w0, #1
	ret
str1_greater:
strings_equal_lt:
	mov w0, #0
	ret


//---- Fin Función para comparar strings (<) ----

//---- Función para comparar strings (<=) ----
strcmp_le:
	mov x2, x0
	mov x3, x1
	bl strcmp_lt
	cmp w0, #1
	beq return_1_le
	mov x0, x2
	mov x1, x3
	bl strcmp_eq
	ret
return_1_le:
	mov w0, #1
	ret


//---- Fin Función para comparar strings (<=) ----

//---- Función para comparar strings (>) ----
strcmp_gt:
	mov x2, x0
	mov x3, x1
	mov x0, x1
	mov x1, x2
	bl strcmp_lt
	ret


//---- Fin Función para comparar strings (>) ----

//---- Función para comparar strings (>=) ----
strcmp_ge:
	mov x2, x0
	mov x3, x1
	mov x0, x1
	mov x1, x2
	bl strcmp_lt
	cmp w0, #1
	beq return_0_ge
	mov x0, x2
	mov x1, x3
	bl strcmp_eq
	ret
return_0_ge:
	mov w0, #0
	ret


//---- Fin Función para comparar strings (>=) ----
