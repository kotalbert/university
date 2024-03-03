package main

import "fmt"

func main() {
	var exams = [3]float64{}

	for i := 0; i < 3; i++ {
		_, _ = fmt.Scan(&exams[i])
	}

	var score float64
	for i := 0; i < 3; i++ {
		score += exams[i]
	}
	score = score / 3

	fmt.Println(score)
	if score >= 60.0 {
		fmt.Println("Congratulations, you are accepted!")
		return
	} else {
		fmt.Println("We regret to inform you that we will not be able to offer you admission.")
		return
	}

}
