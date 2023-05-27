package main

import (
	"fmt"
)

func main() {
	fmt.Println("Шахматная доска.")

	var width int
	for width <= 0 {
		fmt.Println("Введите ширину:")
		fmt.Scan(&width)
		if width <= 0 {
			fmt.Println("Параметр задан неверно. Используйте цифры. Значение не может быть 0 или меньше.")
		}
	}

	var height int
	for height <= 0 {
		fmt.Println("Введите высоту:")
		fmt.Scan(&height)
		if height <= 0 {
			fmt.Println("Параметр задан неверно. Используйте цифры. Значение не может быть 0 или меньше.")
		}
	}

	evenHeight := 0 //упорядоченность строк, для чередования начала строки с пробела или звездочки

	for i := 0; i < height; i++ { //цикл устанавливающий высоту (количество строк = количество итераций)
		evenWidth := 0 //чередование вывода пробела и звездочки в самой строке

		for j := 0; j < width && evenHeight == 0; j++ { //цикл для вывода череды символов с 0 порядком (evenHeight)
			if evenWidth == 0 {
				fmt.Print(" ")
				evenWidth++
			} else {
				fmt.Print("*")
				evenWidth--
			}
		}

		for j := 0; j < width && evenHeight == 1; j++ { //цикл для вывода череды символов с 1 порядком (evenHeight)
			if evenWidth == 0 {
				fmt.Print("*")
				evenWidth++
			} else {
				fmt.Print(" ")
				evenWidth--
			}
		}

		fmt.Println() //после завершения цикла вывода символов - осуществить перенос строки

		if evenHeight == 0 { //и после меняем очередность
			evenHeight++
		} else {
			evenHeight--
		}
	}

}
