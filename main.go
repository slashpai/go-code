package main

import (
	"fmt"
	"github.com/go-code/student"
	"github.com/go-code/subfactorial"
	"os"
)

// call student package methods
func callStudent(){
	var numberOfStudents int
	fmt.Println("Enter Student Details")
	fmt.Printf("Enter Number of Students: ")
	fmt.Scanln(&numberOfStudents)
	students := make([]student.Student, numberOfStudents)

	for i := 0; i < numberOfStudents; i++ {
		student := student.NewStudent()
		students[i] = *student
		students[i].GetStudentDetails()
		students[i].CalcTotal()
		fmt.Println("..................................")
	}

	fmt.Println("Student Details")
	fmt.Printf("%10s%25s\n", "Name", "Total Marks")
	for i := 0; i < numberOfStudents; i++ {
		students[i].PrintDetails()
	}
}

func callSubFactorial(){
	var n int
	fmt.Printf("Enter n: ")
	fmt.Scanf("%d", &n)
	arr := make([]int, n)
	// initialize array
	for i := 0; i < n; i++ {
		arr[i] = i + 1
	}
	subfactorial.Generate(n, arr)
	fmt.Printf("!%d -> %d\n", n, subfactorial.SubFactorial)
}

func main(){
	var option int
	fmt.Println("Select option from below:")
	fmt.Println("1. Student Program")
	fmt.Println("2. Subfactorial Program")
	fmt.Printf("Enter option, which program to execute: ")
	fmt.Scanf("%d",&option)
	switch option {
	case 1:
		fmt.Println("Running Student Program...")
		callStudent()
	case 2:
		fmt.Println("Running Subfactorial Program...")
		callSubFactorial()
	default:
		fmt.Println("Wrong option selected, try again!")
		os.Exit(1)
	}
}
