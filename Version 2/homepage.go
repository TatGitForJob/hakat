package main

import (
	"encoding/json"
	"fmt"
	"github.com/julienschmidt/httprouter"
	"html/template"
	"net/http"
)

type HomePage struct {
	Time string
}

type TimeDataInput struct {
	Direction string
	Date      string
	Class     string
	//Number    string
	StartDate string
	EndDate   string
}

type TimeDataOutput struct {
	Direction string
	Date      string
	Class     string
	//Number    string
	StartDate string
	EndDate   string
}

func getTime(writer http.ResponseWriter, request *http.Request, params httprouter.Params) {
	var data TimeDataInput
	err := json.NewDecoder(request.Body).Decode(&data)
	if err != nil {
		return
	}
	fmt.Println("\n" + data.Direction)
	fmt.Println("\n" + data.Date)
	fmt.Println("\n" + data.Class)
	//fmt.Println("\n" + data.Number)
	fmt.Println("\n" + data.StartDate)
	fmt.Println("\n" + data.EndDate)

	var responseData TimeDataOutput
	responseData.Direction = data.Direction
	responseData.Date = data.Date
	responseData.Class = data.Class
	responseData.StartDate = data.StartDate
	responseData.EndDate = data.EndDate
	writer.Header().Set("Content-Type", "application/json")
	_ = json.NewEncoder(writer).Encode(responseData)
	return
}

func serveHomepage(writer http.ResponseWriter, request *http.Request, params httprouter.Params) {
	var homepage HomePage
	tmpl := template.Must(template.ParseFiles("html/homepage.html"))
	_ = tmpl.Execute(writer, homepage)
}
