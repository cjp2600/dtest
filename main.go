package main

import (
	"fmt"
)

type Lesson struct {
	Name   string
	Rating int32
}

type Profile struct {
	SchoolName string
	Lessons    []Lesson
}

type Human struct {
	Ves       int32
	Age       int32
	Sex       string
	FirstName string
	LastName  string

	Profile Profile
}

func main() {

	students := GetStudents()
	for _, student := range students {
		if student.Age > 35 {
			fmt.Println(student.FirstName + " " + student.LastName)
			if student.Profile.SchoolName == "16 школа" {
				fmt.Println(student.Profile.Lessons)
			}
		}
	}

}

func GetStudents() []Human {
	var humans []Human
	var vladLesson []Lesson
	var olegLesson []Lesson
	var dLesson []Lesson

	vladLesson = append(vladLesson, Lesson{
		Name:   "Физкультура",
		Rating: 5,
	})
	vladLesson = append(vladLesson, Lesson{
		Name:   "История",
		Rating: 2,
	})

	olegLesson = append(vladLesson, Lesson{
		Name:   "Физкультура",
		Rating: 2,
	})
	olegLesson = append(vladLesson, Lesson{
		Name:   "История",
		Rating: 4,
	})

	dLesson = append(vladLesson, Lesson{
		Name:   "Физкультура",
		Rating: 5,
	})
	dLesson = append(vladLesson, Lesson{
		Name:   "История",
		Rating: 3,
	})

	humans = append(humans, Human{
		Ves:       190,
		Age:       36,
		Sex:       "man",
		FirstName: "Влад",
		LastName:  "Супруненко",
		Profile: Profile{
			SchoolName: "13 школа",
			Lessons:    vladLesson,
		},
	})

	humans = append(humans, Human{
		Ves:       40,
		Age:       36,
		Sex:       "man",
		FirstName: "Олег",
		LastName:  "Истомин",
		Profile: Profile{
			SchoolName: "16 школа",
			Lessons:    olegLesson,
		},
	})

	humans = append(humans, Human{
		Ves:       80,
		Age:       35,
		Sex:       "man",
		FirstName: "Дмитрий",
		LastName:  "Попов",
		Profile: Profile{
			SchoolName: "9 школа",
			Lessons:    dLesson,
		},
	})

	return humans
}
