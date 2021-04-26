package main

import (
	"fmt"
)

type Doctor struct {
	name       string
	speicality string
}

type software struct {
	name  string
	skill string
}

type Employee interface {
	getnameAndSpecaility() (string, string)
}

func main() {
	doctor := Doctor{
		name:       "porus",
		speicality: "internal medicine",
	}
	software := software{
		name:  "purushotham",
		skill: "Golang",
	}
	printDetails(doctor)
	printDetails(software)
}
func (d Doctor) getnameAndSpecaility() (string, string) {
	return "Dr." + d.name, d.speicality
}
func (s software) getnameAndSpecaility() (string, string) {
	return "Sw." + s.name, s.skill
}
func printDetails(emp Employee) {
	name, speciality := emp.getnameAndSpecaility()
	fmt.Println("employee name : ", name)
	fmt.Println("employee speciality : ", speciality)
}
