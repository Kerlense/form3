package controllers

import (
	"fmt"
	"log"
	"net/http"

	
	"github.com/jinzhu/gorm"
	"github.com/gorilla/mux"
	_ "github.com/jinzhu/gorm/dialects/postgres" 
	"github.com/Kerlense/form3/api/models"
)

type DBServer struct {
	DB     *gorm.DB
	Router *mux.Router
}

func (s *DBServer) {

	var err error
	connStr := fmt.Sprintf("user=%s password=%s port=%s dbname=%s sslmode=disable", "root", "password", 5432, "postgresql")
	db, err := sql.Open("postgres", connStr)
		if err != nil {
			fmt.Printf("Cannot connect to %s database", postgres)
			log.Fatal("This is the error:", err)
		} else {
			fmt.Printf("We are connected to the %s database", postgres)
		}
	}

	
	s.Router = mux.NewRouter()

	s.initializeRoutes()
}

func (s *Server) Run(addr string) {
	fmt.Println("Listening to port 8080")
	log.Fatal(http.ListenAndServe(addr, s.Router))
}