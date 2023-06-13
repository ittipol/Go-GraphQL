package main

import (
	"fmt"
	"log"
	"restserver/database"
	"restserver/handlers/categoryhandler"
	"restserver/handlers/itemhandler"
	"restserver/repositories/categoryrepository"
	"restserver/repositories/itemrepository"
	"restserver/services/categorysrv"
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

	categoryRepository := categoryrepository.NewCategoryRepository(db)
	categoryService := categorysrv.NewCategoryService(categoryRepository)
	categoryhandler := categoryhandler.NewCategoryHandler(categoryService)

	app := fiber.New()

	app.Get("/allItems", itemHandler.AllItems)

	app.Post("/addItem", func(c *fiber.Ctx) error {

		body := c.Body()

		fmt.Printf("req body: %v", string(body))

		return c.JSON("Item saved")
	})

	app.Get("/getItemBySlug/:slug", itemHandler.GetItemBySlug)

	app.Get("AllCategories", categoryhandler.AllCategories)

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
