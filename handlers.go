package main

import (
	"html/template"
	"net/http"
)

type ProductData struct {
	Products []Product
}

type Product struct {
	Name    string
	Picture string
	Price   string
}

func homeHandler(w http.ResponseWriter, r *http.Request) {
	// get the user's info and present his cart
	templ, err := template.ParseFiles("templates/home.gohtml")
	if err != nil {
		panic(err)
	}

	w.Header().Set("Content-Type", "text/html")
	pd := ProductData{[]Product{
		{Name: "ethiopian", Price: "$22"},
		{Name: "Guatemalan", Price: "$18"},
		{Name: "Honduran", Price: "$15"}}}
	err = templ.Execute(w, pd)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
	}
}
