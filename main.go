// main go
package main

import (
	"errors"
	"fmt"
	"math"
)

func main() {
	for {
		userHeight, userKg := getUserInput()
		IMT, err := calculateIMT(userHeight, userKg)
		if err != nil {
			fmt.Println("Некорректные параметры для рассчета")
			continue
			//panic("Некорректные параметры для рассчета")
		}
		outputRezult(IMT)
		isRepiat := beginContinue()
		if !isRepiat {
			break
		}
	}
}

// outputRezult print Imt
func outputRezult(imt float64) {
	result := fmt.Sprintf("Ваш ИМТ: %.0f", imt)
	fmt.Println(result)
	switch {
	case imt < 16:
		fmt.Println("Сильный дефицит")
	case imt < 18.5:
		fmt.Println("Дефицит")
	case imt < 25:
		fmt.Println("Норма")
	case imt < 30:
		fmt.Println("Избыток")
	default:
		fmt.Println("Ожирение")
	}
}

func calculateIMT(userHeight, userKg float64) (float64, error) {
	if userHeight <= 0 || userKg <= 0 {
		return 0, errors.New("NO_PARAMS_ERROR")
	}
	const IMTPower = 2
	return userKg / math.Pow(userHeight/100, IMTPower), nil
}

func getUserInput() (float64, float64) {
	var userHeight float64
	var userKg float64

	fmt.Println(`__Рссчет ИМТ__
Введите ваш рост (см): `)
	fmt.Scan(&userHeight)

	fmt.Print("Введите ваш вес (см): ")
	fmt.Scan(&userKg)

	return userHeight, userKg
}

func beginContinue() bool {
	var answerCont string
	fmt.Println("Желаете продолжить? (y / n)")
	fmt.Scan(&answerCont)
	if answerCont == "Y" || answerCont == "y" {
		return true
	}
	return false
}
