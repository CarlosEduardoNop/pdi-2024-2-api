package forum

import (
	"fmt"
	"forum-api/pkg/db"
	"time"
)

type LastPost struct {
	ID       *int       `bson:"_id,omiempty" json:"id,omitempty"`
	Name     *string    `bson:"name" json:"name,omitempty"`
	UserName *string    `bson:"user_name" json:"user_name,omitempty"`
	Date     *time.Time `bson:"date" json:"date,omitempty"`
}

func GetAllWithLastPost() ([]*Forum, error) {
	conn, err := db.OpenConnection()

	if err != nil {
		fmt.Println(err)
		return nil, err
	}

	defer conn.Close()

	sql := `
		SELECT forum.*,
		       last_post.id AS last_post_id,
			   last_post.name AS last_post_name,
			   user.name AS last_post_user_name,
			   last_post.created_at AS last_post_date,
			   COUNT(post.forum_id) AS post
		FROM sistema.forums AS forum
				 LEFT JOIN sistema.posts AS post ON forum.id = post.forum_id
				 LEFT JOIN (
			SELECT p.forum_id,
				   p.name,
				   p.created_at,
				   p.user_id,
				   p.id
			FROM sistema.posts p
			WHERE (p.forum_id, p.created_at) IN (
				SELECT forum_id, MAX(created_at)
				FROM sistema.posts
				GROUP BY forum_id
			)
		) AS last_post ON forum.id = last_post.forum_id
			 LEFT JOIN sistema.users AS user ON last_post.user_id = user.id
		GROUP BY forum.id, last_post.name, last_post.created_at, user.name, last_post.id
  	`

	res, err := conn.Query(sql)

	if err != nil {
		fmt.Println(err)
		return nil, err
	}

	if err != nil {
		fmt.Println(err)
		return nil, err
	}

	forums := []*Forum{}

	for res.Next() {
		var forum Forum
		var lastPost LastPost

		err = res.Scan(
			&forum.ID,
			&forum.Name,
			&forum.Description,
			&forum.Image,
			&forum.CreatedAt,
			&forum.UpdatedAt,
			&lastPost.ID,
			&lastPost.Name,
			&lastPost.UserName,
			&lastPost.Date,
			&forum.Posts,
		)

		if err != nil {
			fmt.Println(err)
			return nil, err
		}

		forum.LastPost = &lastPost

		forums = append(forums, &forum)
	}

	return forums, nil
}
