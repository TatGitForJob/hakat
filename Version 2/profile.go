package main

import (
	"database/sql"
	"encoding/json"
	"fmt"
	"github.com/360EntSecGroup-Skylar/excelize"
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

type SecondSort []Profiling

func (s SecondSort) Len() int {
	return len(s)
}

func (s SecondSort) Swap(i, j int) {
	s[i], s[j] = s[j], s[i]
}

func (s SecondSort) Less(i, j int) bool {
	return s[i].sdat_s < s[j].sdat_s
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
	profilingU := make([]Profiling, 0)
	profilingB := make([]Profiling, 0)
	for i := 0; i <= len(profiling)-1; i++ {
		if profiling[i].dtd < 30 {
			profilingB = append(profilingB, profiling[i])
		} else {
			profilingU = append(profilingU, profiling[i])
		}
	}
	sort.Sort(SecondSort(profilingB))
	sort.Sort(SecondSort(profilingU))
	for i := 1; i <= len(profilingB)-1; i++ {
		if profilingB[i].sdat_s == profilingB[i-1].sdat_s {
			profilingB[i].pass_bk += profilingB[i-1].pass_bk
		}
	}
	for i := 1; i <= len(profilingU)-1; i++ {
		if profilingU[i].sdat_s == profilingU[i-1].sdat_s {
			profilingU[i].pass_bk += profilingU[i-1].pass_bk
		}
	}
	profilingUU := make([]Profiling, 0)
	profilingBB := make([]Profiling, 0)
	for i := 1; i <= len(profilingB)-1; i++ {
		if profilingB[i].sdat_s >= data.StartDate && profilingB[i].sdat_s <= data.EndDate && (i == len(profilingB)-1 || profilingB[i].sdat_s != profilingB[i+1].sdat_s) {
			profilingBB = append(profilingBB, profilingB[i])
		}
	}
	for i := 1; i <= len(profilingU)-1; i++ {
		if profilingU[i].sdat_s >= data.StartDate && profilingU[i].sdat_s <= data.EndDate && (i == len(profilingU)-1 || profilingU[i].sdat_s != profilingU[i+1].sdat_s) {
			profilingUU = append(profilingUU, profilingU[i])
		}
	}

	otArr := make([]int, 0)
	raArr := make([]int, 0)
	strArr := make([]string, 0)
	for i := 1; i <= len(profilingBB)-1; i++ {
		raArr = append(raArr, profilingBB[i].pass_bk)
		strArr = append(strArr, profilingBB[i].sdat_s)
	}
	for i := 1; i <= len(profilingUU)-1; i++ {
		otArr = append(otArr, profilingUU[i].pass_bk)
	}

	fmt.Println(len(profilingBB))
	fmt.Println(len(profilingUU))

	dad := struct {
		OtpuskArray []int
		RabotaArray []int
		StringArray []string
	}{
		OtpuskArray: otArr,
		RabotaArray: raArr,
		StringArray: strArr,
	}

	file := excelize.NewFile()

	file.SetCellValue("Sheet1", "A1", "Даты")
	file.SetCellValue("Sheet1", "B1", "Командировка")
	file.SetCellValue("Sheet1", "C1", "Отпуск")
	for i := 0; i < len(raArr); i++ {
		file.SetCellValue("Sheet1", "A"+strconv.Itoa(i+2), dad.StringArray[i])
		file.SetCellValue("Sheet1", "B"+strconv.Itoa(i+2), dad.RabotaArray[i])
	}
	for i := 0; i < len(otArr); i++ {
		file.SetCellValue("Sheet1", "C"+strconv.Itoa(i+2), dad.OtpuskArray[i])
	}

	filePath := "excel/profile.xlsx"
	if err := file.SaveAs(filePath); err != nil {
		fmt.Println("Ошибка сохранения файла:", err)
		writer.WriteHeader(http.StatusInternalServerError)
		return
	}

	writer.Header().Set("Content-Type", "application/json")
	_ = json.NewEncoder(writer).Encode(dad)
	return
}

func downloadProfile(w http.ResponseWriter, r *http.Request, params httprouter.Params) {

	// Отправка файла в ответе
	w.Header().Set("Content-Type", "application/vnd.openxmlformats-officedocument.spreadsheetml.sheet")
	w.Header().Set("Content-Disposition", "attachment; filename=profile.xlsx")
	http.ServeFile(w, r, "excel/profile.xlsx")
}
