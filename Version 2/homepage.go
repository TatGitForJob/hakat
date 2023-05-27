package main

import (
	"database/sql"
	"encoding/json"
	"fmt"
	"github.com/julienschmidt/httprouter"
	"html/template"
	"log"
	"net/http"
	"strings"
)

type User struct {
	Login    string
	Password string
}

type UserOutput struct {
	Text string `json:"message"`
}

type Response struct {
	Rows []string `json:"rows"`
}

type ClassOut struct {
	Direction string
	Date      string
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
	Count []string `json:"count"`
	Date  []string `json:"date"`
}

func getTime(writer http.ResponseWriter, request *http.Request, params httprouter.Params) {
	var data TimeDataInput
	err := json.NewDecoder(request.Body).Decode(&data)
	if err != nil {
		return
	}
	fmt.Println("НАЖАЛИ НА ГРАФИК ДИНАМИКИ")
	fmt.Println(data.Direction)
	fmt.Println(data.Date)
	fmt.Println(data.Class)
	fmt.Println(data.Number)
	fmt.Println(data.StartDate)
	fmt.Println(data.EndDate)

	var response TimeDataOutput

	db, err := sql.Open("postgres", "postgres://postgres:root@localhost:5432/server?sslmode=disable")
	if err != nil {
		log.Fatal(err)
	}
	defer db.Close()

	firstUpdateDate := strings.ReplaceAll(data.Date, "-", "")
	secondUpdateDate := firstUpdateDate[6:] + firstUpdateDate[4:6] + firstUpdateDate[:4]
	name := data.Direction + secondUpdateDate
	fmt.Println(name)
	insertQuery := fmt.Sprintf(`SELECT sdat_s, pass_bk, dtd  FROM %s WHERE flt_num = '%s' AND seg_class_code = '%s'`, name, data.Number, data.Class)
	rows, err := db.Query(insertQuery)
	if err != nil {
		http.Error(writer, err.Error(), http.StatusInternalServerError)
		fmt.Println(err)
		return
	}
	defer rows.Close()

	for rows.Next() {
		var sdats, passbk, dtd string
		err := rows.Scan(&sdats, &passbk, &dtd)
		if err != nil {
			http.Error(writer, err.Error(), http.StatusInternalServerError)
			fmt.Println("ошибка при скане из данных базы")
			return
		}
		fmt.Println(sdats + " " + passbk + " " + dtd)
		response.Count = append(response.Count, passbk)
	}

	writer.Header().Set("Content-Type", "application/json")
	_ = json.NewEncoder(writer).Encode(response)
	return
}

func getClass(writer http.ResponseWriter, request *http.Request, params httprouter.Params) {
	var data ClassOut
	err := json.NewDecoder(request.Body).Decode(&data)
	if err != nil {
		return
	}
	fmt.Println("Приняли дату и направление")
	fmt.Println(data.Direction)
	fmt.Println(data.Date)
	db, err := sql.Open("postgres", "postgres://postgres:root@localhost:5432/server?sslmode=disable")
	if err != nil {
		log.Fatal(err)
	}
	defer db.Close()

	firstUpdateDate := strings.ReplaceAll(data.Date, "-", "")
	secondUpdateDate := firstUpdateDate[6:] + firstUpdateDate[4:6] + firstUpdateDate[:4]
	name := data.Direction + secondUpdateDate
	fmt.Println(name)
	insertQuery := fmt.Sprintf(`SELECT DISTINCT flt_num FROM %s`, name)
	rows, err := db.Query(insertQuery)
	if err != nil {
		http.Error(writer, err.Error(), http.StatusInternalServerError)
		fmt.Println("ошибка при запросе")
		return
	}
	defer rows.Close()

	var response Response
	for rows.Next() {
		var field3 string
		err := rows.Scan(&field3)
		if err != nil {
			http.Error(writer, err.Error(), http.StatusInternalServerError)
			fmt.Println("ошибка при скане из данных базы")
			return
		}
		fmt.Println(field3)
		response.Rows = append(response.Rows, field3)
	}

	writer.Header().Set("Content-Type", "application/json")
	_ = json.NewEncoder(writer).Encode(response)
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
