package main

import (
	"fmt"
	"io/ioutil"
	"log"
	"os"
)

func main() {
	createFile()
	readFile()
	removeFile()
}

func createFile() {
	fmt.Println("creating file.....")

	file, err := os.Create("test.txt")

	if err != nil {
		log.Fatalf("error creating file %s", err)
	}

	defer file.Close()

	len, err := file.WriteString("Hi porus how are you doing in GOlang?")

	if err != nil {
		log.Fatalf("eroor while writing to file %s", err)
	}

	fmt.Println("file name ", file.Name())
	fmt.Printf("size of the file %d bytes ", len)

}

func readFile() {
	fmt.Println("reading file......")

	data, err := ioutil.ReadFile("test.txt")

	if err != nil {
		log.Fatalf("error while reading file %s", err)
	}

	//fmt.Println("File name is : ",fileName)
	fmt.Printf("\nfile size is %d bytes ", len(data))
	fmt.Printf("\ndata in file is : %s ", data)
}

func removeFile() {
	fmt.Println("\nremoving file.....")

	err := os.Remove("test.txt")

	if err != nil {
		fmt.Printf("error while deleting the file %s", err)
	}

	fmt.Println("\nfile deleted successfully")
}
