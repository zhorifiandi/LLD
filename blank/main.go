package main

import (
	"log"

	"github.com/zhorifiandi/LLD/blank/mvpapp"
)

type IApplication interface {
}

func main() {
	input := mvpapp.ApplicationInputs{}
	var app IApplication = mvpapp.NewApplication(
		input,
	)
	log.Printf("App is running..... %+v\n", app)

}
