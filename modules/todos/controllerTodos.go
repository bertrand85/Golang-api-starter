package todos

import (
	db2 "api-starter/db"
	"github.com/labstack/echo"
	"net/http"
)

type ControllerTodos struct {
	Item       Todo
	Collection []Todo
}

/**
	List All
 */
func ( P *ControllerTodos) Index(c echo.Context) error  {
	db := db2.DbManager()

	db.Find( &P.Collection)

	return c.JSON(http.StatusOK, &P.Collection)
}

/**
	One record by ID
 */
func ( P *ControllerTodos) Show(c echo.Context) error  {
	db := db2.DbManager()
	id := c.Param("id")

	if err := db.Where("id = ?", id ).First( &P.Item ).Error; err != nil  {
		return c.JSON( http.StatusNotFound, err.Error() )
	}

	return c.JSON( http.StatusOK, &P.Item )
}

/**
	Create
 */
func ( P *ControllerTodos) Create(c echo.Context) error  {
	db := db2.DbManager()
	if err := c.Bind(&P.Item ); err != nil {
		return err
	}

	if err := db.Create( &P.Item  ).Error; err != nil {
		return c.JSON( http.StatusBadRequest, err )
	}

	return c.JSON( http.StatusCreated, &P.Item  )
}

/**
	Update
 */
func ( P *ControllerTodos) Update(c echo.Context) error  {
	db := db2.DbManager()
	id := c.Param("id")

	if err := db.Where("id = ?", id ).First( &P.Item  ).Error; err != nil  {
		return c.JSON( http.StatusNotFound, err.Error() )
	}

	var uuid = P.Item .Id
	if err := c.Bind(&P.Item ); err != nil {
		return c.JSON( http.StatusBadRequest, err )
	}
	P.Item .Id = uuid

	if err := db.Save( &P.Item  ).Error; err != nil {
		return c.JSON( http.StatusBadRequest, err )
	}

	return c.JSON( http.StatusAccepted, &P.Item  )
}

/**
	Delete
*/
func ( P *ControllerTodos) Delete(c echo.Context) error {
	db := db2.DbManager()
	id := c.Param("id")

	if err := db.Where("id = ?", id ).First( &P.Item  ).Error; err != nil  {
		return c.JSON( http.StatusNotFound, err.Error() )
	}

	if err := db.Delete(&P.Item).Error; err != nil  {
		return c.JSON( http.StatusBadRequest, err.Error() )
	}

	return c.NoContent( http.StatusOK  )
}