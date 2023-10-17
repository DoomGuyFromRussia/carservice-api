package models

import (
	"database/sql"
	"encoding/json"
	"fmt"

	_ "github.com/mattn/go-sqlite3"
)

type Car struct {
	Id       int
	Producer string
	Model    string
	Year     string
	Vin      string
}

func (c *Car) Describe() string {
	return c.Producer + c.Model + c.Year + c.Vin
}

type Client struct {
	id      int
	name    string
	surname string
	address string
	phone   string
}

type Order struct {
	id          int
	carId       int
	clientId    int
	date        string
	description string
	status      string
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

func GetCar(id int) string {
	InitDb()
	row := Db.QueryRow("select * from cars where id=$1", id)
	var c Car
	var enc string
	switch err := row.Scan(&c.Id, &c.Producer, &c.Model, &c.Year, &c.Vin); err {
	case sql.ErrNoRows:
		fmt.Println("No rows were returned!")
		errMap := map[string]string{
			"error": "invalid id",
		}

		errJson, err := json.Marshal(errMap)
		checkErr(err)
		return string(errJson)
	default:
		//panic(err)
		b, err := json.Marshal(c)
		checkErr(err)
		enc = string(b)
	}
	//checkErr(err)
	defer Db.Close()
	return enc
}

func checkErr(err error) {
	if err != nil {
		panic(err)
	}
}
