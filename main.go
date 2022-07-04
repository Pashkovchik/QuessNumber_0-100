package main

import "fmt"

func main() {
	fmt.Println("Я умный компьютер. Я смогу угадать загаданное число максимум с 7 попыток!")
	fmt.Println("Загадай число и отвечай: угадал, больше или меньше")
	var min, max, middle int
	var str string
	var b bool
	min = 0
	max = 101
	middle = (max + min) / 2
	for i := 0; b == false; {
		fmt.Println("Ваше число равно ", middle, "?")
		fmt.Scanln(&str)
		switch str {
		case "больше":
			fmt.Printf("Ваше число больше %v, вычисляю ещё\n", middle)
			min = middle
			middle = (max + min) / 2
			i++
		case "меньше":
			fmt.Printf("Ваше число меньше %v, вычисляю ещё\n", middle)
			max = middle
			middle = (max + min) / 2
			i++
		case "угадал":
			fmt.Println("Ура! Я угадал! Твое число:", middle, "Количество попыток:", i+1)
			b = true
		default:
			fmt.Println("Похоже ты неправильно ввел ответ. Попробуй еще раз)")
		}

	}
	fmt.Println("Пока пока")

}
