package models

import (
	"database/sql"

	"github.com/fatih/structs"
	"github.com/gernest/wuxia/db"
	"github.com/microcosm-cc/bluemonday"
	"github.com/uber-go/zap"
)

const (
	//CtxBuild is the key that stores Context in all the requests.
	CtxBuild = "buildCtx"

	//TaskTable is the name of the database table for tasks.
	TaskTable = "tasks"

	//SessionTable is the name of  the database table for sessions
	SessionTable = "sessions"

	//UserTable is the name of the databse table for users.
	UserTable = "users"

	//DBTag is the name of the struct tag used by the store
	DBTag = "store"
)

var strict *bluemonday.Policy

func init() {
	strict = bluemonday.StrictPolicy()
}

//Context holda important information that can be used by diffenet components of
//the application.
type Context struct {
	Log zap.Logger
	Cfg *Config
}

//Config configuration object.
type Config struct {
	Port       int
	WorkDir    string
	PublishDir string
}

//Sanitize sanitizes src to avoid SQL injections.
func Sanitize(src string) string {
	return strict.Sanitize(src)
}

type fieldInfo struct {
	tag   string
	field *structs.Field
}

func getFieldInfo(model interface{}) []fieldInfo {
	var mFields []fieldInfo
	f := structs.Fields(model)
	for _, v := range f {
		mFields = append(mFields, fieldInfo{v.Tag(DBTag), v})
	}
	return mFields
}

func getArgs(q db.Query, info []fieldInfo) []interface{} {
	var args []interface{}
	params := q.Params()
	for _, v := range params {
		for i := range info {
			if v == info[i].tag {
				f := info[i].field
				args = append(args, f.Value())
			}
		}
	}
	return args
}

//ExecModel executes the given query. If the guery requires data from the model
//then the data is also taken care of.
//
// This provides a simple abstraction for repetitive dababase query execution.
// It supports transactions.
func ExecModel(store *db.DB, model interface{}, query db.Query) (sql.Result, error) {
	info := getFieldInfo(model)
	if query.IsTx() {
		return execTx(store, query, info)
	}
	return execNormal(store, query, info)
}

func execTx(store *db.DB, q db.Query, info []fieldInfo) (sql.Result, error) {
	tx, err := store.Begin()
	if err != nil {
		return nil, err
	}
	args := getArgs(q, info)
	rst, err := tx.Exec(q.String(), args...)
	if err != nil {
		_ = tx.Rollback()
		return rst, err
	}
	err = tx.Commit()
	if err != nil {
		_ = tx.Rollback()
		return rst, err
	}
	return rst, nil
}

func execNormal(store *db.DB, q db.Query, info []fieldInfo) (sql.Result, error) {
	args := getArgs(q, info)
	return store.Exec(q.String(), args...)
}

//QueryRowModel use the model to query for one row.
func QueryRowModel(store *db.DB, model interface{}, query db.Query) *sql.Row {
	info := getFieldInfo(model)
	args := getArgs(query, info)
	return store.QueryRow(query.String(), args...)
}
