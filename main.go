package main

import (
	"encoding/json"
	"fmt"
	"math/rand"
	"os"
	"prj-go/domain"
	"sort"
	"strconv"
	"time"
)

var id uint64 = 1

const (
	pointsPerQuestion = 50
	totalPoints       = 50
)

func main() {
	fmt.Println("Вітаємо у грі MATH-COR!")
	time.Sleep(2 * time.Second)

	var users []domain.User
	users = getUsers()
	for _, user := range users {
		if user.Id >= id {
			id = user.Id + 1
		}
	}

	for {
		menu()

		punct := ""
		fmt.Scan(&punct)

		switch punct {
		case "1":
			u := play()
			users = getUsers()
			users = append(users, u)
			sortAndSave(users)
		case "2":
			users = getUsers()
			for i, user := range users {
				fmt.Printf("i: %v, Id: %v Name: %s, Time: %v\n",
					i, user.Id, user.Name, user.Time)
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
	for i := 3; i > 0; i-- {
		fmt.Println(i)
		time.Sleep(1 * time.Second)
	}

	myPoints := 0
	points := totalPoints
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

func sortAndSave(users []domain.User) {
	sort.Slice(users, func(i, j int) bool {
		return users[i].Time < users[j].Time
	})

	file, err := os.OpenFile("users.json", os.O_RDWR|os.O_CREATE|os.O_TRUNC, 0755)
	if err != nil {
		fmt.Printf("os.OpenFile: %s", err)
		return
	}

	defer func(f *os.File) {
		err = f.Close()
		if err != nil {
			fmt.Printf("f.Close(): %s", err)
		}
	}(file)

	encoder := json.NewEncoder(file)
	err = encoder.Encode(users)
	if err != nil {
		fmt.Printf("encoder.Encode(users): %s", err)
		return
	}
}

func getUsers() []domain.User {
	var users []domain.User

	info, err := os.Stat("users.json")
	if err != nil {
		if os.IsNotExist(err) {
			_, err := os.Create("users.json")
			if err != nil {
				fmt.Printf("os.Create: %s", err)
				return nil
			}
			return nil
		}
	}

	if info.Size() != 0 {
		file, err := os.Open("users.json")
		if err != nil {
			fmt.Printf("os.Open: %s", err)
			return nil
		}

		defer func(f *os.File) {
			err = f.Close()
			if err != nil {
				fmt.Printf("f.Close(): %s", err)
			}
		}(file)

		decoder := json.NewDecoder(file)
		err = decoder.Decode(&users)
		if err != nil {
			fmt.Printf("decoder.Decode: %s", err)
			return nil
		}
	}

	return users
}
