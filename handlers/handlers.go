package handlers

import (
	"fmt"
	"log"
	"net/http"
	"os"

	"github.com/fenriz07/EjemploTransbankGoRest/controllers"
	"github.com/gorilla/mux"
	"github.com/rs/cors"
)

/*Init punto de entrada que define las diferentes rutas del api*/
func Init() {

	router := mux.NewRouter()

	/*View*/
	router.HandleFunc("/", controllers.InitTransaction).Methods("GET")

	PORT := os.Getenv("PORT")

	if PORT == "" {
		PORT = "8080"
	}

	handler := cors.AllowAll().Handler(router)

	fmt.Println("Server Listen")

	log.Fatal(http.ListenAndServe(":"+PORT, handler))
}
