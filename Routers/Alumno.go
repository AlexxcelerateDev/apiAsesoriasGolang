package Routers

import (
	"apiAsesoria/Scanners"
	"apiAsesoria/Struct"
	"database/sql"
	"encoding/json"
	"github.com/gofiber/fiber/v2"
	"log"
)

func SetupAlumno(app *fiber.App, db *sql.DB) {
	alumno := app.Group("/Alumno")
	alumno.Get("/", func(c *fiber.Ctx) error {
		c.Type("json", "utf-8") // => "application/json; charset=utf-8"
		alumnos, err := Scanners.Query(db, "SELECT * FROM alumno", Scanners.ScanAlumno)
		if err != nil {
			// handle error
		}
		j, err := json.Marshal(alumnos)
		if err != nil {
			log.Fatal(err)
		}
		return c.Send(j)
	})
	alumno.Post("/newAlumno", func(c *fiber.Ctx) error {

		var alumno Struct.AlumnoPost
		if err := c.BodyParser(&alumno); err != nil {
			log.Println("Error al analizar el cuerpo de la solicitud: ", err)
			return err
		}
		_, err := db.Exec("CALL InsertarAlumno(?, ?, ?, ?)", alumno.Nombre, alumno.Carrera, alumno.Sexo, alumno.Matricula)
		if err != nil {
			log.Println("Error al ejecutar la consulta SQL: ", err)
			return err
		}
		return c.SendString("Alumno insertado")
	})
}
