package main

import (
	"github.com/Bananad47/go_test_task_echo/pkg/app"
	"log"
)

func main() {
	a := app.CreateApp()
	err := a.Start()
	if err != nil {
		log.Fatal(err)
	}
}
