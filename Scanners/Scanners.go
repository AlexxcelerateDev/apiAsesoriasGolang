package Scanners

import (
	"apiAsesoria/Struct"
	"database/sql"
)

func ScanAsesor(rows *sql.Rows) (interface{}, error) {
	var a Struct.Asesor
	direcciones := obtenerDireccionesCampos(&a)

	if err := rows.Scan(direcciones...); err != nil {
		return nil, err
	}
	return a, nil
}

func ScanAlumno(rows *sql.Rows) (interface{}, error) {
	var a Struct.Alumno
	direcciones := obtenerDireccionesCampos(&a)
	if err := rows.Scan(direcciones...); err != nil {
		return nil, err
	}
	return a, nil
}

func ScanAsesoriaNoTerminada(rows *sql.Rows) (interface{}, error) {
	var a Struct.AsesoriaNoTerminadas
	direcciones := obtenerDireccionesCampos(&a)
	if err := rows.Scan(direcciones...); err != nil {
		return nil, err
	}
	return a, nil
}

func ScanProfesor(rows *sql.Rows) (interface{}, error) {
	var a Struct.Profesor
	direcciones := obtenerDireccionesCampos(&a)
	if err := rows.Scan(direcciones...); err != nil {
		return nil, err
	}
	return a, nil
}

func ScanUA(rows *sql.Rows) (interface{}, error) {
	var a Struct.UnidadAprendizaje
	direcciones := obtenerDireccionesCampos(&a)
	if err := rows.Scan(direcciones...); err != nil {
		return nil, err
	}
	return a, nil
}
