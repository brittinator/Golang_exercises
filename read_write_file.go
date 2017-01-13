package main

import (
	"bufio"
	"fmt"
	"io/ioutil"
	"os"
	"strings"
)

func create(path string) {
	_, err := os.Create(path)
	check(err)
	fmt.Printf("File created with path %s", path)
}

func read(path string) {
	file, err := ioutil.ReadFile(path)
	check(err)
	// print the bytes of the file
	fmt.Print(file)
	// print the content as a string
	fmt.Println(string(file))
}

func write(path string) {
	reader := bufio.NewReader(os.Stdin)
	fmt.Print("write")
	fmt.Print("What would you like to write?")
	toWrite, _ := reader.ReadString('\n')
	// checking if file already exists
	if _, err := os.Stat(path); os.IsNotExist(err) {
		create(path)
	}
	file, err := os.OpenFile(path, os.O_APPEND|os.O_RDWR, 0644)
	// err := ioutil.WriteFile(path, toWrite, 0644)
	check(err)
	defer file.Close()

	_, err = file.WriteString(toWrite)
	check(err)

	fmt.Printf("Wrote %s to %s", toWrite, path)
}

func check(e error) {
	if e != nil {
		fmt.Print(e.Error())
		panic(e)
	}
}

func main() {
	fmt.Print("What is the absolute filepath?")
	reader := bufio.NewReader(os.Stdin)
	path, _ := reader.ReadString('\n')
	path = strings.TrimSpace(path)
	fmt.Println("What would you like to do? Type \n 1) to read from a file, \n 2) to create a file \n 3) to write to a file.")
	choice, _ := reader.ReadString('\n')
	choice = strings.TrimSpace(choice)
	fmt.Print(choice)
	switch choice {
	case "1":
		read(path)
	case "2":
		create(path)
	case "3":
		write(path)
	default:
		fmt.Print("I don't understand")
	}
}
