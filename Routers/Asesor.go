package Routers

import (
	"apiAsesoria/Scanners"
	"apiAsesoria/Struct"
	"database/sql"
	"encoding/json"
	"github.com/gofiber/fiber/v2"
	"log"
)

func SetupAsesor(app *fiber.App, db *sql.DB) {
	asesor := app.Group("/Asesor")
	asesor.Get("/", func(c *fiber.Ctx) error {
		asesores, err := Scanners.Query(db, "SELECT * FROM asesor where Activo = 1", Scanners.ScanAsesor)
		if err != nil {
			log.Fatal(err)
		}
		j, err := json.Marshal(asesores)
		if err != nil {
			log.Fatal(err)
		}
		c.Type("json", "utf-8") // => "application/json; charset=utf-8"

		return c.Send(j)
	})
	asesor.Post("/newAsesor", func(c *fiber.Ctx) error {
		var asesor Struct.AsesorPost
		if err := c.BodyParser(&asesor); err != nil {
			log.Println("Error al analizar el cuerpo de la solicitud: ", err)
			return err
		}
		_, err := db.Exec("CALL InsertarAsesor(?, ?, ?)", asesor.Nombre, asesor.Matricula, asesor.Carrera)
		if err != nil {
			log.Println("Error al ejecutar la consulta SQL: ", err)
			return err
		}
		return c.SendString("Asesor insertado")
	})
}
