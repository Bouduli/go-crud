package routers

import (
	"go-crud/internal/handlers"
	"go-crud/internal/utils/router"
)

var TodoRouter router.Router

// init function allows us to initialize variables
func init() {
	TodoRouter = router.NewRouter("/todo")

	//reference instance- that way our Index and Show get proper function-signatures
	var TodoHandler = &handlers.TodoHandler{}

	TodoRouter.GET("/", TodoHandler.Index)
	TodoRouter.GET("/{id}", TodoHandler.Show)
	TodoRouter.POST("/", TodoHandler.Create)
	TodoRouter.DELETE("/{id}", TodoHandler.Delete)
}
