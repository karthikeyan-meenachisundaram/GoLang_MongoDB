package main

import "fmt"

type Employee struct {
	EmpID  int
	Name   string
	Salary float32
}

func main() {
	employees := []Employee{
		{EmpID: 101, Name: "Karthikeyan", Salary: 12000.00},
		{EmpID: 102, Name: "Mahesh", Salary: 18000.00},
		{EmpID: 103, Name: "Jagan", Salary: 15000.00},
		{EmpID: 104, Name: "Sanjay", Salary: 13000.00},
	}
	for _, e := range employees {
		fmt.Printf("Id: %d\n", e.EmpID)
		fmt.Printf("Name: %s\n", e.Name)
		fmt.Printf("Salary: %.2f\n", e.Salary)
		fmt.Println("-------------------------------")
	}
}
