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
)

type ProfileInput struct {
	Direction string
	Class     string
	Number    string
	StartDate string
	EndDate   string
}

type Profiling struct {
	dd      string
	sdat_s  string
	pass_bk int
	dtd     int
}
type Output struct {
	date  string
	count int
	color string
}

type FirstSort []Profiling

func (s FirstSort) Len() int {
	return len(s)
}

func (s FirstSort) Swap(i, j int) {
	s[i], s[j] = s[j], s[i]
}

func (s FirstSort) Less(i, j int) bool {
	if s[i].dd == s[j].dd {
		return s[i].dtd > s[j].dtd
	}
	return s[i].dd < s[j].dd
}

func GetTruePassBk(prof []Profiling) []Profiling {
	for i := len(prof) - 1; i > 0; i-- {
		if prof[i].dd == prof[i-1].dd {
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
	insertQuery := fmt.Sprintf(`SELECT dd, pass_bk, sdat_s, dtd  FROM %s WHERE seg_class_code = '%s'`, name, data.Class)
	rows, err := db.Query(insertQuery)
	if err != nil {
		http.Error(writer, err.Error(), http.StatusInternalServerError)
		fmt.Println(err)
		return
	}
	defer rows.Close()
	profiling := make([]Profiling, 0)
	for rows.Next() {
		var dd, pass_bk, sdat_s, dtd string
		err := rows.Scan(&dd, &pass_bk, &sdat_s, &dtd)
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
		sdateNew := strings.ReplaceAll(sdat_s, ".", "")
		sdateNewNew := sdateNew[4:] + "-" + sdateNew[2:4] + "-" + sdateNew[:2]
		ddNew := strings.ReplaceAll(dd, ".", "")
		ddNewNew := ddNew[4:] + "-" + ddNew[2:4] + "-" + ddNew[:2]
		profiling = append(profiling, Profiling{dd: ddNewNew, sdat_s: sdateNewNew, pass_bk: j, dtd: i})
	}
	sort.Sort(FirstSort(profiling))
	profiling = GetTruePassBk(profiling)
	/*
		for i := len(profiling) - 1; i >= 0; i-- {
			if profiling[i].dtd < 30 {
				profilingB = append(profilingB, profiling[i])
			} else {
				profilingU = append(profilingU, profiling[i])
			}
		}
	*/
	dad := struct {
		OtpuskArray []int
		RabotaArray []int
		StringArray []string
	}{
		OtpuskArray: []int{1, 1, 1, 1, 1, 1, 1, 1, 1, 1, 1, 113, 1, 1, 12, 1, 1, 1, 1, 1, 1, 1, 1, 1, 1},
		RabotaArray: []int{4, 4, 5, 4, 4, 4, 10, 4, 4, 4, 4, 4, 4, 4, 4, 4, 4, 4, 4, 4, 4, 4, 4, 4, 4},
		StringArray: []string{"1day", "2day", "3day", "4day", "5day", "1day", "2day", "3day", "4day", "5day", "1day", "2day", "3day", "4day", "5day", "1day", "2day", "3day", "4day", "5day", "1day", "2day", "3day", "4day", "5day"},
	}
	writer.Header().Set("Content-Type", "application/json")
	_ = json.NewEncoder(writer).Encode(dad)
	return
}
