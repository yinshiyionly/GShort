package main

import (
	"fmt"
	"os"
)

func main() {

	file, err := os.Open("books.xml")
	if err != nil {
		fmt.Printf("error: %v", err)
		return
	}
	defer file.Close()
	//fmt.Println(file)
	fmt.Printf("%T", file)
	//v := BookStore{}
	//err = xml.Unmarshal(data, &v)
	//if err != nil {
	//	fmt.Printf("error: %v", err)
	//	return
	//}

	//fmt.Println(v)
}
