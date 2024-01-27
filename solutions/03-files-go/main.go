package main

import (
	"fmt"
	"os"
)

func main() {

	file, err := os.Create("hello.txt")
	if err != nil {
		fmt.Println("😡", err)
		return
	}
	defer file.Close()

	_, err = file.WriteString("👋 Hello, World! 🌍")
	if err != nil {
		fmt.Println("😡", err)
		return
	}

	data, err := os.ReadFile("hello.txt")
	if err != nil {
		fmt.Println("😡", err)
	}

	fmt.Print(string(data))

}
