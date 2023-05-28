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
	db, err := sql.Open("postgres", "postgres://postgres:root@localhost:5432/class?sslmode=disable")
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

		name := record[6] + record[7]

		createTableQuery := fmt.Sprintf(`CREATE TABLE IF NOT EXISTS %s
(   FLT_NUM   VARCHAR(100));`, name)
		_, err = db.Exec(createTableQuery)
		if err != nil {
			fmt.Println("Ошибка при создании")
		}

		// Пропущены остальные поля, которые необходимо добавить в запрос INSERT.

		insertQuery := fmt.Sprintf(`INSERT INTO %s (FLT_NUM)
		VALUES ('%s')`,
			name, record[3])

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
