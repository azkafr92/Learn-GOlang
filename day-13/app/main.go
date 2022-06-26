package main

import (
	"day-13/app/infrastructure"
)

func init() {
	infrastructure.InitSqliteDB()
}

func main() {
	infrastructure.Dispatch(infrastructure.DB)
}