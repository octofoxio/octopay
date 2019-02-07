package main

import (
	"softnet/cmd/sandbox/lib"
	"softnet/pkg/repository"
)

func main() {
	db, _ := repository.NewGORMSqlLiteConnection("./agent.sqlite")
	lib.StartSandboxServer(db)
}
