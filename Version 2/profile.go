package main

import (
	"database/sql"
	"encoding/json"
	"fmt"
	"github.com/julienschmidt/httprouter"
	"log"
	"net/http"
	"sort"
	"strconv"
	"strings"
	"time"
)

type ProfileInput struct {
	Direction string
	Class     string
	Number    string
	StartDate string
	EndDate   string
}

type Profiling struct {
	dd    int
	date  string
	count int
}
type Sort_Profiling []Profiling

func (s Sort_Profiling) Len() int {
	return len(s)
}

func (s Sort_Profiling) Swap(i, j int) {
	s[i], s[j] = s[j], s[i]
}

func (s Sort_Profiling) Less(i, j int) bool {
	return s[i].dd < s[j].dd // сравниваем в обратном порядке для сортировки по убыванию
}

func getProfile(writer http.ResponseWriter, request *http.Request, params httprouter.Params) {
	var data ProfileInput
	err := json.NewDecoder(request.Body).Decode(&data)
	if err != nil {
		return
	}
	fmt.Println("НАЖАЛИ НА ГРАФИК СЕЗОННОСТИ")
	fmt.Println(data.Direction)
	fmt.Println(data.Class)
	fmt.Println(data.Number)
	fmt.Println(data.StartDate)
	fmt.Println(data.EndDate)

	db, err := sql.Open("postgres", "postgres://postgres:root@localhost:5432/profiles?sslmode=disable")
	if err != nil {
		log.Fatal(err)
	}
	defer db.Close()

	name := data.Direction + data.Number
	fmt.Println(name)
	insertQuery := fmt.Sprintf(`SELECT dd, pass_dep  FROM %s WHERE seg_class_code = '%s'`, name, data.Class)
	rows, err := db.Query(insertQuery)
	if err != nil {
		http.Error(writer, err.Error(), http.StatusInternalServerError)
		fmt.Println(err)
		return
	}
	defer rows.Close()
	seasoning := make([]Profiling, 0)
	for rows.Next() {
		var dd, pass_dep string
		err := rows.Scan(&dd, &pass_dep)
		if err != nil {
			http.Error(writer, err.Error(), http.StatusInternalServerError)
			fmt.Println("ошибка при скане из данных базы")
			return
		}
		i, err := strconv.Atoi(pass_dep)
		if err != nil {
			fmt.Println("жопа с приведение к int")
		}
		layout := "2006-01-02"
		ddNew := strings.ReplaceAll(dd, ".", "")
		ddNewNew := ddNew[4:] + "-" + ddNew[2:4] + "-" + ddNew[:2]
		t1, _ := time.Parse(layout, ddNewNew)
		t2, _ := time.Parse(layout, data.StartDate)
		t3, _ := time.Parse(layout, data.EndDate)
		if int(t3.Sub(t1).Hours()/24) >= 0 && int(t1.Sub(t2).Hours()/24) >= 0 {
			fmt.Println(dd + " " + pass_dep)
			seasoning = append(seasoning, Seasoning{count: i, dd: int(t1.Sub(t2).Hours() / 24), date: ddNewNew})
		}
	}
	sort.Sort(Sort_Seasons(seasoning))
	fmt.Println(seasoning)

	writer.Header().Set("Content-Type", "application/json")
	_ = json.NewEncoder(writer).Encode(seasoning)
	return
}
