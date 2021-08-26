package main

import (
	"api-echo/db"
	"api-echo/routes"
)

func main() {
	db.Init()
	e := routes.Init()

	e.Logger.Fatal(e.Start(":8888"))
}
