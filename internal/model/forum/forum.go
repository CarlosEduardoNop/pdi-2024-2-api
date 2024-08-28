package forum

import (
	"time"
)

type Forum struct {
	ID          int       `bson:"_id,omiempty" json:"id"`
	Name        string    `bson:"name" json:"name"`
	Description string    `bson:"description" json:"description"`
	Image       string    `bson:"image" json:"image"`
	CreatedAt   time.Time `bson:"created_at" json:"created_at"`
	UpdatedAt   time.Time `bson:"updated_at" json:"updated_at"`
	Posts       *int      `bson:"posts" json:"posts,omitempty"`
	LastPost    *LastPost `bson:"last_post" json:"last_post,omitempty"`
}
