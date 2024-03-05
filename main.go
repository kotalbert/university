package main

import (
	"fmt"
	"sort"
)

type Department int

const (
	Physics = iota
	Mathematics
	Engineering
	Biotech
	Chemistry
)

type Student struct {
	firstName           string
	lastName            string
	gpa                 float64
	dept1, dept2, dept3 Department
}

func main() {
	//n := readInt() // total numbers of seats per department

	//students := readApplications()
	//sortStudents(students)
	//
	//fmt.Println("Successful applicants:")
	//for i := 0; i < m; i++ {
	//	fmt.Println(students[i].firstName, students[i].lastName)
	//}
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

func translateDepartment(department string) Department {
	switch department {
	case "Biotech":
		return Biotech
	case "Chemistry":
		return Chemistry
	case "Engineering":
		return Engineering
	case "Mathematics":
		return Mathematics
	case "Physics":
		return Physics
	}
	return -1
}

func readStudent(line string) Student {
	var student Student
	_, _ = fmt.Scan(&student.firstName, &student.lastName, &student.gpa)
	return student
}

func readApplications() []Student {
	var applicants []Student

	return applicants
}
