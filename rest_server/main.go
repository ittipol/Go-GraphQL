package main

import (
	"fmt"
	"log"
	"restserver/database"
	"restserver/handlers/itemhandler"
	"restserver/repositories/itemrepository"
	"restserver/services/itemsrv"
	"strings"

	"github.com/gofiber/fiber/v2"
	"github.com/spf13/viper"
	"gorm.io/gorm"
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

	db := initDbConnection()

	itemRepository := itemrepository.NewItemRepository(db)
	itemService := itemsrv.NewItemService(itemRepository)
	itemHandler := itemhandler.NewItemHandler(itemService)

	app := fiber.New()

	// app.Get("/users", func(c *fiber.Ctx) error {

	// 	return c.JSON("USER")
	// })

	// app.Post("/user", func(c *fiber.Ctx) error {

	// 	body := c.Body()

	// 	fmt.Printf("Body: %v", string(body))

	// 	return c.JSON("User saved")
	// })

	app.Get("/allItems", itemHandler.AllItems)

	app.Get("/health", func(c *fiber.Ctx) error {
		return c.JSON("OK")
	})

	log.Fatal(app.Listen(fmt.Sprintf(":%v", viper.GetInt("app.port"))))

}

func initDbConnection() *gorm.DB {
	return database.GetDbConnection(
		viper.GetString("database.username"),
		viper.GetString("database.password"),
		viper.GetString("database.host"),
		viper.GetInt("database.port"),
		viper.GetString("database.db"),
		false,
	)
}
