package main

import "net/http"

//Route holds the format of a navigatable route
type Route struct {
	Name        string
	Method      string
	Pattern     string
	HandlerFunc http.HandlerFunc
}

//Routes is a list of routes for the router to use
type Routes []Route

var routes = Routes{
	Route{
		"Login",
		"PUT",
		"/login",
		Login,
	},
	Route{
		"AirlineShow",
		"GET",
		"/{id}",
		ShowAirline,
	},
	Route{
		"AirlineToggle",
		"PUT",
		"/{id}",
		ActivateAirline,
	},
}
