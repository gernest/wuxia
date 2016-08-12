package models

import "time"

//BuildArtifact is an interface which defines a buildable command.
type BuildArtifact interface {
	User() string
	Project() string
	Source() string
}

//BuildTask is a task for building a project.
type BuildTask struct {
	ID        int64     `store:"id"`
	UUID      string    `store:"uuid"`
	Done      bool      `store:"done"`
	User      int64     `store:"user_id"`
	Project   int64     `store:"project_id"`
	Source    string    `store:"source"`
	CreatedAt time.Time `store:"created_on"`
	UpdateAt  time.Time `store:"updated_on"`
}
