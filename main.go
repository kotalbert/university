package main

import (
	"fmt"
	"sort"
)

type Student struct {
	firstName string
	lastName  string
	gpa       float64
}

func main() {
	n := readInt() // total numbers of applicants
	m := readInt() // number of available seats

	students := readApplications(n)
	sortStudents(students)

	fmt.Println("Successful applicants:")
	for i := 0; i < m; i++ {
		fmt.Println(students[i].firstName, students[i].lastName)
	}
}
func sortStudents(students []Student) {
	sort.Slice(students, func(i, j int) bool {
		if students[i].gpa != students[j].gpa {
			return students[i].gpa > students[j].gpa
		}
		if students[i].firstName != students[j].firstName {
			return students[i].firstName < students[j].firstName
		}

		return students[i].lastName < students[j].lastName
	})
}

func readInt() int {
	var input int
	_, _ = fmt.Scan(&input)
	return input
}

func readStudent() Student {
	var student Student
	_, _ = fmt.Scan(&student.firstName, &student.lastName, &student.gpa)
	return student
}

func readApplications(n int) []Student {
	applicants := make([]Student, n)
	for i := 0; i < n; i++ {
		applicants[i] = readStudent()
	}
	return applicants
}
