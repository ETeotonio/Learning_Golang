package main

import (
	"fmt"
	"os"
)

func main() {
	fmt.Println("2nd program in go")
	menu := 1
	for menu != 0 {
		fmt.Println("Start Menu")
		fmt.Println(" 1)Open a file\r\n 2)Get all the values from the file \r\n 3)Copy to a new file\r\n 0) Exit")
		var menu_item string
		fmt.Scanln(&menu_item)
		file_path := ""
		if menu_item != "0" {
			fmt.Println("Type the full file path:")
			fmt.Scan(&file_path)
		}
		switch menu_item {
		case "1":
			openFile(file_path)
		case "2":
			_ = readFile(file_path)
		case "3":
			copyFile(file_path)
		default:
			menu = 0
		}
	}
}

func openFile(file_path string) {
	file, err := os.Open(file_path)
	if err != nil {
		panic(err)
	}
	fmt.Println("File open and locked")
	file.Close()
}

func readFile(file_path string) []byte {
	fmt.Println("Opening the file and reading it")
	dat, err := os.ReadFile(file_path)
	if err != nil {
		panic(err)
	}
	fmt.Print(string(dat))
	return dat
}
func copyFile(file_path string) {
	fmt.Println("Opening the file and reading it")
	file_data := readFile(file_path)
	new_file_name := ""
	fmt.Println("Type the new file name")
	fmt.Scanln(&new_file_name)
	err := os.WriteFile(new_file_name, file_data, 0644)
	if err != nil {
		panic(err)
	}
	fmt.Println("File created")
}
