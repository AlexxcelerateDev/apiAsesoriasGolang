package main

import (
	"apiAsesoria/Routers"
	"database/sql"
	"fmt"
	"github.com/go-sql-driver/mysql"
	"github.com/gofiber/fiber/v2"
	"log"
)

var db *sql.DB

func main() {

	// Capture connection properties.
	cfg := mysql.Config{
		User:                 "root",
		Passwd:               "1234",
		Net:                  "tcp",
		Addr:                 "localhost:3306",
		DBName:               "asesorias",
		AllowNativePasswords: true,
	}
	var err error
	db, err = sql.Open("mysql", cfg.FormatDSN())
	if err != nil {
		log.Fatal(err)
	}
	pingErr := db.Ping()
	if pingErr != nil {
		log.Fatal(pingErr)
	}
	fmt.Println("Connected!")

	app := fiber.New()

	Routers.SetupAsesoria(app, db)
	Routers.SetupAlumno(app, db)
	Routers.SetupAsesor(app, db)
	Routers.SetupProfesor(app, db)
	Routers.SetupUA(app, db)
	app.Get("/", func(c *fiber.Ctx) error {
		return c.SendString("Hello, World!")
	})

	log.Fatal(app.Listen(":3000"))
}
