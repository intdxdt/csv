package csv

import (
	"io"
	"fmt"
	"os"
	"log"
	"encoding/csv"
	"strings"
	"regexp"
)

//Parse csv file
func ParseCSV(fileName string, delim rune, hasHeading bool) []map[string]string {
	rows := make([]map[string]string, 0)
	header := make(map[int]string)
	file, err := os.Open(fileName)

	if err != nil {
		log.Println(err)
		return rows
	}
	defer file.Close()

	reader := csv.NewReader(file)
	reader.Comma = delim

	for {
		record, err := reader.Read()

		if err == io.EOF {
			break
		} else if err != nil {
			log.Println(err)
			return rows
		}

		if len(header) == 0 {
			header = MakeHeading(record, hasHeading)
			if hasHeading {
				continue
			}
		}
		rec := MakeRecord(header, record)
		if len(rec) > 0 {
			rows = append(rows, rec)
		}
	}
	return rows
}

//Make heading
func MakeHeading(record []string, hasHeading bool) map[int]string {
	sanitizer, err := regexp.Compile("[^a-zA-Z0-9]+")
	if err != nil {
		log.Fatalln("failed to compile regex sanitizer")
	}
	heading := make(map[int]string)

	if len(record) == 0 {
		return heading
	}

	for i := range record {
		record[i] = sanitizer.ReplaceAllString(record[i], "")
	}

	for r := range record {
		val := fmt.Sprintf("%v", strings.TrimSpace(record[r]))
		if hasHeading {
			val = record[r]
		}
		heading[r] = val
	}
	return heading
}

//Make record
func MakeRecord(header map[int]string, record []string) map[string]string {
	rec := make(map[string]string)
	row := record[:len(header)]
	for i := 0; i < len(row); i++ {
		rec[header[i]] = row[i]
	}
	return rec
}
