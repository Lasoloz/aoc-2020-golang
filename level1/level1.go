package main

import "fmt"

const SUM = 2020

func main() {
	expenses := readExpenses()

	printAnswer(calculate1(expenses))
	printAnswer(calculate2(expenses))
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

func calculate1(expenses []uint) (answer uint, err error) {
	for index, expense1 := range expenses {
		for _, expense2 := range expenses[index+1:] {
			if expense1+expense2 == SUM {
				answer = expense1 * expense2

				return
			}
		}
	}

	err = fmt.Errorf("no answer for part 1")
	return
}

func calculate2(expenses []uint) (uint, error) {
	for index1, expense1 := range expenses {
		for index2, expense2 := range expenses[index1+1:] {
			for _, expense3 := range expenses[index2+1:] {
				if expense1+expense2+expense3 == SUM {
					return expense1 * expense2 * expense3, nil
				}
			}
		}
	}

	return 0, fmt.Errorf("no answer for part 2")
}

func printAnswer(answer uint, err error) {
	if err != nil {
		fmt.Println("No answer", err)
	} else {
		fmt.Println(answer)
	}
}
