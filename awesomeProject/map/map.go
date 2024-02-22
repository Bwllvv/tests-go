package main

import "fmt"
import "github.com/mitchellh/mapstructure"

type point struct {
	x int
	y int
}

func main() {
	pointsmap := map[string]point{}      // стринг ключ
	otherpointmap := make(map[int]point) //инт ключ
	//var morepointmap map[string]point    // обьявлен но не инициализирован поэтому == nil

	pointsmap["a"] = point{ // стринг ключ
		x: 1,
		y: 2,
	}
	pointsmap["b"] = point{ // стринг ключ
		x: 2,
		y: 3,
	}

	//вывод по ключу и оформление
	fmt.Println(pointsmap["a"], pointsmap)
	fmt.Printf("длина: %d, высота: %d", pointsmap["a"].x, pointsmap["a"].y)

	//карта с инт ключем
	otherpointmap[1] = point{ // инт ключ
		x: 52,
		y: 31,
	}
	fmt.Println(otherpointmap[1])

	// проверка на имеющийся ключ

	value, ok := pointsmap["a"]
	if ok == true {
		fmt.Println("yes, i have")
		fmt.Print(value)
	} else {
		fmt.Println("no, i haven't")
	}

	// Создадим структуру из карты
	p1 := point{}
	mapstructure.Decode(pointsmap, &p1) //передаем именно ссылку на точку
	fmt.Println(p1)
	// до конца не понял как это работает
}
