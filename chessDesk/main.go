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

	for i := 0; i < height; i++ { //цикл, задающий высоту доски (выоста = количество итераций)

		for j := 0; j < width; j++ { //цикл, задающий ширину доски; на основе чётности строки выводим цвет клетки
			if (i+j)%2 == 0 {
				fmt.Print("*")
			} else {
				fmt.Print(" ")
			}
		}

		fmt.Println() //после завершения цикла вывода символов - осуществить перенос строки
	}

}
