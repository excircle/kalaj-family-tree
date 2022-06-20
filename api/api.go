package main

import "log"
import "fmt"
import "net/http"
import "encoding/json"
//import "strconv"

import "database/sql"

import _ "github.com/go-sql-driver/mysql"
import "github.com/gorilla/mux"

type House struct {
    House   string  `json:"house"`
}

type Member struct {
    ID          int     `json:"id"`
    House_id    string  `json:"house_id"`
    Firstname   string  `json:"firstname"`
    Age         int     `json:"age"`
    Gender      string  `json:"gender"`
}

func getHouses(w http.ResponseWriter, r *http.Request) {
	var houses []House
    var house House

	db, err := sql.Open("mysql", "api:Catch22@tcp(127.0.0.1:3306)/famtree")
	if err != nil {
		log.Fatal(err)
	}

	defer db.Close()

	rows, err := db.Query("SELECT house FROM famtree.house")
	if err != nil {
		log.Fatal(err)
	}

	for rows.Next() {
		err = rows.Scan(&house.House)
		if err != nil {
			log.Fatal(err)
		}

		houses = append(houses, house)
	}
	json.NewEncoder(w).Encode(houses)
}

func submitHouse(w http.ResponseWriter, r *http.Request) {
    log.Println("submitHouse is called")

    var house House

    json.NewDecoder(r.Body).Decode(&house)

    db, err := sql.Open("mysql", "api:Catch22@tcp(127.0.0.1:3306)/famtree")
    if err != nil {
        log.Fatal(err)
    }

    defer db.Close()

    res, err := db.Exec(fmt.Sprintf("INSERT INTO famtree.house (id, house) VALUES(NULL, '%s')", house.House))
    if err != nil {
        log.Fatal(err)
    }

    rowCnt, err := res.RowsAffected()
    if err != nil {
        log.Fatal(err)
    }

    log.Printf("Wrote %d new record(s) - (HOUSE: '%s')", rowCnt, house.House)
}

func main() {
    r := mux.NewRouter()


    r.HandleFunc("/houses", getHouses).Methods("GET")
    r.HandleFunc("/houses", submitHouse).Methods("POST")


    log.Fatal(http.ListenAndServe(":8000", r))
}

