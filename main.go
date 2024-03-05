package main

import (
	"bufio"
	"fmt"
	"log"
	"os"
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

// translateDepartment translates the department name to the corresponding Department constant, in order to read
// students applications and convert string department names and create a Student struct.
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

// readStudent reads a student's application and returns a Student struct.
//
//	The student application is a line read from the list, with the following format:
//	<first name> <last name> <GPA> <department 1> <department 2> <department 3>
func readStudent(line string) (Student, error) {
	var student Student
	var dept1, dept2, dept3 string

	_, err := fmt.Sscanf(line, "%s %s %f %s %s %s", &student.firstName, &student.lastName, &student.gpa, &dept1, &dept2, &dept3)

	if err != nil {
		return student, err
	}

	student.dept1 = translateDepartment(dept1)
	student.dept2 = translateDepartment(dept2)
	student.dept3 = translateDepartment(dept3)

	return student, nil
}

// readApplications reads the students applications from a file `applicant_list.txt` and returns a slice of Student structs.
func readApplications() []Student {
	var applicants []Student

	file, err := os.Open("applicant_list.txt")
	if err != nil {
		log.Fatal(err)
	}
	defer func(file *os.File) {
		_ = file.Close()
	}(file)

	scanner := bufio.NewScanner(file)
	for scanner.Scan() {
		student, err := readStudent(scanner.Text())
		if err != nil {
			log.Fatal(err)
		}
		applicants = append(applicants, student)
	}

	return applicants
}
