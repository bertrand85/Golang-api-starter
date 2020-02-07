package todos

import (
	"github.com/jinzhu/gorm"
	"github.com/satori/go.uuid"
	"time"
)

type Todo struct {
	Id            uuid.UUID `json:"id";gorm:"type:uuid;primary_key;"`
	Title         string `json:"title"`
	Description   string `json:"description"`
	ParentTask    string `json:"parent_task"`
	Done          string `json:"done"`
	Due			  time.Time `json: "due"`
}

// BeforeCreate will set a UUID rather than numeric ID.
func (base *Todo) BeforeCreate(scope *gorm.Scope) error {
	uid := uuid.NewV4()

	return scope.SetColumn("Id", uid)
}
