package main

import (
	"bufio"
	"encoding/csv"
	"fmt"
	"os"
)

func main() {
	file, err := os.Open("./tete1.csv")

	if err != nil {
		fmt.Println("@@@@@", err)
	}
	rdr := csv.NewReader(bufio.NewReader(file))

	rows, err := rdr.ReadAll()
	if err != nil {
		fmt.Println("@@@@@2", err)
	}
	fmt.Println("@@@@", rows)
	for i, row := range rows {
		for j := range row {
			fmt.Printf("%s", rows[i][j])
		}
		fmt.Println()
	}
}
