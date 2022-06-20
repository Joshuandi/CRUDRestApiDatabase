package user

import (
	"time"
)

type User struct {
	Id         int
	Username   string
	Email      string
	Password   string
	Age_user   int
	Division   string
	Created_at time.Time
	Updated_at time.Time
}

type photo struct {
	id         int
	title      string
	caption    string
	photo_url  string
	user_id    int
	created_at time.Time
	updated_at time.Time
}

type comment struct {
	id         int
	user_id    int
	photo_id   int
	message    string
	created_at time.Time
	updated_at time.Time
}
type socialMedia struct {
	id              int
	name            string
	socialMedia_url string
	user_id         int
}
