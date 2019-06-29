package main

import (
	"fmt"
)

func main() {

	fmt.Println("Программа страны: ")
	fmt.Println("")

	var coutries []string

	coutries = []string{"Россия", "Китай", "США", "Германия", "Франция", "Италия", "Испания", "Англия"}
	countCountries := len(coutries)

	fmt.Println("")
	fmt.Printf("Всего %v", countCountries)
	fmt.Println("")
	fmt.Println("")

	for i := 0; i < countCountries; i++ {
		currentName := coutries[i]

		if currentName == "Италия" {
			continue
		}

		fmt.Println(currentName + " !")
	}

}
