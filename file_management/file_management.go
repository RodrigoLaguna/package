package file_management

import (
	"fmt"
	"io/ioutil"
	"strings"
)
//Funcion que lee un archivo y devuelve su contenido en arreglo de strings
func Read(archivo string) []string {
	var contenido []string
	//El contenido del archivo se convierte a un arreglo de bytes
	bytesLeidos, err := ioutil.ReadFile(archivo)
	if err != nil {
		fmt.Printf("Error leyendo archivo: %v", err)
	}
	//Convierte los bytes leidos en string
	texto_contenido := string(bytesLeidos)
	//Guardamos cada string en una posicion del arreglo
	contenido = strings.Fields(texto_contenido)
	return contenido
}
