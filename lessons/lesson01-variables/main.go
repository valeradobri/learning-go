package main

import "fmt"

const OS = "windows"

func main() {

	name := "Valeriy"
	var age int = 24
	var salary = 2500.50
	var isReady bool

	fmt.Println("=== Моя визитка ===")
	fmt.Println("Имя:", name)
	fmt.Println("Возраст:", age)
	fmt.Println("Зарплата:", salary)
	fmt.Println("Готов к учебе:", isReady)
	fmt.Println("Операционная система:", OS)

	tempC := 25
	tempF := (tempC * 9.0 / 5.0) + 32.0

	fmt.Println("\n=== Конвертер температуры ===")
	fmt.Println("Градусы Цельсия:", tempC)
	fmt.Println("Градусы Фаренгейта:", tempF)

	var price int = 100
	var tax float64 = 1.20 // 20% налог
	// Попробуй перемножить две переменные:
	total := float64(price) * tax
	fmt.Println(total)

	secretcode := 12345
	_ = secretcode
}
