package main

import (
	"fmt"
)

func main() {
	fmt.Println("Вывод ёлочки")

	var height int
	for height <= 0 {
		fmt.Println("Введите высоту ёлочки:")
		fmt.Scan(&height)
		if height <= 0 {
			fmt.Println("Параметр задан неверно. Используйте цифры. Значение не может быть 0 или меньше.")
		}
	}

	for i := 1; i <= height; i++ { //цикл определяющий высоту елочки по количеству выводимых строк

		for k := 0; k < height-i; k++ { //определение количества пробелов на строку
			fmt.Print(" ")
		}

		for m := 0; m < 2*i-1; m++ { //определение количества звездочек на строку (нечётный пропорциональный рост)
			fmt.Print("*")
		}

		fmt.Println() //перенос строки

	}
}
