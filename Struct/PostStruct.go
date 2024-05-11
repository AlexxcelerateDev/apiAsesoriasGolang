package Struct

type AlumnoPost struct {
	Nombre    string `json:"nombre"`
	Carrera   string `json:"carrera"`
	Sexo      string `json:"sexo"`
	Matricula int64  `json:"matricula"`
}

type AsesoriaPost struct {
	Tema            string `json:"tema"`
	IdProfesor      int    `json:"idProfesor"`
	MatriculaAlumno int64  `json:"matriculaAlumno"`
	MatriculaAsesor int64  `json:"matriculaAsesor"`
	NombreUA        string `json:"nombreUA"`
	Oportunidad     string `json:"oportunidad"`
}

type AsesorPost struct {
	Nombre    string `json:"nombre"`
	Matricula int64  `json:"matricula"`
	Carrera   string `json:"carrera"`
}

type ProfesorPost struct {
	Nombre string `json:"nombre"`
}

type UnidadAprendizajePost struct {
	Nombre string `json:"nombre"`
}

type TerminarAsesoria struct {
	IdAsesoria   int    `json:"idAsesoria"`
	DudaResuelta string `json:"dudaResuelta"`
}
