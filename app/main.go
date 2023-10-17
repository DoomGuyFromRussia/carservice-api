package main

import (
	//"mux"

	"app/models"
	"encoding/json"
	"fmt"
	"io"
	"net/http"

	//"strconv"

	//"github.com/gorilla/mux"

	"github.com/gorilla/mux"
	_ "github.com/mattn/go-sqlite3"
)

func main() {
	router := mux.NewRouter()
	router.HandleFunc("/api/getCars", getCars)
	router.HandleFunc("/api/createCar", createCar)
	router.HandleFunc("/api/updateCar", updateCar)
	router.HandleFunc("/api/deleteCar", deleteCar)
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

func createCar(w http.ResponseWriter, r *http.Request) {
	fmt.Println("got createCar request")
	decoder := json.NewDecoder(r.Body)
	var c models.Car
	err := decoder.Decode(&c)
	checkErr(err)
	if !carValid(c) {
		ret := map[string]string{
			"error": "null values",
		}
		r, err := json.Marshal(ret)
		checkErr(err)
		io.WriteString(w, string(r))
	} else {

		io.WriteString(w, models.CreateCar(c)+"\n")
	}
}

func updateCar(w http.ResponseWriter, r *http.Request) {
	fmt.Println("got createCar request")
	decoder := json.NewDecoder(r.Body)
	var c models.Car
	err := decoder.Decode(&c)
	checkErr(err)
	if c.Model == "" || c.Producer == "" || c.Vin == "" || c.Year == "" {
		ret := map[string]string{
			"error": "null values",
		}
		r, err := json.Marshal(ret)
		checkErr(err)
		io.WriteString(w, string(r))
	} else {
		io.WriteString(w, models.UpdateCar(c)+"\n")
	}
}

func deleteCar(w http.ResponseWriter, r *http.Request) {
	fmt.Println("got deleteCar request")
	decoder := json.NewDecoder(r.Body)
	var id int
	err := decoder.Decode(&id)
	checkErr(err)
	io.WriteString(w, models.DeleteCar(id)+"\n")
}

// i need validator

func carValid(c models.Car) bool {
	if c.Model == "" || c.Producer == "" || c.Vin == "" || c.Year == "" {
		return false
	} else {
		return true
	}

}
