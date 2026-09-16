// =========================================================
// EJEMPLO BÁSICO ARM64 / AArch64
// =========================================================

// -------------------------
// DATOS INICIALIZADOS
// -------------------------
.data

numero:
    .word 42

mensaje:
    .asciz "Hola desde ARM64\n"

mensaje_len = . - mensaje


// -------------------------
// MEMORIA RESERVADA
// -------------------------
.bss

buffer:
    .skip 64


// -------------------------
// CÓDIGO EJECUTABLE
// -------------------------
.text
.global _start


// -------------------------
// PUNTO DE ENTRADA
// -------------------------
_start:

    // Cargar dirección de numero
    ldr x0, =numero

    // Cargar su contenido
    ldr w1, [x0]

    // Llamar una función
    bl ejemplo_stack

    // Imprimir mensaje
    mov x0, #1
    ldr x1, =mensaje
    mov x2, #mensaje_len
    mov x8, #64
    svc #0

    // Finalizar programa
    mov x0, #0
    mov x8, #93
    svc #0


// -------------------------
// EJEMPLO DE STACK FRAME
// -------------------------
ejemplo_stack:

    // Reservar stack frame y preservar FP/LR
    stp x29, x30, [sp, #-16]!
    mov x29, sp

    // Ejemplo de operación local
    mov x2, #10
    add x3, x2, #5

    // Restaurar stack
    ldp x29, x30, [sp], #16
    ret