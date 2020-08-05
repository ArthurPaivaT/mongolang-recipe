package main

import (
	"fmt"
	"net/http"

	"github.com/gorilla/mux"
)

func main() {

	router := mux.NewRouter()
	router.HandleFunc("/hi", func(w http.ResponseWriter, r *http.Request) {
		fmt.Fprintf(w, "hi, it worked")
	})

	fmt.Println("listening on :1212")
	err := http.ListenAndServe(":1212", router)
	if err != nil {
		fmt.Println("Error on listen and serve: ", err)
	}

	return
}
