package main

import (
	"day-14/app/infrastructure"
)

func init() {
	infrastructure.InitSqliteDB()
}

func main() {
	infrastructure.Dispatch(infrastructure.DB)
}
