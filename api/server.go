package api

import (
	"fmt"
	"log"
	"os"

	"github.com/Kerlense/form3/api/controllers"
	"github.com/joho/godotenv"
)

var s = controllers.Server{}

func Run() {

	var err error
	err = godotenv.Load()
	if err != nil {
		log.Fatalf("Error for the env, not comming through %v", err)
	} else {
		fmt.Println("The env values is")
	}

	s.Initialize(os.Getenv("postgres"), os.Getenv("root"), os.Getenv("password"), os.Getenv("5432"), os.Getenv("postgresql"), os.Getenv("postgresql"))

}
