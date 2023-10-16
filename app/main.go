package main

import (
	//"mux"

	"fmt"
	"io"
	"net/http"

	//"github.com/gorilla/mux"
	"database/sql"

	"github.com/gorilla/mux"
	_ "github.com/mattn/go-sqlite3"
)

func main() {
	db, err := sql.Open("sqlite3", "db.db")
	checkErr(err)
	fmt.Println(db)
	router := mux.NewRouter()
	router.HandleFunc("/api/test", getRoot)
	http.ListenAndServe(":8010", router)

}

func checkErr(err error) {
	if err != nil {
		panic(err)
	}
}

func getRoot(w http.ResponseWriter, r *http.Request) {
	fmt.Printf("got / request\n")
	io.WriteString(w, "test\n")
}
