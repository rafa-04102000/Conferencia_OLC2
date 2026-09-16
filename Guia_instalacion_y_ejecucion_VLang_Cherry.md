# Guía de instalación y ejecución — V-Lang Cherry en Ubuntu

Esta guía resume el orden recomendado para preparar el entorno, generar el parser y lexer con ANTLR, ejecutar el backend, iniciar la interfaz gráfica con Fyne y, cuando corresponda, ensamblar y ejecutar el código ARM64 generado.

> Basado en el `Manual_tecnico.md` del proyecto y en el flujo de trabajo mostrado para Ubuntu.

---

## 1. Requisitos principales

El proyecto utiliza:

- Ubuntu 22.04 o superior.
- Go 1.22 o superior.
- Java, necesario para ejecutar ANTLR.
- ANTLR4 para generar lexer, parser, listener y visitor en Go.
- Fyne v2 para la interfaz gráfica.
- Herramientas nativas de compilación requeridas por Fyne.
- Binutils para AArch64/ARM64.
- QEMU para ejecutar binarios ARM64 cuando el equipo anfitrión no es ARM64.

---

# PARTE I — INSTALACIÓN

## 2. Actualizar Ubuntu

```bash
sudo apt update
sudo apt upgrade -y
```

---

## 3. Verificar Go

Antes de instalar nada, comprobar si Go ya está disponible:

```bash
go version
```

Debe aparecer una versión igual o superior a Go 1.22.

También puede verificarse su ubicación:

```bash
which go
```

---

## 4. Instalar dependencias del sistema

```bash
sudo apt install -y \
    openjdk-17-jre \
    gcc \
    git \
    pkg-config \
    libgl1-mesa-dev \
    xorg-dev \
    binutils-aarch64-linux-gnu \
    qemu-user \
    wget
```

### ¿Para qué sirve cada paquete?

- `openjdk-17-jre`: ejecuta el archivo `.jar` de ANTLR.
- `gcc`: requerido por Fyne y algunas dependencias gráficas.
- `git`: control de versiones.
- `pkg-config`: ayuda a Go/CGo a localizar librerías del sistema.
- `libgl1-mesa-dev`: soporte OpenGL para Fyne.
- `xorg-dev`: librerías gráficas X11 usadas por Fyne en Linux.
- `binutils-aarch64-linux-gnu`: proporciona `aarch64-linux-gnu-as` y `aarch64-linux-gnu-ld`.
- `qemu-user`: incluye `qemu-aarch64`, usado para ejecutar binarios ARM64.
- `wget`: facilita la descarga de ANTLR.

---

## 5. Verificar Java

```bash
java -version
```

ANTLR necesita Java porque ANTLR se distribuye como un archivo `.jar`.

---

## 6. Descargar ANTLR4

Crear una carpeta para almacenar ANTLR:

```bash
mkdir -p "$HOME/Lib"
```

Descargar ANTLR 4.13.2:

```bash
wget -O "$HOME/Lib/antlr-4.13.2-complete.jar" \
https://www.antlr.org/download/antlr-4.13.2-complete.jar
```

---

## 7. Configurar ANTLR como comando

Si se utiliza Zsh:

```bash
nano ~/.zshrc
```

Agregar:

```bash
export CLASSPATH=".:$HOME/Lib/antlr-4.13.2-complete.jar:$CLASSPATH"
alias antlr4='java -jar $HOME/Lib/antlr-4.13.2-complete.jar'
```

Guardar con:

```text
Ctrl + O
Enter
Ctrl + X
```

Recargar la configuración:

```bash
source ~/.zshrc
```

Si se utiliza Bash, realizar el mismo procedimiento en:

```bash
nano ~/.bashrc
```

y luego:

```bash
source ~/.bashrc
```

Verificar ANTLR:

```bash
antlr4
```

Si muestra la ayuda de ANTLR, la configuración funciona correctamente.

---

## 8. Descargar las dependencias Go del backend

Entrar al backend:

```bash
cd Backend
```

Descargar las dependencias declaradas en `go.mod`:

```bash
go mod download
```

También puede usarse:

```bash
go mod tidy
```

`go mod tidy` además revisa y ajusta las dependencias realmente utilizadas por el proyecto.

---

## 9. Descargar las dependencias Go del frontend

Regresar a la carpeta raíz y entrar a la interfaz gráfica:

```bash
cd ../cherry-vlang-gui
```

Ejecutar:

```bash
go mod download
```

o:

```bash
go mod tidy
```

Fyne no se instala normalmente mediante `apt`; el módulo de Fyne queda administrado por Go mediante el archivo `go.mod`.

---

# PARTE II — GENERACIÓN DEL PARSER Y LEXER

## 10. Archivo de gramática

La gramática del lenguaje se encuentra en:

```text
Backend/Gramatica.g4
```

ANTLR lee este archivo y genera automáticamente el código Go necesario para realizar el análisis léxico y sintáctico.

El flujo conceptual es:

```text
Gramatica.g4
     │
     ▼
   ANTLR4
     │
     ├── Lexer
     ├── Parser
     ├── Listener
     └── Visitor
           │
           ▼
     código Go generado
           │
           ▼
     AST / Intérprete
```

---

## 11. Generar parser, lexer, listener y visitor

Desde la carpeta `Backend`:

```bash
cd Backend
```

Ejecutar:

```bash
antlr4 -Dlanguage=Go -visitor -listener -o parser Gramatica.g4
```

### Significado del comando

- `-Dlanguage=Go`: genera el código en Go.
- `-visitor`: genera las interfaces y clases del patrón Visitor.
- `-listener`: genera las interfaces Listener.
- `-o parser`: guarda los archivos generados dentro de la carpeta `parser`.
- `Gramatica.g4`: archivo fuente de la gramática.

Después del comando, la carpeta `parser/` contendrá varios archivos `.go` generados por ANTLR, incluyendo el lexer y parser.

---

## 12. ¿Cuándo debo volver a ejecutar ANTLR?

No es necesario ejecutar ANTLR cada vez que se inicia el proyecto.

Debe regenerarse el parser cuando se modifica:

```text
Gramatica.g4
```

Flujo:

```text
Modificar Gramatica.g4
        │
        ▼
Ejecutar ANTLR
        │
        ▼
Se actualiza parser/
        │
        ▼
Ejecutar backend
```

Si `Gramatica.g4` no cambió y la carpeta `parser/` ya contiene los archivos generados, puede pasar directamente a ejecutar el backend.

---

# PARTE III — EJECUCIÓN DEL PROYECTO

## 13. Orden recomendado

El orden normal de trabajo es:

```text
1. Generar parser con ANTLR
   solamente si cambió Gramatica.g4

2. Iniciar Backend

3. Iniciar Frontend / GUI

4. Escribir o cargar código V-Lang Cherry

5. Ejecutar el código desde la interfaz

6. El backend analiza el código:
      Lexer
        ↓
      Parser
        ↓
      Árbol sintáctico / CST (no AST)
        ↓
      Visitor
        ↓
      Intérprete o generador ARM64
```

---

## 14. Iniciar el backend

Abrir una terminal.

Entrar al backend:

```bash
cd Backend
```

Ejecutar:

```bash
go run main.go
```

El backend utiliza Go y el código generado por ANTLR.

Mantener esta terminal abierta mientras se utiliza la aplicación.

---

## 15. Iniciar el frontend

Abrir una segunda terminal.

Entrar al proyecto de la interfaz:

```bash
cd cherry-vlang-gui
```

Ejecutar:

```bash
go run main.go
```

Esto inicia la interfaz gráfica desarrollada con Fyne.

---

# PARTE IV — FLUJO COMPLETO DEL PROYECTO

## 16. Qué ocurre al ejecutar código

De forma simplificada:

```text
Código fuente V-Lang Cherry
          │
          ▼
        Lexer
          │
          ▼
        Parser
          │
          ▼
         CST
          │
          ▼
    Visitor / lógica Go
          │
          ▼
        (AST se podria generar uno Simpleficando el CST) 
          │
          ├───────────────┐
          ▼               ▼
     Intérprete      Generación ARM64
                          │
                          ▼
                       ARM64.s
```

La gramática define qué construcciones del lenguaje son válidas. ANTLR genera las herramientas de análisis, mientras que la lógica del proyecto en Go recorre las producciones mediante el patrón Visitor.

---

# PARTE V — ARM64

## 17. Generar código ARM64

Cuando el backend genera código ensamblador, el resultado puede ser un archivo como:

```text
ARM64.s
```

Este archivo todavía es texto ensamblador. No es un programa ejecutable.

El flujo es:

```text
ARM64.s
   │
   ▼
Assembler
   │
   ▼
ARM64.o
   │
   ▼
Linker
   │
   ▼
ARM64.elf
   │
   ▼
QEMU
```

---

## 18. Ensamblar ARM64

Desde la carpeta donde se encuentra `ARM64.s`:

```bash
aarch64-linux-gnu-as -o ARM64.o ARM64.s
```

Esto convierte el código ensamblador a un archivo objeto.

---

## 19. Enlazar el archivo objeto

```bash
aarch64-linux-gnu-ld -o ARM64.elf ARM64.o
```

Esto genera el ejecutable ARM64:

```text
ARM64.elf
```

---

## 20. Ejecutar con QEMU

```bash
qemu-aarch64 ./ARM64.elf
```

QEMU no ensambla ni enlaza el código.

Su función en este flujo es ejecutar un binario ARM64 en un equipo cuya arquitectura puede ser x86-64.

---

# PARTE VI — VERIFICACIÓN RÁPIDA

## 21. Comprobar todas las herramientas

```bash
go version
java -version
antlr4
aarch64-linux-gnu-as --version
aarch64-linux-gnu-ld --version
qemu-aarch64 --version
```

Si todos los comandos responden correctamente, el entorno principal está preparado.

---

# PARTE VII — RESUMEN DE COMANDOS

## Primera instalación

```bash
sudo apt update

sudo apt install -y \
    openjdk-17-jre \
    gcc \
    git \
    pkg-config \
    libgl1-mesa-dev \
    xorg-dev \
    binutils-aarch64-linux-gnu \
    qemu-user \
    wget

mkdir -p "$HOME/Lib"

wget -O "$HOME/Lib/antlr-4.13.2-complete.jar" \
https://www.antlr.org/download/antlr-4.13.2-complete.jar
```

Configurar el alias de ANTLR y después:

```bash
cd Backend
go mod download

cd ../cherry-vlang-gui
go mod download
```

---

## Cuando se modifica la gramática

```bash
cd Backend

antlr4 -Dlanguage=Go -visitor -listener -o parser Gramatica.g4
```

---

## Para ejecutar normalmente el proyecto

### Terminal 1 — Backend

```bash
cd Backend
go run main.go
```

### Terminal 2 — Frontend

```bash
cd cherry-vlang-gui
go run main.go
```

---

## Para ejecutar manualmente un archivo ARM64 generado

```bash
cd Backend

aarch64-linux-gnu-as -o ARM64.o ARM64.s

aarch64-linux-gnu-ld -o ARM64.elf ARM64.o

qemu-aarch64 ./ARM64.elf
```

---

# 22. Orden mental recomendado

La idea más importante es distinguir entre **instalación**, **generación** y **ejecución**:

```text
INSTALAR
Ubuntu + Go + Java + ANTLR + dependencias Fyne + binutils + QEMU
        │
        ▼
GENERAR
Gramatica.g4 ──ANTLR──> parser/lexer/visitor/listener
        │
        ▼
EJECUTAR
Backend Go
        │
        ▼
Frontend Fyne
        │
        ▼
Código V-Lang Cherry
        │
        ▼
Lexer → Parser → AST/Visitor
        │
        ▼
ARM64.s
        │
        ▼
Assembler → Linker → QEMU
```

Así, **ANTLR se utiliza para preparar el analizador del lenguaje**, mientras que **Go ejecuta el backend y la interfaz**, y **QEMU solamente entra en juego cuando ya existe un binario ARM64 que se desea ejecutar**.
