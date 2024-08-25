package user

type User struct {
	ID        int     `bson:"_id,omiempty" json:"id"`
	Email     string  `bson:"email" json:"email"`
	Name      string  `bson:"name" json:"name"`
	Password  string  `bson:"password" json:"password"`
	CreatedAt []uint8 `bson:"created_at" json:"created_at"`
	UpdatedAt []uint8 `bson:"updated_at" json:"updated_at"`
}
