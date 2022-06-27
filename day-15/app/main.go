package main

import (
	"day-15/app/infrastructure"
)

func init() {
	infrastructure.InitSqliteDB()
	infrastructure.InitialMigration()
}

func main() {
	infrastructure.Dispatch(infrastructure.DB)
}
