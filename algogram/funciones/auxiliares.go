package auxiliares

import (
	red "algogram/redUsuario/red"
	usuario "algogram/redUsuario/usuario"
	"bufio"
	"os"
	"strings"
)

// agregar usuarios a la red
func AgregarUsuarios(ruta string) (red.Red, error) {
	archivo, err := os.Open(ruta)
	if err != nil {
		return nil, err
	}
	defer archivo.Close()

	// leer archivo
	red := red.CrearRed()

	usuarios := bufio.NewScanner(archivo)
	posicion := 0

	for usuarios.Scan() {
		registro := red.Registrados()
		usuario := usuario.CrearUsuario(usuarios.Text(), posicion, &registro)
		red.Registrados().Guardar(usuarios.Text(), usuario)
		posicion++
	}
	return red, nil
}

// leer comando
func LeerComando(entrada string) (string, string) {
	dato := strings.Split(entrada, " ")
	comando := dato[0]
	return comando, strings.Join(dato[1:], " ")
}
