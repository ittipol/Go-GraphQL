package main

import (
	"fmt"
	"log"
	"strings"

	"github.com/gofiber/fiber/v2"
	"github.com/spf13/viper"
)

func init() {
	viper.SetConfigName("config")
	viper.SetConfigType("yaml")
	viper.AddConfigPath(".")
	viper.AutomaticEnv()
	viper.SetEnvKeyReplacer(strings.NewReplacer(".", "_"))

	if err := viper.ReadInConfig(); err != nil {
		panic(fmt.Errorf("Fatal error config file: %s \n", err))
	}
}

func main() {

	app := fiber.New()

	app.Get("/users", func(c *fiber.Ctx) error {

		return c.JSON("USER")
	})

	app.Post("/user", func(c *fiber.Ctx) error {

		body := c.Body()

		fmt.Printf("Body: %v", string(body))

		return c.JSON("User saved")
	})

	app.Get("/health", func(c *fiber.Ctx) error {
		return c.JSON("OK")
	})

	log.Fatal(app.Listen(fmt.Sprintf(":%v", viper.GetInt("app.port"))))

}
