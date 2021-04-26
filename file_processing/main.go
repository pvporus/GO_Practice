package main

import (
	"fmt"
	"io/ioutil"
	"log"
	"os"
)

func main() {
	fileName := "test.txt"
	createFile(fileName)
	readFile(fileName)
	removeFile(fileName)
}

func createFile(fileName string) {
	fmt.Printf("creating file.....%s", fileName)

	file, err := os.Create(fileName)

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

func readFile(fileName string) {
	fmt.Printf("reading file......%s", fileName)

	data, err := ioutil.ReadFile(fileName)

	if err != nil {
		log.Fatalf("error while reading file %s", err)
	}

	//fmt.Println("File name is : ",fileName)
	fmt.Printf("\nfile size is %d bytes ", len(data))
	fmt.Printf("\ndata in file is : %s ", data)
}

func removeFile(fileName string) {
	fmt.Printf("\nremoving file.....%s", fileName)

	if _, err := os.Stat(fileName); !os.IsNotExist(err) {

		err := os.Remove(fileName)

		if err != nil {
			fmt.Printf("error while deleting the file %s", err)
		}

		fmt.Println("\nfile deleted successfully")

	} else {
		fmt.Printf("\nNo file named : %s available.", fileName)
	}

}
