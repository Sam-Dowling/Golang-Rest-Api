package main

import "net/http"

type jsonErr struct {
	Code int    `json:"code"`
	Text string `json:"text"`
}

func writeError(w http.ResponseWriter, status int, response jsonErr) error {
	return writeJSON(w, status, response)
}

func checkErr(err error) {
	if err != nil {
		panic(err)
	}
}
