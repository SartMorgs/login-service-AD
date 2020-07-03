package main

import(
	"fmt"
	"net/http"
	"log"
	
	// Internal Packages
	"github.com/sartmorgs/app-login/api"

	// External Packages
	"github.com/gorilla/context"
)


func main(){
	port := ":8000"

	// Static directory
	fs := http.FileServer(http.Dir("static"))
	http.Handle("/css/login/", fs)
    http.Handle("/js/login/", fs)
    http.Handle("/css/acesso_autorizado/", fs)
    http.Handle("/js/acesso_autorizado/", fs)

	// Api
	http.HandleFunc("/api/v1/access/", api.Access)
	http.HandleFunc("/api/v1/logout/", api.Logout)


	fmt.Println("Server listening at" + port)
	log.Fatal(http.ListenAndServe(port, context.ClearHandler(http.DefaultServeMux)))
}
