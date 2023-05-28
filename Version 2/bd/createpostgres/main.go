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
	db, err := sql.Open("postgres", "postgres://postgres:root@localhost:5432/postgres?sslmode=disable")
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
		name := record[6] + record[7] + record[3]

		createTableQuery := fmt.Sprintf(`CREATE TABLE IF NOT EXISTS %s
(   DD  VARCHAR(100),
    SEG_CLASS_CODE  VARCHAR(100),
    PASS_DEP VARCHAR(100));`, name)
		_, err = db.Exec(createTableQuery)
		if err != nil {
			fmt.Println("Ошибка при создании")
		}

		// Пропущены остальные поля, которые необходимо добавить в запрос INSERT.
		if record[17] == "-1" {
			insertQuery := fmt.Sprintf(`INSERT INTO %s (DD, SEG_CLASS_CODE, PASS_DEP)
		VALUES ('%s','%s','%s')`,
				name, record[4], record[9], record[15])

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
