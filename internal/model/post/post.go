package post

type Post struct {
	ID        int     `bson:"_id,omiempty" json:"id"`
	Name      string  `bson:"name" json:"name"`
	UserId    int     `bson:"user_id,omiempty" json:"user_id"`
	ForumID   int     `bson:"forum_id,omiempty" json:"forum_id"`
	Text      string  `bson:"text" json:"text"`
	CreatedAt []uint8 `bson:"created_at" json:"created_at"`
	UpdatedAt []uint8 `bson:"updated_at" json:"updated_at"`
}
