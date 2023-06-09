package main

import (
	"encoding/json"
	"fmt"
	"github.com/360EntSecGroup-Skylar/excelize"
	"github.com/julienschmidt/httprouter"
	"net/http"
)

type PredictInput struct {
	Direction string
	Class     string
	Number    string
	StartDate string
	EndDate   string
}

func getPredict(writer http.ResponseWriter, request *http.Request, params httprouter.Params) {
	var data PredictInput
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

	file := excelize.NewFile()

	file.SetCellValue("Sheet1", "A1", "Даты")
	file.SetCellValue("Sheet1", "B1", "Командировка")
	file.SetCellValue("Sheet1", "C1", "Отпуск")

	filePath := "excel/predict.xlsx"
	if err := file.SaveAs(filePath); err != nil {
		fmt.Println("Ошибка сохранения файла:", err)
		writer.WriteHeader(http.StatusInternalServerError)
		return
	}
	da := struct {
		IntArray    []int
		StringArray []string
	}{
		IntArray:    []int{1, 2, 3, 4, 5, 6, 7, 8, 9, 10},
		StringArray: []string{"10", "9", "8", "7", "6", "5", "4", "3", "2", "1"},
	}

	writer.Header().Set("Content-Type", "application/json")
	_ = json.NewEncoder(writer).Encode(da)
	return
}

func downloadPredict(w http.ResponseWriter, r *http.Request, params httprouter.Params) {

	// Отправка файла в ответе
	w.Header().Set("Content-Type", "application/vnd.openxmlformats-officedocument.spreadsheetml.sheet")
	w.Header().Set("Content-Disposition", "attachment; filename=predict.xlsx")
	http.ServeFile(w, r, "excel/predict.xlsx")
}
