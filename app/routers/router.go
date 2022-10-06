/* This file is auto-generated, manual edits in this file will be overwritten! */
package routers

import (
	"github.com/gin-contrib/cors"
	"github.com/gin-gonic/gin"
	"github.com/nightborn-be/blink/skipr-test/app/controllers"
	"github.com/nightborn-be/blink/skipr-test/app/middlewares"
)

type Router struct {
	engine         *gin.Engine
	authMiddleware middlewares.IAuthMiddleware
	controller     controllers.Controller
}

func Initialise(engine *gin.Engine, authMiddleware middlewares.IAuthMiddleware, controller controllers.Controller) Router {
	return Router{
		engine:         engine,
		authMiddleware: authMiddleware,
		controller:     controller,
	}
}

// Initiliases the router with the entire sub-tree and runs it
func (router Router) Run() error {

	// Setup CORS
	router.engine.Use(cors.New(cors.Config{
		AllowOrigins:     []string{"*"},
		AllowMethods:     []string{"*"},
		AllowHeaders:     []string{"*"},
		ExposeHeaders:    []string{"*"},
		AllowCredentials: true,
	}))

	// Creates the api-group
	api := router.engine.Group("")

	// Middlewares
	middlewares.UseMiddlewares(api)

	////////////////////////////////////////
	//     Initialises all the routers    //
	////////////////////////////////////////

// Expenses
api.GET("/expenses", router.controller.ExpenseController.GetExpenses)
api.POST("/expenses", router.controller.ExpenseController.CreateExpense)
api.GET("/expenses/:expenseId", router.controller.ExpenseController.GetExpense)
api.PUT("/expenses/:expenseId", router.controller.ExpenseController.UpdateExpense)

	// Runs the engine
	if err := router.engine.Run(); err != nil {
		return err
	}

	return nil
}
