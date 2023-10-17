package models

import (
	"database/sql"
	"encoding/json"
	"fmt"

	_ "github.com/mattn/go-sqlite3"
)

type Error struct {
	Message string
}
type Car struct {
	Id       int `json:"id,string,omitempty"`
	Producer string
	Model    string
	Year     string
	Vin      string
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

func CreateCar(c Car) string {
	InitDb()
	res, err := Db.Exec("insert into cars (producer, model, year, vin) values ($1, $2, $3, $4)", c.Producer, c.Model, c.Year, c.Vin)
	checkErr(err)
	fmt.Println(res)
	r, err := json.Marshal(c)
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
func DeleteCar(id int) string {
	InitDb()
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

func checkErr(err error) {
	if err != nil {
		panic(err)
	}
}
