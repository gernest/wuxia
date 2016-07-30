package models

import (
	"fmt"
	"time"

	"github.com/gernest/wuxia/db"
	"github.com/gernest/wuxia/metric"
	"github.com/microcosm-cc/bluemonday"
	"github.com/satori/go.uuid"
	"github.com/uber-go/zap"
)

const (
	//CtxBuild is the key that stores Context in all the requests.
	CtxBuild = "buildCtx"

	//TaskTable is the name of the database table for tasks.
	TaskTable = "tasks"

	//SessionTable is the name of  the database table for sessions
	SessionTable = "sessions"
)

var strict *bluemonday.Policy

func init() {
	strict = bluemonday.StrictPolicy()
}

//Context holda important information that can be used by diffenet components of
//the application.
type Context struct {
	Log    zap.Logger
	Metric metric.Metric
	Cfg    *Config
}

//Config configuration object.
type Config struct {
	Port       int
	WorkDir    string
	PublishDir string
}

//BuildArtifact is an interface which defines a buildable command.
type BuildArtifact interface {
	User() string
	Project() string
	Source() string
}

//BuildTask is a task for building a project.
type BuildTask struct {
	ID        int64
	UUID      string
	Done      bool
	User      string
	Project   string
	Source    string
	CreatedAt time.Time
	UpdateAt  time.Time
}

//CreateBuildTask creates a new task record and stores it into the database.
func CreateBuildTask(store *db.DB, t *BuildTask) (*BuildTask, error) {
	var query = `
	BEGIN TRANSACTION;
	  INSERT INTO %s VALUES ($1,$2,$3,$4,$5, $6);
	COMMIT;
	`
	query = fmt.Sprintf(query, TaskTable)
	tx, err := store.Begin()
	if err != nil {
		return nil, err
	}
	now := time.Now()
	uid := uuid.NewV4()
	_, err = tx.Exec(query,
		uid.String(),
		t.Done,
		t.User,
		t.Project,
		t.Source,
		now, now)
	if err != nil {
		return nil, err
	}
	err = tx.Commit()
	if err != nil {
		return nil, err
	}
	t.UUID = uid.String()
	t.CreatedAt = now
	t.UpdateAt = now
	return t, nil
}

//Sanitize sanitizes src to avoid SQL injections.
func Sanitize(src string) string {
	return strict.Sanitize(src)
}
