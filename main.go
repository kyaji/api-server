package main

import (
	"database/sql"
	"fmt"
	_ "github.com/go-sql-driver/mysql"
	"github.com/julienschmidt/httprouter"
	"github.com/rs/cors"
	"github.com/shakdwipeea/shadowfax/server"
	"log"
	"net/http"
)

//Index Server root
func Index(w http.ResponseWriter, r *http.Request, _ httprouter.Params) {
	fmt.Fprintf(w, "VCR MUSIC \n")
}

//Hello route
func Hello(w http.ResponseWriter, r *http.Request, pm httprouter.Params) {
	fmt.Fprintf(w, "Hello %s \n", pm.ByName("name"))
}


func main() {
	fmt.Println("Starting the web server")

	fmt.Println("Connecting to MYSQL database")

	db, err := sql.Open("mysql", "root:@/shadowfax")
	if err != nil {
		panic(err.Error()) // Just for example purpose. You should use proper error handling instead of panic
	}
	defer db.Close()

	// Set up environment
	env := server.Env{Db: db}

	router := httprouter.New()
	router.GET("/", Index)
	router.GET("/hello/:name", Hello)

	//router.GET("/write", WriteToDb)

	server.RegisterHandlers(router, env)

	handler := cors.Default().Handler(router)

	err = http.ListenAndServe(":8080", handler)

	log.Fatal(err)
}