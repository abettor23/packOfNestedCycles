package main

import (
	"fmt"
)

func main() {
	minTicket := 100000
	maxTicket := 999999
	mirrorTickets := 0

	for minTicket <= maxTicket {

		left := minTicket / 1000
		right := minTicket % 1000

		reversedLeft := 0

		for temp := left; temp > 0; temp /= 10 {
			reversedLeft = reversedLeft*10 + temp%10
		}

		if right == reversedLeft {
			mirrorTickets++
		}

		minTicket++
	}

	fmt.Println("Количество зеркальных билетов среди всех шестизначных номеров от 100000 до 999999:")
	fmt.Println(mirrorTickets)
}
