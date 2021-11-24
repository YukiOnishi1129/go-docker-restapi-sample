package main

import (
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
	db := db.Init()

	// logic層
	authLogic :=  logic.NewAuthLogic()
	userLogic :=  logic.NewUserLogic()
	todoLogic := logic.NewTodoLogic()
	responseLogic := logic.NewResponseLogic()
	jwtLogic := logic.NewJWTLogic()

	// validation層
	authValidate := validation.NewAuthValidation()
	todoValidate := validation.NewTodoValidation()

	// repository層
	userRepo := repositories.NewUserRepository(db)
	todoRepo := repositories.NewTodoRepository(db)
	// service層
	authService := services.NewAuthService(userRepo, authLogic, userLogic, responseLogic, jwtLogic, authValidate)
	todoService := services.NewTodoService(todoRepo, todoLogic, responseLogic, todoValidate)
	// controller層
	appController := controllers.NewAppController()
	authController := controllers.NewAuthController(authService)
	todoContoroller := controllers.NewTodoController(todoService, authService)

	// router設定
	appRouter := router.NewAppRouter(appController)
	authRouter := router.NewAuthRouter(authController)
	todoRouter := router.NewTodoRouter(todoContoroller)
	mainRouter := router.NewMainRouter(appRouter, authRouter, todoRouter)

	// API起動
	mainRouter.StartWebServer()

	// server := http.Server{
	// 	Addr: ":4000",
	// }

	// server.ListenAndServe()

	// API起動
	// router.StartWebServer()
}