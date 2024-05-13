package Routers

import (
	"apiAsesoria/Scanners"
	"apiAsesoria/Struct"
	"database/sql"
	"encoding/json"
	"fmt"
	"github.com/gofiber/fiber/v2"
	"log"
)

func SetupAsesoria(app *fiber.App, db *sql.DB) {
	asesoria := app.Group("/Asesoria")
	asesoria.Get("/", func(c *fiber.Ctx) error {
		c.Type("json", "utf-8") // => "application/json; charset=utf-8"
		asesorias, err := Scanners.Query(db, "call VerAsesoriasNoterminadas()", Scanners.ScanAsesoriaNoTerminada)
		if err != nil {
			fmt.Println("Error:", err)
			return c.Status(fiber.StatusInternalServerError).SendString(err.Error())
		}
		j, err := json.Marshal(asesorias)
		if err != nil {
			log.Fatal(err)
		}
		return c.Send(j)
	})
	asesoria.Post("/newAsesoria", func(c *fiber.Ctx) error {
		var asesoria Struct.AsesoriaPost
		if err := c.BodyParser(&asesoria); err != nil {
			log.Println("Error al analizar el cuerpo de la solicitud: ", err)
			return err
		}
		_, err := db.Exec("CALL InsertarAsesoria(?, ?, ?, ?, ?, ?)", asesoria.Tema, asesoria.IdProfesor,
			asesoria.MatriculaAlumno, asesoria.MatriculaAsesor, asesoria.NombreUA, asesoria.Oportunidad)
		if err != nil {
			log.Println("Error al ejecutar la consulta SQL: ", err)
			return err
		}
		return c.JSON(fiber.Map{"message": "Asesoria insertado"})
	})
	asesoria.Post("/endAsesoria", func(c *fiber.Ctx) error {
		var salidAsesoria Struct.TerminarAsesoria
		if err := c.BodyParser(&salidAsesoria); err != nil {
			log.Println("Error al analizar el cuerpo de la solicitud: ", err)
			return err
		}
		_, err := db.Exec("CALL SalidaAsesoria(?, ?)", salidAsesoria.IdAsesoria, salidAsesoria.DudaResuelta)
		if err != nil {
			log.Println("Error al ejecutar la consulta SQL: ", err)
			return err
		}
		return c.JSON(fiber.Map{"message": "Asesoria terminada!"})
	})

}
