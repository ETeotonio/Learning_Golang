package main

import (
	"fmt"
	"regexp"
)

func main() {

	email_addr := ""
	fmt.Println(email_addr)
	i := 0
	for email_addr != "exit" {
		fmt.Println("Type here your email or type exit to close the program")
		fmt.Scanln(&email_addr)

		if email_addr != "exit" {
			validate_Email(email_addr)
		}
		i++
	}
	fmt.Print("I have run ", i, " times")
}
func validate_Email(email_addr string) {
	regex_val := "^[\\w-\\.]+@([\\w-]+\\.)+[\\w-]{2,4}$"
	match, _ := regexp.MatchString(regex_val, string(email_addr))

	if match {
		fmt.Println("Email matches")
	} else {
		fmt.Println("Email not matching")
	}
}
