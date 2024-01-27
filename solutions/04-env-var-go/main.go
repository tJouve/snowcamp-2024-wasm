package main

import (
	"fmt"
	"os"
)

func main() {
	fileName := os.Getenv("FILE_NAME")

	file, err := os.Create(fileName)
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
