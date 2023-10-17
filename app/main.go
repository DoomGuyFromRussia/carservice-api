package main

import (
	//"mux"

	"app/models"
	"fmt"
	"io"
	"net/http"
	"strconv"

	//"github.com/gorilla/mux"

	"github.com/gorilla/mux"
	_ "github.com/mattn/go-sqlite3"
)

func main() {
	router := mux.NewRouter()
	router.HandleFunc("/api/getCars", getCars)
	router.HandleFunc("/api/getCar", getCar)

	http.ListenAndServe(":8010", router)

}

func checkErr(err error) {
	if err != nil {
		panic(err)
	}
}

func getCars(w http.ResponseWriter, r *http.Request) {
	fmt.Println("got getCars request")
	io.WriteString(w, models.GetCars()+"\n")
}

func getCar(w http.ResponseWriter, r *http.Request) {
	fmt.Println("got getCar request, id=", r.URL.Query().Get("id"))
	id, err := strconv.Atoi(r.URL.Query().Get("id"))
	checkErr(err)
	io.WriteString(w, models.GetCar(id))
}
