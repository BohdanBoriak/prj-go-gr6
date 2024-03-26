package main

import (
	"fmt"
	"math/rand"
	"prj-go/domain"
	"strconv"
	"time"
)

var points int = 50
var id uint64 = 1

const pointsPerQuestion = 50

func main() {
	fmt.Println("Вітаємо у грі MATH-COR!")
	time.Sleep(2 * time.Second)

	var users []domain.User

	for {
		menu()

		punct := ""
		fmt.Scan(&punct)

		switch punct {
		case "1":
			u := play()
			users = append(users, u)
		case "2":
			for i, user := range users {

			}
		case "3":
			return
		default:
			fmt.Println("Зробіть коректний вибір")
		}
	}

}

func menu() {
	fmt.Println("1. Почати гру")
	fmt.Println("2. Переглянути рейтинг")
	fmt.Println("3. Вийти")
}

func play() domain.User {
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

	fmt.Printf("Який ти молодець! Впорався всього то за %v\n", gameTime)
	fmt.Println("Введіть ім'я:")

	name := ""
	fmt.Scan(&name)

	user := domain.User{
		Id:   id,
		Name: name,
		Time: gameTime,
	}
	id++

	return user
}
