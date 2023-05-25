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
	Number    string
	StartDate string
	EndDate   string
}

type TimeDataOutput struct {
	Count string
	Date  string
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
	fmt.Println(data.Number)
	fmt.Println(data.StartDate)
	fmt.Println(data.EndDate)

	var responseData TimeDataOutput
	responseData.Count = "даты здесь"
	responseData.Date = "дата тут"
	/*
		db, err := sql.Open("postgres", "postgres://postgres:root@localhost:5432/server?sslmode=disable")
		if err != nil {
			log.Fatal(err)
		}
		defer db.Close()

		var query string
		var args []interface{}
		query = "SELECT id, name, cit, univ, dates, format, link, photo FROM %s"

		rows, err := db.Query(query, args...)
		if err != nil {
			http.Error(w, err.Error(), http.StatusInternalServerError)
			return
		}

		var events []Event
		for rows.Next() {
			var u Event
			err := rows.Scan(&u.ID, &u.Name, &u.City, &u.Univ, &u.Dates, &u.Format, &u.Link, &u.Photo)
			if err != nil {
				http.Error(w, err.Error(), http.StatusInternalServerError)
				return
			}
			events = append(events, u)
		}
		err = rows.Err()
		if err != nil {
			http.Error(w, err.Error(), http.StatusInternalServerError)
			return
		}
	*/

	writer.Header().Set("Content-Type", "application/json")
	_ = json.NewEncoder(writer).Encode(responseData)
	return
}

func serveHomepage(writer http.ResponseWriter, request *http.Request, params httprouter.Params) {
	tmpl := template.Must(template.ParseFiles("html/homepage.html"))
	_ = tmpl.Execute(writer, nil)
}
func serveSeasons(writer http.ResponseWriter, request *http.Request, params httprouter.Params) {
	tmpl := template.Must(template.ParseFiles("html/seasons.html"))
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
