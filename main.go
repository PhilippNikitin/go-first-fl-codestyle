package main

import (
	"fmt"
	"math/rand"
	"strings"
)

func randint(min, max int) int {
	return rand.Intn(max-min) + min
}

func attack(charName, charClass string) string {
	attackRange := map[string][2]int{
		"warrior": {8, 10},
		"mage":    {10, 15},
		"healer":  {2, 4},
	}

	attackLimits, found := attackRange[charClass] // записываем в переменную attackLimits массив, содержащий минимальную и максимальную величины атаки юнита
	if !found {
		return "неизвестный класс персонажа"
	}
	return fmt.Sprintf("%s нанес урон противнику равный %d.", charName, randint(attackLimits[0], attackLimits[1]))
}

// обратите внимание на "if else" и на "else"
func defence(charName, charClass string) string {
	blockRange := map[string][2]int{
		"warrior": {15, 20},
		"mage":    {8, 12},
		"healer":  {12, 15},
	}

	blockLimits, found := blockRange[charClass] // записываем в переменную blockLimits массив, содержащий минимальный и максимальный размеры блока
	if !found {
		return "неизвестный класс персонажа"
	}
	return fmt.Sprintf("%s нанес урон противнику равный %d.", charName, randint(blockLimits[0], blockLimits[1]))
}

// обратите внимание на "if else" и на "else"
func special(charName, charClass string) string {
	// создаем словарь с дополнительными способностями персонажей
	specialAbility := map[string]string{
		"warrior": fmt.Sprintf("`Выносливость %d`", 105),
		"mage":    fmt.Sprintf("`Атака %d`", 45),
		"healer":  fmt.Sprintf("`Защита %d`", 40),
	}
	special, found := specialAbility[charClass]
	if !found {
		return "неизвестный класс персонажа"
	}
	return fmt.Sprintf("%s применил специальное умение %s", charName, special)
}

// здесь обратите внимание на имена параметров
func startTraining(charName, charClass string) string {
	// создаем словарь с приветствиями в зависимости от класса персонажа
	greeting := map[string]string{
		"warrior": ", ты Воитель - отличный боец ближнего боя.\n",
		"mage":    ", ты Маг - превосходный укротитель стихий.\n",
		"healer":  ", ты Лекарь - чародей, способный исцелять раны.\n",
	}
	// печатаем приветствие
	fmt.Printf("%s%s", charName, greeting[charClass])
	// печатаем приглашение к тренировке
	fmt.Println(`Потренируйся управлять своими навыками.
Введи одну из команд: attack — чтобы атаковать противника,
defence — чтобы блокировать атаку противника,
special — чтобы использовать свою суперсилу.
Если не хочешь тренироваться, введи команду skip.`)

	var cmd string
	var action string
	for cmd != "skip" {
		fmt.Print("Введи команду: ")
		fmt.Scanf("%s\n", &cmd)
		switch cmd {
		case "attack":
			action = attack(charName, charClass)
		case "defence":
			action = defence(charName, charClass)
		case "special":
			action = special(charName, charClass)
		}
		fmt.Println(action)
	}
	return "тренировка окончена"
}

// обратите внимание на имя функции и имена переменных
func choiseCharClass() string {
	var approveChoice string
	var charClass string
	// создаем мапу description, в которой будем хранить описание персонажей
	description := map[string]string{
		"warrior": "Воитель — дерзкий воин ближнего боя. Сильный, выносливый и отважный.",
		"mage":    "Маг — находчивый воин дальнего боя. Обладает высоким интеллектом.",
		"healer":  "Лекарь — могущественный заклинатель. Черпает силы из природы, веры и духов.",
	}

	// пользователь выбирает персонажа
	for approveChoice != "y" {
		fmt.Print("Введи название персонажа, за которого хочешь играть: Воитель — warrior, Маг — mage, Лекарь — healer: ")
		fmt.Scanf("%s\n", &charClass)
		fmt.Println(description[charClass])
		fmt.Print("Нажми (Y), чтобы подтвердить выбор, или любую другую кнопку, чтобы выбрать другого персонажа: ")
		fmt.Scanf("%s\n", &approveChoice)
		approveChoice = strings.ToLower(approveChoice)
	}
	return charClass
}

// обратите внимание на имена переменных
func main() {
	fmt.Println(`Приветствую тебя, искатель приключений!
Прежде чем начать игру...`)
	// пользователь вводит свое имя
	var charName string
	fmt.Print("...назови себя: ")
	fmt.Scanf("%s\n", &charName)

	fmt.Printf("Здравствуй, %s\n", charName)
	fmt.Println(`Сейчас твоя выносливость — 80, атака — 5 и защита — 10.
Ты можешь выбрать один из трёх путей силы:
Воитель, Маг, Лекарь`)

	charClass := choiseCharClass()

	fmt.Println(startTraining(charName, charClass))
}
