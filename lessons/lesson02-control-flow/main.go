package main

import "fmt"

func main() {
	// var age int
	// fmt.Print("Введите ваш возраст: ")
	// fmt.Scan(&age)

	// if age < 0 || age > 120 {
	// 	fmt.Println("Ошибка: неккоректный возраст!")
	// } else if age < 18 {
	// 	fmt.Println("Вам еще нет 18-ти лет, вход воспрещен!")
	// } else {
	// 	fmt.Println("Добро пожаловать в клуб!")
	// }

	// var light string
	// fmt.Print("\nКакой цвет горит на светофоре? (зеленый/желтый/красный): ")
	// fmt.Scan(&light)

	// switch light {

	// case "зеленый":
	// 	fmt.Println("Проезд разрешен! Поехали!")

	// case "желтый":
	// 	fmt.Println("Внимание! Приготовьтесь.")

	// case "красный":
	// 	fmt.Println("СТОЙ! Проезд запрещен!")
	// }

	// // === МИНИ-ИГРА: Угадай число ===
	// secret := 42  // Загаданное число
	// var guess int // Переменная для догадки пользователя
	// attempts := 0 // Счетчик попыток

	// fmt.Println("\n=== Я загадал число от 1 до 100. Попробуй угадать! ===")

	// // Бесконечный цикл: крутится до тех пор, пока мы не скажем "break"!
	// for {
	// 	fmt.Print("Введи число: ")
	// 	fmt.Scan(&guess)
	// 	attempts++ // увеличиваем счетчик попыток на 1 (attempts = attempts + 1)

	// 	if guess < secret {
	// 		fmt.Println("Загаданное число БОЛЬШЕ! ⬆️")
	// 	} else if guess > secret {
	// 		fmt.Println("Загаданное число МЕНЬШЕ! ⬇️")
	// 	} else {
	// 		fmt.Printf("🎉 ПОБЕДА! Ты угадал число %d за %d попыток!\n", secret, attempts)
	// 		break // ВЫХОД ИЗ ЦИКЛА! Игра закончена.
	// 	}
	// }

	balance := 1000
	var choice int
	var amount int

	for {
		fmt.Println("\n=== Банкомат ===")
		fmt.Println("1. Проверить баланс")
		fmt.Println("2. Пополнить счет")
		fmt.Println("3. Снять наличные")
		fmt.Println("4. Выход")
		fmt.Print("Выбери действие: ")

		fmt.Scan(&choice)

		if choice == 1 {
			fmt.Println("Ваш текущий баланс:", balance, "рублей.")

		} else if choice == 2 {
			fmt.Println("Сколько хотите внести?")
			fmt.Scan(&amount)

			if amount > 0 {
				balance = balance + amount
				fmt.Println("Пополнение успешно! \nВаш баланс: ", balance)
			} else {
				fmt.Println("Сумма должна быть больше 0")
			}
		} else if choice == 3 {
			fmt.Println("Сколько хотите снять?")
			fmt.Scan(&amount)

			if amount > 0 {
				balance = balance - amount
				fmt.Println("Заберите наличные! \nВаш баланс ", balance)
			} else {
				fmt.Println("Сумма должна быть больше 0")
			}
		} else if choice == 4 {
			fmt.Println("До свидания!")
			break
		}
	}
}
