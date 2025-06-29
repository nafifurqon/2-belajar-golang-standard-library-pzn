package main

import (
	"encoding/csv"
	"os"
)

func main() {
	writer := csv.NewWriter(os.Stdout)

	_ = writer.Write([]string{"name", "age", "city"})
	_ = writer.Write([]string{"Nafi", "20", "Jakarta"})
	_ = writer.Write([]string{"Furqon", "21", "Bekasi"})
	_ = writer.Write([]string{"Diani", "22", "Tangerang Selatan"})

	writer.Flush()
}
