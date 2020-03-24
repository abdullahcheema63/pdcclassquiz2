package pdcclassquiz2

import (
	"fmt"
)

var status = make(map[string]bool)

// PrintCasesForFrance ... Function prints covid19 cases in france as of 24.03.2020 3.41M
func PrintCasesForFrance() {

	status["france"] = true
	fmt.Println("COVID-19 statistics for France")
	fmt.Println("Total Number of COVID-19 Cases: 15,821")
	fmt.Println("No Of Deaths: 674")
}

// PrintCasesForSouthKorea ... Function prints covid19 cases in south korea as of 24.03.2020 3.41PM
func PrintCasesForSouthKorea() {

	status["southkorea"] = true

	fmt.Println("COVID-19 statistics for South Korea")
	fmt.Println("Total Number of COVID-19 Cases: 8,961")
	fmt.Println("No Of Deaths: 111")
}

// PrintCasesForPakistan ... Function prints covid19 cases in pakistan as of 24.03.2020 3.41PM
func PrintCasesForPakistan() {
	status["pakistan"] = true
	fmt.Println("COVID-19 statistics for Pakistan")
	fmt.Println("Total Number of COVID-19 Cases: 784")
	fmt.Println("No Of Deaths: 5")
}

func ChooseOptions() {
	status["pakistan"] = false
	status["southkorea"] = false
	status["france"] = false
	for !(status["pakistan"] && status["southkorea"] && status["france"]) {
		var choosen string
		fmt.Println("1. Print Stats for Pakistan")
		fmt.Println("2. Print Stats for South Korea")
		fmt.Println("3. Print Stats for France")
		fmt.Println("0. Exit")
		fmt.Scanf("%s", &choosen)
		if choosen == "1" {
			fmt.Println()
			PrintCasesForPakistan()
			fmt.Println()
		} else if choosen == "2" {
			fmt.Println()
			PrintCasesForSouthKorea()
			fmt.Println()
		} else if choosen == "3" {
			fmt.Println()
			PrintCasesForFrance()
			fmt.Println()
		} else if choosen == "0" {
			return
		} else {
			fmt.Println()
			fmt.Println("Incorrect Input, Please enter from the given options only")
		}

	}

}
