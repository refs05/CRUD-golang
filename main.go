package main

import(
	"crudgoeg/config"
	"crudgoeg/routes"
)

func main() {
	config.InitDB()
	config.MigrateDB()
	e := routes.NewRoutes()
	e.Start(":1234")
}