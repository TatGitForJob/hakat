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

type SeasonInput struct {
	Direction string
	Class     string
	Number    string
	StartDate string
	EndDate   string
}

type SeasonClassInput struct {
	Direction string
}

type Seasoning struct {
	date  string
	Count int
}
type Sort_Seasons []Seasoning

func (s Sort_Seasons) Len() int {
	return len(s)
}

func (s Sort_Seasons) Swap(i, j int) {
	s[i], s[j] = s[j], s[i]
}

func (s Sort_Seasons) Less(i, j int) bool {
	return s[i].date < s[j].date // сравниваем в обратном порядке для сортировки по убыванию
}

func getSeasonClass(writer http.ResponseWriter, request *http.Request, params httprouter.Params) {
	var data SeasonClassInput
	err := json.NewDecoder(request.Body).Decode(&data)
	if err != nil {
		return
	}
	fmt.Println("Приняли дату для сезонов")
	fmt.Println(data.Direction)

	response := []string{}
	if data.Direction == "aersvo" {
		response = []string{"1117", "1119", "1121", "1123", "1125", "1127", "1129", "1131",
			"1133", "1135", "1137", "1139", "1141", "1151", "1153", "1741", "1771", "1773",
			"1781", "1783", "1785", "1787", "1789", "1791", "1793", "1795", "1797", "1799",
			"2957", "2981", "2985", "6180", "6182", "6186"}
	}
	if data.Direction == "svoaer" {
		response = []string{"1116", "1118", "1120", "1122", "1124", "1126", "1128", "1130",
			"1132", "1134", "1136", "1138", "1140", "1148", "1152", "1740", "1772", "1780", "1782", "1784", "1786", "1788", "1790", "1792", "1794", "1796", "1798", "2980", "2990", "6179", "6181", "6185"}
	}
	if data.Direction == "svoasf" {
		response = []string{"1172", "1174", "1642"}
	}
	if data.Direction == "asfsvo" {
		response = []string{"1173", "1175", "1643", "1775"}
	}

	writer.Header().Set("Content-Type", "application/json")
	_ = json.NewEncoder(writer).Encode(response)
	return
}

func getSeason(writer http.ResponseWriter, request *http.Request, params httprouter.Params) {
	var data SeasonInput
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

	db, err := sql.Open("postgres", "postgres://postgres:root@localhost:5432/postgres?sslmode=disable")
	if err != nil {
		log.Fatal(err)
	}
	defer db.Close()

	name := data.Direction + data.Number
	fmt.Println(name)
	insertQuery := fmt.Sprintf(`SELECT dd, pass_dep FROM %s WHERE seg_class_code = '%s'`, name, data.Class)
	rows, err := db.Query(insertQuery)
	if err != nil {
		fmt.Println(err)
		return
	}
	defer rows.Close()
	seasoning := make([]Seasoning, 0)
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

		ddNew := strings.ReplaceAll(dd, ".", "")
		ddNewNew := ddNew[4:] + "-" + ddNew[2:4] + "-" + ddNew[:2]

		seasoning = append(seasoning, Seasoning{Count: i, date: ddNewNew})

	}

	if len(seasoning) > 1 {
		sort.Sort(Sort_Seasons(seasoning))
	}

	canvas := make([]int, 0)
	datte := make([]string, 0)
	for i := 0; i < len(seasoning); i++ {
		if seasoning[i].date >= data.StartDate && seasoning[i].date <= data.EndDate {
			canvas = append(canvas, seasoning[i].Count)
			datte = append(datte, seasoning[i].date)
		}
	}
	da := struct {
		IntArray    []int
		StringArray []string
	}{
		IntArray:    canvas,
		StringArray: datte,
	}

	//fmt.Println(datte[0:10])
	fmt.Println("sssssssssssss")
	//fmt.Println(canvas[0:10])

	writer.Header().Set("Content-Type", "application/json")
	err = json.NewEncoder(writer).Encode(da)
	if err != nil {
		fmt.Println(err)
	}
	return
}
