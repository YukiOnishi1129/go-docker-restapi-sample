package main

import (
	"log"
	"myapp/controllers"
	"myapp/db"
	"myapp/repositories"
	"myapp/router"
	"myapp/services"
	"myapp/utils/logic"
	"myapp/utils/validation"
)

func main() {
	// DB接続
	dbCon := db.Init()

	// logic層
	authLogic := logic.NewAuthLogic()
	userLogic := logic.NewUserLogic()
	todoLogic := logic.NewTodoLogic()
	responseLogic := logic.NewResponseLogic()
	jwtLogic := logic.NewJWTLogic()

	// validation層
	authValidate := validation.NewAuthValidation()
	todoValidate := validation.NewTodoValidation()

	// repository層
	userRepo := repositories.NewUserRepository(dbCon)
	todoRepo := repositories.NewTodoRepository(dbCon)
	// service層
	authService := services.NewAuthService(userRepo, authLogic, userLogic, responseLogic, jwtLogic, authValidate)
	todoService := services.NewTodoService(todoRepo, todoLogic, responseLogic, todoValidate)
	// controller層
	appController := controllers.NewAppController()
	authController := controllers.NewAuthController(authService)
	todoController := controllers.NewTodoController(todoService, authService)

	// router設定
	appRouter := router.NewAppRouter(appController)
	authRouter := router.NewAuthRouter(authController)
	todoRouter := router.NewTodoRouter(todoController)
	mainRouter := router.NewMainRouter(appRouter, authRouter, todoRouter)

	// API起動
	if err := mainRouter.StartWebServer(); err != nil {
		log.Printf("error start web server")
	}
}
