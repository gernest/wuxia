package models

import "time"

//Project represent a project.
type Project struct {
	ID        int       `store:"id"`
	Name      string    `store:"name"`
	UserID    int       `store:"user_id"`
	CreatedAt time.Time `store:"created_at"`
	UpdatedAt time.Time `store:"updated_at"`
}
