package main

import (
	"fmt"
	"log"
	"net/http"
	"os"

	"github.com/joho/godotenv"
	"github.com/thompsonmss/minificador-url-golang-with-gorm/database"
	"github.com/thompsonmss/minificador-url-golang-with-gorm/routers"
)

func main() {

	err := godotenv.Load()
	if err != nil {
		log.Fatalf("Error loading .env file")
	}

	port := os.Getenv("PORT")
	if len(port) == 0 {
		port = "8090"
	}

	database.StartDatabase()

	// Call to routers
	routers.Handler()
	http.Handle("/static/", http.StripPrefix("/static/", http.FileServer(http.Dir("views/static"))))

	println(fmt.Sprintf("Server started in host: http://localhost:%v", port))
	log.Fatal(http.ListenAndServe(":"+port, nil))
}
