package main

/*
curl -X PUT \
     -H "Content-Type: application/json" \
     -d '{"email": "sam@email.com", "password": "password"}' \
     --cacert server.cert \
     https://192.168.1.37:8080/login


curl -H 'Authorization: Bearer <SESSION_KEY>' \
     --cacert server.cert \
     https://192.168.1.37:8080/<CARRIERCODE>


curl -X PUT \
		 -H 'Authorization: Bearer <SESSION_KEY>' \
		 -H 'Content-Type: application/json' \
		 -d '{"activate":true}' \
     --cacert server.cert \
     https://192.168.1.37:8080/<CARRIERCODE>

*/

import (
	"log"
	"net/http"
	"time"
)

func doEvery(d time.Duration) {
	for range time.Tick(d) {
		refreshRepo()
	}
}

func main() {

	router := NewRouter()

	go doEvery(time.Minute)

	log.Fatal(http.ListenAndServeTLS(":8080", serverCertPath, signingKeyPath, router))

	db.Close()
}
