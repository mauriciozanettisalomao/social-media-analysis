package main

import (
	"fmt"
	"net/http"

	"github.com/gorilla/mux"

	//socialmedia "github.com/mauriciozanettisalomao/social-media-analysis/internal/app/factory"

	handler "github.com/mauriciozanettisalomao/social-media-analysis/api"
)

const (
	port = ":80"
)

func main() {

	r := mux.NewRouter()

	r.HandleFunc("/socialmedia/{socialmedia}/{param}", handler.SearchMention).Methods("GET")
	r.HandleFunc("/socialmedia/{socialmedia}/user/{username}", handler.SearchUser).Methods("GET")
	r.HandleFunc("/socialmedia/{socialmedia}/user/{username}/followers", handler.SearchFollowers).Methods("GET")

	fmt.Printf("listening on port %s", port)

	http.ListenAndServe(port, r)

}
