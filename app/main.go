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

type IdInt struct {
	Id int `json:",string"`
}

func main() {
	router := mux.NewRouter()
	router.HandleFunc("/api/getCars", getCars)
	router.HandleFunc("/api/getClients", getClients)
	router.HandleFunc("/api/getOrders", getOrders)
	router.HandleFunc("/api/createCar", createCar)
	router.HandleFunc("/api/createClient", createClient)
	router.HandleFunc("/api/createOrder", createOrder)
	router.HandleFunc("/api/updateCar", updateCar)
	router.HandleFunc("/api/updateClient", updateClient)
	router.HandleFunc("/api/updateOrder", updateOrder)
	router.HandleFunc("/api/deleteCar", deleteCar)
	router.HandleFunc("/api/deleteClient", deleteClient)
	router.HandleFunc("/api/deleteOrder", deleteOrder)

	fmt.Println(models.GetClientOrders(2))
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

func getClients(w http.ResponseWriter, r *http.Request) {
	fmt.Println("got getClients request")
	io.WriteString(w, models.GetClients()+"\n")
}

func getOrders(w http.ResponseWriter, r *http.Request) {
	fmt.Println("got getOrders request")
	io.WriteString(w, models.GetOrders()+"\n")
}

func createCar(w http.ResponseWriter, r *http.Request) {
	fmt.Println("got createCar request")
	decoder := json.NewDecoder(r.Body)
	var c models.Car
	err := decoder.Decode(&c)
	checkErr(err)
	if !c.Validate() {
		sendErrorJson(w, r, "validation fail")
	} else {

		io.WriteString(w, models.CreateCar(c)+"\n")
	}
}

func createClient(w http.ResponseWriter, r *http.Request) {
	fmt.Println("got createClient request")
	decoder := json.NewDecoder(r.Body)
	var c models.Client
	err := decoder.Decode(&c)
	checkErr(err)
	if !c.Validate() {
		sendErrorJson(w, r, "validation fail")
	} else {

		io.WriteString(w, models.CreateClient(c)+"\n")
	}
}

func createOrder(w http.ResponseWriter, r *http.Request) {
	fmt.Println("got createOrder request")
	decoder := json.NewDecoder(r.Body)
	var o models.Order
	err := decoder.Decode(&o)
	checkErr(err)
	fmt.Println("main desc", o.Description)
	if !o.Validate() {
		sendErrorJson(w, r, "validation fail")
	} else {
		io.WriteString(w, models.CreateOrder(o)+"\n")
	}
}

func updateCar(w http.ResponseWriter, r *http.Request) {
	fmt.Println("got updateCar request")
	decoder := json.NewDecoder(r.Body)
	var c models.Car
	err := decoder.Decode(&c)
	checkErr(err)
	if !c.Validate() {
		sendErrorJson(w, r, "validation fail")
	} else {
		io.WriteString(w, models.UpdateCar(c)+"\n")
	}
}

func updateClient(w http.ResponseWriter, r *http.Request) {
	fmt.Println("got updateClient request")
	decoder := json.NewDecoder(r.Body)
	var c models.Client
	err := decoder.Decode(&c)
	checkErr(err)
	if !c.Validate() {
		sendErrorJson(w, r, "validation fail")
	} else {
		io.WriteString(w, models.UpdateClient(c)+"\n")
	}

}

func updateOrder(w http.ResponseWriter, r *http.Request) {
	fmt.Println("got updateOrder request")
	decoder := json.NewDecoder(r.Body)
	var o models.Order
	err := decoder.Decode(&o)
	checkErr(err)
	if !o.Validate() {
		sendErrorJson(w, r, "validation fail")
	} else {
		io.WriteString(w, models.UpdateOrder(o)+"\n")
	}
}

func deleteCar(w http.ResponseWriter, r *http.Request) {
	fmt.Println("got deleteCar request")
	decoder := json.NewDecoder(r.Body)
	var j IdInt
	err := decoder.Decode(&j)
	checkErr(err)
	io.WriteString(w, models.DeleteCar(j.Id)+"\n")
}

func deleteClient(w http.ResponseWriter, r *http.Request) {
	fmt.Println("got deleteClient request")
	decoder := json.NewDecoder(r.Body)
	var j IdInt
	err := decoder.Decode(&j)
	checkErr(err)
	io.WriteString(w, models.DeleteClient(j.Id)+"\n")
}

func deleteOrder(w http.ResponseWriter, r *http.Request) {
	fmt.Println("got deleteOrder request")
	decoder := json.NewDecoder(r.Body)
	var j IdInt
	err := decoder.Decode(&j)
	checkErr(err)
	io.WriteString(w, models.DeleteOrder(j.Id)+"\n")
}

// i need validator

/*func carValid(c models.Car) bool {
	if c.Model == "" || c.Producer == "" || c.Vin == "" || c.Year == "" {
		return false
	} else {
		return true
	}

}*/

func getClientCars(w http.ResponseWriter, r *http.Request) {
	fmt.Println("got getClientCars request")
	decoder := json.NewDecoder(r.Body)
	var j IdInt
	err := decoder.Decode(&j)
	checkErr(err)
	io.WriteString(w, models.GetClientCars(j.Id)+"\n")
}

func getClientOrders(w http.ResponseWriter, r *http.Request) {
	fmt.Println("got getClientOrders request")
	decoder := json.NewDecoder(r.Body)
	var j IdInt
	err := decoder.Decode(&j)
	checkErr(err)
	io.WriteString(w, models.GetClientOrders(j.Id)+"\n")
}

func sendErrorJson(w http.ResponseWriter, r *http.Request, msg string) {
	ret := map[string]string{
		"errorStatus": msg,
	}
	r_, err := json.Marshal(ret)
	checkErr(err)
	io.WriteString(w, string(r_))
}
