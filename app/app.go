package main

import (
	"fmt"
	"github.com/julienschmidt/httprouter"
	"net/http"
)

func indexHandler(w http.ResponseWriter, r *http.Request, _ httprouter.Params) {
	fmt.Fprintf(w, "Hello index!")
}

func main() {
	router := httprouter.New()
	router.GET("/", indexHandler)

	http.ListenAndServe(":8080", router)
}
