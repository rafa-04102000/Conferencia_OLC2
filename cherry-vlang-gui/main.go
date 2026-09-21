package main

import (
	"bytes"
	"encoding/json"
	"fmt"
	"image/color"
	"io/ioutil"
	"net/http"
	"os"
	"os/exec"
	"path/filepath"
	"runtime"

	"fyne.io/fyne/v2"
	"fyne.io/fyne/v2/app"
	"fyne.io/fyne/v2/container"
	"fyne.io/fyne/v2/dialog"
	"fyne.io/fyne/v2/storage"
	"fyne.io/fyne/v2/theme"
	"fyne.io/fyne/v2/widget"
)

func main() {
	myApp := app.New()
	myApp.Settings().SetTheme(&myTheme{Theme: theme.DefaultTheme()})

	myWindow := myApp.NewWindow("Conferencia Compiladores 2 - Vlang Cherry Editor")

	// Área de consola (no editable)
	consola := widget.NewMultiLineEntry()
	consola.SetText("Bienvenido.\n")
	consola.Disable()
	consola.SetMinRowsVisible(10)

	// contador de pestañas
	contador := 1

	// Contenedor de pestañas
	tabs := container.NewAppTabs()

	// Crear primer archivo por defecto
	crearArchivo(tabs, &contador)

	// Botón ejecutar
	botonEjecutar := widget.NewButton("Ejecutar", func() {
		reqBody, _ := json.Marshal(map[string]string{
			"code": tabs.Selected().Content.(*widget.Entry).Text,
		})

		// Enviamos al backend
		resp, err := http.Post(
			"http://localhost:8080/interprete",
			// "http://localhost:8080/generar-arm64",
			// "http://localhost:8080/interprete_y_generar-arm64",
			"application/json",
			bytes.NewBuffer(reqBody),
		)
		if err != nil {
			consola.SetText(consola.Text + "\n Error conectando al backend \n")
			return
		}
		defer resp.Body.Close()

		// Procesamos la respuesta
		body, _ := ioutil.ReadAll(resp.Body)
		var result struct {
			Output string `json:"output"`
			Error  bool   `json:"error"`
		}
		json.Unmarshal(body, &result)

		// Mostramos resultados
		// limpiar consola
		consola.SetText("") // Limpiamos la consola antes de mostrar el resultado
		consola.SetText(consola.Text + result.Output)
		// if result.Error {
		// 	consola.SetText(consola.Text + "\nEjecución completada con errores\n")
		// } else {
		// 	consola.SetText(consola.Text + "\nEjecución exitosa\n")
		// }
	})
	// botonCentrado := container.NewCenter(botonEjecutar)

	// Menú "Archivo"
	archivoMenu := fyne.NewMenu("Archivo",
		fyne.NewMenuItem("Crear archivo", func() {
			crearArchivo(tabs, &contador)
		}),
		fyne.NewMenuItem("Abrir archivo", func() {
			abrirArchivo(myWindow, tabs, consola)
		}),
		fyne.NewMenuItem("Guardar archivo", func() {
			guardarArchivo(myWindow, tabs)
		}),
		fyne.NewMenuItem("Cerrar archivo", func() {
			if len(tabs.Items) > 0 {
				tabs.Remove(tabs.Selected())
			}
		}),
	)

	// Menú "Reportes"
	reportesMenu := fyne.NewMenu("Reportes",
		fyne.NewMenuItem("Reporte de Errores", func() {
			consola.SetText(consola.Text + "\n[Reporte de errores]\n")
			mostrarHtml("../Backend/Reports/tabla_errores.html", consola)
		}),
		fyne.NewMenuItem("Tabla de símbolos", func() {
			consola.SetText(consola.Text + "\n[Reporte de tabla de símbolos]\n")
			mostrarHtml("../Backend/Reports/tabla_simbolos.html", consola)
		}),
		fyne.NewMenuItem("CST", func() {
			consola.SetText(consola.Text + "\n[Reporte de CST]\n")
			generateReportCST(tabs, consola)
			VerArbol("../Backend/arbol.svg", consola)
		}),
	)

	myWindow.SetMainMenu(fyne.NewMainMenu(archivoMenu, reportesMenu))

	// Título del área de código
	tituloEditor := widget.NewLabelWithStyle(
		"Código fuente",
		fyne.TextAlignLeading,
		fyne.TextStyle{Bold: true},
	)

	// Panel izquierdo: editor
	panelEditor := container.NewBorder(
		tituloEditor, // superior
		nil,          // inferior
		nil,          // izquierda
		nil,          // derecha
		tabs,         // centro
	)

	// Título del área de salida
	tituloSalida := widget.NewLabelWithStyle(
		"Salida ARM64",
		fyne.TextAlignLeading,
		fyne.TextStyle{Bold: true},
	)

	// Panel derecho: consola de salida
	panelSalida := container.NewBorder(
		tituloSalida, // superior
		nil,          // inferior
		nil,          // izquierda
		nil,          // derecha
		consola,      // centro
	)

	// División horizontal:
	// editor a la izquierda y salida a la derecha
	divisionPrincipal := container.NewHSplit(
		panelEditor,
		panelSalida,
	)

	// El editor ocupará 50 % y la salida 50 %.
	divisionPrincipal.SetOffset(0.50)

	// Botón centrado en la parte inferior
	botonCentrado := container.NewCenter(
		botonEjecutar,
	)

	// Diseño principal
	content := container.NewBorder(
		nil,               // superior
		botonCentrado,     // inferior
		nil,               // izquierda
		nil,               // derecha
		divisionPrincipal, // centro
	)
	myWindow.SetContent(content)

	// Ventana más ancha para mostrar ambos paneles
	myWindow.Resize(fyne.NewSize(1500, 700))

	// Permite maximizar y cambiar el tamaño de la ventana
	myWindow.SetFixedSize(false)

	myWindow.ShowAndRun()
}

// Funcion para Crear nueva pestania
func crearArchivo(tabs *container.AppTabs, contador *int) {
	nombre := fmt.Sprintf("Archivo %d", *contador)
	*contador++
	editor := widget.NewMultiLineEntry()
	editor.Wrapping = fyne.TextWrapWord
	nueva := container.NewScroll(editor)
	tab := container.NewTabItem(nombre, nueva)
	tab.Content = editor
	tabs.Append(tab)
	tabs.Select(tab)
}

// Funcion para Abrir un archivo con terminacion .vch
func abrirArchivo(win fyne.Window, tabs *container.AppTabs, consola *widget.Entry) {
	filter := storage.NewExtensionFileFilter([]string{".vch"})
	openDialog := dialog.NewFileOpen(func(archivo fyne.URIReadCloser, err error) {
		if err != nil || archivo == nil {
			return
		}
		defer archivo.Close()

		contenido, err := ioutil.ReadAll(archivo)
		if err != nil {
			consola.SetText(consola.Text + "\nError al leer archivo\n")
			return
		}

		editor := widget.NewMultiLineEntry()
		editor.SetText(string(contenido))

		tab := container.NewTabItem(archivo.URI().Name(), editor)
		tabs.Append(tab)
		tabs.Select(tab)
		consola.SetText(consola.Text + "\nArchivo abierto: " + archivo.URI().Name() + "\n")
	}, win)
	openDialog.SetFilter(filter)
	openDialog.Show()
}

// Funcion para guardar archivo con terminacion .vch
func guardarArchivo(win fyne.Window, tabs *container.AppTabs) {
	if len(tabs.Items) == 0 {
		return
	}

	contenido := tabs.Selected().Content.(*widget.Entry).Text

	saveDialog := dialog.NewFileSave(func(archivo fyne.URIWriteCloser, err error) {
		if err != nil || archivo == nil {
			return
		}
		defer archivo.Close()
		archivo.Write([]byte(contenido))
	}, win)
	saveDialog.SetFileName("archivo.vch")
	saveDialog.Show()

}

// Funcion que llamara al navegador para mostrar el arbol
func VerArbol(ruta string, consola *widget.Entry) {
	rutaAbsoluta, err := filepath.Abs(ruta)
	if err != nil {
		consola.SetText(consola.Text + "\nError obteniendo la ruta del archivo.\n")
		return
	}

	if _, err := os.Stat(rutaAbsoluta); os.IsNotExist(err) {
		consola.SetText(consola.Text + "\nEl archivo no existe: " + rutaAbsoluta + "\n")
		return
	}

	var cmd *exec.Cmd
	switch runtime.GOOS {
	case "linux":
		cmd = exec.Command("firefox", rutaAbsoluta)
	case "windows":
		cmd = exec.Command("rundll32", "url.dll,FileProtocolHandler", rutaAbsoluta)
	case "darwin": // macOS
		cmd = exec.Command("open", rutaAbsoluta)
	default:
		fmt.Println("Sistema operativo no soportado")
		return
	}

	err = cmd.Start()
	if err != nil {
		consola.SetText(consola.Text + "\nError al abrir el navegador\n")
	} else {
		consola.SetText(consola.Text + "\nAbriendo reporte CST...\n")
	}
}

type myTheme struct {
	fyne.Theme // Embed el tema por defecto
}

// Sobrescribe el color para texto deshabilitado
func (m myTheme) Color(name fyne.ThemeColorName, variant fyne.ThemeVariant) color.Color {
	if name == theme.ColorNameDisabled {
		return color.NRGBA{R: 200, G: 200, B: 200, A: 255} // Gris claro sin transparencia
	}
	return m.Theme.Color(name, variant)
}

func generateReportCST(tabs *container.AppTabs, consola *widget.Entry) {
	reqBody, _ := json.Marshal(map[string]string{
		"code": tabs.Selected().Content.(*widget.Entry).Text,
	})
	//verificamos que el body se haya creado correctamente
	fmt.Println(reqBody)
	// Enviamos al backend
	resp, err := http.Post(
		"http://localhost:8080/cst",
		"application/json",
		bytes.NewBuffer(reqBody),
	)
	if err != nil {
		consola.SetText(consola.Text + "\n Error conectando al backend \n")
		return
	}
	defer resp.Body.Close()

	// Procesamos la respuesta
	body, _ := ioutil.ReadAll(resp.Body)
	var result struct {
		Output string `json:"output"`
		Error  bool   `json:"error"`
	}
	json.Unmarshal(body, &result)

	// Mostramos resultados
	consola.SetText(consola.Text + result.Output)
}

// Funcion para mostrar la tabla de simbolos
func mostrarHtml(ruta string, consola *widget.Entry) {
	rutaAbsoluta, err := filepath.Abs(ruta)
	if err != nil {
		consola.SetText(consola.Text + "\nError obteniendo la ruta del archivo.\n")
		return
	}

	if _, err := os.Stat(rutaAbsoluta); os.IsNotExist(err) {
		consola.SetText(consola.Text + "\nEl archivo no existe: " + rutaAbsoluta + "\n")
		return
	}

	var cmd *exec.Cmd
	switch runtime.GOOS {
	case "linux":
		cmd = exec.Command("firefox", rutaAbsoluta)
	default:
		fmt.Println("Sistema operativo no soportado")
		return
	}

	err = cmd.Start()
	if err != nil {
		consola.SetText(consola.Text + "\nError al abrir el navegador\n")
	} else {
		consola.SetText(consola.Text + "\nAbriendo en el Navegador ...\n")
	}
}
