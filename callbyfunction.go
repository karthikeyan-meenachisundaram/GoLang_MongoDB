package main

import "fmt"

type Employee struct {
	Id     int
	Name   string
	Salary float32
}

// 1. Array passed by value (copy only, no effect on main)
func updateArray(arr [3]Employee) {
	arr[0].Salary += 5000
	fmt.Println("\nInside updateArray (copy of array):")
	for _, emp := range arr {
		fmt.Printf("Id: %d, Name: %s, Salary: %.2f\n", emp.Id, emp.Name, emp.Salary)
	}
}

// 2. Array passed by pointer (can update original array)
func updateArrayPtr(arr *[3]Employee) {
	arr[0].Salary += 5000
	fmt.Println("\nInside updateArrayPtr (pointer to array):")
	for _, emp := range arr {
		fmt.Printf("Id: %d, Name: %s, Salary: %.2f\n", emp.Id, emp.Name, emp.Salary)
	}
}

// 3. Slice passed (points to same underlying array, so updates main)
func updateSlice(slc []Employee) {
	slc[0].Salary += 5000
	fmt.Println("\nInside updateSlice (same underlying array):")
	for _, emp := range slc {
		fmt.Printf("Id: %d, Name: %s, Salary: %.2f\n", emp.Id, emp.Name, emp.Salary)
	}
}

// 4. Slice copied with copy() (independent copy, no effect on main)
func updateSliceSafe(slc []Employee) {
	newSlc := make([]Employee, len(slc))
	copy(newSlc, slc) // deep copy of slice

	newSlc[0].Salary += 5000
	fmt.Println("\nInside updateSliceSafe (copy of slice):")
	for _, emp := range newSlc {
		fmt.Printf("Id: %d, Name: %s, Salary: %.2f\n", emp.Id, emp.Name, emp.Salary)
	}
}

func main() {
	// Array of employees
	empArray := [3]Employee{
		{Id: 101, Name: "Karthi", Salary: 20000},
		{Id: 102, Name: "Mahesh", Salary: 25000},
		{Id: 103, Name: "Jagan", Salary: 30000},
	}

	// Slice of employees
	empSlice := []Employee{
		{Id: 201, Name: "Srini", Salary: 40000},
		{Id: 202, Name: "Sanjay", Salary: 45000},
		{Id: 203, Name: "Ilavarasi", Salary: 50000},
	}

	// 1. Array by value
	updateArray(empArray)
	fmt.Println("\nBack in main after updateArray (original unchanged):")
	for _, emp := range empArray {
		fmt.Printf("Id: %d, Name: %s, Salary: %.2f\n", emp.Id, emp.Name, emp.Salary)
	}

	// 2. Array by pointer
	updateArrayPtr(&empArray)
	fmt.Println("\nBack in main after updateArrayPtr (original changed):")
	for _, emp := range empArray {
		fmt.Printf("Id: %d, Name: %s, Salary: %.2f\n", emp.Id, emp.Name, emp.Salary)
	}

	// 3. Slice (reference-like)
	updateSlice(empSlice)
	fmt.Println("\nBack in main after updateSlice (original changed):")
	for _, emp := range empSlice {
		fmt.Printf("Id: %d, Name: %s, Salary: %.2f\n", emp.Id, emp.Name, emp.Salary)
	}

	// 4. Slice safe copy
	updateSliceSafe(empSlice)
	fmt.Println("\nBack in main after updateSliceSafe (original unchanged):")
	for _, emp := range empSlice {
		fmt.Printf("Id: %d, Name: %s, Salary: %.2f\n", emp.Id, emp.Name, emp.Salary)
	}
}
