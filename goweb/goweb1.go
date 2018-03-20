package main

import (
	"net/http"
)

type person struct {
	fName string
}

func (p *person) ServeHTTP(w http.ResponseWriter, r *http.Request)  {
	w.Write([]byte("First Name: " + p.fName))
}

func main()  {
	//personOne := &person{"xielchuan"}
	//http.ListenAndServe(":8080", personOne)

	personOne := &person{"xielchuan"}
	mux := http.NewServeMux()
	mux.Handle("/", personOne)
	http.ListenAndServe(":8080", mux)
}
