package main

import (
	"day-12/app/infrastructure"
)

func init() {
	infrastructure.InitSqliteDB()
}

func main() {
	infrastructure.Dispatch(infrastructure.DB)
}
