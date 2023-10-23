package data

import (
	"database/sql"
	"fmt"

	"github.com/Mth-Ryan/waveaction/pkg/conf"
	_ "github.com/lib/pq"
)

type Database struct {
	Ctx *sql.DB
}

func NewDatabase(appConf *conf.AppConf) (*Database, error) {
	ctx, err := newPostgresConnection(appConf)
	database := Database { Ctx: ctx }

	return &database, err
}

func newPostgresConnection(appConf *conf.AppConf) (*sql.DB, error) {
	dbConf := appConf.Data.Database

	connectionStr := fmt.Sprintf(
		"host=%s port=%s user=%s password=%s dbname=%s sslmode=disable",
		dbConf.Host,
		dbConf.Port,
		dbConf.User,
		dbConf.Pass,
		dbConf.Name,
	)

	return sql.Open("postgres", connectionStr)
}


