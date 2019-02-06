package test

import (
	"softnet/pkg/database"
	"testing"
)

func TestBootstrap(t *testing.T) {
	db, _ := repository.NewGORMSqlLiteConnection("./db.sqlite")
	repository.Clean(db)
	repository.DefaultGenesis(db)
}
