package main

import "fmt"

func main() {
	expenses := readExpenses()
	answer, err := calculateAnswer(expenses)

	if err != nil {
		fmt.Println("No answer", err)
	} else {
		fmt.Println(answer)
	}
}

func calculateAnswer(expenses []uint) (answer uint, err error) {
	for index, expense1 := range expenses {
		for _, expense2 := range expenses[index+1:] {
			if expense1+expense2 == 2020 {
				answer = expense1 * expense2

				return
			}
		}
	}

	err = fmt.Errorf("no answer")
	return
}

func readExpenses() []uint {
	list := make([]uint, 0)

	for {
		var expense uint
		_, ok := fmt.Scanln(&expense)

		if ok != nil {
			break
		}

		list = append(list, expense)
	}
	return list
}
