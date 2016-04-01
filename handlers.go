package main

import (
	"encoding/json"
	"net/http"

	"github.com/gorilla/mux"
)

//Login takes the email and password supplied by the user,
//if authentics returns a session token
func Login(w http.ResponseWriter, r *http.Request) {
	u, err := readUser(r)
	if err != nil {
		writeError(w, http.StatusBadRequest, jsonErr{Code: 404, Text: "Invalid arguments"})
		return
	}
	user := RepoFindUser(u.Email)
	if (User{}) == user || user.Password != sha256Hash(u.Password, user.Salt) {
		writeError(w, http.StatusBadRequest, jsonErr{Code: 404, Text: "Invalid Email or Password"})
		return
	}
	sessionID := startSession()
	tokenString, err := generateToken(sessionID)
	if err != nil {
		writeError(w, http.StatusInternalServerError, jsonErr{Code: 404, Text: "Error Creating Token"})
		return
	}
	err = writeJSON(w, http.StatusOK, SessionTokenResponse{Token: tokenString})
	if err != nil {
		writeError(w, http.StatusInternalServerError, jsonErr{Code: 404, Text: "Unable to craft response"})
	}
}

//ShowAirline returns a list of fleet data for a given airline carriercode
func ShowAirline(w http.ResponseWriter, r *http.Request) {

	token, err := readToken(r)

	if err != nil || !isValidToken(token) {
		writeError(w, http.StatusBadRequest, jsonErr{Code: 404, Text: "Invalid or expired Token"})
	} else {

		vars := mux.Vars(r)
		var id string
		id = vars["id"]
		airline := RepoFindAirline(id)
		if airline != (Airline{}) {
			w.Header().Set("Content-Type", "application/json; charset=UTF-8")
			w.WriteHeader(http.StatusOK)
			if err := json.NewEncoder(w).Encode(airline); err != nil {
				panic(err)
			}
		} else {
			return404(w)
		}
	}
}

//ActivateAirline toggles the IsActivated for an airline
func ActivateAirline(w http.ResponseWriter, r *http.Request) {

	token, err := readToken(r)

	if err != nil || !isValidToken(token) {
		writeError(w, http.StatusBadRequest, jsonErr{Code: 404, Text: "Invalid or expired Token"})
	} else {

		airlineUpdate, err := readAirlineActivation(r)

		if err != nil {
			writeError(w, http.StatusBadRequest, jsonErr{Code: 404, Text: "Invalid arguments"})
			return
		}

		vars := mux.Vars(r)
		var id string
		id = vars["id"]
		airline := RepoFindAirline(id)

		if airline != (Airline{}) {
			airline.IsActivated = airlineUpdate.Activate
			RepoOverwriteAirline(airline)
			w.Header().Set("Content-Type", "application/json; charset=UTF-8")
			w.WriteHeader(http.StatusOK)
			if err := json.NewEncoder(w).Encode(airline); err != nil {
				panic(err)
			}
		} else {
			// 404
			return404(w)
		}
	}
}

func return404(w http.ResponseWriter) {
	w.Header().Set("Content-Type", "application/json; charset=UTF-8")
	w.WriteHeader(http.StatusNotFound)
	if err := json.NewEncoder(w).Encode(jsonErr{Code: http.StatusNotFound, Text: "Airline Not Found"}); err != nil {
		panic(err)
	}
}
