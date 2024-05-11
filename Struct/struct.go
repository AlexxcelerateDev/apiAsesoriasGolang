package Struct

type Asesor struct {
	Nombre    string
	Matricula int64
	Carrera   string
	Activo    bool
}

type Alumno struct {
	Nombre    string
	Carrera   string
	Sexo      string
	Matricula int64
}

type AsesoriaNoTerminadas struct {
	Nombre               string
	Sexo                 string
	Matricula            string
	Carrera              string
	Fecha                string
	HoraInicio           string
	HoraTermino          *string
	UnidadAprendizaje    string
	Tema                 string
	DudaResuelta         *string
	ProfesorDeLaUA       string
	NombreCompletoAsesor string
}

type Profesor struct {
	Id     int64
	Nombre string
}

type UnidadAprendizaje struct {
	Nombre string
}
