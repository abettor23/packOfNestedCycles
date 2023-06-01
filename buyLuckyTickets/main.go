package main

import (
	"fmt"
)

func main() {

	minTicket := 100000
	maxTicket := 999999

	luckyPrevious := 0 //переменные предыдущего найденного счастливого билета и текущего
	luckyCurrent := 0
	distance := 0 //переменная наибольшего найденного расстояния между счастливыми билетами

	for minTicket <= maxTicket {

		left := minTicket / 1000 //разделение номера билета на две части, для поиска счастливого по сумме чисел
		right := minTicket % 1000

		leftSum := 0 //сумма чисел частей билета для определения счастливого (сравнения друг с другом)
		rightSum := 0

		for left > 0 || right > 0 { //цикл для определения суммы чисел каждой из частей билета (сложение по очереди)
			numOne := left % 10
			leftSum += numOne
			left /= 10

			numTwo := right % 10
			rightSum += numTwo
			right /= 10
		}

		if leftSum == rightSum { //определение счастливого билета и порядковая запись для определения расстояния между ними
			luckyPrevious = luckyCurrent
			luckyCurrent = minTicket
		}

		if luckyCurrent-luckyPrevious > distance && luckyPrevious != 0 { //определение расстояния. если больше - заменить значение.
			distance = luckyCurrent - luckyPrevious
		}

		minTicket++ //переход к новому номеру билета
	}
	fmt.Println("Максимальный диапазон между счастливыми билетами составляет:", distance)
}
