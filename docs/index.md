---
hide:
  - toc
---

# Compilación a ARM64/AArch64

## Implementación de un compilador con Go y ANTLR

Bienvenido al material de apoyo para la conferencia. Estas guías acompañan el recorrido desde la preparación del entorno y la sintaxis del lenguaje **V-Lang Cherry** hasta la generación, el ensamblado y la ejecución de código **ARM64/AArch64**.

!!! info "Propósito de la documentación"

    Servir como referencia práctica para instalar las herramientas, comprender el lenguaje del proyecto y relacionar las estructuras de alto nivel con instrucciones, registros y memoria en ARM64.

<div class="grid cards" markdown>

-   :material-tools:{ .lg .middle } **Preparación del entorno**

    ---

    Instala Go, Java, ANTLR, Fyne, las herramientas AArch64 y QEMU; luego genera el parser y ejecuta el proyecto.

    [:octicons-arrow-right-24: Abrir guía de instalación](guia-instalacion.md)

-   :material-code-braces:{ .lg .middle } **Sintaxis de V-Lang Cherry**

    ---

    Consulta tipos de datos, variables, expresiones, ciclos, funciones, slices, estructuras y ejemplos del intérprete.

    [:octicons-arrow-right-24: Consultar la sintaxis](sintaxis-vlang-cherry.md)

-   :material-memory:{ .lg .middle } **Fundamentos de ARM64**

    ---

    Estudia las secciones de memoria, el stack, el heap, los registros, las instrucciones y las llamadas al sistema.

    [:octicons-arrow-right-24: Estudiar ARM64](arm64-memoria-registros.md)

</div>

## Ruta de aprendizaje sugerida

1. **Prepara el entorno:** verifica e instala las herramientas necesarias.
2. **Comprende el lenguaje fuente:** revisa qué construcciones reconoce V-Lang Cherry.
3. **Relaciona fuente y destino:** identifica cómo las estructuras del lenguaje se traducen a ARM64.
4. **Prueba el resultado:** ensambla, enlaza y ejecuta el código generado con QEMU.

!!! tip "Para aprovechar la conferencia"

    Ten Ubuntu preparado, ejecuta los ejemplos de forma incremental y compara siempre el código fuente con el ensamblador generado. No es necesario memorizar todas las instrucciones: interesa comprender el proceso de traducción.

## Tecnologías utilizadas

| Tecnología | Función en el proyecto |
|---|---|
| **Go** | Implementación del analizador, intérprete y traductor |
| **ANTLR** | Generación del lexer y parser a partir de la gramática |
| **Fyne** | Interfaz gráfica del proyecto |
| **ARM64/AArch64** | Arquitectura y lenguaje ensamblador de destino |
| **GNU Binutils** | Ensamblado y enlazado del código ARM64 |
| **QEMU** | Ejecución del binario ARM64 en un equipo no ARM64 |

!!! note "ARM64 y AArch64"

    En esta documentación ambos nombres se refieren a la arquitectura ARM de 64 bits. **AArch64** es el nombre oficial del estado de ejecución; **ARM64** es el nombre de uso común.
