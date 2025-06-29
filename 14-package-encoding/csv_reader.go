package main

import (
	"encoding/csv"
	"fmt"
	"io"
	"strings"
)

func main() {
	csvString :=
		"name,age,city\n" +
			"Nafi,20,Jakarta\n" +
			"Furqon,21,Bekasi\n" +
			"Diani,22,Tangerang Selatan\n"

	reader := csv.NewReader(strings.NewReader(csvString))
	for {
		record, err := reader.Read()
		if err == io.EOF {
			break
		}

		fmt.Println(record)
	}
}
