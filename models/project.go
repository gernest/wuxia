package models

import "time"

//Project represent a project.
type Project struct {
	ID        int
	Name      string
	UserID    int
	CreatedAt time.Time
	UpdatedAt time.Time
}
