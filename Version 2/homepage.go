package main

import (
	"encoding/json"
	"fmt"
	"github.com/julienschmidt/httprouter"
	"html/template"
	"net/http"
)

type User struct {
	Login    string
	Password string
}

type UserOutput struct {
	Text string `json:"message"`
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
	fmt.Println(data.Direction)
	fmt.Println(data.Date)
	fmt.Println(data.Class)
	//fmt.Println("\n" + data.Number)
	fmt.Println(data.StartDate)
	fmt.Println(data.EndDate)

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
	tmpl := template.Must(template.ParseFiles("html/homepage.html"))
	_ = tmpl.Execute(writer, nil)
}

func authHandler(writer http.ResponseWriter, request *http.Request, params httprouter.Params) {
	if request.Method == "POST" {
		var data User
		err := json.NewDecoder(request.Body).Decode(&data)
		if err != nil {
			return
		}
		fmt.Println(data.Login)
		fmt.Println(data.Password)

		var responseData UserOutput
		if data.Login == "admin@admin" && data.Password == "admin" {
			fmt.Println("true")
			responseData.Text = "Корректно"
		} else {
			responseData.Text = "Некорректно"
		}
		writer.Header().Set("Content-Type", "application/json")
		_ = json.NewEncoder(writer).Encode(responseData)
	} else if request.Method == "GET" {
		tmpl := template.Must(template.ParseFiles("html/auth.html"))
		_ = tmpl.Execute(writer, nil)
	}
}
