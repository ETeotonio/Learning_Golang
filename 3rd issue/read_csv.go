package main

import (
	"encoding/csv"
	"fmt"
	"io"
	"os"
)

func main() {
	fmt.Println("Type the path to the csv file")
	var file_path string
	fmt.Scanln(&file_path)
	//file_path := ""
	f, err := os.Open(file_path)

	if err != nil {
		panic(err)
	}

	csv_file := csv.NewReader(f)

	for {
		record, err := csv_file.Read()
		if err == io.EOF {
			break
		} else if err != nil {
			panic(err)
		}
		fmt.Println(record[0])
		// for value := range record {
		// 	fmt.Println(record[value])
		// }

	}
}
