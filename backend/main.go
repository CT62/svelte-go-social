package main

import(
  "github.com/gofiber/fiber/v2"
  "github.com/gofiber/fiber/v2/middleware/cors"

  "gorm.io/driver/sqlite"
  "gorm.io/gorm"

  "backend/models"
  "time"
)

func main(){
  db, err := gorm.Open(sqlite.Open("user.db"))
  if err != nil{
    panic("failed to connect to database")
  }

  app := fiber.New()

	app.Use(cors.New(cors.Config{
		AllowOrigins: "*",
		AllowMethods: "GET,POST,HEAD,PUT,DELETE,PATCH",
		AllowHeaders: "Origin, Content-Type, Accept",
	}))

  api := app.Group("/api")
  
  api.Post("/signup",func(c *fiber.Ctx) error {
      user := new(models.User)
      if err := c.BodyParser(user); err!= nil {
        panic(err)
      }

      db.Create(&user)
      return c.JSON(user)
  })

  app.Listen(":3000")
}
