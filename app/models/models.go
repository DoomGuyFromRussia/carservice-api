package models

import (
	"database/sql"
	"encoding/json"
	"fmt"
	"strings"

	_ "github.com/mattn/go-sqlite3"
)

type Car struct {
	Id       int `json:"id,string,omitempty"`
	Producer string
	Model    string
	Year     string
	Vin      string
}

func (c *Car) Validate() bool {
	if strings.TrimSpace(c.Model) == "" || strings.TrimSpace(c.Producer) == "" || strings.TrimSpace(c.Vin) == "" || strings.TrimSpace(c.Year) == "" {
		return false
	} else {
		return true
	}
}

type Client struct {
	Id      int `json:"id,string,omitempty"`
	Name    string
	Surname string
	Address string
	Phone   string
}

func (c *Client) Validate() bool {
	if strings.TrimSpace(c.Name) == "" || strings.TrimSpace(c.Surname) == "" || strings.TrimSpace(c.Address) == "" || strings.TrimSpace(c.Phone) == "" {
		return false
	} else {
		return true
	}
}

type Order struct {
	Id          int `json:"id,string,omitempty"`
	CarId       int
	ClientId    int
	Date        string
	Description string
	Status      string
}

func (o *Order) Validate() bool {
	if strings.TrimSpace(o.Date) == "" || strings.TrimSpace(o.Description) == "" || strings.TrimSpace(o.Status) == "" {
		return false
	} else {
		return true
	}
}

var Db *sql.DB
var err error

func InitDb() {
	Db, err = sql.Open("sqlite3", "db.db")
	checkErr(err)

}

func GetCars() string {
	InitDb()
	rows, err := Db.Query("select * from cars")
	checkErr(err)
	var cars []Car
	for rows.Next() {
		var c Car
		err := rows.Scan(&c.Id, &c.Producer, &c.Model, &c.Year, &c.Vin)
		checkErr(err)
		cars = append(cars, c)
	}
	defer Db.Close()
	b, err := json.Marshal(cars)
	enc := string(b)
	checkErr(err)
	return enc
}

func GetClients() string {
	InitDb()
	rows, err := Db.Query("select * from clients")
	checkErr(err)
	var clients []Client
	for rows.Next() {
		var c Client
		err := rows.Scan(&c.Id, &c.Name, &c.Surname, &c.Address, &c.Phone)
		checkErr(err)
		clients = append(clients, c)
	}
	defer Db.Close()
	b, err := json.Marshal(clients)
	enc := string(b)
	checkErr(err)
	return enc
}

func GetOrders() string {
	InitDb()
	rows, err := Db.Query("select * from orders")
	checkErr(err)
	var orders []Order
	for rows.Next() {
		var o Order
		err := rows.Scan(&o.Id, &o.CarId, &o.ClientId, &o.Date, &o.Description, &o.Status)
		checkErr(err)
		orders = append(orders, o)
	}
	defer Db.Close()
	b, err := json.Marshal(orders)
	enc := string(b)
	checkErr(err)
	return enc
}

func CreateCar(c Car) string {
	InitDb()
	res, err := Db.Exec("insert into cars (producer, model, year, vin) values ($1, $2, $3, $4)", c.Producer, c.Model, c.Year, c.Vin)
	checkErr(err)
	fmt.Println(res)
	r, err := json.Marshal(c)
	checkErr(err)
	return string(r)
}

func CreateClient(c Client) string {
	InitDb()
	res, err := Db.Exec("insert into clients (name, surname, address, phone) values ($1, $2, $3, $4)", c.Name, c.Surname, c.Address, c.Phone)
	checkErr(err)
	fmt.Println(res)
	r, err := json.Marshal(c)
	checkErr(err)
	return string(r)
}

func CreateOrder(o Order) string {
	InitDb()
	fmt.Println("Desc", o.Description)
	res, err := Db.Exec("insert into orders (carId, clientId, date, description, status) values ($1, $2, $3, $4, $5)", o.CarId, o.ClientId, o.Date, o.Description, o.Status)
	checkErr(err)
	fmt.Println(res)
	r, err := json.Marshal(o)
	checkErr(err)
	return string(r)
}

func UpdateCar(c Car) string {
	InitDb()
	res, err := Db.Exec("update cars set producer = $1, model = $2, year = $3, vin = $4 where id = $5", c.Producer, c.Model, c.Year, c.Vin, c.Id)
	checkErr(err)
	fmt.Println(res.RowsAffected())
	ret := map[string]string{
		"result": "ok",
	}
	r, err := json.Marshal(ret)
	checkErr(err)
	return string(r)
}

func UpdateClient(c Client) string {
	InitDb()
	res, err := Db.Exec("update clients set name = $1, surname = $2, address = $3, phone = $4 where id = $5", c.Name, c.Surname, c.Address, c.Phone, c.Id)
	checkErr(err)
	fmt.Println(res.RowsAffected())
	ret := map[string]string{
		"result": "ok",
	}
	r, err := json.Marshal(ret)
	checkErr(err)
	return string(r)
}

func UpdateOrder(o Order) string {
	InitDb()
	res, err := Db.Exec("update orders set carId = $1, clientId = $2, date = $3, description = $4, status = $5 where id = $6", o.CarId, o.ClientId, o.Date, o.Description, o.Status, o.Id)
	checkErr(err)
	fmt.Println(res.RowsAffected())
	ret := map[string]string{
		"result": "ok",
	}
	r, err := json.Marshal(ret)
	checkErr(err)
	return string(r)
}

func DeleteCar(id int) string {
	InitDb()
	fmt.Println(id)
	res, err := Db.Exec("delete from cars where id = $1", id)
	checkErr(err)
	fmt.Println(res.RowsAffected())
	ret := map[string]string{
		"result": "ok",
	}
	r, err := json.Marshal(ret)
	checkErr(err)
	return string(r)
}

func DeleteClient(id int) string {
	InitDb()
	res, err := Db.Exec("delete from clients where id = $1", id)
	checkErr(err)
	fmt.Println(res.RowsAffected())
	ret := map[string]string{
		"result": "ok",
	}
	r, err := json.Marshal(ret)
	checkErr(err)
	return string(r)
}

func DeleteOrder(id int) string {
	InitDb()
	res, err := Db.Exec("delete from orders where id = $1", id)
	checkErr(err)
	fmt.Println(res.RowsAffected())
	ret := map[string]string{
		"result": "ok",
	}
	r, err := json.Marshal(ret)
	checkErr(err)
	return string(r)
}

func GetClientCars(id int) string {
	InitDb()
	rows, err := Db.Query("select * from clientsCars where clientId = $1", id)
	checkErr(err)
	var carId, clientId int
	var cars []Car
	for rows.Next() {
		err := rows.Scan(&clientId, &carId)
		checkErr(err)
		rowsCar, err := Db.Query("select * from cars where id = $1", carId)
		checkErr(err)
		for rowsCar.Next() {
			var c Car
			err := rowsCar.Scan(&c.Id, &c.Producer, &c.Model, &c.Year, &c.Vin)
			checkErr(err)
			cars = append(cars, c)
		}
	}
	defer Db.Close()
	b, err := json.Marshal(cars)
	enc := string(b)
	checkErr(err)
	if enc == "null" {
		return sendErrorJson("db null error")
	} else {
		return enc
	}

}

func GetClientOrders(id int) string {
	InitDb()
	rows, err := Db.Query("select * from clientsOrders where clientId = $1", id)
	checkErr(err)
	var orderId, clientId int
	var orders []Order
	for rows.Next() {
		err := rows.Scan(&clientId, &orderId)
		checkErr(err)
		rowsOrder, err := Db.Query("select * from orders where id = $1", orderId)
		checkErr(err)
		for rowsOrder.Next() {
			var o Order
			err := rowsOrder.Scan(&o.Id, &o.CarId, &o.ClientId, &o.Date, &o.Description, &o.Status)
			checkErr(err)
			orders = append(orders, o)
		}
	}
	defer Db.Close()
	b, err := json.Marshal(orders)
	enc := string(b)
	checkErr(err)
	if enc == "null" {
		return sendErrorJson("db null error")
	} else {
		return enc
	}
}

func checkErr(err error) {
	if err != nil {
		panic(err)
	}
}

func sendErrorJson(msg string) string {
	ret := map[string]string{
		"errorStatus": msg,
	}
	r_, err := json.Marshal(ret)
	checkErr(err)
	return string(r_)
}
