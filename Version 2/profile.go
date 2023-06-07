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
	sdat_s  string
	pass_bk int
	dtd     int
}
type Output struct {
	date  string
	count int
	color string
}

type Sort_Profiling []Profiling

func (s Sort_Profiling) Len() int {
	return len(s)
}

func (s Sort_Profiling) Swap(i, j int) {
	s[i], s[j] = s[j], s[i]
}

func (s Sort_Profiling) Less(i, j int) bool {
	if s[i].sdat_s == s[j].sdat_s {
		return s[i].dtd > s[j].dtd
	}
	return s[i].sdat_s < s[j].sdat_s
}

func GetTruePass(prof []Profiling) []Profiling {
	for i := len(prof) - 1; i > 0; i-- {
		if prof[i].sdat_s == prof[i-1].sdat_s {
			prof[i].pass_bk = prof[i].pass_bk - prof[i-1].pass_bk
		}
	}
	return prof
}
func GetResult(prof []Profiling) []Output {
	count := make([]Output, 0)
	for i := len(prof) - 1; i >= 0; i-- {
		if prof[i].sdat_s == prof[i-1].sdat_s {
			prof[i].pass_bk = prof[i].pass_bk - prof[i-1].pass_bk
		}
	}
	return prof
}

func getProfile(writer http.ResponseWriter, request *http.Request, params httprouter.Params) {
	var data ProfileInput
	err := json.NewDecoder(request.Body).Decode(&data)
	if err != nil {
		return
	}
	fmt.Println("НАЖАЛИ НА ГРАФИК ПРОФИЛЕЙ")
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
	insertQuery := fmt.Sprintf(`SELECT pass_bk, sdat_s, dtd  FROM %s WHERE seg_class_code = '%s'`, name, data.Class)
	rows, err := db.Query(insertQuery)
	if err != nil {
		http.Error(writer, err.Error(), http.StatusInternalServerError)
		fmt.Println(err)
		return
	}
	defer rows.Close()
	profilingU := make([]Profiling, 0)
	profilingB := make([]Profiling, 0)
	for rows.Next() {
		var pass_bk, sdat_s, dtd string
		err := rows.Scan(&pass_bk, &sdat_s, &dtd)
		if err != nil {
			http.Error(writer, err.Error(), http.StatusInternalServerError)
			fmt.Println("ошибка при скане из данных базы")
			return
		}
		i, err := strconv.Atoi(dtd)
		if err != nil {
			fmt.Println("жопа с приведение к int")
		}
		j, err := strconv.Atoi(pass_bk)
		if err != nil {
			fmt.Println("жопа с приведение к int")
		}
		layout := "2006-01-02"
		ddNew := strings.ReplaceAll(sdat_s, ".", "")
		ddNewNew := ddNew[4:] + "-" + ddNew[2:4] + "-" + ddNew[:2]
		t1, _ := time.Parse(layout, ddNewNew)
		t2, _ := time.Parse(layout, data.StartDate)
		t3, _ := time.Parse(layout, data.EndDate)
		if int(t3.Sub(t1).Hours()/24) >= 0 && int(t1.Sub(t2).Hours()/24) >= 0 {
			profiling = append(profiling, Profiling{dd: dd, sdat_s: sdat_s, pass_bk: j, dtd: i})
		}
	}
	fmt.Println(profiling)
	sort.Sort(Sort_Profiling(profiling))
	fmt.Println(profiling)
	profiling = GetTruePass(profiling)
	for i := len(profiling) - 1; i >= 0; i-- {
		if profiling[i].dtd < 30 {
			profilingB = append(profilingB, profiling[i])
		} else {
			profilingU = append(profilingU, profiling[i])
		}
	}
	writer.Header().Set("Content-Type", "application/json")
	_ = json.NewEncoder(writer).Encode(profiling)
	return
}
