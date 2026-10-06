package main

import "fmt"

type student struct {
	name  string
	age   int
	marks []map[string]int
}

func (s *student) String() string {
	result := fmt.Sprintf("Name: %s | Age: %d | Marks:\n", s.name, s.age)
	for _, m := range s.marks {
		for subject, score := range m {
			result += fmt.Sprintf("  %-8s: %d\n", subject, score)
		}
	}
	return result
}

func (s *student) updateStudent(name string, age int, marks []map[string]int) {

	fmt.Printf("Updating the student %s\n", name)
	s.name = name
	s.age = age
	s.marks = append(s.marks, marks...)
}

func main() {

	fmt.Println("This is where the actual updates happen")
	//initialising the student struct

	var s1 student

	s1.updateStudent("tanay", 22, []map[string]int{{"phy": 22, "eng": 23, "maths": 33}})
	fmt.Println("The student information is\n", &s1)

}
