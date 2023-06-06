package main

import (
	"database/sql"
	"encoding/json"
	"fmt"
	"github.com/julienschmidt/httprouter"
	"html/template"
	"log"
	"net/http"
	"sort"
	"strconv"
	"strings"
	"time"
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

type Dinamica struct {
	count string
	dtd   int
}
type Sorting []Dinamica

func (s Sorting) Len() int {
	return len(s)
}

func (s Sorting) Swap(i, j int) {
	s[i], s[j] = s[j], s[i]
}

func (s Sorting) Less(i, j int) bool {
	return s[i].dtd > s[j].dtd // сравниваем в обратном порядке для сортировки по убыванию
}

func ProcessDinamica(date, startdate, enddate string, dinamic []Dinamica) []int {
	result := make([]int, 0)
	layout := "2006-01-02"
	t1, _ := time.Parse(layout, date)
	t2, _ := time.Parse(layout, startdate)
	t3, _ := time.Parse(layout, enddate)
	fmt.Println(int(t1.Sub(t2).Hours() / 24))

	// Вычисление разницы в днях между датами
	daysstart := int(t1.Sub(t2).Hours() / 24)
	dayssend := int(t1.Sub(t3).Hours() / 24)
	sort.Sort(Sorting(dinamic))
	for _, item := range dinamic {
		if item.dtd <= daysstart && item.dtd >= dayssend {

			i, err := strconv.Atoi(item.count)
			if err != nil {
				fmt.Println("превращение в инт попизде")
			}
			result = append(result, i)
			fmt.Println(i)
		}
	}
	if len(result) <= 1 {
		return result
	}
	for i := len(result) - 1; i > 0; i-- {
		result[i] = result[i] - result[i-1]
	}

	return result
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

	db, err := sql.Open("postgres", "postgres://postgres:root@localhost:5432/server?sslmode=disable")
	if err != nil {
		log.Fatal(err)
	}
	defer db.Close()

	firstUpdateDate := strings.ReplaceAll(data.Date, "-", "")
	secondUpdateDate := firstUpdateDate[6:] + firstUpdateDate[4:6] + firstUpdateDate[:4]
	name := data.Direction + secondUpdateDate
	fmt.Println(name)
	insertQuery := fmt.Sprintf(`SELECT pass_bk, dtd  FROM %s WHERE flt_num = '%s' AND seg_class_code = '%s'`, name, data.Number, data.Class)
	rows, err := db.Query(insertQuery)
	if err != nil {
		http.Error(writer, err.Error(), http.StatusInternalServerError)
		fmt.Println(err)
		return
	}
	defer rows.Close()
	dinamica := make([]Dinamica, 0)
	for rows.Next() {
		var passbk, dtd string
		err := rows.Scan(&passbk, &dtd)
		if err != nil {
			http.Error(writer, err.Error(), http.StatusInternalServerError)
			fmt.Println("ошибка при скане из данных базы")
			return
		}
		fmt.Println(passbk + " " + dtd)
		i, err := strconv.Atoi(dtd)
		if err != nil {
			fmt.Println("жопа с приведение к int")
		}
		dinamica = append(dinamica, Dinamica{count: passbk, dtd: i})
	}
	aaa := ProcessDinamica(data.Date, data.EndDate, data.StartDate, dinamica)
	fmt.Println(aaa)

	writer.Header().Set("Content-Type", "application/json")
	_ = json.NewEncoder(writer).Encode(aaa)
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
func servePro(writer http.ResponseWriter, request *http.Request, params httprouter.Params) {
	tmpl := template.Must(template.ParseFiles("html/profil1.html"))
	err := tmpl.Execute(writer, nil)
	if err != nil {
		fmt.Println("не переводит на профил1")
	}
}
func serveProTwo(writer http.ResponseWriter, request *http.Request, params httprouter.Params) {
	tmpl := template.Must(template.ParseFiles("html/profil2.html"))
	err := tmpl.Execute(writer, nil)
	if err != nil {
		fmt.Println("не переводит на профил1")
	}
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
