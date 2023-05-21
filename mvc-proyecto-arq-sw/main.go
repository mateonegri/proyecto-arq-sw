package main

import (
	"mvc-proyecto-arq-sw/app"
	"mvc-proyecto-arq-sw/db"
)

func main() {
	db.StartDbEngine()
	app.StartRoute()
}
