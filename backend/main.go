package main

import(
  "github.com/gofiber/fiber/v2"
  "github.com/gofiber/fiber/v2/middleware/cors"

  "gorm.io/gorm"
  "gorm.io/driver/sqlite"

  "fmt"
)

type User struct{
  gorm.Model
  Username    string `json:"username"`
  Password    string `json:"password"`
}

func main(){
  db, err := gorm.Open(sqlite.Open("user.db"))
  if err != nil{
    panic("failed to connect to database")
  }
  var user User
  db.First(&user)
  fmt.Println(user)
  app := fiber.New()


	app.Use(cors.New(cors.Config{
		AllowOrigins: "*",
		AllowMethods: "GET,POST,HEAD,PUT,DELETE,PATCH",
		AllowHeaders: "Origin, Content-Type, Accept",
	}))

  api := app.Group("/api")
  
  api.Post("/signup",func(c *fiber.Ctx) error {
      user := new(User)
      if err := c.BodyParser(user); err!= nil {
        return err
      }

      return c.JSON(user)
  })

  app.Listen(":3000")
}
