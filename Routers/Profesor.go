package Routers

import (
	"apiAsesoria/Scanners"
	"apiAsesoria/Struct"
	"database/sql"
	"encoding/json"
	"github.com/gofiber/fiber/v2"
	"log"
	"net/url"
)

func SetupProfesor(app *fiber.App, db *sql.DB) {
	profesor := app.Group("/Profesor")
	profesor.Get("/", func(c *fiber.Ctx) error {
		c.Type("json", "utf-8") // => "application/json; charset=utf-8"
		profesores, err := Scanners.Query(db, "SELECT * FROM profesor", Scanners.ScanProfesor)
		if err != nil {
			log.Fatal(err)
		}
		j, err := json.Marshal(profesores)
		if err != nil {
			log.Fatal(err)
		}
		return c.Send(j)
	})
	profesor.Get("/buscar/:nombre", func(c *fiber.Ctx) error {
		nombre, err := url.QueryUnescape(c.Params("nombre"))

		c.Type("json", "utf-8") // => "application/json; charset=utf-8"
		profesor, err := Scanners.Query(db, "Select * from profesor where Nombre = ?", Scanners.ScanProfesor, nombre)
		if err != nil {
			return c.Status(fiber.StatusInternalServerError).SendString(err.Error())
		}
		j, err := json.Marshal(profesor)

		return c.Send(j)

	})
	profesor.Post("/newProfesor", func(c *fiber.Ctx) error {

		var profesor Struct.ProfesorPost
		if err := c.BodyParser(&profesor); err != nil {
			log.Println("Error al analizar el cuerpo de la solicitud: ", err)
			return err
		}
		_, err := db.Exec("CALL InsertarProfesor(?)", profesor.Nombre)
		if err != nil {
			log.Println("Error al ejecutar la consulta SQL: ", err)
			return err
		}
		return c.SendString("Profesor insertado")
	})
}
