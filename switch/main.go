package main

import "fmt"

func weekDay(number int) string {
	switch number {
	case 1:
		return "Domingo"
	case 2:
		return "Segunda"
	case 3:
		return "Terça"
	case 4:
		return "Quarta"
	case 5:
		return "Quinta"
	case 6:
		return "Sexta"
	case 7:
		return "Sabado"
	default:
		return "Numero invalido"
	}
}

func outerWeekDay(number int) string {
	switch {
	case number == 1:
		return "Domingo"
	case number == 2:
		return "Segunda"
	case number == 3:
		return "Terça"
	case number == 4:
		return "Quarta"
	case number == 5:
		return "Quinta"
	case number == 6:
		return "Sexta"
	case number == 7:
		return "Sabado"
	default:
		return "Numero invalido"
	}
}

func outerWeekDay2(number int) string {
	var dayWeek string
	switch {
	case number == 1:
		dayWeek = "Domingo"
	case number == 2:
		dayWeek = "Segunda"
	case number == 3:
		dayWeek = "Terça"
	case number == 4:
		dayWeek = "Quarta"
	case number == 5:
		dayWeek = "Quinta"
	case number == 6:
		dayWeek = "Sexta"
	case number == 7:
		dayWeek = "Sabado"
	default:
		dayWeek = "Numero invalido"
	}
	return dayWeek
}

func main() {
	fmt.Println(weekDay(7))
	fmt.Println(outerWeekDay(2))
	fmt.Println(outerWeekDay2(1))
}
