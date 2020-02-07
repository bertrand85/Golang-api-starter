package todos

import (
	"github.com/jinzhu/gorm"
	"github.com/labstack/echo"
)

func InitModule( e *echo.Echo, db *gorm.DB  ) {
	SetRoutes( e )
	makeMigration( db )
}

/**
create or update table(s) in DB
 */
func makeMigration(db *gorm.DB) {
	db.AutoMigrate( &Todo{} )
}


