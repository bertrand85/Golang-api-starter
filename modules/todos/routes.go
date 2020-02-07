package todos

import (
	"github.com/labstack/echo"
)

/**
Set routes for module
*/
func SetRoutes( e *echo.Echo ) {

	controllerTodos := ControllerTodos{}

	e.GET   ("/todos",      controllerTodos.Index )
	e.GET   ("/todos/:id",  controllerTodos.Show )
	e.POST  ("/todos",      controllerTodos.Create )
	e.PUT   ("/todos/:id",  controllerTodos.Update )
	e.DELETE("/todos/:id",  controllerTodos.Delete )
}
