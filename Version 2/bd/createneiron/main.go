package main

import (
	"database/sql"
	"encoding/csv"
	"fmt"
	_ "github.com/lib/pq"
	"log"
	"os"
)

func main() {
	db, err := sql.Open("postgres", "postgres://postgres:root@localhost:5432/neiron?sslmode=disable")
	if err != nil {
		log.Fatal(err)
	}
	defer db.Close()

	file, err := os.Open("202.csv")
	if err != nil {
		panic(err)
	}
	defer file.Close()

	reader := csv.NewReader(file)
	reader.FieldsPerRecord = 6
	reader.Comment = '#'
	count := 0
	for {
		count++
		record, e := reader.Read()
		if e != nil {
			fmt.Println("csv not open")
			break
		}
		name := record[2][:3] + record[3][:3] + record[1]

		createTableQuery := fmt.Sprintf(`CREATE TABLE IF NOT EXISTS %s
(   DD  VARCHAR(100),
    FLT_NUM VARCHAR(100),
    SDAT_S  VARCHAR(100),
    SORG  VARCHAR(100),
    SDST VARCHAR(100));`, name)
		_, err = db.Exec(createTableQuery)
		if err != nil {
			fmt.Println("Ошибка при создании")
		}

		// Пропущены остальные поля, которые необходимо добавить в запрос INSERT.
		if record[1] != "FLT_NUM" {
			insertQuery := fmt.Sprintf(`INSERT INTO %s (DD, FLT_NUM, SDAT_S, SORG, SDST)
		VALUES ('%s','%s','%s','%s','%s')`,
				name, record[4], record[1], record[5], record[2], record[3])

			_, err = db.Exec(insertQuery)
			if err != nil {
				fmt.Println("Ошибка при инсерте")
			}
		}
		if count%1000000 == 0 {
			fmt.Println(count)
		}
	}
}
