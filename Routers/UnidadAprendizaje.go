package Routers

import (
	"apiAsesoria/Scanners"
	"apiAsesoria/Struct"
	"database/sql"
	"encoding/json"
	"github.com/gofiber/fiber/v2"
	"net/url"
)

func SetupUA(app *fiber.App, db *sql.DB) {
	ua := app.Group("/UA")
	ua.Get("/", func(c *fiber.Ctx) error {
		c.Type("json", "utf-8") // => "application/json; charset=utf-8"
		ua, err := Scanners.Query(db, "Select * from unidadaprendizaje", Scanners.ScanUA)
		if err != nil {
			return c.Status(fiber.StatusInternalServerError).SendString(err.Error())
		}
		j, err := json.Marshal(ua)

		return c.Send(j)
	})
	ua.Get("/buscar/:nombre", func(c *fiber.Ctx) error {
		nombre, err := url.QueryUnescape(c.Params("nombre"))

		c.Type("json", "utf-8") // => "application/json; charset=utf-8"
		ua, err := Scanners.Query(db, "Select * from unidadaprendizaje where Nombre = ?", Scanners.ScanUA, nombre)
		if err != nil {
			return c.Status(fiber.StatusInternalServerError).SendString(err.Error())
		}
		j, err := json.Marshal(ua)

		return c.Send(j)

	})
	ua.Post("/newUA", func(c *fiber.Ctx) error {
		var ua Struct.UnidadAprendizajePost
		if err := c.BodyParser(&ua); err != nil {
			return err
		}
		_, err := db.Exec("CALL InsertarUA(?)", ua.Nombre)
		if err != nil {
			return err
		}
		return c.JSON(fiber.Map{"message": "UA Insertada!"})
	})
}
