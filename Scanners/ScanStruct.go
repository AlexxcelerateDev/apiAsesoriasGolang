package Scanners

import (
	"reflect"
)

func obtenerDireccionesCampos(estructura interface{}) []interface{} {
	val := reflect.ValueOf(estructura)

	if val.Kind() != reflect.Ptr || val.Elem().Kind() != reflect.Struct {
		panic("Se esperaba un puntero a una estructura como entrada")
	}

	val = val.Elem() // Desreferenciamos para obtener el valor real
	numCampos := val.NumField()
	direcciones := make([]interface{}, numCampos)

	for i := 0; i < numCampos; i++ {
		field := val.Field(i)
		direcciones[i] = field.Addr().Interface()
	}

	return direcciones
}
