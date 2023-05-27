package main

import (
	"fmt"
	"math/rand"
	"time"
)

func main() {
	rand.Seed(time.Now().UnixNano())

	minTicket := 100000
	maxTicket := 999999

	luckyTicketOne := 0
	distanceOne := 0

	luckyTicketTwo := 0
	distanceTwo := 0

	for {

		for luckyTicketOne == 0 {
			fmt.Println("**1")
			currentTicketOne := rand.Intn(maxTicket-minTicket+1) + minTicket

			itemOne := currentTicketOne / 100000
			itemTwo := (currentTicketOne / 10000) % 10
			itemThree := (currentTicketOne / 1000) % 10
			itemFour := (currentTicketOne / 100) % 10
			itemFive := (currentTicketOne / 10) % 10
			itemSix := currentTicketOne % 10

			if itemOne+itemTwo+itemThree == itemFour+itemFive+itemSix {
				luckyTicketOne = currentTicketOne
			} else {
				currentTicketOne++
				distanceOne++
			}
		}

		for distanceOne > 0 && luckyTicketTwo == 0 {
			fmt.Println("**2")
			currentTicketTwo := rand.Intn(maxTicket-minTicket+1) + minTicket

			itemOne2 := currentTicketTwo / 100000
			itemTwo2 := (currentTicketTwo / 10000) % 10
			itemThree2 := (currentTicketTwo / 1000) % 10
			itemFour2 := (currentTicketTwo / 100) % 10
			itemFive2 := (currentTicketTwo / 10) % 10
			itemSix2 := currentTicketTwo % 10

			if itemOne2+itemTwo2+itemThree2 == itemFour2+itemFive2+itemSix2 {
				luckyTicketTwo = currentTicketTwo
			} else {
				currentTicketTwo++
				distanceTwo++
			}
		}

		if distanceOne == 0 {
			fmt.Println("Вы сразу купили сачстливый билет, не надо ничего больше покупать!")
			break
		}

		if distanceTwo > distanceOne {
			distanceOne = distanceTwo
			luckyTicketOne = luckyTicketTwo

			distanceTwo = 0
			luckyTicketTwo = 0
		} else {
			fmt.Println("Билетов до счастливого в 1 случае:", distanceOne)
			fmt.Println("Билетов до счастливого в 2 случае:", distanceTwo)
			fmt.Println("Первый купленный билет имеет наилучшее значение для покупки билетов до счастливого")
			break
		}
	}
}
