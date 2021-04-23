package main

import (
	"github.com/gofiber/fiber/v2"
	"github.com/joho/godotenv"
	"github.com/liuguangw/zx_account/routes"
	"github.com/liuguangw/zx_account/service/middlewares"
	"log"
	"os"
)

func main() {
	if err := godotenv.Load(); err != nil {
		log.Fatal(err)
	}
	app := SetupApp()
	listenAddress := os.Getenv("LISTEN_ADDRESS")
	if err := app.Listen(listenAddress); err != nil {
		log.Fatal(err)
	}
}

//SetupApp 初始化fiber服务(此函数仅用于command内部调用和单元测试)
func SetupApp() *fiber.App {
	app := fiber.New()
	app.Use(middlewares.RecoverHandle)
	//加载api路由
	routes.LoadAPIRoutes(app)
	routes.LoadCommonRoutes(app)
	app.Use(middlewares.Error404Handle)
	return app
}
