package main

import (
	"database/sql"
	"encoding/csv"
	"fmt"
	_ "github.com/lib/pq"
	"log"
	"os"
	"strings"
)

func main() {
	db, err := sql.Open("postgres", "postgres://postgres:root@localhost:5432/server?sslmode=disable")
	if err != nil {
		log.Fatal(err)
	}
	defer db.Close()

	file, err := os.Open("class_table.csv")
	if err != nil {
		panic(err)
	}
	defer file.Close()

	reader := csv.NewReader(file)
	reader.FieldsPerRecord = 18
	reader.Comment = '#'
	count := 0
	for {
		count++
		record, e := reader.Read()
		if e != nil {
			fmt.Println("csv not open")

			break
		}
		if record[1] == "SDAT_S" {
			continue
		}

		newDate := strings.ReplaceAll(record[4], ".", "")
		name := record[6] + record[7] + newDate

		createTableQuery := fmt.Sprintf(`CREATE TABLE IF NOT EXISTS %s (FLT_NUM   VARCHAR(100),
    SEG_CLASS_CODE  VARCHAR(100),
    PASS_BK  VARCHAR(100),
    DTD VARCHAR(100));`, name)
		_, err = db.Exec(createTableQuery)
		if err != nil {
			fmt.Println("Ошибка при создании")
		}

		// Пропущены остальные поля, которые необходимо добавить в запрос INSERT.

		insertQuery := fmt.Sprintf(`INSERT INTO %s (FLT_NUM, SEG_CLASS_CODE, PASS_BK, DTD)
		VALUES ('%s', '%s', '%s', '%s')`,
			name, record[3], record[9], record[12], record[17])

		_, err = db.Exec(insertQuery)
		if err != nil {
			fmt.Println("Ошибка при инсерте")
		}
		fmt.Println(count)
		if count%1000000 == 0 {
			fmt.Println(count)
		}
	}
}
