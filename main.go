package main

import (
	"fmt"
	"math/rand"
	"strconv"
	"time"
)

var points int = 50

const pointsPerQuestion = 50

func main() {
	fmt.Println("Вітаємо у грі MATH-COR!")
	time.Sleep(2 * time.Second)

}

func play() {
	for i := 5; i >= 0; i-- {
		fmt.Println(i)
		time.Sleep(1 * time.Second)
	}

	myPoints := 0
	startTime := time.Now()
	for points > 0 {
		x, y := rand.Intn(100), rand.Intn(100)
		res := x + y

		fmt.Printf("%v + %v = ", x, y)

		ans := ""
		fmt.Scan(&ans)

		ansInt, err := strconv.Atoi(ans)
		if err != nil {
			fmt.Println("Не правильно! Спробуй ще!")
		} else {
			if res == ansInt {
				points -= pointsPerQuestion
				myPoints += pointsPerQuestion
				fmt.Println("Правильно!")
				fmt.Printf("У тебе %v балів, залишилось зібрати %v\n", myPoints, points)
			} else {
				fmt.Println("Вчи математику!!!11")
			}
		}
	}
	endTime := time.Now()
	gameTime := endTime.Sub(startTime)

	fmt.Printf("Який ти молодець! Впорався всього то за %v", gameTime)
	time.Sleep(10 * time.Second)
}
