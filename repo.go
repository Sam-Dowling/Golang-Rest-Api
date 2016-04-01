package main

import (
	"database/sql"
	"encoding/json"
	"fmt"
	_ "github.com/lib/pq"
	"net/http"
	"strconv"
	"time"
)

const (
	//DBUSER is the postgres username
	DBUSER = "default_user"
	//DBPASSWORD is the postgres password
	DBPASSWORD = "abc123"
	//DBNAME is the postgres database name
	DBNAME = "flights"
)

var airlines Airlines
var db *sql.DB

func init() {
	dbinfo := fmt.Sprintf("user=%s password=%s dbname=%s sslmode=disable",
		DBUSER, DBPASSWORD, DBNAME)
	conn, err := sql.Open("postgres", dbinfo)
	checkErr(err)
	db = conn

	refreshRepo()
}

func refreshRepo() {

	airlineRows, err := db.Query("SELECT id, full_name, is_activated, last_updated FROM airline")
	checkErr(err)

	airlines = Airlines{}

	for airlineRows.Next() {

		var id string
		var fullname string
		var isactivated bool
		var lastupdated time.Time

		err = airlineRows.Scan(&id, &fullname, &isactivated, &lastupdated)
		checkErr(err)
		RepoSetAirline(Airline{ID: id, FullName: fullname, IsActivated: isactivated, LastUpdated: lastupdated})
	}
}

func RepoOverwriteAirline(airline Airline) {
	stmt, err := db.Prepare("update airline set is_activated = $1 where id = $2")
	checkErr(err)

	_, err = stmt.Exec(strconv.FormatBool(airline.IsActivated), airline.ID)
	checkErr(err)

	RepoSetAirline(airline)

}

func RepoFindUser(email string) User {
	query := fmt.Sprintf("SELECT encrypted_password, password_salt FROM users where email = '%s'", email)

	var password string
	var salt string

	db.QueryRow(query).Scan(&password, &salt)

	return User{Password: password, Salt: salt}
}

//RepoFindAirline searches the airline repo for a given id
//and returns all records with that id
func RepoFindAirline(id string) Airline {
	return airlines[id]
}

func RepoSetAirline(airline Airline) {
	airlines[airline.ID] = airline
}

func readAirlineActivation(r *http.Request) (Activate, error) {
	var a Activate
	decoder := json.NewDecoder(r.Body)
	err := decoder.Decode(&a)
	return a, err
}

func readUser(r *http.Request) (User, error) {
	var u User
	decoder := json.NewDecoder(r.Body)
	err := decoder.Decode(&u)
	return u, err
}
