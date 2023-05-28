package main

import (
	"encoding/json"
	"fmt"
	"github.com/julienschmidt/httprouter"
	"net/http"
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

type SeasonClassOutput struct {
	Rows []string `json:"rows"`
}

/*
type TimeDataOutput struct {
	Count []string `json:"count"`
	Date  []string `json:"date"`
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
*/

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

/*
func getSeason(writer http.ResponseWriter, request *http.Request, params httprouter.Params) {
	var data SeasonInput
	err := json.NewDecoder(request.Body).Decode(&data)
	if err != nil {
		return
	}
	fmt.Println("НАЖАЛИ НА ГРАФИК ДИНАМИКИ")
	fmt.Println(data.Direction)
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
*/
