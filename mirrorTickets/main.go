package main

import (
	"fmt"
)

func main() {
	minTicket := 100000
	maxTicket := 999999
	mirrorTickets := 0

	for minTicket <= maxTicket {
		itemOne := minTicket / 100000
		itemTwo := (minTicket / 10000) % 10
		itemThree := (minTicket / 1000) % 10
		itemFour := (minTicket / 100) % 10
		itemFive := (minTicket / 10) % 10
		itemSix := minTicket % 10

		if itemOne == itemSix && itemTwo == itemFive && itemThree == itemFour {
			mirrorTickets++
		}

		minTicket++
	}

	fmt.Println("Количество зеркальных билетов среди всех шестизначных номеров от 100000 до 999999:")
	fmt.Println(mirrorTickets)
}
